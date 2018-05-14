package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	timeproUrl string
	loginID    string
	password   string
)

const (
	ExitOk = iota
	ExitArgMissing
	ExitWrongArg
	ExitPostError
	ExitNoCredential
)

func init() {
	timeproUrl = os.Getenv("TIMEPRO_URL")
	loginID = os.Getenv("TIMEPRO_ID")
	password = os.Getenv("TIMEPRO_PASSWORD")

	if loginID == "" || password == "" {
		fmt.Println("missing login ID or password or both.")
		os.Exit(ExitNoCredential)
	}
}

func main() {
	res := kintai(os.Args[0])
	os.Exit(res)
}

func kintai(arg string) int {
	if arg == "" {
		fmt.Println("second argument is missed. (in/out)")
		return ExitArgMissing
	}

	if arg == "in" {
		_, err := http.PostForm(timeproUrl,
			url.Values{"PAGESTATUS": {"PUNCH1"}, "PROCESS": {"PUNCH1"}, "LoginID": {loginID}, "PassWord": {password}})

		if err != nil {
			fmt.Println(err)
			return ExitPostError
		}

		fmt.Printf("[%v] 出勤しました！", time.Now())
	} else if arg == "out" {
		_, err := http.PostForm(timeproUrl,
			url.Values{"PAGESTATUS": {"PUNCH1"}, "PROCESS": {"PUNCH1"}, "LoginID": {loginID}, "PassWord": {password}})

		if err != nil {
			fmt.Println(err)
			return ExitPostError
		}

		fmt.Printf("[%v] 退勤しました。お疲れ様でした！", time.Now())
	} else {
		fmt.Println("second argument is wrong. (in/out)")
		return ExitWrongArg
	}

	return ExitOk
}
