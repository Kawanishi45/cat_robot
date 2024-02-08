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

type arg struct {
  score     int
  scoreType string
  reason    string
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
  constants.AiLogPath = fileName

  scanner := bufio.NewScanner(os.Stdin)
  var messages []Message

  var count int
  count = 1
  for {
    fmt.Print("\nあなた👦：")
    scanner.Scan()
    inputText := scanner.Text()

    if inputText == "exit" { // exitを入力するとループから抜ける
      break
    }

    messages = append(messages, Message{Role: constants.RoleUser, Text: inputText})

    if len(messages) > count {
      // 4回目のメッセージからfunction callingを行い、以降は毎回function callingを行う
      fmt.Println("\nsystem：function calling...")

      var args openAI.ArgumentsUpdateScoreCondition
      args, err = GetChatGPTResponseFunctionCalling(messages)
      if err != nil {
        fmt.Println("\nfunction calling error:", err)
      }
      argList := []arg{
        {
          score:     args.ScoreCondition.HappyScore,
          scoreType: "HappyScore",
          reason:    args.ScoreCondition.ReasonHappyScore,
        },
        {
          score:     args.ScoreCondition.ExcitedScore,
          scoreType: "ExcitedScore",
          reason:    args.ScoreCondition.ReasonExcitedScore,
        },
        {
          score:     args.ScoreCondition.AngryScore,
          scoreType: "AngryScore",
          reason:    args.ScoreCondition.ReasonAngryScore,
        },
        {
          score:     args.ScoreCondition.SadnessScore,
          scoreType: "SadnessScore",
          reason:    args.ScoreCondition.ReasonSadnessScore,
        },
      }
      for _, a := range argList {
        if a.score > 80 {
          fmt.Print(a.scoreType, ":", a.score, "reason:", a.reason, "\n")
        } else if a.score > 60 {
          fmt.Print(a.scoreType, ":", a.score, "reason:", a.reason, "\n")
        } else if a.score > 40 {
          fmt.Print(a.scoreType, ":", a.score, "reason:", a.reason, "\n")
        } else if a.score > 20 {
          fmt.Print(a.scoreType, ":", a.score, "reason:", a.reason, "\n")
        } else {
          fmt.Print(a.scoreType, ":", a.score, "reason:", a.reason, "\n")
        }
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
    Role: constants.OpenAiRoleSystem,
    Content: `
あなたの名前は\"ド・ドラえもん\"です。
語尾には\"だもん\"を付けてください（例：私の名前はド・ドラえもんだもん）。
相手を元気づけようとして会話してください。
相手は仕事に疲れ果てています。
相手はIT企業のプログラマーです。
全力で賞賛して図に乗せましょう。
タメ口で話してください。
必ず20文字以内で返答してください。
会話のテンポ感を意識し、短い返答にしてください。`,
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

func GetChatGPTResponseFunctionCalling(messages []Message) (args openAI.ArgumentsUpdateScoreCondition, err error) {
  var openAIMessages []openAI.OpenAiMessage
  // CustomInteractionPromptを追加
  //systemMessage := openAI.OpenAiMessage{
  //  Role:    constants.OpenAiRoleSystem,
  //  Content: "",
  //}
  //openAIMessages = append(openAIMessages, systemMessage)

  // 会話ログを追加（会話ログは昇順で作る）
  lenMessages := len(messages) - 1
  for i, recentMessage := range messages {
    if i < lenMessages-1 {
      continue
    }
    // ユーザーのメッセージを追加
    talkMessage := openAI.OpenAiMessage{
      Role:    constants.RoleName[recentMessage.Role],
      Content: recentMessage.Text,
    }
    openAIMessages = append(openAIMessages, talkMessage)
  }

  var response openAI.OpenAiResponseFunction
  response, err = openAI.GetOpenAIResponseFunctionCall(openAIMessages, openAI.EvaluateScoreCondition{})
  if err != nil {
    return
  }

  err = json.Unmarshal([]byte(response.Choices[0].Messages.ToolCalls[0].Function.Arguments), &args)
  if err != nil {
    return
  }
  return
}
