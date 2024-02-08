package constants

var OpenAiModel = OpenAiModelGpt4

// const openAI API
const (
  OpenAiModelGpt3Turbo = "gpt-3.5-turbo"
  OpenAiModelGpt4      = "gpt-4-0613"
  OpenAiRoleSystem     = "system"
  OpenAiRoleUser       = "user"
  OpenAiRoleAssistant  = "assistant"
  OpenAIURL            = "https://api.openai.com/v1/chat/completions"
)

// const role
const (
  RoleAI     = 0
  RoleUser   = 1
  RoleSystem = 2
)

var RoleName = []string{
  RoleAI:     OpenAiRoleAssistant,
  RoleUser:   OpenAiRoleUser,
  RoleSystem: OpenAiRoleSystem,
}

var AiLogPath = "./logs/" // 会話ログの保存先。main.goで「./fileName」で設定する

type TalkMethodNum int

type HostType string

var ApiKey string

const (
  ErrNoChoicesInResponse   = "NoChoicesInResponse"
  ErrNoToolCallsInResponse = "NoToolCallsInResponse"
  ErrContentIsEmpty        = "ContentIsEmpty"
)

var CheckAgainErrors = []string{
  ErrNoChoicesInResponse,
  ErrNoToolCallsInResponse,
  ErrContentIsEmpty,
}

// OpenApi type
const (
  Function = "function"
  Object   = "object"
  Boolean  = "boolean"
  Number   = "number"
  String   = "string"
)

// OsType device
type OsType int

// const functions of what is actions
const (
  FunctionNameScoreCondition            = "score_condition"
  FunctionDescriptionCondition          = "相手の感情スコアを判定し、その結果を返します。"
  PropertyNameScoreCondition            = "score_condition"
  PropertyDescriptionScoreCondition     = "相手の感情スコアを判定し、その結果を返します。「幸せそうか」をhappy_score、「興奮しているか」をexcited_score、「怒っているか」をangry_score、「悲しいか」をsadness_scoreとして返します。"
  PropertyDescriptionHappyScore         = "相手はどのくらい幸せそう? (最も不幸せなら0、最も幸せなら10)"
  PropertyDescriptionReasonHappyScore   = "happy_scoreがその数値である根拠を20文字以内で説明してください。"
  PropertyDescriptionExcitedScore       = "相手はどのくらい興奮している? (最も冷静なら0、最も興奮しているなら10)"
  PropertyDescriptionReasonExcitedScore = "excited_scoreがその数値である根拠を20文字以内で説明してください。"
  PropertyDescriptionAngryScore         = "相手はどのくらい怒っている? (最も怒っているなら0、最も穏やかなら10)"
  PropertyDescriptionReasonAngryScore   = "angry_scoreがその数値である根拠を20文字以内で説明してください。"
  PropertyDescriptionSadnessScore       = "相手はどのくらい悲しい? (最も楽観的なら0、最も悲観的なら10)"
  PropertyDescriptionReasonSadnessScore = "sadness_scoreがその数値である根拠を20文字以内で説明してください。"
)

const (

  // ---------------------------function_calling用---------------------------

  FunctionCallScoreCondition = `
ここまでの会話ログから、userの感情スコアがどのような状態にあるかを判定し、その結果を「score_condition」functionをcallして送信してください。
score_conditionでは、ここまでの会話でuserの感情スコアがどのような状態にあるかを判定し、その結果を「HappyScore」「ExcitedScore」「AngryScore」「SadnessScore」の各パラメータに示してください。
各パラメータの値は0から10の間の整数で、それぞれの感情スコアが高いほど10に近づくようにしてください。
`
)
