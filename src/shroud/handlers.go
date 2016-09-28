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
)

func Index(w http.ResponseWriter, r *http.Request) {
	placeholder := `Placeholder{}`
	t := template.New("test")
	t, err := t.Parse(getSecretTmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, placeholder)
	if err != nil {
		log.Fatal(err)
	}
}

func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, VERSION)
}

func GetSecretWeb(w http.ResponseWriter, r *http.Request) {
	key := []byte(keyStr)
	token := r.FormValue("token")
	decodedsecret, deldate, delviews := GetSecret(token, key)
	type secretdict struct {
		PASS string
		DELDATE string
		DELVIEWS string
	}
	params := &secretdict{PASS: decodedsecret, DELDATE: deldate, DELVIEWS: delviews}
	t := template.New("Secret HTML")
	t, err := t.Parse(SecretHTML)
	if err != nil {
		fmt.Println("error parsing SecretURL")
		fmt.Println(err)
	}
	err = t.Execute(w, params)
	if err != nil {
		fmt.Println("error with execute")
	}
}

func PutSecretWeb(w http.ResponseWriter, r *http.Request) {
	key := []byte(keyStr)
	expireDate := r.FormValue("date")
	numViews := r.FormValue("views")
	secret := r.FormValue("secret")
	token, err := makeRandomString(40)
	if err != nil {
		log.Fatal(err)
	}
	ciphertext, err := encryptstring(key, secret)
	if err != nil {
		log.Fatal(err)
	}
//	fmt.Printf(ciphertext)

	insertComplete := PutSecret(ciphertext, token, expireDate, numViews)
	if insertComplete == false {
		log.Println("bad insert using ciphertext" + "," + expireDate + "," + numViews)
	}
	URLString := "https://" + r.Host + "/getsecret?token=" + token
	type urldict struct {
		URL     string
		EXPDATE string
		VIEWS   string
	}
	params := &urldict{URL: URLString, EXPDATE: expireDate, VIEWS: numViews}
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
	values := "'" + encryptedSecret + "','" + token + "','" + expireDate + "','" + numViews + "'"
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
	selectstatement := "SELECT secret, date, views from secrets where token ='" + token + "';"
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
