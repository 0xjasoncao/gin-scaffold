package validatorutil

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

type ValidationErrors validator.ValidationErrors

var zhTrans ut.Translator
var enTrans ut.Translator

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		chinese := zh.New()
		uni := ut.New(chinese, en.New())
		zhTrans, _ = uni.GetTranslator("zh")
		enTrans, _ = uni.GetTranslator("en")

		_ = zh_translations.RegisterDefaultTranslations(v, zhTrans)
		_ = en_translations.RegisterDefaultTranslations(v, enTrans)

	}
}

func ZhTranslate(err error) string {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		var errMessages []string
		for _, err := range errs {
			errMessages = append(errMessages, err.Translate(zhTrans))

		}
		return strings.Join(errMessages, ";")
	}

	return err.Error()
}
