package main

const getSecretTmpl string =`<!DOCTYPE html>
<html>
    <head>
        <style>
       		form li .datepickerimg
            {
        	cursor:pointer !important;
        	float:left;
        	height:16px;
        	margin:.1em 5px 0 0;
        	padding:0;
        	width:16px;
            }

            .form_description{
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
            body{
                background-color:{{.BACKGROUND}};
            }
            html,
            body{
                height:100%;
            }
            #tableContainer-1{
                height:100%;
                width:100%;
                display:table;
            }
            #tableContainer-2{
                vertical-align:middle;
                display:table-cell;
                height:100%;
            }
            #myTable{
                margin:0 auto;
            }
            #form_container
            {
                background:{{.FORMBACKGROUND}};
                margin:0 auto;
                text-align:left;
                width:640px;
            }
            h1
            {
                background-color:{{.FOREGROUND}};
                margin:10;
                min-height:0;
                padding:0;
                text-decoration:none;
                text-indent:-8000px;
	   {
	   img.loading{
  		position: absolute;
  		left: 50%;
  		margin-left: -(half ot the image width)px
	   }
	</style>
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
    <body>
        <div id="form_container">

            <h1>
                <a>Create a new password to share.</a>
            </h1>
            <form id="form_1158694" class="appnitro" method="post" action="{{.PATH}}">
                <div class="form_description">
                    <h2>Create a new password to share</h2>
                    <p></p>
                </div>
                <ul> <div>
                        <label class="description" for="secret">Password</label>
                            <input id="secret" name="secret" class="element text medium" type="text"
                                maxlength="255" value="" />
                        </div>
			<div>
                        <label class="description" for="views">Number of views before secret is
                            removed:</label>
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
			<div>
                        <label class="description" for="element_2">Date when secret will be removed:</label>
                        <input id="datepicker" name="date" />
			</div>
                        <input id="saveForm" class="button_text" type="submit" name="submit" value="Submit" />
                </ul>
            </form>
            <div id="footer"> </div>
        </div>
<script>
var possible = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_)(*&^%$#@!';
var stringLength = Math.floor(Math.random() * 16) + 8 ;

var randomString = Array.apply(null, new Array(stringLength)).map(function () {
    return possible[Math.floor(Math.random() * possible.length)];
}).join('');

document.getElementById('secret').value = randomString;
</script>


    </body>
</html>`


const SecretURL string =`<!DOCTYPE html>
<html>
    <head>
        <style>
            body{
                background-color:{{.BACKGROUND}};
            }
            html,
            body{
                height:100%;
            }
            #tableContainer-1{
                height:100%;
                width:100%;
                display:table;
            }
            #tableContainer-2{
                vertical-align:middle;
                display:table-cell;
                height:100%;
            }
            #myTable{
                margin:0 auto;
            }
            #form_container
            {
                background:{{.FORMBACKGROUND}};
                margin:0 auto;
                text-align:left;
                width:1240px;
            }
            h1
            {
                background-color:{{.FOREGROUND}};
                margin:10;
                min-height:0;
                padding:0;
                text-decoration:none;
                text-indent:-8000px;
            }</style>
    </head>
    <body>
        <div id="form_container">
            <h1><a>Create a secret</a></h1>
            <div class="form_description">
                <h2>Secret URL</h2>
                <p></p>
            </div>
            <label class="description" for="secret">URL: </label>
            <a href="{{.URL}}">{{.URL}}</a>
            <p>
                <label class="description" for="secret">This secret expires on: {{.EXPDATE}}
                </label>
            </p>
            <p>
                <label class="description" for="secret">Maxium number of views: {{.VIEWS}}</label>
            </p>
            <br></br>
        </div>
    </body>
</html>
`


const SecretHTML string =`<!DOCTYPE html>
<html>
    <head>
        <style>
            body{
                background-color:{{.BACKGROUND}};
            }
            html,
            body{
                height:100%;
            }
            #tableContainer-1{
                height:100%;
                width:100%;
                display:table;
            }
            #tableContainer-2{
                vertical-align:middle;
                display:table-cell;
                height:100%;
            }
            #myTable{
                margin:0 auto;
            }
            #form_container
            {
                background:{{.FORMBACKGROUND}};
                margin:0 auto;
                text-align:left;
                width:740px;
            }
            h1
            {
                background-color:{{.FOREGROUND}};
                margin:10;
                min-height:0;
                padding:0;
                text-decoration:none;
                text-indent:-8000px;
            }</style>
    </head>
    <body>
        <div id="form_container">
            <h1><a>Create a secret</a></h1>
            <div class="form_description">
            <h2>Your password is "<span style="background-color: #99ccff;">{{.PASS}}</span>"</h2>
            <h4>This password will be removed after {{.DELVIEWS}} more views or midnight on
                {{.DELDATE}}, whichever comes first, so please save it in a safe place.</h4>
            <p></p>
        </div>
    </body>
</html>`
