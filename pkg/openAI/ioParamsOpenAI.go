package openAI

type OpenAiRequest struct {
  Model    string          `json:"model"`
  Messages []OpenAiMessage `json:"messages"`
}

type OpenAiRequestFunctionCall struct {
  Model       string          `json:"model"`
  Messages    []OpenAiMessage `json:"messages"`
  Tools       []ToolInterface `json:"tools"`
  ToolChoice  string          `json:"tool_choice"` // auto is default, but we'll be explicit
  Temperature float64         `json:"temperature"`
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

type PropertyConditionObject struct {
  Type        string                   `json:"type"`
  Description string                   `json:"description"`
  Properties  PropertiesScoreCondition `json:"properties"`
}

type FunctionCondition struct {
  Name        string              `json:"name"`
  Description string              `json:"description"`
  Parameters  ParametersCondition `json:"parameters"`
}

type ParametersCondition struct {
  Type       string              `json:"type"`
  Properties PropertiesCondition `json:"properties"`
  Required   []string            `json:"required"`
}

type PropertiesCondition struct {
  Condition PropertyConditionObject `json:"score_condition"`
}

type PropertiesScoreCondition struct {
  HappyScore         Property `json:"happy_score"`
  ReasonHappyScore   Property `json:"reason_happy_score"`
  ExcitedScore       Property `json:"excited_score"`
  ReasonExcitedScore Property `json:"reason_excited_score"`
  AngryScore         Property `json:"angry_score"`
  ReasonAngryScore   Property `json:"reason_angry_score"`
  SadnessScore       Property `json:"sadness_score"`
  ReasonSadnessScore Property `json:"reason_sadness_score"`
}
