package openAI

import (
  "bytes"
  "encoding/json"
  "errors"
  "github.com/Kawanishi45/cat_robot/pkg/constants"
  loging "github.com/Kawanishi45/cat_robot/pkg/loging"
  "io"
  "net/http"
  "strconv"
)

func GetOpenAIResponseTalk(messages []OpenAiMessage) (responseMessage string, err error) {
  var requestJSON []byte

  // OpenAI APIリクエストのJSONを作成
  requestBody := OpenAiRequest{
    Model:    constants.OpenAiModel,
    Messages: messages,
  }

  //// AIに送信したメッセージをログに残す
  //logAi(messages, "リクエスト")

  if requestJSON, err = json.Marshal(requestBody); err != nil {
    return "", errors.Join(err, errors.New("json.Marshal failed"))
  }

  // OpenAI APIリクエストを送信
  var responseBody []byte
  if responseBody, err = postHttpRequest(requestJSON); err != nil {
    return "", errors.Join(err, errors.New("postHttpRequest failed"))
  }

  //// 生のレスポンスボディを出力
  //loging.SaveLog("Raw response:" + string(responseBody))

  var response OpenAiResponse
  err = json.Unmarshal(responseBody, &response)
  if err != nil {
    return "", errors.Join(err, errors.New("json.Unmarshal failed"))
  }

  // Choicesが空かどうかをチェック
  if len(response.Choices) == 0 {
    return "", errors.New(constants.ErrNoChoicesInResponse)
  }

  // Messages.Contentが空かどうかをチェック
  if response.Choices[0].Messages.Content == "" {
    return "", errors.New(constants.ErrContentIsEmpty)
  }

  // AIに送信したメッセージをログに残す
  logAi(requestBody.Messages, response.Choices[0].Messages.Content)

  return response.Choices[0].Messages.Content, nil
}

func GetOpenAIResponseFunctionCall(messages []OpenAiMessage, evaluateInterface EvaluateToolInterface) (response OpenAiResponseFunction, err error) {
  var requestJSON []byte

  // OpenAI APIリクエストのJSONを作成
  requestBody := OpenAiRequestFunctionCall{
    Model:       constants.OpenAiModel,
    Messages:    messages,
    Tools:       evaluateInterface.getTools(),
    ToolChoice:  "auto",
    Temperature: 0,
  }

  // AIに送信したメッセージをログに残す
  var responseContent string
  defer logAi(requestBody.Messages, responseContent)

  requestJSON, err = json.Marshal(requestBody)
  if err != nil {
    return response, errors.Join(err, errors.New("json.Marshal failed"))
  }

  // OpenAI APIリクエストを送信
  var responseBody []byte
  if responseBody, err = postHttpRequest(requestJSON); err != nil {
    return response, errors.Join(err, errors.New("postHttpRequest failed"))
  }

  //// 生のレスポンスボディを出力
  //loging.SaveLog("Raw response:" + string(responseBody))

  err = json.Unmarshal(responseBody, &response)
  if err != nil {
    return response, errors.Join(err, errors.New("json.Unmarshal failed"))
  }

  // Choicesが空かどうかをチェック
  if len(response.Choices) == 0 {
    return response, errors.New(constants.ErrNoChoicesInResponse)
  }

  // ToolCallsが空かどうかをチェック
  if len(response.Choices[0].Messages.ToolCalls) == 0 {
    // ToolCallsが空の場合は非常に多く、想定内であるため、ロギングしないエラーを返す
    return response, errors.New(constants.ErrNoToolCallsInResponse)
  }

  // Messages.Contentが空かどうかをチェック
  if response.Choices[0].Messages.ToolCalls[0].Function.Name == "" || response.Choices[0].Messages.ToolCalls[0].Function.Arguments == "" {
    return response, errors.New(constants.ErrContentIsEmpty)
  }

  for i, choice := range response.Choices {
    responseContent += "Choices[" + strconv.Itoa(i) + "].Messages.Content:" + choice.Messages.Content + "\n"
    for j, toolcall := range choice.Messages.ToolCalls {
      responseContent += "ToolCalls[" + strconv.Itoa(j) + "] " + toolcall.Function.Name + ":" + toolcall.Function.Arguments + "\n"
    }
  }
  responseContent += "TotalTokens:" + strconv.Itoa(response.Usages.TotalTokens) + "\n"

  if response.Choices[0].Messages.ToolCalls[0].Function.Name != evaluateInterface.getFunctionName() {
    return response, errors.New("function name is not " + evaluateInterface.getFunctionName())
  }

  return
}

func postHttpRequest(requestJSON []byte) (body []byte, err error) {
  // OpenAI APIリクエストを送信
  var req *http.Request
  req, err = http.NewRequest("POST", constants.OpenAIURL, bytes.NewBuffer(requestJSON))
  if err != nil {
    return nil, errors.Join(err, errors.New("http.NewRequest failed"))
  }
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("Authorization", "Bearer "+constants.ApiKey)

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return nil, errors.Join(err, errors.New("client.Do failed"))
    // panic(err)
  }
  defer func(Body io.ReadCloser) {
    err = Body.Close()
    if err != nil {
      return
      // panic(err)
    }
  }(resp.Body)

  // レスポンスボディを読み込む
  body, err = io.ReadAll(resp.Body)
  if err != nil {
    return nil, errors.Join(err, errors.New("io.ReadAll failed"))
    //panic(err)
  }
  return
}

func logAi(messages []OpenAiMessage, responseMessage string) {
  var logText string
  for _, message := range messages {
    logText += message.Role + ":" + message.Content + "\n"
  }
  logText += "\n------------------\n"
  logText += responseMessage
  logText += "\n------------------\n------------------\n\n\n"

  loging.SaveLog(logText)
}
