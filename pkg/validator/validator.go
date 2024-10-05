package validator

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/ahargunyllib/hackathon-fiber-starter/pkg/log"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	idtranslations "github.com/go-playground/validator/v10/translations/id"
)

type Validator struct {
	validator  *validator.Validate
	translator ut.Translator
}

type ValidationError struct {
	Tag         string `json:"tag"`
	Param       string `json:"param"`
	Translation string `json:"translation"`
}

type ValidationErrorsResponse []map[string]ValidationError

func (v ValidationErrorsResponse) Error() string {
	j, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(j)
}

func (v ValidationErrorsResponse) Serialize() any {
	return v
}

func NewValidator() Validator {
	idInstance := id.New()
	uni := ut.New(idInstance, idInstance)

	translator, _ := uni.GetTranslator("id")

	val := validator.New()
	err := idtranslations.RegisterDefaultTranslations(val, translator)
	if err != nil {
		log.Error(log.LogInfo{
			"error": err.Error(),
		}, "[VALIDATOR][NewValidator] Failed to register default translations")
		return Validator{}
	}

	return Validator{
		validator:  val,
		translator: translator,
	}
}

func getJSONFieldName(field reflect.StructField) string {
	checkTags := []string{"json", "query", "param"}
	for _, tag := range checkTags {
		jsonTag := field.Tag.Get(tag)
		if jsonTag != "" {
			return jsonTag
		}
	}

	return field.Name
}

func (v Validator) Validate(data interface{}) ValidationErrorsResponse {
	err := v.validator.Struct(data)
	if err != nil {
		var valErrs validator.ValidationErrors
		if errors.As(err, &valErrs) {
			// Get the type of the data struct
			dataType := reflect.TypeOf(data)
			if dataType.Kind() == reflect.Ptr {
				dataType = dataType.Elem()
			}

			length := len(valErrs)
			res := make(ValidationErrorsResponse, length)
			for i, err := range valErrs {
				field, _ := dataType.FieldByName(err.StructField())
				jsonTag := getJSONFieldName(field)

				res[i] = map[string]ValidationError{
					jsonTag: {
						Tag:         err.Tag(),
						Param:       err.Param(),
						Translation: err.Translate(v.translator),
					},
				}
			}

			return res
		} else {
			log.Error(log.LogInfo{
				"error": err.Error(),
			}, "[VALIDATOR][Validate] Failed to validate data")
		}
	}

	return nil
}
