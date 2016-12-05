package main

const getSecretTmpl string =`<!DOCTYPE html>
<html>
<head>
<link href="https://fonts.googleapis.com/css?family=Gloria+Hallelujah" rel="stylesheet"> 
<style>
.basicbox {
    margin: 0 auto;
    padding: 10px 10px;
    width: 700px;
    height: 280px;
    text-align: center;
    font-family: 'Gloria Hallelujah', Arial;
    font-size: 1.4em;
    color: white;
    border:1px #E8Ds47 solid;
    -moz-box-shadow:0px 0px 6px 1px #070707;
    -webkit-box-shadow:0px 0px 6px 1px #070707;
    box-shadow:0px 0px 6px 1px #070707;
background:#6594b6;
}
.pw {
font-family:  Arial;
}
.red {
color: red;
}
.basicbox ul {
    margin: 0px;
text-align: left;
}
.basicbox p {
    text-align: center;
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
<body bgcolor=#5388ae>
</br>
</br>
</br>
</br>
<div class="basicbox">
    <p>
    <strong>Create a new password to share:</strong>
    </p>

            <form id="form_1158694" class="appnitro" method="post" action="{{.PATH}}">
                <div class="form_description">
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
		</ul>
                <input id="saveForm" class="button_text" type="submit" name="submit" value="Submit" />
            </form>
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







const SecretURL string = `<!DOCTYPE html>
<html>
    <head>
<link href="https://fonts.googleapis.com/css?family=Gloria+Hallelujah" rel="stylesheet">
        <style>
body{
                background-color:#5388ae;
}

.basicbox {
    /* General */
    margin: 0 auto;
    padding: 10px 10px;
    width: 900px;
    height: 280px;
    text-align: center;
    /* Font */
    font-family: 'Gloria Hallelujah', Arial;
    font-size: 1.4em;
    color: white;
    /* Border */
    border:1px #E8Ds47 solid;
    /* Shadow */
    -moz-box-shadow:0px 0px 6px 1px #070707;
    -webkit-box-shadow:0px 0px 6px 1px #070707;
    box-shadow:0px 0px 6px 1px #070707;
    background:#6594b6;
}
 
.pw {
font-family:  Arial;
font-size: 0.6em;
}
.basicbox ul {
    margin: 0px;
text-align: left;
}
 
.basicbox p {
    text-align: left;

}

    </style>
    </head>
    <body>
	<div class="basicbox">
		<h2>Share this secret info.</h2>
		<p>URL:<a href="{{.URL}}" class= "pw">{{.URL}}</a></br>This secret expires on: {{.EXPDATE}}</br>Maxium number of views: {{.VIEWS}}</p>
        </div>
    </body>
</html>`



const SecretURLOld string =`<!DOCTYPE html>
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
</html>`


const NotFoundHTML string =`<!DOCTYPE html>
<html>
<head>
<link href="https://fonts.googleapis.com/css?family=Gloria+Hallelujah" rel="stylesheet"> 
<style>
.backgroundBlue {
background:#5388ae;
}
.paper {
    margin: 0 auto;
    padding: 1px 2px;
    width: 500px;
    height: 300px;
    color: #575757;
    font-family: 'Gloria Hallelujah', Arial;
    font-size: 2em;
    border:4px #E8Ds47 solid;
    -moz-box-shadow:0px 0px 30px 5px #070707;
    -webkit-box-shadow:0px 0px 30px 5px #070707;
    box-shadow:0px 0px 30px 5px #070707;
    -webkit-transform:rotate(-4deg);
    -o-transform:rotate(-4deg);
    -moz-transform:rotate(-4deg);
    background-image: url('data:image/jpeg;base64,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');
    background-size: 500px 300px;
    padding-left: 90px ;
}
.pw {
font-family:  Arial;
}
.red {
color: red;
}
.paper p{
font-size: .6em;
}
</style>
</head>
<body bgcolor=#5388ae>
</br>
</br>
</br>
</br>
</br>
<div class="paper">
    <p><strong>Token Not Found</strong></br></p>
    <p>Although this link no longer works, the password should still work.</p>
    <p> Did you save it in a safe place?</p>
