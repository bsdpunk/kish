package twil

import (
	"../commands"
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var TwilSubs = commands.Commands{
	{
		Name:      "send",
		ShortName: "send",
		Usage:     "Send Text Message",
		Action:    SendText,
		Category:  "twil",
	},
}

var accountSid = string(os.Getenv("tsid"))
var authToken = string(os.Getenv("ttoken"))
var urlStr = string(os.Getenv("turl"))
var number = string(os.Getenv("tnumber"))

func SendText() {
	// Set account keys & information

	reader := bufio.NewReader(os.Stdin)
	//fmt.Println(urlStr)
	fmt.Print("To$ ")
	to, _ := reader.ReadString('\n')
	to = strings.Replace(to, "\n", "", -1)

	fmt.Print("Message Text$ ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", number)
	msgData.Set("Body", text)
	msgDataReader := *strings.NewReader(msgData.Encode())
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
	return
}
