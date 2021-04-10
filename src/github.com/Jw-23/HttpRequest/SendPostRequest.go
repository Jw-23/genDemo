package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	client := http.Client{Timeout: time.Second * 5}
	user := "wnm"
	message := "王尼玛"
	password := "sb2223"
	//postData := fmt.Sprintf(`user=%s&password=%s&message=%s`, user, password, message)
	var postData = fmt.Sprintf(`"user":"%s","password":"%s","message":"%s"`, user, password, message)
	req, err := http.NewRequest("POST", "http://localhost:8000/login", strings.NewReader(postData))
	if err != nil {
		fmt.Println("create httpRequest failed,err:", err)
		return
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("Content-Type", "application/json,charset=utf-8")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("send post request failed, error:", err)
		return
	}
	fmt.Println("发送成功")
	buffs, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Println("send post request failed, error:", err)
		return
	}
	fmt.Println(string(buffs))
}