</br>
</br>
</br>
</div>
</body>
</html>
`

const SecretHTML string =`<!DOCTYPE html>
<html>
<head>
<link href="https://fonts.googleapis.com/css?family=Gloria+Hallelujah" rel="stylesheet"> 
<style>

.paper {
    /* General */
    margin: 0 auto;
    padding: 8px 24px;
    width: 500px;
    height: 300px;
    color: #575757;
    /* Font */
    font-family: 'Gloria Hallelujah', Arial;
    font-size: 2em;
    /* Border */
    border:1px #E8Ds47 solid;
    /* Shadow */
    -moz-box-shadow:0px 0px 6px 1px #070707;
    -webkit-box-shadow:0px 0px 6px 1px #070707;
    box-shadow:0px 0px 6px 1px #070707;
 -webkit-transform:rotate(4deg);
  -o-transform:rotate(4deg);
  -moz-transform:rotate(4deg);
    background-image: url('data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBw8PDw8PDxAPDw8PDw8PDw8PEg8PDw8PFRUWFhUVFRUYHyggGBolGxUVITEhJSkrLi4uFx8zODMtNygtLisBCgoKDg0NFQ8PFysdFRktNy0tLS0rNystLSsrKy0tLS0tLSsrLS0tLS0rLSs3LSstKy0rLS0tLS0tLTc3LS0rLf/AABEIAKgBLAMBEQACEQEDEQH/xAAYAAEBAQEBAAAAAAAAAAAAAAAAAwEEB//EADIQAAIABAUDBAAGAgIDAAAAAAACAQNSkhMUUdHSERKRU2KToiIyQUKBsXGjBHIhQ/D/xAAXAQEBAQEAAAAAAAAAAAAAAAAAAQMC/8QAKBEAAgEDAwQDAQEBAQEAAAAAAAFRAhKRERPwIVJh4YGh0SIy8TED/9oADAMBAAIRAxEAPwD24AAAAAAAAEJ351/j+yGlP+WXKZgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEJ351/8Av1BpT/llwZgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHPlIVzfkY7v8ACwZbKl5GUhXN+Rxf4WBsqXkZSFc35GF/hYGypeRlIVzfkcXuFgmypeRlIVzfkcX+FguypeRlIVzfkcXuFgbKl5ITf+JDvX8c39P/AGNqS/wsGlH/AMla+ryXykK5vyOW9wsGeypeRlIVzfkcX+FgbKl5GUhXN+Rxe4WBsqXkZSFc35HF/hYGypeRlIVzfkcXuFgbKl5GUhXN+Rhe4WBsqXkZSFc35HF7hYGypeRlIVzfkcXuFgbKl5GUhXN+Rhe4WBsqXkZSFc35HF/hYGypeRlIVzfkcXuFgbKl5GUhXN+Rxf4WBsqXkZSFc35HF/hYGypeRlIVzfkcXuFgbKl5GUhXN+Rxf4WBsqXkZSFc35HF/hYJsqXkZSFc35HF/hYLsqXkZSFc35HF/hYGypeRlIVzfkcXuFgbKl5GUhXN+Rxe4WBsqXkZSFc35GF7hYGypeRlIVzfkYXuFgbKl5GUhXN+Rxf4WCbKl5GUhXN+Rxf4WC7Kl5GUhXN+Rhe4WCbKl5GUhXN+Rxe4WC7Kl5GUhXN+Rxe4WBsqXkZSFc35HF7hYGypeRlIVzfkcX+FgbKl5GUhXN+Rhf4WBsqXkZSFc35GF/hYGypeSmNDRrH2ObTq9ecMYy+6x9haL1xDHX3WPsLWL1xDHh7rH2FrF64hjw91j7C1i9cQxoaPY+xbWL15wyE2dDvX/wAN+n7H1/wTQ1prVr/9wXxoaPY+xbWZXrzhjGho9j7EtF684YxoaPY+xbWL15wMaGj2PsS1i9ecDHho9j7FtF684GNDR7H2JaxevOGMaGj2PsLWL15wMaGj2PsW1i9ecMY0NHsfYlovXnDGNDR7H2FrF684YxoaPY+wtYvXnDGNDR7H2FrF684YxoaPY+wtYvXnAxoaPY+w0YvXnAxoaPY+xbWL15wMaGj2PsS0XrzgY0NHsfYWsXriGNDR7H2FovXnAxoaPY+wtYvXnAxoaPY+wtYvXEMaGj2PsW1i9ecDGho9j7EtYvXnDGPDR7H2FrF684YxoaPY+w0YvXnDGNDR7H2FrF684YxoaPY+wtF684GPDR7H2FovXnAxoaPY+xbRevOGMaGj2PsS1i9ecDGho9j7FtYvXEMeGj2PsS0XriKEOwAAAAAAAQnfnX+P7Ia0/wCWXKZAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE+j1JZHkXpBxpXKx7Ha9SWR5F6RzA0rlY9jo9SWR5DpA0rlY9jo9SWR5E6QNK5WPY6PUlkeRekcwNK5WPY7XqSxuROkDSuVj2QnK/ev4l/T9sdf+w6GtKqtfVY9l+16ksjyGqgy0rlY9jtepLI8h0gaVysex0epLG5DpA0rlY9jtepLI8h0gaVysex2vUljch0gaVysex0epLI8i9I5gaVysex2vUlkeROkDSuVj2Oj1JZHkOg0rlY9jtepLI8h0gaVSsex0epLI8hqoGlcrHsdHqSxuQ6QNK5WPY7XqSyPIdIGlUrHsdHqSyPIdIGlcrHsdHqSyPIvQaVysex2vUlkeROkDSuVj2O16ksjyHQaVysex0epLI8h0gaVysex2vUlkeQ6QNK5WPY7XqSyPIdIGlcrHsdHqSyPIfzA0rlY9jtepLI8h0gaVysex2vUtkeRf5jmBpXKx7Ha9SWR5E6DSuVj2O16ksjyHSBpVKx7Ha9SWR5DoNK5WPY6PUlkeRekcwNK5WPY7XqSxuQ6QNK5WPY7XqSyPImqgaVysex2vUljci9I5gaVysex2vUlkeQ1UDSuVj2OkypLG5E1pjmBpXKx7KkOzAAAADQDACE786/x/YNKf8suDM0AwA0AwAdAAAaAYAAAAaAYAAAAOgAAAAAABoBgBoAAM6AAAAG9ADOgBoBgBLLy6EtU6ucnG3RCNy8uhLVJdVI26IQy8uhLVLdVI26IQy8uhLVFzkbdEIZeXQlqkucjbohDLy6EtUXVSNuiEc83/AI6d6/gT9P2rqLnJrR/86LX0R0ZeXQlqlucmW3RCGXl0JaouqkbdEIZeXQlqkuqkbdELAy8uhLVFzkbdHahl5dCWqLnI26O1DLy6EtUXVSNuiEMvLoS1Rc5G3R2oZeXQlqluqkbdEIzLy6EtUl1UjbohYGXl0Jaoucjbo7Vg3Ly6EtUXORt0QsDLy6EtUXORt0QjMvLoS1Rc5G3RCNy8uhLVF1UjbohDLy6EtUXVSNuiFgzLy6EtUt1UjbohDLy6EtUl1UjbohG5eXQlqlucjbohDLy6EtUlzkbdEIZeXQlqi5yNuiEMvLoS1Rc5G3RCGXl0JaoucjbohYMy8uhLVF1UjbohG5eXQlqi5yNuiFgZeXQlqlucjbohYGXl0JaoucjbohDLy6EtUXORt0Qhl5dCWqS6qRt0Qhl5dCWqLnI26IQy0uhLVF1UjbohGZaXQlqi6qRt0QsG4kfTfzL5DRSLn2v6/RiR9N/MvkXRSLn2v6/RiN6b+ZfImikXPtf1+jEj6b+ZfIuikXPtf1+jFb038y+Q0Ui59r+v0YkfTfzL5E0Ui59r+v0hNmR71/A36frL1/7DRSa01O19H9fpfEj6b+ZfIuikyufa/r9GK3pv5l8iaKRc+1/X6MRvTfzL5F0Ui59r+v0YkfTfzL5E0Ui59r+v0Yjem/mXyGikXPtf1+jEj6b+ZfIaKRc+1/X6MRvTfzL5F0Ui59r+v0YkfTfzL5DRSLn2v6/RiN6b+ZfIaKRc+1/X6MSPpv5l8iaKRc+1/X6MRvTfzL5F0Ui59r+v0Yjem/mXyGikXPtf1+jFb038y+RNFIufa/r9GI1D+ZfIaKRc+1/X6MSPpv5l8hopFz7X9foxG9N/MvkXRSLn2v6/Rit6b+ZfImikXPtf1+jEb038y+RdFIufa/r9GI3pv5l8iaKRc+1/X6MVvTfzL5DRSLn2v6/RiN6b+ZfIaKRc+1/X6MSPpv5l8hopFz7X9foxG9N/MvkXRSLn2v6/RiR9N/MvkTRSLn2v6/RiR9N/MvkXRSLn2v6/RiN6b+ZfImikXPtf1+jEb038y+Q0Ui59r+v0YkfTfzL5DRSLn2v6/RiR9N/MvkNPIufa/r9GLH038y+RdPIufa/r9GJH038y+RNPIufa/r9KEOwAAAAAAAQnfnX+P7IaU/5ZcpmAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAATwo1v9OJ1r4OLX3PnwMONb/6+I18C19z58DDjW/02Jr4Fr7nz4GFGt/CcRr4Fr7nz4GHGt/pxLr4Fr7nz4GFGt/pxJr4Fr7nz4ITZce9fxv+lGv+Br4NaaXa+r58F8KNb/TiNfBla+58+BhRrf6cRr4Fr7nz4GFGt/8AXxLr4Fr7nz4GFGt/pxJr4Fr7nz4GFGt/9fEa+Ba+58+BhRrf6cRr4Fr7nz4GFGt/pxGqgWvufPgYca3+nEaqBa+58+Bhxrf6cRr4Fr7nz4GHGt/pxGqgWvufPgYca3/18RqoFrl/X4MONb/TYaqBa+58+BhRrf6cRqoFr7nz4GFGt/psNVAtfc+fAw41v9OI18C19z58DCjW/wBOJdVAtfc+fAw41v8ATiTXwLX3PnwMONb/AE4jVQLX3PnwMKNb/TiNVAtfc+fAwo1v9OI1UC19z58DDjW/02GvgWvufPgYUa3+mw18C19z58DCjW/02GqgWvufPgYUa38JxGvgWvufPgYUa3+mw18C19z58DCjW/04jVQLXL58DCjW/wBNi6qBa+58+Bhxrf6bE1UC19z58DDjW/02GvgWvufPgYca3+nEaqBa5fPgYcfUf6cRr4Fr7nz4GYStfMBa4G5TIzCVr5gLXA3KZGYStfMBa4G5TIzCVr5gW1wNymRmErXzAlrgblMjMJWt0Ba4F9MkJs9O9fxL+n6w1Frg1orptfUvmErW6AtcGW5TIzCVr5gLXA3KZGYStfMBa4G5TIzCVrdAWuBuUyMwla+YC1wNymRmErXzAWuBuUyMdK18wLa4G5TIzCVr5gS1wNymRmErXzAWuBuUyMwla3QFrgblMjMJWvmAtcDcpkZhK18wFrgblMjMJWvmAtcDcpkZhK1ugLXA3KZGYStfMBa4G5TIzCVr5gLXAvpkZhK1ugLXA3KZGYStfMBa4G5TIzCVr5gLXA3KZGYStfMBa4G5TIzCVr5gLXA3KZGYStfMBa4G5TIzCVr5gLXAvpkZhK18wLa4G5TIzCVr5gS1wNymRmErW6AtcDcpkZhK18wFrgblMjMJWvmAtcDcpkZhK18wFrgblMjMJWvmAtcDcpkZhK18wFrgblMlCHYAAAAAAAITvzr/AB/ZDSn/ACy5TMAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA0AwAAAAAAAAAAAAAn1elLo8Trocf1Cz6HV6UujxJ0GtULPodXpS6PEdBrVCz6HV6UujxHQa1Qs+h1elLm4joNa4WfQ6vSl0eI6c/6Na4WfRCdF+9fwr+n7o6/wCB0NabrX0WfRfuelLm4joZa1ws+h1elLo8R05/0a1Qs+h1elLm4joNa4WfQ6vSl0eI6c/6NaoWfQ6vSlzcR0GtULPodXpS6PEdB/ULPodXpS5uJdEP6hZ9Dq9KXR4k0Q1rhZ9Dq9KXNxHQa1Qs+h1elLo8R0GtULPodXpS6PEuiH9ws+h1elLo8R0GtcLPodXpS6PEmi5/0a1ws+h1elLm4joNaoWfQ6vSl0eI6D+oWfQ6vSl0eI6D+4WfQ6vSl0eI6DWqFn0Or0pdHiOnP+jWuFn0Or0pdHiOg1qhZ9Dq9KXNxHQa1Qs+h1elLm4jpz/o1qhZ9Dq9KXNxHQf1Cz6HV9EubiOg1qhZ9Dq9KXR4joNa4WfQ6vSl0eI6c/6NaoWfQ6vSlzcR0H9Qs+h1elLm4joNaoWfQ6vSl0eI6DWuFn0Or0pdHiOg1rhZ9Dq9KXNxHQa1ws+h3TKUujxHQf1Cz6KEOwAAAAAAAQnfnX+P7Ia0/wCWXKZAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE8Bfde+5bmcWLzljAh7r33FzFi85YwF9177luYsXGxgL7r33JcxYuMYK+699xcxYvORgL7r33FzFi4yE6TDvX836fvfX/IuZrTQrX+l8Bfde+4uZlYuNjAh7r33LcxYvOWMBffe+5LmLFxsYC+699y3PiFi85YwIate+5LmLFxsYC+699xcxYvOWMBfde+5bmLFxjBX33vuS5ixcbGAvuvfctzFi85YwIe6+ZuLnxCxcbGAvuvfcXMWLzljAh7r33FzFi42MBfde+4uYsXGMBdWvfcXMba85YwV9177kuY215yxgQ9177i5ixecsYC+699y3Mba85GCvuvfclzFi85YwV9177i5ixcbGAvuvfcXMba4xgQ9177i5ixecsYC+699xcxYvORgL7r33FwsXnLGBD3XvuLmNtecsYK+699xcxYuMYC+699xcxYvOWMCHuvfctzFi85YwF9177kuYsXGxgL7r33LcxYvOWMFfde+5LmNtcbGAvuvfcXMWLzljGho9kzYWi9ecMY0NHsmbC1i9ecMY0NHsmbC18YvXnDGNDR7JmwtfGL15wxjQ0eyZsLWL15wxjw0eyZsLXxi9ecMhNnQ71/8AD/p+x9f8C01prVr/APcMvjQ0eyZsLXxmV684YxoaPZM2Fr4xevOGMaGj2TNhbzUXrzhjHho9kzYWsXrzhjGho9j7C18YvXnDGNDR7JmwtfGL15wxjQ0eyZsLXxi9ecMY0NHsmbC3movXnDGNDR7JmwtF684YxoaPZM2Fr4xevOGMaGj2TNha+MXrzhjGho9kzYWvjF684YxoaPZM2LaL15wxjQ0eyZsS18YvXnDGNDR7JmxbRevOGMaGj2TNiWvjF684YxoaPZM2Fr4xevOGMaGj2TNhaxevOGMaGj2TNi2sXrzhjGho9kzYlovXnDGNDR7JmxbRevOGMaGj2TNiWvjF684YxoaPZM2FrF684YxoaPY+wtfGL15wxjw0eyZsLXxi9ecMY0NHsfYWvjF684YxoaPZM2Fr4xevOGMaGj2TNhbzUXrzhjGho9kzYWvjF684YxoaPZM2Fr4xevOGMaGj2TNha+MXrzgoQ7AAAAAAABCd+df4/shrT/llymQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABPtepLI8jrVR9+jjSuVj2O2ZUljchqo+/Q0qlY9jtepLI8ia0xzA0rlY9jtepLI8i9IGlcrHsdr1JZHkNaY+/Q0rlY9jtmVJZHkP5jmBpXKx7ITlfvX8Sfp+yOv8A2J0g1oVVr6/Xsv2zKksjyLqo5gy0rlY9jtepLI8h/McwNK5WPY7XqSyPIaqOYGlcrHsdr1JZHkP5jmBpXKx7Ha9SWR5D+Y+/Q0rlY9jtepLG5DVRzA0qlY9jtmVJZHkTpHMDSuVj2O16ksjyHSBpXKx7Ha9SWNyL/McwNK5WPY7ZlSWR5Dpz/g0rlY9jtepLI8hrTHMDSuVj2O16ksjyGtMcwNK5WPY7XqSyPIdI5gaVysex2vUljchqo5gaVysex2zKksjyJ0gaVysex0epLG5F6RzA0rlY9jtmVJZHkP5jmBpXKx7Ha9SWNyGtMcwNK5WPY7ZlSWNyH8xzA0rlY9jtepLI8h/McwNK5WPY6PUlkeQ1UcwNK5WPY6TKksjyHSOYGlcrHsdsypLI8h0jmBpXKx7HbMqSyPIaqOYGlcrHsdr1JZHkOkffoaVysex2vUlkeQ6RzA0rlY9jtepLI8h0jmBpXKx7Ha9SWNyH8xzA0rlY9jtmVJY3IdI5gaVysex2vUlkeQ/mOYGlcrHsdsypLG5DVRzA0rlY9lDk7AAAAAAABCd+df4/shpT/llymYAAGoAAIAUAAAAAAAAAAAAAAAAAAAAAAADUAagAAAAAgBdQAAAATy8uhLVLc5ONumFgZeXQlqi5yNumFgZdKEtUXORt0QsDLy6EtUXVSNumFgZeXQlqi5yNumEMvLoS1RdVJbKYRCd/x071/An6ftXUXOTSmim19EXy8uhLVFzky26YQy8uhLVFzkbdMIZeXQlqi6qRt0whl5dCWqLnI26YWBl0oS1Rc5G3TCGXl0JaouqkbdEIZeXQlqlucjbphGZdKEtUXORt0wsG5eXQlqi6qRt0Qhl5dCWqS5yNuiFgZeXQlqi5yLKYQy6UJaoucjbphDLy6EtUXORt0whl0oS1RdVI26YQy6UJapbqpG3RCwMvLoS1SXORt0wsDLy6EtUt1UiymEMvLoS1SXORt0whl0oS1Rc5Lt0Qhl5dCWqW5yTbphDLy6EtUl1UjbphDLpQlqi5yNumFgZeXQlqi5yNumEMvLoS1S3ORZTCGXl0JapLnI26YQy8uhLVFzkbdELAy6UJaoucjbohYGXShLVFzkbdEIZeXQlqi5yNumFgZdKEtUtzkbdELBmXShLVJc5G3TCwf//Z');
background-size: 530px 300px;
padding-left: 90px ;
}


