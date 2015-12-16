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
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"
)

const e0 = `<!DOCTYPE html>
<html lang="en">
<head><meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1">

<meta http-equiv="X-UA-Compatible" content="IE=edge" />
<title>ORA-00000 to ORA-00851</title>
<meta name="generator" content="Oracle DARB XHTML Converter (Mode = document) - Version 5.1 Build 125" />
<meta name="dcterms.created" content="2008-11-24T14:48:31Z" />
<meta name="robots" content="all" />
<meta name="dcterms.title" content="Database Error Messages" />
<meta name="dcterms.identifier" content="B28278-02" />
<meta name="dcterms.isVersionOf" content="ERRMG" />
<link rel="Start" href="../../index.htm" title="Home" type="text/html" />
<link rel="Copyright" href="../../dcommon/html/cpyr.htm" title="Copyright" type="text/html" />

<script type="application/javascript"   src="../../dcommon/js/headfoot.js"></script>
<script type="application/javascript"   src="../../nav/js/doccd.js"></script>
<link rel="Contents" href="toc.htm" title="Contents" type="text/html" />
<link rel="Index" href="index.htm" title="Index" type="text/html" />
<link rel="Prev" href="intro.htm" title="Previous" type="text/html" />
<link rel="Next" href="e900.htm" title="Next" type="text/html" />
<link rel="alternate" href="../B28278-02.epub" title="ePub version" type="application/epub+zip" />
<link rel="alternate" href="../B28278-02.mobi" title="Mobipocket version" type="application/x-mobipocket-ebook" />
<link rel="schema.dcterms" href="http://purl.org/dc/terms/" />
<link rel="stylesheet" href="../../dcommon/css/fusiondoc.css">
<link rel="stylesheet" type="text/css"  href="../../dcommon/css/header.css">
<link rel="stylesheet" type="text/css"  href="../../dcommon/css/footer.css">
<link rel="stylesheet" type="text/css"  href="../../dcommon/css/fonts.css">
<link rel="stylesheet" href="../../dcommon/css/foundation.css">
<link rel="stylesheet" href="../../dcommon/css/codemirror.css">
<link rel="stylesheet" type="text/css" title="Default" href="../../nav/css/html5.css">
<link rel="stylesheet" href="../../dcommon/css/respond-480-tablet.css">
<link rel="stylesheet" href="../../dcommon/css/respond-768-laptop.css">
<link rel="stylesheet" href="../../dcommon/css/respond-1140-deskop.css">
<script type="application/javascript"  src="../../dcommon/js/modernizr.js"></script>
<script type="application/javascript"  src="../../dcommon/js/codemirror.js"></script>
<script type="application/javascript"  src="../../dcommon/js/jquery.js"></script>
<script type="application/javascript"  src="../../dcommon/js/foundation.min.js"></script>
<script type="application/javascript"  src="//s7.addthis.com/js/300/addthis_widget.js#pubid=ra-552992c80ef99c8d" async="async"></script>
<script type="application/javascript"  src="../../dcommon/js/jqfns.js"></script>
<script type="application/javascript"  src="../../dcommon/js/ohc-inline-videos.js"></script>
<!-- Add fancyBox -->
<link rel="stylesheet" href="../../dcommon/fancybox/jquery.fancybox.css?v=2.1.5" type="text/css" media="screen" />
<script type="application/javascript"  src="../../dcommon/fancybox/jquery.fancybox.pack.js?v=2.1.5"></script>
<!-- Optionally add helpers - button, thumbnail and/or media -->
<link rel="stylesheet"  href="../../dcommon/fancybox/helpers/jquery.fancybox-buttons.css?v=1.0.5"  type="text/css" media="screen" />
<script type="application/javascript"  src="../../dcommon/fancybox/helpers/jquery.fancybox-buttons.js?v=1.0.5"></script>
<script type="application/javascript"  src="../../dcommon/fancybox/helpers/jquery.fancybox-media.js?v=1.0.6"></script>
<link rel="stylesheet"  href="../../dcommon/fancybox/helpers/jquery.fancybox-thumbs.css?v=1.0.7"  type="text/css" media="screen" />
<script type="application/javascript"  src="../../dcommon/fancybox/helpers/jquery.fancybox-thumbs.js?v=1.0.7"></script>
<link rel="schema.dcterms" href="http://purl.org/dc/terms/" />
</head>
<body>
<header><!--
<div class="zz-skip-header"><a id="top" href="#BEGIN">Skip Headers</a>--></header>
<div class="row" id="CONTENT">
<div class="IND large-9 medium-8 columns"><span id="PAGE" style="display:none;">5/86</span> <!-- End Header --><a id="sthref17"></a>
<h1 class="chapter"><span class="secnum">2</span> ORA-00000 to ORA-00851</h1>
<div class="msgset">
<div class="msgentry">
<dl>
<dt><span class="msg"><a id="sthref18"></a><a id="ORA-00000"></a>ORA-00000: normal, successful completion</span> <!-- class="msg" --></dt>
<dd>
<div class="msgexplan"><span class="msgexplankw">Cause:</span> Normal exit.</div>
<!-- class="msgexplan" --></dd>
<dd>
<div class="msgaction"><span class="msgactionkw">Action:</span> None</div>
<!-- class="msgaction" --></dd>
</dl>
</div>
<!-- class="msgentry" -->
<div class="msgentry">
<dl>
<dt><span class="msg"><a id="sthref19"></a><a id="ORA-00001"></a>ORA-00001: unique constraint (<span class="italic">string</span>.<span class="italic">string</span>) violated</span> <!-- class="msg" --></dt>
<dd>
<div class="msgexplan"><span class="msgexplankw">Cause:</span> An UPDATE or INSERT statement attempted to insert a duplicate key. For Trusted Oracle configured in DBMS MAC mode, you may see this message if a duplicate entry exists at a different level.</div>
<!-- class="msgexplan" --></dd>
<dd>
<div class="msgaction"><span class="msgactionkw">Action:</span> Either remove the unique restriction or do not insert the key.</div>
<!-- class="msgaction" --></dd>
</dl>
</div>
<!-- class="msgentry" -->
<div class="msgentry">
<dl>
<dt><span class="msg"><a id="sthref20"></a><a id="ORA-00017"></a>ORA-00017: session requested to set trace event</span> <!-- class="msg" --></dt>
<dd>
<div class="msgexplan"><span class="msgexplankw">Cause:</span> The current session was requested to set a trace event by another session.</div>
<!-- class="msgexplan" --></dd>
<dd>
<div class="msgaction"><span class="msgactionkw">Action:</span> This is used internally; no action is required.</div>
<!-- class="msgaction" --></dd>
</dl>
</div>
<!-- class="msgentry" -->
<div class="msgentry">
<dl>
<dt><span class="msg"><a id="sthref21"></a><a id="ORA-00018"></a>ORA-00018: maximum number of sessions exceeded</span> <!-- class="msg" --></dt>
<dd>
<div class="msgexplan"><span class="msgexplankw">Cause:</span> All session state objects are in use.</div>
<!-- class="msgexplan" --></dd>
<dd>
<div class="msgaction"><span class="msgactionkw">Action:</span> Increase the value of the SESSIONS initialization parameter.</div>
<!-- class="msgaction" --></dd>
</dl>
</div>
<!-- class="msgentry" -->
<div class="msgentry">
<dl>
<dt><span class="msg"><a id="sthref22"></a><a id="ORA-00019"></a>ORA-00019: maximum number of session licenses exceeded</span> <!-- class="msg" --></dt>
<dd>
<div class="msgexplan"><span class="msgexplankw">Cause:</span> All licenses are in use.</div>
<!-- class="msgexplan" --></dd>
<dd>
<div class="msgaction"><span class="msgactionkw">Action:</span> Increase the value of the LICENSE MAX SESSIONS initialization parameter.</div>
<!-- class="msgaction" --></dd>
</dl>
</div>
<!-- class="msgentry" -->
<div class="msgentry">
<dl>
<dt><span class="msg"><a id="sthref436"></a><a id="ORA-00850"></a>ORA-00850: PGA_AGGREGATE_TARGET <span class="italic">string</span> cannot be set to more than MEMORY_MAX_TARGET <span class="italic">string</span>.</span> <!-- class="msg" --></dt>
<dd>
<div class="msgexplan"><span class="msgexplankw">Cause:</span> PGA_AGGREGATE_TARGET value was more than MEMORY_MAX_TARGET value.</div>
<!-- class="msgexplan" --></dd>
<dd>
<div class="msgaction"><span class="msgactionkw">Action:</span> Set PGA_AGGREGATE_TARGET to be less than MEMORY_MAX_TARGET.</div>
<!-- class="msgaction" --></dd>
</dl>
</div>
<!-- class="msgentry" -->
<div class="msgentry">
<dl>
<dt><span class="msg"><a id="sthref437"></a><a id="ORA-00851"></a>ORA-00851: SGA_MAX_SIZE <span class="italic">string</span> cannot be set to more than MEMORY_TARGET <span class="italic">string</span>.</span> <!-- class="msg" --></dt>
<dd>
<div class="msgexplan"><span class="msgexplankw">Cause:</span> SGA_MAX_SIZE value was more than MEMORY_TARGET value.</div>
<!-- class="msgexplan" --></dd>
<dd>
<div class="msgaction"><span class="msgactionkw">Action:</span> Set SGA_MAX_SIZE to be less than MEMORY_TARGET.</div>
<!-- class="msgaction" --></dd>
</dl>
</div>
<!-- class="msgentry" --></div>
<!-- class="msgset" --></div>
<!-- class="ind" -->
<!-- Start Footer -->
</div>
<!-- add extra wrapper close div-->
<footer><!--
<hr />
<table class="cellalignment1228">
<tr>
<td class="cellalignment1235">
<table class="cellalignment1233">
<tr>
<td class="cellalignment1232"><a href="intro.htm"><img width="24" height="24" src="../../dcommon/gifs/leftnav.gif" alt="Go to previous page" /><br />
<span class="icon">Previous</span></a></td>
<td class="cellalignment1232"><a href="e900.htm"><img width="24" height="24" src="../../dcommon/gifs/rightnav.gif" alt="Go to next page" /><br />
<span class="icon">Next</span></a></td>
</tr>
</table>
</td>
<td class="cellalignment-copyrightlogo"><img width="144" height="18" src="../../dcommon/gifs/oracle.gif" alt="Oracle" /><br />
Copyright&nbsp;&copy;&nbsp;2008,&nbsp;Oracle.&nbsp;All&nbsp;rights&nbsp;reserved.<br />
<a href="../../dcommon/html/cpyr.htm">Legal Notices</a></td>
<td class="cellalignment1237">
<table class="cellalignment1231">
<tr>
<td class="cellalignment1232"><a href="../../index.htm"><img width="24" height="24" src="../../dcommon/gifs/doclib.gif" alt="Go to Documentation Home" /><br />
<span class="icon">Home</span></a></td>
<td class="cellalignment1232"><a href="../../nav/portal_booklist.htm"><img width="24" height="24" src="../../dcommon/gifs/booklist.gif" alt="Go to Book List" /><br />
<span class="icon">Book List</span></a></td>
<td class="cellalignment1232"><a href="toc.htm"><img width="24" height="24" src="../../dcommon/gifs/toc.gif" alt="Go to Table of Contents" /><br />
<span class="icon">Contents</span></a></td>
<td class="cellalignment1232"><a href="index.htm"><img width="24" height="24" src="../../dcommon/gifs/index.gif" alt="Go to Index" /><br />
<span class="icon">Index</span></a></td>
<td class="cellalignment1232"><a href="../../nav/mindx.htm"><img width="24" height="24" src="../../dcommon/gifs/masterix.gif" alt="Go to Master Index" /><br />
<span class="icon">Master Index</span></a></td>
<td class="cellalignment1232"><a href="../../dcommon/html/feedback.htm"><img width="24" height="24" src="../../dcommon/gifs/feedbck2.gif" alt="Go to Feedback page" /><br />
<span class="icon">Contact Us</span></a></td>
</tr>
</table>
</td>
</tr>
</table>
--></footer>
<noscript>
<p>Scripting on this page enhances content navigation, but does not change the content in any way.</p>
</noscript>
</body>
</html>`

func TestParse(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	out := make(chan Message, 1)
	go func() {
		defer close(out)
		if err := parseMessages(ctx, out, strings.NewReader(e0)); err != nil {
			t.Error(err)
		}
	}()

	n := 0
	for msg := range out {
		t.Logf("msg=%s", msg)
		n++
		if msg.Code == 1 {
			await := "unique constraint (string.string) violated"
			if msg.Description != await {
				t.Errorf("ORA-00001: got %q, waited %q.", msg.Description, await)
			}
		}
	}
	if n < 7 {
		t.Errorf("Parsed only %d!", n)
	}
}
