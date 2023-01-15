package main

import (
	"encoding/json"
	"log"
)

type Response struct {
	Header struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"header"`

	Data []struct {
		Type       string `json:"type"`
		ID         int    `json:"id"`
		Attributes struct {
			Email      string `json:"email"`
			ArticleIDS []int  `json:"article_ids"`
		} `json:"attributes"`
	} `json:"data"`
}

func ReadResponse(rawResp string) (Response, error) {
	var resp Response
	err := json.Unmarshal([]byte(rawResp), &resp)
	if err != nil {
		return Response{}, err
	}

	return resp, nil
}

func main() {
	raw := `{
	"header": {
		"code": 0,
		"message": ""
	},
	"data": [{
		"type": "user",
		"id": 100,
	"attributes": {
		"email": "bob@yandex.ru",
		"article_ids": [10, 11, 12]
	}
	}]
}`
	resp, err := ReadResponse(raw)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)
}
