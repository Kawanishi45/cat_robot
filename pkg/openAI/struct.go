package openAI

import (
  "github.com/Kawanishi45/cat_robot/pkg/constants"
)

type EvaluateToolInterface interface {
  getCustomPrompt() string
  getTools() []ToolInterface
  getFunctionName() string
}

// EvaluateScoreCondition は、スコアによる条件を評価するための構造体
type EvaluateScoreCondition struct{}

func (e EvaluateScoreCondition) getCustomPrompt() string {
  return constants.FunctionCallScoreCondition
}
func (e EvaluateScoreCondition) getTools() []ToolInterface {
  return []ToolInterface{
    {
      Type: constants.Function,
      Function: FunctionCondition{
        Name:        constants.FunctionNameScoreCondition,
        Description: constants.FunctionDescriptionCondition,
        Parameters: ParametersCondition{
          Type: constants.Object,
          Properties: PropertiesCondition{
            Condition: PropertyConditionObject{
              Type:        constants.Object,
              Description: constants.PropertyDescriptionScoreCondition,
              Properties: PropertiesScoreCondition{
                HappyScore:         Property{Type: constants.Number, Description: constants.PropertyDescriptionHappyScore},
                ReasonHappyScore:   Property{Type: constants.String, Description: constants.PropertyDescriptionReasonHappyScore},
                ExcitedScore:       Property{Type: constants.Number, Description: constants.PropertyDescriptionExcitedScore},
                ReasonExcitedScore: Property{Type: constants.String, Description: constants.PropertyDescriptionReasonExcitedScore},
                AngryScore:         Property{Type: constants.Number, Description: constants.PropertyDescriptionAngryScore},
                ReasonAngryScore:   Property{Type: constants.String, Description: constants.PropertyDescriptionReasonAngryScore},
                SadnessScore:       Property{Type: constants.Number, Description: constants.PropertyDescriptionSadnessScore},
                ReasonSadnessScore: Property{Type: constants.String, Description: constants.PropertyDescriptionReasonSadnessScore},
              },
            },
          },
          Required: []string{
            constants.PropertyNameScoreCondition,
          },
        },
      },
    },
  }
}
func (e EvaluateScoreCondition) getFunctionName() string {
  return constants.FunctionNameScoreCondition
}

type ArgumentsUpdateScoreCondition struct {
  ScoreCondition ScoreConditionObject `json:"score_condition"`
}
type ScoreConditionObject struct {
  HappyScore         int    `json:"happy_score"`
  ReasonHappyScore   string `json:"reason_happy_score"`
  ExcitedScore       int    `json:"excited_score"`
  ReasonExcitedScore string `json:"reason_excited_score"`
  AngryScore         int    `json:"angry_score"`
  ReasonAngryScore   string `json:"reason_angry_score"`
  SadnessScore       int    `json:"sadness_score"`
  ReasonSadnessScore string `json:"reason_sadness_score"`
}
