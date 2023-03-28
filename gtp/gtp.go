package gtp

// package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// const BASEURL = "https://api.openai.com/v1/"
const BASEURL = "http://region-41.seetacloud.com:20487/"

// ChatGPTResponseBody 请求体
type ChatGPTResponseBody struct {
	Response string     `json:"response"`
	History  [][]string `json:"history"`
	Status   int        `json:"status"`
	Time     string     `json:"time"`
}

type ChoiceItem struct {
}

// ChatGPTRequestBody 响应体
type ChatGLMRequestBody struct {
	// Model            string  `json:"model"`
	Prompt string `json:"prompt"`
	// MaxTokens        int     `json:"max_tokens"`
	// Temperature      float32 `json:"temperature"`
	// TopP             int     `json:"top_p"`
	// FrequencyPenalty int     `json:"frequency_penalty"`
	// PresencePenalty  int     `json:"presence_penalty"`
	History []string `json:"history"`
}

// Completions gtp文本模型回复
// curl https://api.openai.com/v1/completions
// -H "Content-Type: application/json"
// -H "Authorization: Bearer your chatGPT key"
// -d '{"model": "text-davinci-003", "prompt": "give me good song", "temperature": 0, "max_tokens": 7}'
func Completions(msg string) (string, error) {

	requestBody := ChatGLMRequestBody{
		Prompt:  msg,
		History: []string{},
	}
	requestData, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}
	log.Printf("request gtp json string : %v", string(requestData))
	req, err := http.NewRequest("POST", BASEURL, bytes.NewBuffer(requestData))
	if err != nil {
		return "", err
	}

	// apiKey := config.LoadConfig().ApiKey
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	gptResponseBody := &ChatGPTResponseBody{}
	log.Println(string(body))
	err = json.Unmarshal(body, gptResponseBody)
	if err != nil {
		return "", err
	}

	// if len(gptResponseBody.Choices) > 0 {
	// 	for _, v := range gptResponseBody.Choices {
	// 		reply = v["text"].(string)
	// 		break
	// 	}
	// }
	reply := gptResponseBody.Response
	log.Printf("gpt response text: %s \n", reply)
	return reply, nil
}

// func main() {
// 	reply, err := Completions("你好")
// 	if err != nil {
// 		fmt.Printf("gtp request error: %v \n", err)
// 	}
// 	if reply != "" {
// 		fmt.Printf(reply)
// 	}
// }
