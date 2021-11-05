package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/codec"
	"github.com/zeromicro/zero-examples/http/signature/internal"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)



var (
	crypt = flag.Bool("crypt", false, "encrypt body or not")
)

func hs256(key []byte, body string) string {
	h := hmac.New(sha256.New, key)
	io.WriteString(h, body)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func main() {
	flag.Parse()

	var err error
	body := "hello world!"
	if *crypt {
		bodyBytes, err := codec.EcbEncrypt(internal.Key, []byte(body))
		if err != nil {
			log.Fatal(err)
		}
		body = base64.StdEncoding.EncodeToString(bodyBytes)
	}

	r, err := http.NewRequest(http.MethodPost, "http://localhost:3333/a/b?c=first&d=second", strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	timestamp := time.Now().Unix()
	sha := sha256.New()
	sha.Write([]byte(body))
	bodySign := fmt.Sprintf("%x", sha.Sum(nil))
	contentOfSign := strings.Join([]string{
		strconv.FormatInt(timestamp, 10),
		http.MethodPost,
		r.URL.Path,
		r.URL.RawQuery,
		bodySign,
	}, "\n")
	sign := hs256(internal.Key, contentOfSign)
	var mode string
	if *crypt {
		mode = "1"
	} else {
		mode = "0"
	}
	content := strings.Join([]string{
		"version=v1",
		"type=" + mode,
		fmt.Sprintf("key=%s", base64.StdEncoding.EncodeToString(internal.Key)),
		"time=" + strconv.FormatInt(timestamp, 10),
	}, "; ")

	encrypter, err := codec.NewRsaEncrypter(internal.PubKey)
	if err != nil {
		log.Fatal(err)
	}

	output, err := encrypter.Encrypt([]byte(content))
	if err != nil {
		log.Fatal(err)
	}

	encryptedContent := base64.StdEncoding.EncodeToString(output)
	r.Header.Set("X-Content-Security", strings.Join([]string{
		fmt.Sprintf("key=%s", internal.Fingerprint),
		"secret=" + encryptedContent,
		"signature=" + sign,
	}, "; "))
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	io.Copy(os.Stdout, resp.Body)
}
