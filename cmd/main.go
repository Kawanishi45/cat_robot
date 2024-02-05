package main

import (
  "bufio"
  "encoding/json"
  "fmt"
  "github.com/Kawanishi45/cat_robot/pkg/constants"
  "github.com/Kawanishi45/cat_robot/pkg/loging"
  "github.com/Kawanishi45/cat_robot/pkg/openAI"
  "os"
)

type Message struct {
  Role int
  Text string
}

func main() {
  var err error
  constants.ApiKey = os.Getenv("OPENAI_API_KEY")
  if constants.ApiKey == "" {
    fmt.Println("環境変数OPENAI_API_KEYが設定されていません。")
    return
  }

  var fileName string
  fileName, err = loging.CreateLog()
  if err != nil {
    fmt.Println("\nエラー:", err)
    return
  }
  constants.AiLogPath += fileName

  scanner := bufio.NewScanner(os.Stdin)
  var messages []Message

  var count int
  count = 2
  for {
    fmt.Print("\nあなた👦：")
    scanner.Scan()
    inputText := scanner.Text()

    if inputText == "exit" { // exitを入力するとループから抜ける
      break
    }

    messages = append(messages, Message{Role: constants.RoleUser, Text: inputText})

    if len(messages)/2 > count {
      // 4回目のメッセージからfunction callingを行い、以降は毎回function callingを行う
      fmt.Println("\nsystem：function calling...")

      var args openAI.ArgumentsUpdateUserAction
      args, err = GetChatGPTResponseFunctionCalling(messages)
      if err != nil {
        fmt.Println("\nfunction calling error:", err)
      }
      fmt.Println("\nargs:", args)
    }

    var responseMessage string
    responseMessage, err = GetChatGPTResponseMessage(messages)
    if err != nil {
      fmt.Println("\nエラー:", err)
      return
    }
    replyMessage := responseMessage
    messages = append(messages, Message{Role: constants.RoleAI, Text: replyMessage})
    fmt.Println("\n猫ロボ🤖：", replyMessage)
  }

  fmt.Println("\n会話ログ:")
  for _, message := range messages {
    fmt.Println(message)
  }
}

func GetChatGPTResponseMessage(messages []Message) (responseMessage string, err error) {
  var openAIMessages []openAI.OpenAiMessage
  // CustomInteractionPromptを追加
  systemMessage := openAI.OpenAiMessage{
    Role:    constants.OpenAiRoleSystem,
    Content: "",
  }
  openAIMessages = append(openAIMessages, systemMessage)

  // 会話ログを追加（会話ログは昇順で作る）
  for _, recentMessage := range messages {
    // ユーザーのメッセージを追加
    talkMessage := openAI.OpenAiMessage{
      Role:    constants.RoleName[recentMessage.Role],
      Content: recentMessage.Text,
    }
    openAIMessages = append(openAIMessages, talkMessage)
  }

  responseMessage, err = openAI.GetOpenAIResponseTalk(openAIMessages)
  if err != nil {
    return
  }
  return
}

func GetChatGPTResponseFunctionCalling(messages []Message) (args openAI.ArgumentsUpdateUserAction, err error) {
  var openAIMessages []openAI.OpenAiMessage
  // CustomInteractionPromptを追加
  systemMessage := openAI.OpenAiMessage{
    Role:    constants.OpenAiRoleSystem,
    Content: "",
  }
  openAIMessages = append(openAIMessages, systemMessage)

  // 会話ログを追加（会話ログは昇順で作る）
  for _, recentMessage := range messages {
    // ユーザーのメッセージを追加
    talkMessage := openAI.OpenAiMessage{
      Role:    constants.RoleName[recentMessage.Role],
      Content: recentMessage.Text,
    }
    openAIMessages = append(openAIMessages, talkMessage)
  }

  var response openAI.OpenAiResponseFunction
  response, err = openAI.GetOpenAIResponseFunctionCall(openAIMessages, openAI.EvaluateIsExistsAction{})
  if err != nil {
    return
  }

  err = json.Unmarshal([]byte(response.Choices[0].Messages.ToolCalls[0].Function.Arguments), &args)
  if err != nil {
    return
  }
  return
}
