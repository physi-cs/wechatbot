package glm

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const BASEURL = "http://direct.virtaicloud.com:22207/"

// ChatGLMResponseBody 请求体
type ChatGPTResponseBody struct {
	Response string     `json:"response"`
	History  [][]string `json:"history"`
	Status   int        `json:"status"`
	Time     string     `json:"time"`
}

// ChatGLMRequestBody 响应体
type ChatGLMRequestBody struct {
	// Model            string  `json:"model"`
	Prompt string `json:"prompt"`
	// MaxTokens        int     `json:"max_tokens"`
	// Temperature      float32 `json:"temperature"`
	// TopP             int     `json:"top_p"`
	// FrequencyPenalty int     `json:"frequency_penalty"`
	// PresencePenalty  int     `json:"presence_penalty"`
	History [][]string `json:"history"`
}

// Completions gtp文本模型回复
// curl https://api.openai.com/v1/completions
// -H "Content-Type: application/json"
// -H "Authorization: Bearer your chatGPT key"
// -d '{"model": "text-davinci-003", "prompt": "give me good song", "temperature": 0, "max_tokens": 7}'
func Completions_with_history(msg string, history_stack *History_stack) (string, error) {

	// 校验是否已超出轮数
	err := history_stack.check_rounds()
	if err != nil {
		return "", err
	}

	requestBody := ChatGLMRequestBody{
		Prompt:  msg,
		History: *history_stack.History,
	}
	requestData, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}
	log.Printf("request glm json string : %v", string(requestData))
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

	*history_stack.History = gptResponseBody.History

	if len(gptResponseBody.History) > 0 {
		for _, v := range gptResponseBody.History {
			log.Printf("gpt response history: 问%s ,答%s \n", v[0], v[1])
			break
		}
	}

	reply := gptResponseBody.Response
	log.Printf("gpt response text: %s \n", reply)
	return reply, nil
}

func Completions(sender string, msg string) (string, error) {
	// 读取存储的历史记录
	// TODO 根据wx_id获取历史对话

	history,err := GetHistoryStack(sender)
	if err != nil {
		return "", err
	}
	reply, err := Completions_with_history(msg, history)

	//history := HISTORY_STACK
	//history_stack := New_History_stack(sender, history, Max_boxes)
	//reply, err := Completions_with_history(msg, history_stack)
	return reply, err
}
