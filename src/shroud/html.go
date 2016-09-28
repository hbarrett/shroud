package main

const getSecretTmpl string =`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml"><head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8"></meta>
<title>Create a Password with Shroud</title>
<style rel="stylesheet" type="text/css" media="all" data-href="http://www.phpform.org/formbuilder/colors/color2/view.css">body
{
	background:#336699;
	font-family:"Lucida Grande", Tahoma, Arial, Verdana, sans-serif;
	font-size:small;
	margin:8px 0 16px;
	text-align:center;
}

#form_container
{
	background:#fff;
	
	margin:0 auto;
	text-align:left;
	width:640px;
}

#top
{
	display:block;
	height:10px;
	margin:10px auto 0;
	width:650px;
}

#footer
{
	width:640px;
	clear:both;
	color:#999999;
	text-align:center;
	width:640px;
	padding-bottom: 15px;
	font-size: 85%;
}

#footer a{
	color:#999999;
	text-decoration: none;
	border-bottom: 1px dotted #999999;
}

#bottom
{
	display:block;
	height:10px;
	margin:0 auto;
	width:650px;
}

form.appnitro
{
	margin:20px 20px 0;
	padding:0 0 20px;
}

/**** Logo Section  *****/
h1
{
	background-color:#6699CC;
	margin:0;
	min-height:0;
	padding:0;
	text-decoration:none;
	text-indent:-8000px;
	
}

h1 a
{
	
	display:block;
	height:100%;
	min-height:40px;
	overflow:hidden;
}


img
{
	behavior:url(data:text/html;base64,PCFET0NUWVBFIGh0bWwgUFVCTElDICItLy9XM0MvL0RURCBYSFRNTCAxLjAgVHJhbnNpdGlvbmFsLy9FTiIgImh0dHA6Ly93d3cudzMub3JnL1RSL3hodG1sMS9EVEQveGh0bWwxLXRyYW5zaXRpb25hbC5kdGQiPgo8aHRtbCB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94aHRtbCIgbGFuZz0iZW4iIHhtbDpsYW5nPSJlbiI+IAo8aGVhZD4KCTx0aXRsZT5QYWdlIG5vdCBmb3VuZCB8IFBIUEZvcm0gLSBVbHRpbWF0ZSBGb3JtIEJ1aWxkaW5nIFJlc291cmNlPC90aXRsZT4KCTxtZXRhIGh0dHAtZXF1aXY9IkNvbnRlbnQtVHlwZSIgY29udGVudD0idGV4dC9odG1sOyBjaGFyc2V0PXV0Zi04IiAvPgo8bGluayByZWw9InNob3J0Y3V0IGljb24iIGhyZWY9Ii90aGVtZXMvYXBwbml0cm8vZmF2aWNvbi5pY28iIHR5cGU9ImltYWdlL3gtaWNvbiIgLz4KICAJPHN0eWxlIHR5cGU9InRleHQvY3NzIiBtZWRpYT0iYWxsIj5AaW1wb3J0ICIvbW9kdWxlcy9ub2RlL25vZGUuY3NzIjs8L3N0eWxlPgo8c3R5bGUgdHlwZT0idGV4dC9jc3MiIG1lZGlhPSJhbGwiPkBpbXBvcnQgIi9tb2R1bGVzL3N5c3RlbS9kZWZhdWx0cy5jc3MiOzwvc3R5bGU+CjxzdHlsZSB0eXBlPSJ0ZXh0L2NzcyIgbWVkaWE9ImFsbCI+QGltcG9ydCAiL21vZHVsZXMvc3lzdGVtL3N5c3RlbS5jc3MiOzwvc3R5bGU+CjxzdHlsZSB0eXBlPSJ0ZXh0L2NzcyIgbWVkaWE9ImFsbCI+QGltcG9ydCAiL21vZHVsZXMvdXNlci91c2VyLmNzcyI7PC9zdHlsZT4KPHN0eWxlIHR5cGU9InRleHQvY3NzIiBtZWRpYT0iYWxsIj5AaW1wb3J0ICIvdGhlbWVzL2FwcG5pdHJvL3N0eWxlLmNzcyI7PC9zdHlsZT4KICAJPG1ldGEgbmFtZT0idmlld3BvcnQiIGNvbnRlbnQ9IndpZHRoPWRldmljZS13aWR0aCwgaW5pdGlhbC1zY2FsZT0xIj4KPGxpbmsgcmVsPSJzdHlsZXNoZWV0IiB0eXBlPSJ0ZXh0L2NzcyIgaHJlZj0iL21vYmlsZS5jc3MiIG1lZGlhPSJhbGwiIC8+CjwvaGVhZD4KCjxib2R5IGlkPSJkZWZhdWx0X3BhZ2UiIGNsYXNzPSIiPgoJCTxkaXYgaWQ9Im1haW5fY29udGFpbmVyIj4KCQkJPGRpdiBjbGFzcz0id3JhcHBlciI+CgkJCTxkaXYgaWQ9ImhlYWRlciI+Cgk8aDE+PGEgaHJlZj0iaHR0cDovL3d3dy5waHBmb3JtLm9yZyIgdGl0bGU9IkFwcG5pdHJvIHBGb3JtIj4mbmJzcDs8L2E+PC9oMT4KCTx1bCA+CgkJCQk8bGkgY2xhc3M9ImJ1dHRvbl9zdGVwMyI+CgkJCTxhPiZuYnNwOzwvYT4KCQk8L2xpPgoJCQoJPC91bD4KPC9kaXY+CQoJCQk8ZGl2IGlkPSJjb250ZW50X3Jvb3QiPgoJCQkJPGRpdiBjbGFzcz0id3JhcHBlciI+CgkJCQkJPGRpdiBpZD0iY29udGVudF9oZWFkZXIiPgoJCQkJCQkmbmJzcDsKCQkJCQk8L2Rpdj4KCQkJCQk8ZGl2IGlkPSJjb250ZW50X2xlZnQiIHN0eWxlPSJ3aWR0aDogNzYwcHgiPgoJCQkJCQk8aDEgY2xhc3M9InRpdGxlIj5QYWdlIG5vdCBmb3VuZDwvaDE+CQoJCQkJCQkgCgkJCQkJCQkJCQkJCSAgICAgICAgCQkJCQkJCQkJCQkJCQkJCSAgICAgICAgCQkJCQkJCQkJPC9kaXY+CgkJCQkJCQkJCTwvZGl2PgoJCQk8L2Rpdj4KCQkJICAgICAgICAJCjxkaXYgaWQ9ImZvb3RlciI+CgkJCQk8aDQ+CgkJCQkJJmNvcHk7IENvcHlyaWdodCAyMDA3LTIwMTUgLSZuYnNwOzxhIGhyZWY9Imh0dHA6Ly93d3cuYXBwbml0cm8uY29tIiBjbGFzcz0iZG90dGVkX2xpbmsiIHN0eWxlPSJjb2xvcjogI2ZmZiI+Rm9ybSBCdWlsZGVyPC9hPiAtIEFwcG5pdHJvIFNvZnR3YXJlPC9zdHJvbmc+CgkJCQk8L2g0PgoJCQkKCQkJCSZuYnNwOwoJCQkJCQkJCjwvZGl2Pgo8IS0tIFN0YXJ0IG9mIFN0YXRDb3VudGVyIENvZGUgZm9yIERlZmF1bHQgR3VpZGUgLS0+CjxzY3JpcHQgdHlwZT0idGV4dC9qYXZhc2NyaXB0Ij4KdmFyIHNjX3Byb2plY3Q9MzAwMjcyNjsgCnZhciBzY19pbnZpc2libGU9MTsgCnZhciBzY19zZWN1cml0eT0iZTlkZDM0ZDEiOyAKdmFyIHNjX2h0dHBzPTE7IAp2YXIgc2NfcmVtb3ZlX2xpbms9MTsgCnZhciBzY0pzSG9zdCA9ICgoImh0dHBzOiIgPT0gZG9jdW1lbnQubG9jYXRpb24ucHJvdG9jb2wpID8KImh0dHBzOi8vc2VjdXJlLiIgOiAiaHR0cDovL3d3dy4iKTsKZG9jdW1lbnQud3JpdGUoIjxzYyIrInJpcHQgdHlwZT0ndGV4dC9qYXZhc2NyaXB0JyBzcmM9JyIgKwpzY0pzSG9zdCsKInN0YXRjb3VudGVyLmNvbS9jb3VudGVyL2NvdW50ZXIuanMnPjwvIisic2NyaXB0PiIpOwo8L3NjcmlwdD4KPG5vc2NyaXB0PjxkaXYgY2xhc3M9InN0YXRjb3VudGVyIj48aW1nIGNsYXNzPSJzdGF0Y291bnRlciIKc3JjPSJodHRwOi8vYy5zdGF0Y291bnRlci5jb20vMzAwMjcyNi8wL2U5ZGQzNGQxLzEvIgphbHQ9IndlYiBhbmFseXRpY3MiPjwvZGl2Pjwvbm9zY3JpcHQ+CjwhLS0gRW5kIG9mIFN0YXRDb3VudGVyIENvZGUgZm9yIERlZmF1bHQgR3VpZGUgLS0+Cgo8c2NyaXB0IHNyYz0iaHR0cDovL3d3dy5nb29nbGUtYW5hbHl0aWNzLmNvbS91cmNoaW4uanMiIHR5cGU9InRleHQvamF2YXNjcmlwdCI+Cjwvc2NyaXB0Pgo8c2NyaXB0IHR5cGU9InRleHQvamF2YXNjcmlwdCI+Cl91YWNjdCA9ICJVQS0yNDQ0MzMwLTIiOwp1cmNoaW5UcmFja2VyKCk7Cjwvc2NyaXB0PgoJCTwvZGl2PgoJPC9kaXY+CjwvYm9keT4KCgo8L2h0bWw+);
	border:none;
}


/**** Form Section ****/
.appnitro
{
	font-family:Lucida Grande, Tahoma, Arial, Verdana, sans-serif;
	font-size:small;
}

.appnitro li
{
	width:61%;
}

form ul
{
	font-size:100%;
	list-style-type:none;
	margin:0;
	padding:0;
	width:100%;
}

form li
{
	display:block;
	margin:0;
	padding:4px 5px 2px 9px;
	position:relative;
}

form li:after
{
	clear:both;
	content:".";
	display:block;
	height:0;
	visibility:hidden;
}

.buttons:after
{
	clear:both;
	content:".";
	display:block;
	height:0;
	visibility:hidden;
}

.buttons
{
	clear:both;
	display:block;
	margin-top:10px;
}

* html form li
{
	height:1%;
}

* html .buttons
{
	height:1%;
}

* html form li div
{
	display:inline-block;
}

form li div
{
	color:#444;
	margin:0 4px 0 0;
	padding:0 0 8px;
}

form li span
{
	color:#444;
	float:left;
	margin:0 4px 0 0;
	padding:0 0 8px;
}

form li div.left
{
	display:inline;
	float:left;
	width:48%;
}

form li div.right
{
	display:inline;
	float:right;
	width:48%;
}

form li div.left .medium
{
	width:100%;
}

form li div.right .medium
{
	width:100%;
}

.clear
{
	clear:both;
}

form li div label
{
	clear:both;
	color:#444;
	display:block;
	font-size:9px;
	line-height:9px;
	margin:0;
	padding-top:3px;
}

form li span label
{
	clear:both;
	color:#444;
	display:block;
	font-size:9px;
	line-height:9px;
	margin:0;
	padding-top:3px;
}

form li .datepickerimg
{
	cursor:pointer !important;
	float:left;
	height:16px;
	margin:.1em 5px 0 0;
	padding:0;
	width:16px;
}

.form_description
{
	border-bottom:1px dotted #ccc;
	clear:both;
	display:inline-block;
	margin:0 0 1em;
}

.form_description[class]
{
	display:block;
}

.form_description h2
{
	clear:left;
	font-size:160%;
	font-weight:400;
	margin:0 0 3px;
}

.form_description p
{
	font-size:95%;
	line-height:130%;
	margin:0 0 12px;
}

form hr
{
	display:none;
}

form li.section_break
{
	border-top:1px dotted #ccc;
	margin-top:9px;
	padding-bottom:0;
	padding-left:9px;
	padding-top:13px;
	width:97% !important;
}

form ul li.first
{
	border-top:none !important;
	margin-top:0 !important;
	padding-top:0 !important;
}

form .section_break h3
{
	font-size:110%;
	font-weight:400;
	line-height:130%;
	margin:0 0 2px;
}

form .section_break p
{
	font-size:85%;

	margin:0 0 10px;
}

/**** Buttons ****/
input.button_text
{
	overflow:visible;
	padding:0 7px;
	width:auto;
}

.buttons input
{
	font-size:120%;
	margin-right:5px;
}

/**** Inputs and Labels ****/
label.description
{
	border:none;
	color:#222;
	display:block;
	font-size:95%;
	font-weight:700;
	line-height:150%;
	padding:0 0 1px;
}

span.symbol
{
	font-size:115%;
	line-height:130%;
}

input.text
{
	background:#fff url(data:image/gif;base64,R0lGODlhBwACAIABAMzMzP///yH5BAEAAAEALAAAAAAHAAIAAAIFhA+hGwUAOw==) repeat-x top;
	border-bottom:1px solid #ddd;
	border-left:1px solid #c3c3c3;
	border-right:1px solid #c3c3c3;
	border-top:1px solid #7c7c7c;
	color:#333;
	font-size:100%;
	margin:0;
	padding:2px 0;
}

input.file
{
	color:#333;
	font-size:100%;
	margin:0;
	padding:2px 0;
}

textarea.textarea
{
	background:#fff url(data:image/gif;base64,R0lGODlhBwACAIABAMzMzP///yH5BAEAAAEALAAAAAAHAAIAAAIFhA+hGwUAOw==) repeat-x top;
	border-bottom:1px solid #ddd;
	border-left:1px solid #c3c3c3;
	border-right:1px solid #c3c3c3;
	border-top:1px solid #7c7c7c;
	color:#333;
	font-family:"Lucida Grande", Tahoma, Arial, Verdana, sans-serif;
	font-size:100%;
	margin:0;
	width:99%;
}

select.select
{
	color:#333;
	font-size:100%;
	margin:1px 0;
	padding:1px 0 0;
	background:#fff url(data:image/gif;base64,R0lGODlhBwACAIABAMzMzP///yH5BAEAAAEALAAAAAAHAAIAAAIFhA+hGwUAOw==) repeat-x top;
	border-bottom:1px solid #ddd;
	border-left:1px solid #c3c3c3;
	border-right:1px solid #c3c3c3;
	border-top:1px solid #7c7c7c;
}


input.currency
{
	text-align:right;
}

input.checkbox
{
	display:block;
	height:13px;
	line-height:1.4em;
	margin:6px 0 0 3px;
	width:13px;
}

input.radio
{
	display:block;
	height:13px;
	line-height:1.4em;
	margin:6px 0 0 3px;
	width:13px;
}

label.choice
{
	color:#444;
	display:block;
	font-size:100%;
	line-height:1.4em;
	margin:-1.55em 0 0 25px;
	padding:4px 0 5px;
	width:90%;
}

select.select[class]
{
	margin:0;
	padding:1px 0;
}

*:first-child+html select.select[class]
{
	margin:1px 0;
}

.safari select.select
{
	font-size:120% !important;
	margin-bottom:1px;
}

input.small
{
	width:25%;
}

select.small
{
	width:25%;
}

input.medium
{
	width:50%;
}

select.medium
{
	width:50%;
}

input.large
{
	width:99%;
}

select.large
{
	width:100%;
}

textarea.small
{
	height:5.5em;
}

textarea.medium
{
	height:10em;
}

textarea.large
{
	height:20em;
}

/**** Errors ****/
#error_message
{
	background:#fff;
	border:1px dotted red;
	margin-bottom:1em;
	padding-left:0;
	padding-right:0;
	padding-top:4px;
	text-align:center;
	width:99%;
}

#error_message_title
{
	color:#DF0000;
	font-size:125%;
	margin:7px 0 5px;
	padding:0;
}

#error_message_desc
{
	color:#000;
	font-size:100%;
	margin:0 0 .8em;
}

#error_message_desc strong
{
	background-color:#FFDFDF;
	color:red;
	padding:2px 3px;
}

form li.error
{
	background-color:#FFDFDF !important;
	border-bottom:1px solid #EACBCC;
	border-right:1px solid #EACBCC;
	margin:3px 0;
}

form li.error label
{
	color:#DF0000 !important;
}

form p.error
{
	clear:both;
	color:red;
	font-size:10px;
	font-weight:700;
	margin:0 0 5px;
}

form .required
{
	color:red;
	float:none;
	font-weight:700;
}

/**** Guidelines and Error Highlight ****/
form li.highlighted
{
	background-color:#fff7c0;
}

form .guidelines
{
	background:#f5f5f5;
	border:1px solid #e6e6e6;
	color:#444;
	font-size:80%;
	left:100%;
	line-height:130%;
	margin:0 0 0 8px;
	padding:8px 10px 9px;
	position:absolute;
	top:0;
	visibility:hidden;
	width:42%;
	z-index:1000;
}

form .guidelines small
{
	font-size:105%;
}

form li.highlighted .guidelines
{
	visibility:visible;
}

form li:hover .guidelines
{
	visibility:visible;
}

.no_guidelines .guidelines
{
	display:none !important;
}

.no_guidelines form li
{
	width:97%;
}

.no_guidelines li.section
{
	padding-left:9px;
}

/*** Success Message ****/
.form_success 
{
	clear: both;
	margin: 0;
	padding: 90px 0pt 100px;
	text-align: center
}

.form_success h2 {
    clear:left;
    font-size:160%;
    font-weight:normal;
    margin:0pt 0pt 3px;
}

/*** Password ****/
ul.password{
    margin-top:60px;
    margin-bottom: 60px;
    text-align: center;
}
.password h2{
    color:#DF0000;
    font-weight:bold;
    margin:0pt auto 10px;
}

.password input.text {
   font-size:170% !important;
   width:380px;
   text-align: center;
}
.password label{
   display:block;
   font-size:120% !important;
   padding-top:10px;
   font-weight:bold;
}

#li_captcha{
   padding-left: 5px;
}


#li_captcha span{
	float:none;
}

/** Embedded Form **/

.embed #form_container{
	border: none;
}

.embed #top, .embed #bottom, .embed h1{
	display: none;
}

.embed #form_container{
	width: 100%;
}

.embed #footer{
	text-align: left;
	padding-left: 10px;
	width: 99%;
}

.embed #footer.success{
	text-align: center;
}

.embed form.appnitro
{
	margin:0px 0px 0;
	
}



/*** Calendar **********************/
div.calendar { position: relative; }

.calendar table {
cursor:pointer;
border:1px solid #ccc;
font-size: 11px;
color: #000;
background: #fff;
font-family:"Lucida Grande", Tahoma, Arial, Verdana, sans-serif;
}

.calendar .button { 
text-align: center;    
padding: 2px;          
}

.calendar .nav {
background:#f5f5f5;
}

.calendar thead .title { 
font-weight: bold;      
text-align: center;
background: #dedede;
color: #000;
padding: 2px 0 3px 0;
}

.calendar thead .headrow { 
background: #f5f5f5;
color: #444;
font-weight:bold;
}

.calendar thead .daynames { 
background: #fff;
color:#333;
font-weight:bold;
}

.calendar thead .name { 
border-bottom: 1px dotted #ccc;
padding: 2px;
text-align: center;
color: #000;
}

.calendar thead .weekend { 
color: #666;
}

.calendar thead .hilite { 
background-color: #444;
color: #fff;
padding: 1px;
}

.calendar thead .active { 
background-color: #d12f19;
color:#fff;
padding: 2px 0px 0px 2px;
}


.calendar tbody .day { 
width:1.8em;
color: #222;
text-align: right;
padding: 2px 2px 2px 2px;
}
.calendar tbody .day.othermonth {
font-size: 80%;
color: #bbb;
}
.calendar tbody .day.othermonth.oweekend {
color: #fbb;
}

.calendar table .wn {
padding: 2px 2px 2px 2px;
border-right: 1px solid #000;
background: #666;
}

.calendar tbody .rowhilite td {
background: #FFF1AF;
}

.calendar tbody .rowhilite td.wn {
background: #FFF1AF;
}

.calendar tbody td.hilite { 
padding: 1px 1px 1px 1px;
background:#444 !important;
color:#fff !important;
}

.calendar tbody td.active { 
color:#fff;
background: #529214 !important;
padding: 2px 2px 0px 2px;
}

.calendar tbody td.selected { 
font-weight: bold;
border: 1px solid #888;
padding: 1px 1px 1px 1px;
background: #f5f5f5 !important;
color: #222 !important;
}

.calendar tbody td.weekend { 
color: #666;
}

.calendar tbody td.today { 
font-weight: bold;
color: #529214;
background:#D9EFC2;
}

.calendar tbody .disabled { color: #999; }

.calendar tbody .emptycell { 
visibility: hidden;
}

.calendar tbody .emptyrow { 
display: none;
}

.calendar tfoot .footrow { 
text-align: center;
background: #556;
color: #fff;
}

.calendar tfoot .ttip { 
background: #222;
color: #fff;
font-size:10px;
border-top: 1px solid #dedede;
padding: 3px;
}

.calendar tfoot .hilite { 
background: #aaf;
border: 1px solid #04f;
color: #000;
padding: 1px;
}

.calendar tfoot .active { 
background: #77c;
padding: 2px 0px 0px 2px;
}

.calendar .combo {
position: absolute;
display: none;
top: 0px;
left: 0px;
width: 4em;
border: 1px solid #ccc;
background: #f5f5f5;
color: #222;
font-size: 90%;
z-index: 100;
}

.calendar .combo .label,
.calendar .combo .label-IEfix {
text-align: center;
padding: 1px;
}

.calendar .combo .label-IEfix {
width: 4em;
}

.calendar .combo .hilite {
background: #444;
color:#fff;
}

.calendar .combo .active {
border-top: 1px solid #999;
border-bottom: 1px solid #999;
background: #dedede;
font-weight: bold;
}

</style>


<link type="image/x-icon" rel="shortcut icon" href="data:image/gif;base64,R0lGODlhAQABAIAAAP///////yH5BAEKAAEALAAAAAABAAEAAAICTAEAOw=="></link>
<link rel="stylesheet" href="https://ajax.googleapis.com/ajax/libs/jqueryui/1.11.4/themes/smoothness/jquery-ui.css"></link>
<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.3/jquery.min.js"></script>
<script src="https://ajax.googleapis.com/ajax/libs/jqueryui/1.11.4/jquery-ui.min.js"></script>
  
  <script>
  var tomorrow = new Date();
  tomorrow.setDate(tomorrow.getDate() + 1);

  $.datepicker.setDefaults({ dateFormat: 'yy-mm-dd', minDate: +2, maxDate: "+1W" });
  $(document).ready(function() { $("#datepicker").datepicker();$("#datepicker").datepicker("setDate", new Date());});
  </script>

</head>
<body id="main_body" class="no_guidelines safari">
	
	<img id="top" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAwIAAAAKCAYAAAAHB+lIAAAABGdBTUEAANbY1E9YMgAAABl0RVh0U29mdHdhcmUAQWRvYmUgSW1hZ2VSZWFkeXHJZTwAAAEzSURBVHja7NvrbsIwDAbQpBT2/q87xrIVtZPrXrQhlU3aOZJJIOW/Pxxqa60AAAD/Sz+81Fr3nqkPngEAAMdqj5wNw4D+G819XdlXgQAAAP5MAGg/DQb9RkO/tsbaCwUAAMDzwkBcY9V0Ft/PgsBe89+ldSsUAAAAzw0Bsd7TmkPAVxjYmwjE5n+oU9jnUCAMAADA74SA2PwPdRt78xwKNicCJYWALoWAPoWBriwnBAAAwPFBoKXmP4eAGAZKmU8IVoNAvg50CiHgnALBqSynAwAAwPFBIP7afxvrLQSAKQxMz9ewnwWBfMUnTgLOK5UnBNN3AACAY+UQMAWA61g1Nf75D8T3z9cmAvk60ND4Xz7rZVwvZT4hiGHAVAAAAI7TynIScB3DwOvYm2+FgKnuPgQYAGd6YIIkTAoCAAAAAElFTkSuQmCC" alt="">
	<div id="form_container">
	
		<h1><a>Create a new password to share.</a></h1>
		<form id="form_1158694" class="appnitro" method="post" action="/putsecret">
					<div class="form_description">
			<h2>Create a new password to share</h2>
			<p></p>
		</div>
			<ul>

					<li id="li_1">
		<label class="description" for="secret">Password</label>
		<div>
			<input id="secret" name="secret" class="element text medium" type="text" maxlength="255" value=""></input>
		</div>

		</li>		<li id="li_3">
		<label class="description" for="views">Number of views before secret is removed:</label>
		<div>
		<select class="element select medium" id="views" name="views">
			<!--<option value=5 selected="selected"></option>-->
			<option value="10">10</option>
			<option value="9">9</option>
			<option value="8">8</option>
			<option value="7">7</option>
			<option value="6">6</option>
			<option value="5">5</option>
			<option value="4">4</option>
			<option value="3">3</option>
			<option value="2">2</option>
			<option value="1">1</option>
		</select>
		</div>
		</li>		<li id="li_2">

                <label class="description" for="element_2">Date when secret will be removed:</label>
                <input id="datepicker" name="date"/>

		</form>
		</li>
					<li class="buttons">
			    <!--<input type="hidden" name="form_id" value="1158694">--></input>

				<input id="saveForm" class="button_text" type="submit" name="submit" value="Submit"></input>
		</li>
			</ul>
		</form>
		<div id="footer">
		</div>
	</div>
	<img id="bottom" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAwIAAAAKCAYAAAAHB+lIAAAABGdBTUEAANbY1E9YMgAAABl0RVh0U29mdHdhcmUAQWRvYmUgSW1hZ2VSZWFkeXHJZTwAAAFBSURBVHja7N1RT4MwGAXQMuv8/z9XGVRI2uSzlmU+DE08J7kBBnzPvRtkUyklTZuU0p7LlpeavOW15rrlrW6v4Vy79lLv3wMAADxHqVm3LDXzltuWj5r3up3DuXbtfl/Z5cHgtS7ol8Hifq2DclcCLvW8IgAAAM8tAm1dHsvArWauJeAWCkC7tsRBuRs4KgOxeSyhAMQSMCkCAABwWhEogzIQC8GoBMSkfDC8De6Pl0EB8GsAAACcXwbWO4VgVAK+yHeGpq41tOGxBMQAAADnlYF+rT5KOSoDuRs2hf32eFDpjvsXg2MJUAgAAOC5BaDfL4MFf/840Lf78wPDp7A9eolYCQAAgPPLQOoW+2VQAMpoyNGjQdNBc/DtPwAA/N1i8Mj+sAike63h4HOFAAAAfrcA/Pj8tP+hGAAA8L98CjAAlCWZhVgMBgwAAAAASUVORK5CYII=" alt=""></img>
<script>
var possible = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_)(*&^%$#@!';
var stringLength = Math.floor(Math.random() * 16) + 8 ;

var randomString = Array.apply(null, new Array(stringLength)).map(function () {
    return possible[Math.floor(Math.random() * possible.length)];
}).join('');

document.getElementById('secret').value = randomString;
</script>
</body></html>
`