.sticky {
    /* General */
    margin: 0 auto;
    padding: 10px 10px;
    width: 400px;
    height: 280px;
 text-align: center;
    /* Font */
    font-family: 'Gloria Hallelujah', Arial;
    font-size: 1.4em;
    color: white;
    /* Border */
    border:1px #E8Ds47 solid;
    /* Shadow */
    -moz-box-shadow:0px 0px 6px 1px #070707;
    -webkit-box-shadow:0px 0px 6px 1px #070707;
    box-shadow:0px 0px 6px 1px #070707;
/* -webkit-transform:rotate(-2deg);
  -o-transform:rotate(-2deg);
  -moz-transform:rotate(-2deg);*/
background:#6594b6;

  

}
 
.pw {
font-family:  Arial;

}

.red {
color: red;

}
.sticky ol {
    margin: 0px;
}
 
.sticky p {
    text-align: center;

}

.paper p{
font-size: .6em;

}

</style>
</head>

<body bgcolor=#5388ae>



</br>
</br>
</br>
</br>
<div class="sticky">
    <p>
    <strong>Your password is:</strong>
    </p>
<div class="pw">
        <p>{{.PASS}}</p>
</div>

</div>   
</br>
</br>
</br>
</br>
<div class="paper">
    <p>
    <strong>IMPORTANT!</strong></br>
    Save this password! This URL will expire on {{.DELDATE}} or {{.DELVIEWS}} more views.  (Whichever comes first.) 
    </p>





</br>
</br>
</br>

</div>
</body>
</html>`
