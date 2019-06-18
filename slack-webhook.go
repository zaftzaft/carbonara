package main

import (
	"net/http"
	"strings"
)

func SendWebhook(url string, text string) error {
	payload := strings.NewReader("{\n\t\"text\": \"" + text + "\"\n}")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
	//body, err := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))

}
