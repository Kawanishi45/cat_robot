package openAI

type OpenAiRequest struct {
  Model    string          `json:"model"`
  Messages []OpenAiMessage `json:"messages"`
}

type OpenAiRequestFunctionCall struct {
  Model      string          `json:"model"`
  Messages   []OpenAiMessage `json:"messages"`
  Tools      []ToolInterface `json:"tools"`
  ToolChoice string          `json:"tool_choice"` // auto is default, but we'll be explicit
}

type OpenAiResponse struct {
  Id      string         `json:"id"`
  Object  string         `json:"object"`
  Created int            `json:"created"`
  Choices []OpenAiChoice `json:"choices"`
  Usages  OpenAiUsage    `json:"usage"`
}

type OpenAiResponseFunction struct {
  Id      string                 `json:"id"`
  Object  string                 `json:"object"`
  Created int                    `json:"created"`
  Choices []OpenAiChoiceFunction `json:"choices"`
  Usages  OpenAiUsage            `json:"usage"`
}

type OpenAiChoice struct {
  Index        int           `json:"index"`
  Messages     OpenAiMessage `json:"message"`
  FinishReason string        `json:"finish_reason"`
}

type OpenAiChoiceFunction struct {
  Index        int                   `json:"index"`
  Messages     OpenAiMessageFunction `json:"message"`
  FinishReason string                `json:"finish_reason"`
}

type OpenAiMessage struct {
  Role    string `json:"role"`
  Content string `json:"content"`
}

type OpenAiMessageFunction struct {
  Role      string     `json:"role"`
  Content   string     `json:"content"`
  ToolCalls []ToolCall `json:"tool_calls"`
}

type ToolCall struct {
  Id       string           `json:"id"`
  Type     string           `json:"type"`
  Function FunctionResponse `json:"function"`
}

type FunctionResponse struct {
  Name      string `json:"name"`
  Arguments string `json:"arguments"`
}

type OpenAiUsage struct {
  PromptTokens     int `json:"prompt_tokens"`
  CompletionTokens int `json:"completion_tokens"`
  TotalTokens      int `json:"total_tokens"`
}

type ToolInterface struct {
  Type     string      `json:"type"`
  Function interface{} `json:"function"`
}

type Property struct {
  Type        string `json:"type"`
  Description string `json:"description"`
}

type PropertyEnum struct {
  Type        string   `json:"type"`
  Description string   `json:"description"`
  Enum        []string `json:"enum"`
}

type PropertyObject struct {
  Type        string                   `json:"type"`
  Description string                   `json:"description"`
  Properties  PropertiesBoolDaysOfWeek `json:"properties"`
}

type FunctionIsExistsAction struct {
  Name        string                   `json:"name"`
  Description string                   `json:"description"`
  Parameters  ParametersIsExistsAction `json:"parameters"`
}

type ParametersIsExistsAction struct {
  Type       string                   `json:"type"`
  Properties PropertiesIsExistsAction `json:"properties"`
  Required   []string                 `json:"required"`
}

type PropertiesIsExistsAction struct {
  IsDecidedAction            Property       `json:"is_decided_action"`
  ActionTitle                Property       `json:"action_title"`
  ActionDetail               Property       `json:"action_detail"`
  IsDecidedActionFrequency   Property       `json:"is_decided_action_frequency"`
  ActionFrequency            PropertyEnum   `json:"action_frequency"`
  IsDecidedActionDayOfWeek   Property       `json:"is_decided_action_day_of_week"`
  ActionDayOfWeek            PropertyObject `json:"action_day_of_week"`
  IsDecidedActionTimeToStart Property       `json:"is_decided_action_time_to_start"`
  ActionTimeToStart          Property       `json:"action_time_to_start"`
}

type PropertiesBoolDaysOfWeek struct {
  IsSunday    Property `json:"is_sunday"`
  IsMonday    Property `json:"is_monday"`
  IsTuesday   Property `json:"is_tuesday"`
  IsWednesday Property `json:"is_wednesday"`
  IsThursday  Property `json:"is_thursday"`
  IsFriday    Property `json:"is_friday"`
  IsSaturday  Property `json:"is_saturday"`
}
