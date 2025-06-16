package validator

import (
	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		uni := ut.New(zh.New())
		trans, _ := uni.GetTranslator("zh")
		err := zh_translations.RegisterDefaultTranslations(v, trans)
		if err != nil {
			return
		}
		Trans = trans

		// 适用 struct 标签 `label` 来作为字段名
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			label := field.Tag.Get("label")
			if label == "" {
				return field.Name
			}
			return label
		})
	}
}

func BindAndValidateFirst(c *gin.Context, obj interface{}) (string, bool) {
	if err := c.ShouldBindJSON(obj); err != nil {
		errorTrans := TranslateError(err)
		return errorTrans[0], false
	}
	return "", true
}

func BindAndValidate(c *gin.Context, obj interface{}) ([]string, bool) {
	if err := c.ShouldBindJSON(obj); err != nil {
		errorTrans := TranslateError(err)
		return errorTrans, false
	}
	return nil, true
}

func TranslateError(err error) []string {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		result := make([]string, 0, len(errs))
		for _, e := range errs.Translate(Trans) {
			result = append(result, e)
		}
		return result
	}
	return []string{err.Error()}
}
