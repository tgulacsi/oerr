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

// Copyright Tamás Gulácsi 2015. All rights reserved.

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	oerr "github.com/tgulacsi/oerr/lib"
)

//go:generate sh -c "go install && oerr download oerr.db"

func main() {
	dbPath := "oerr.db"
	if os.Getenv("BRUNO_HOME") != "" {
		dbPath = os.ExpandEnv("$BRUNO_HOME/data/ws/oerr.db")
	}
	URL := oerr.URL

	mainCmd := &cobra.Command{
		Use: "oerr",
	}
	mainCmd.PersistentFlags().StringVarP(&dbPath, "db", "D", dbPath, "path of the Bolt DB of Oracle Error Messages")

	downloadCmd := &cobra.Command{
		Use: "download",
		Run: func(_ *cobra.Command, args []string) {
			if err := oerr.DownloadInto(dbPath, URL); err != nil {
				log.Fatalf("DownloadInto(%q, %q): %v", dbPath, URL, err)
			}
		},
	}
	downloadCmd.Flags().StringVarP(&URL, "url", "", URL, "URL of TOC")
	mainCmd.AddCommand(downloadCmd)

	getCmd := &cobra.Command{
		Use: "get",
		Run: func(_ *cobra.Command, args []string) {
			id := oerr.MsgID{Prefix: "ORA"}
			txt := strings.TrimSpace(args[0])
			i := strings.IndexByte(txt, '-')
			if i >= 3 {
				id.Prefix = strings.ToUpper(txt[:3])
			}
			code, err := strconv.Atoi(txt[i+1:])
			if err != nil {
				log.Fatalf("Parse %q as integer: %v", txt[i+1:], err)
			}
			id.Code = uint32(code)

			db, err := oerr.Open(dbPath)
			if err != nil {
				log.Fatalf("Open %q: %v", dbPath, err)
			}
			data, err := db.Get(id)
			if err != nil {
				log.Printf("get %s: %v", id, err)
			} else {
				fmt.Fprintf(os.Stderr, "%s: %v\n", id, data)
			}
		},
	}
	mainCmd.AddCommand(getCmd)

	if _, _, err := mainCmd.Find(os.Args[1:]); err != nil {
		mainCmd.SetArgs(append([]string{"get"}, os.Args[1:]...))
	}
	mainCmd.Execute()
}
