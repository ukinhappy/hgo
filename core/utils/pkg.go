package utils

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var (
	va    *validator.Validate
	trans ut.Translator
)

type Register struct {
	Func func(validator.FieldLevel) bool
	Msg  string
}

// 初始化 验证器 以及 翻译器
func Init(locale string, register map[string]*Register) (err error) {
	va = validator.New()
	va.RegisterTagNameFunc(func(fld reflect.StructField) string {
		var name string
		switch locale {
		case "en":
			name = strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		case "zh":
			name = fld.Tag.Get("label")
		default:
			name = strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		}
		if name == "-" {
			return ""
		}
		return name
	})

	zhT := zh.New() // 中文翻译器
	enT := en.New() // 英文翻译器
	uni := ut.New(enT, zhT, enT)
	var ok bool
	trans, ok = uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
	}
	switch locale {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(va, trans)
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(va, trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(va, trans)
	}
	for key, value := range register {
		if err := va.RegisterValidation(key, value.Func); err != nil {
			return err
		}

		if err := va.RegisterTranslation(
			key,
			trans,
			registerTranslator(key, value.Msg),
			tFunc,
		); err != nil {
			return err
		}
	}
	return
}

func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

func tFunc(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}

// 把数据验证错误，转为放回错误
func GetMsg(err error) string {
	errors, _ := err.(validator.ValidationErrors)
	var errs []string
	for _, v := range errors.Translate(trans) {
		errs = append(errs, v)

	}
	return strings.Join(errs, ",")
}

// 数据验证
func Struct(data interface{}) error {
	return va.Struct(data)
}
