package openAI

import (
  "github.com/Kawanishi45/cat_robot/pkg/constants"
)

type EvaluateToolInterface interface {
  getCustomPrompt() string
  getTools() []ToolInterface
  getFunctionName() string
}

// EvaluateIsExistsAction アクションが存在するかどうかを判定する
type EvaluateIsExistsAction struct{}

func (e EvaluateIsExistsAction) getCustomPrompt() string {
  return constants.FunctionCallWhatIsAction
}
func (e EvaluateIsExistsAction) getTools() []ToolInterface {
  return []ToolInterface{
    {
      Type: constants.Function,
      Function: FunctionIsExistsAction{
        Name:        constants.FunctionNameUpdateIsExistsAction,
        Description: constants.FunctionDescriptionUpdateIsExistsAction,
        Parameters: ParametersIsExistsAction{
          Type: constants.Object,
          Properties: PropertiesIsExistsAction{
            IsDecidedAction: Property{
              Type:        constants.Boolean,
              Description: constants.PropertyDescriptionIsDecidedAction,
            },
            ActionTitle: Property{
              Type:        constants.String,
              Description: constants.PropertyDescriptionActionTitle,
            },
            ActionDetail: Property{
              Type:        constants.String,
              Description: constants.PropertyDescriptionActionDetail,
            },
            IsDecidedActionFrequency: Property{
              Type:        constants.Boolean,
              Description: constants.PropertyDescriptionIsDecidedActionFrequency,
            },
            ActionFrequency: PropertyEnum{
              Type:        constants.String,
              Description: constants.PropertyDescriptionActionFrequency,
              Enum:        constants.ActionFrequencyEnum,
            },
            IsDecidedActionDayOfWeek: Property{
              Type:        constants.Boolean,
              Description: constants.PropertyDescriptionIsDecidedActionDayOfWeek,
            },
            ActionDayOfWeek: PropertyObject{
              Type:        constants.Object,
              Description: constants.PropertyDescriptionActionDayOfWeek,
              Properties: PropertiesBoolDaysOfWeek{
                IsSunday:    Property{Type: constants.Boolean, Description: constants.PropertyDescriptionIsSunday},
                IsMonday:    Property{Type: constants.Boolean, Description: constants.PropertyDescriptionIsMonday},
                IsTuesday:   Property{Type: constants.Boolean, Description: constants.PropertyDescriptionIsTuesday},
                IsWednesday: Property{Type: constants.Boolean, Description: constants.PropertyDescriptionIsWednesday},
                IsThursday:  Property{Type: constants.Boolean, Description: constants.PropertyDescriptionIsThursday},
                IsFriday:    Property{Type: constants.Boolean, Description: constants.PropertyDescriptionIsFriday},
                IsSaturday:  Property{Type: constants.Boolean, Description: constants.PropertyDescriptionIsSaturday},
              },
            },
            IsDecidedActionTimeToStart: Property{
              Type:        constants.Boolean,
              Description: constants.PropertyDescriptionIsDecidedActionTimeToStart,
            },
            ActionTimeToStart: Property{
              Type:        constants.String,
              Description: constants.PropertyDescriptionActionTimeToStart,
            },
          },
          Required: []string{
            constants.PropertyNameIsDecidedAction,
            constants.PropertyNameActionTitle,
            constants.PropertyNameActionDetail,
            constants.PropertyNameIsDecidedActionFrequency,
            constants.PropertyNameActionFrequency,
            constants.PropertyNameIsDecidedActionDayOfWeek,
            constants.PropertyNameActionDayOfWeek,
            constants.PropertyNameIsDecidedActionTimeToStart,
            constants.PropertyNameActionTimeToStart,
          },
        },
      },
    },
  }
}
func (e EvaluateIsExistsAction) getFunctionName() string {
  return constants.FunctionNameUpdateIsExistsAction
}

type ArgumentsUpdateUserAction struct {
  IsDecidedAction            bool                  `json:"is_decided_action"`
  ActionTitle                string                `json:"action_title"`
  ActionDetail               string                `json:"action_detail"`
  IsDecidedActionFrequency   bool                  `json:"is_decided_action_frequency"`
  ActionFrequency            string                `json:"action_frequency"`
  IsDecidedActionDayOfWeek   bool                  `json:"is_decided_action_day_of_week"`
  ActionDayOfWeek            ActionDayOfWeekObject `json:"action_day_of_week"`
  IsDecidedActionTimeToStart bool                  `json:"is_decided_action_time_to_start"`
  ActionTimeToStart          string                `json:"action_time_to_start"`
}
type ActionDayOfWeekObject struct {
  IsSunday    bool `json:"is_sunday"`
  IsMonday    bool `json:"is_monday"`
  IsTuesday   bool `json:"is_tuesday"`
  IsWednesday bool `json:"is_wednesday"`
  IsThursday  bool `json:"is_thursday"`
  IsFriday    bool `json:"is_friday"`
  IsSaturday  bool `json:"is_saturday"`
}
