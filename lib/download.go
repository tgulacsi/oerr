// Copyright 2015 Tamás Gulácsi
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package oerr

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/yhat/scrape"
	"go4.org/syncutil"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"golang.org/x/net/html/charset"
)

var URL = `http://docs.oracle.com/cd/B28359_01/server.111/b28278/toc.htm`

const bucketName = "oerr"

// DownloadInto fills the DB by downloading the messages.
func DownloadInto(dbPath, tocURL string) error {
	os.Remove(dbPath)
	db, err := bolt.Open(dbPath, 0664, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	msgCh := make(chan Message, 8)
	go func() {
		db.NoSync = true
		defer db.Sync()
		if err := db.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
			if err != nil {
				return err
			}

			for msg := range msgCh {
				key, err := msg.MsgID.MarshalBinary()
				if err != nil {
					log.Printf("cannot marshal %#v: %v", msg.MsgID, err)
					continue
				}
				val, err := msg.MsgData.MarshalBinary()
				if err != nil {
					log.Printf("cannot marshal %#v: %v", msg.MsgData, err)
					continue
				}
				if err := bucket.Put(key, val); err != nil {
					log.Printf("Put(%#v): %v", msg, err)
					return err
				}
			}
			return nil
		}); err != nil {
			log.Printf("Update: %v", err)
		}
	}()

	if err := Download(context.Background(), msgCh, tocURL); err != nil {
		return err
	}
	return nil
}

// Download into the given channel, from the given URL.
func Download(ctx context.Context, out chan<- Message, tocURL string) error {
	defer func() { close(out) }()
	body, err := dl(ctx, URL)
	if err != nil {
		return err
	}
	links, err := parseLinks(ctx, body)
	body.Close()
	if err != nil {
		return err
	}

	baseURL := URL[:strings.LastIndex(URL, "/")]
	gate := syncutil.NewGate(8)
	var grp syncutil.Group
	for _, lnk := range links {
		if i := strings.IndexByte(lnk, '#'); i >= 0 {
			lnk = lnk[:i]
		}
		switch lnk {
		case "title.htm", "preface.htm", "intro.htm", "index.htm":
			continue
		}
		lnk := lnk
		grp.Go(func() error {
			gate.Start()
			defer gate.Done()
			body, err := dl(ctx, baseURL+"/"+lnk)
			if err != nil {
				return err
			}
			defer body.Close()
			return parseMessages(ctx, out, body)
		})
	}
	return grp.Err()
}
func parseMessages(ctx context.Context, out chan<- Message, body io.Reader) error {
	doc, err := html.Parse(body)
	if err != nil {
		return err
	}

	for _, n := range scrape.FindAll(doc, func(n *html.Node) bool {
		return n.DataAtom == atom.Div && scrape.Attr(n, "class") == "msgentry"
	}) {
		dt, ok := scrape.Find(n, scrape.ByTag(atom.Dt))
		if !ok {
			continue
		}
		var msg Message
		line := scrape.TextJoin(dt, func(x []string) string { return strings.Join(x, "") })
		i := strings.IndexByte(line, ':')
		j := strings.IndexByte(line[:i], '-')
		var code string
		msg.Prefix, code, msg.Description = strings.ToUpper(line[:j]), line[j+1:i], strings.TrimSpace(line[i+1:])
		codeI, err := strconv.Atoi(code)
		if err != nil {
			log.Printf("parse %q: %v", code, err)
		}
		msg.Code = uint32(codeI)

		for _, s := range scrape.FindAll(n, func(n *html.Node) bool {
			if n.DataAtom != atom.Div {
				return false
			}
			cls := scrape.Attr(n, "class")
			return cls == "msgexplan" || cls == "msgaction"
		}) {
			if scrape.Attr(s, "class") == "msgaction" {
				msg.Action = strings.TrimPrefix(scrape.Text(s), "Action: ")
			} else {
				msg.Cause = strings.TrimPrefix(scrape.Text(s), "Cause: ")
			}
		}
		out <- msg
	}

	return nil
}

func parseLinks(ctx context.Context, body io.Reader) ([]string, error) {
	links := make([]string, 0, 1024)
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			if z.Err() == io.EOF {
				break
			}
			log.Printf("tokenizer returned %v", z.Err())
			return links, z.Err()
		}
		if tt != html.StartTagToken {
			continue
		}
		nm, hasAttr := z.TagName()
		if hasAttr && bytes.Equal(nm, []byte("a")) {
			for {
				k, v, more := z.TagAttr()
				if bytes.Equal(k, []byte("href")) {
					links = append(links, html.UnescapeString(string(v)))
					break
				}
				if !more {
					break
				}
				k, v, more = z.TagAttr()
			}
		}
	}
	return links, nil
}

func getAttr(tok html.Token, name string) string {
	for _, attr := range tok.Attr {
		if attr.Key == name {
			return attr.Val
		}
	}
	return ""
}

func dl(ctx context.Context, URL string) (io.ReadCloser, error) {
	resp, err := ctxhttp.Get(ctx, http.DefaultClient, URL)
	if err != nil {
		return nil, err
	}
	r, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	return struct {
		io.Reader
		io.Closer
	}{r, resp.Body}, nil
}
