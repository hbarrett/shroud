package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	b64 "encoding/base64"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {

        type secretdict struct {
                BACKGROUND string
                FOREGROUND string
		FORMBACKGROUND string
	 	PATH string
        }
        path := altpath + "/putsecret"

        params := &secretdict{BACKGROUND: background, FOREGROUND: foreground, FORMBACKGROUND: formbackground, PATH: path}
	t := template.New("test")
	t, err := t.Parse(getSecretTmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, params)
	if err != nil {
		log.Fatal(err)
	}
}

func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, VERSION)
}

func GetSecretWeb(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
        key := []byte(token)
	decodedsecret, deldate, delviews := GetSecret(token, key)
	if decodedsecret==""{
        nf := template.New("Not Found")
        nf, err := nf.Parse(NotFoundHTML)
        if err != nil {
                fmt.Println("error parsing NotFoundURL")
                fmt.Println(err)
        }
        type nfdict struct {
                TOKEN string
                BACKGROUND string
                FOREGROUND string
                FORMBACKGROUND string
        }
	nfparams := &nfdict{TOKEN: token, BACKGROUND: background, FOREGROUND: foreground, FORMBACKGROUND: formbackground}
        err = nf.Execute(w, nfparams)
        if err != nil {
                fmt.Println("error with Not Found execute")
        }

	}else{
	type secretdict struct {
		PASS string
		DELDATE string
		DELVIEWS string
		BACKGROUND string
		FOREGROUND string
		FORMBACKGROUND string
	}
        delviewsint, err := strconv.Atoi(delviews)
	logErr(err)

	if delviewsint > 1{
	decViews(token[:8])
	}else if delviewsint==1{
	delViews(token[:8])
	}
	if delviewsint != 0 {
	delviewsint--
	}
	if delviewsint == 0 {
        params := &secretdict{PASS: decodedsecret}
        t := template.New("Last Secret HTML")
        t, err = t.Parse(LastSecretHTML)
        if err != nil {
                fmt.Println("error parsing SecretURL")
                fmt.Println(err)
        }
        err = t.Execute(w,params)
        if err != nil {
                fmt.Println("error with execute")
        }

	}else{

	delviewsStr := strconv.Itoa(delviewsint)
	params := &secretdict{PASS: decodedsecret, DELDATE: deldate, DELVIEWS: delviewsStr, BACKGROUND: background, FOREGROUND: foreground, FORMBACKGROUND: formbackground}
	t := template.New("Secret HTML")
	t, err = t.Parse(SecretHTML)
	if err != nil {
		fmt.Println("error parsing SecretURL")
		fmt.Println(err)
	}
	err = t.Execute(w, params)
	if err != nil {
		fmt.Println("error with execute")
	}
	}
	}
}

func RemoveSecretWeb(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	fmt.Println(token)
}

func PutSecretWeb(w http.ResponseWriter, r *http.Request) {
	expireDate := r.FormValue("date")
	numViews := r.FormValue("views")
	secret := r.FormValue("secret")
	token, err := makeRandomString(16)
        key := []byte(token)
	if err != nil {
		log.Fatal(err)
	}
	ciphertext, err := encryptstring(key, secret)
	if err != nil {
		log.Fatal(err)
	}

	insertComplete := PutSecret(ciphertext, token, expireDate, numViews)
	if insertComplete == false {
		log.Println("bad insert using ciphertext" + "," + expireDate + "," + numViews)
	}
	URLString := "https://" + r.Host + altpath +"/getsecret?token=" + token
	type urldict struct {
		URL     string
		EXPDATE string
		VIEWS   string
                BACKGROUND string
                FOREGROUND string
		FORMBACKGROUND string
	}
	params := &urldict{URL: URLString, EXPDATE: expireDate, VIEWS: numViews, BACKGROUND: background, FOREGROUND: foreground, FORMBACKGROUND: formbackground}
	t := template.New("Secret URL")
	t, err = t.Parse(SecretURL)

	if err != nil {
		fmt.Println("error parsing SecretURL")
		fmt.Println(err)
	}
	err = t.Execute(w, params)
	if err != nil {
		fmt.Println("error with execute")
	}
}

var iv = []byte{33, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func encodeBase64(b []byte) string {
	return b64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := b64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func makeRandomString(s int) (string, error) {
	b := make([]byte, s)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	randStr := fmt.Sprintf("%0x", b)
	return randStr, nil

}

func GenerateRandomStr(s int) (string, error) {
	b, err := GenRandomBytes(s)
	return b64.URLEncoding.EncodeToString(b), err
}

func GenRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func PutSecret(encryptedSecret string, token string, expireDate string, numViews string) bool {
	values := "'" + encryptedSecret + "','" + token[:8] + "','" + expireDate + "','" + numViews + "'"
	insert := "INSERT INTO secrets (secret, token, date, views) VALUES(" + values + ");"
	var sucess bool
	_, err := db.Query(insert)
	if err != nil {
		log.Fatal(err)
		sucess = false
	} else {
		sucess = true
	}
	return sucess
}

func GetSecret(token string, key []byte) (secret2, date, views string) {
	var secret string
	var count int
	countstatement := "SELECT COUNT(*) from secrets where token ='" + token[:8] + "';"
	rows, err := db.Query(countstatement)
 	logErr(err)
        for rows.Next() {
        err:= rows.Scan(&count)
        logErr(err)
	}
	if count > 0 {
	selectstatement := "SELECT secret, date, views from secrets where token ='" + token[:8] + "';"

	foundsecret, err := db.Query(selectstatement)

	if err != nil {
		log.Fatal(err)
	}

	for foundsecret.Next() {
		err = foundsecret.Scan(&secret, &date, &views)
		if err != nil {
			log.Fatal(err)
		}

	}
	secret2, err = decryptstring(key, secret)
	if err != nil {
		log.Fatal(err)
	}
	}else{
	secret2 = ""
	date = "0000-00-00"
	views = "0"
	}
	return
}

func encryptstring(key []byte, plaintext string) (string, error) {

	text := []byte(plaintext)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "1", err
	}
	b := b64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "2", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return fmt.Sprintf("%0x", ciphertext), nil
}

func decryptstring(key []byte, ciphertext string) (string, error) {
	var text []byte
	fmt.Sscanf(ciphertext, "%x", &text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "3", err
	}
	if len(text) < aes.BlockSize {
		return "4", errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := b64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return "5", err
	}
	return fmt.Sprintf("%s", data), nil

}


func decViews(token string) {

        stmt, err := db.Prepare("UPDATE secrets SET views = views -1 WHERE token=? and views > 0")
        logErr(err)
        result, err := stmt.Exec(token)
        logErr(err)
        affect, err := result.RowsAffected()
        logErr(err)
        log.Println(affect)

}


func delViews(token string) {

        stmt, err := db.Prepare("DELETE FROM secrets WHERE token=?")
        logErr(err)
        result, err := stmt.Exec(token)
        logErr(err)
        affect, err := result.RowsAffected()
        logErr(err)
        log.Println(affect)

}

