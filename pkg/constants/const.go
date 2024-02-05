package constants

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
  String   = "string"
)

// frequency
const (
  FrequencyEveryDay       = "every_day"
  FrequencyEveryWeek      = "every_week"
  FrequencyEveryOtherWeek = "every_other_week"
  FrequencyEveryMonth     = "every_month"
  FrequencyNone           = "none"
)

// OsType device
type OsType int

// const functions of what is actions
const (
  FunctionNameUpdateIsExistsAction              = "update_user_action"
  FunctionDescriptionUpdateIsExistsAction       = "update user's action. if user decide action, return true. if not, return false."
  PropertyNameIsDecidedAction                   = "is_decided_action"
  PropertyNameActionTitle                       = "action_title"
  PropertyNameActionDetail                      = "action_detail"
  PropertyNameIsDecidedActionFrequency          = "is_decided_action_frequency"
  PropertyNameActionFrequency                   = "action_frequency"
  PropertyNameIsDecidedActionDayOfWeek          = "is_decided_action_day_of_week"
  PropertyNameActionDayOfWeek                   = "action_day_of_week"
  PropertyNameIsDecidedActionTimeToStart        = "is_decided_action_time_to_start"
  PropertyNameActionTimeToStart                 = "action_time_to_start"
  PropertyDescriptionIsDecidedAction            = "Has the user decided on his or her action? (true or false)"
  PropertyDescriptionActionTitle                = "about what user's action. action's title. very short sentence."
  PropertyDescriptionActionDetail               = "about what user's action. action's detail. long sentence."
  PropertyDescriptionIsDecidedActionFrequency   = "Has the user decided on his or her action's frequency? (true or false)"
  PropertyDescriptionActionFrequency            = "about how often user do action. every day? every week? every other week? every month? etc..."
  PropertyDescriptionIsDecidedActionDayOfWeek   = "Has the user decided on his or her action's day of week? (true or false)"
  PropertyDescriptionActionDayOfWeek            = "about when day of week user do action. on Monday? on Tuesday? on Wednesday and Friday? etc..."
  PropertyDescriptionIsSunday                   = "is user do action when sunday? (true or false)"
  PropertyDescriptionIsMonday                   = "is user do action when monday? (true or false)"
  PropertyDescriptionIsTuesday                  = "is user do action when tuesday? (true or false)"
  PropertyDescriptionIsWednesday                = "is user do action when wednesday? (true or false)"
  PropertyDescriptionIsThursday                 = "is user do action when thursday? (true or false)"
  PropertyDescriptionIsFriday                   = "is user do action when friday? (true or false)"
  PropertyDescriptionIsSaturday                 = "is user do action when saturday? (true or false)"
  PropertyDescriptionIsDecidedActionTimeToStart = "Has the user decided on his or her action's time to start? (true or false)"
  PropertyDescriptionActionTimeToStart          = "about when time to start user do action. format is[hour]:[minute]. if user don't decided, none. else if user already decided, 8:00 or 9:30 or 20:15 etc..."
)

var ActionFrequencyEnum = []string{FrequencyEveryDay, FrequencyEveryWeek, FrequencyEveryOtherWeek, FrequencyEveryMonth, FrequencyNone}

const (

  // ---------------------------function_calling用---------------------------

  // FunctionCallWhatIsAction 【function_calling用】目標解像度が1でアクションをAIに判定させるプロンプト
  FunctionCallWhatIsAction = `
ここまでの会話ログから、userが自分自身の目標を実現するための日々の行動・アクションを決められているかを判定し、その結果を「update_user_action」functionをcallして送信してください。
update_user_actionでは、ここまでの会話でuserが自身の行動を決める内容が見つかれば「IsDecidedAction」パラメータのbooleanを「true」に、見つからなければ「false」を指定し、
「IsDecidedAction」パラメータのbooleanを「true」にする時は、行動の内容を「what_action」パラメータのstringに示してください。そうでない時は「what_action」パラメータは空文字で返してください。
`
)