const SecretURL string =`<!DOCTYPE html>
<html>
<head>
<style>
body {
    background-color: #336699
}
    html, body {
    height: 100%;
  }
  #tableContainer-1 {
    height: 100%;
    width: 100%;
    display: table;
  }
  #tableContainer-2 {
    vertical-align: middle;
    display: table-cell;
    height: 100%;
  }
  #myTable {
    margin: 0 auto;
  }
  #form_container
{
	background:#fff;
	margin:0 auto;
	text-align:left;
	width:1240px;
}
 h1
{
	background-color:#6699CC;
	margin:10;
	min-height:0;
	padding:0;
	text-decoration:none;
	text-indent:-8000px;

}

</style>
</head>
<body>

	<div id="form_container">

		<h1><a>Create a secret</a></h1>
					<div class="form_description">
			<h2>Secret URL</h2>
			<p></p>
		</div>
			<ul>
								<li id="li_1">
		<label class="description" for="secret">URL: </label>
                       <a href="{{.URL}}">{{.URL}}</a> 

		<div>
		</div>
					<li id="li_1">
		<label class="description" for="secret">This secret expires on: {{.EXPDATE}} </label>
		<div>
          		</div>
					<li id="li_1">
		<label class="description" for="secret">Maxium number of views: {{.VIEWS}}</label>

		<div>
		<div>
</body>
</html>`


const SecretHTML string =`<!DOCTYPE html>
<html>
<head>
<style>
body {
    background-color: #336699
}
    html, body {
    height: 100%;
  }
  #tableContainer-1 {
    height: 100%;
    width: 100%;
    display: table;
  }
  #tableContainer-2 {
    vertical-align: middle;
    display: table-cell;
    height: 100%;
  }
  #myTable {
    margin: 0 auto;
  }
  #form_container
{
	background:#fff;
	
	margin:0 auto;
	text-align:left;
	width:740px;
}
 h1
{
	background-color:#6699CC;
	margin:10;
	min-height:0;
	padding:0;
	text-decoration:none;
	text-indent:-8000px;
	
}

</style>
</head>
<body>

	<div id="form_container">
	
		<h1><a>Create a secret</a></h1>
					<div class="form_description">
			<h2>Your password is "<span style="background-color: #99ccff;">{{.PASS}}</span>"</h2>
			<p></p>
		</div>						
			<ul>
			
		<label>Please keep your password in a safe place</label>
		<div>
          
         					
		
 
		<div>
		<div>
</body>
</html>`
