package valid

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enLang "github.com/go-playground/validator/v10/translations/en"
	zhLang "github.com/go-playground/validator/v10/translations/zh"
)

type Validate struct {
	validator  *validator.Validate
	translator *ut.UniversalTranslator
}

func NewValidator() (v *Validate) {
	v = &Validate{}
	v.validator = validator.New()
	v.translator = ut.New(en.New(), zh.New())

	english, _ := v.translator.GetTranslator("en")
	if err := enLang.RegisterDefaultTranslations(v.validator, english); nil != err {
		panic(err)
	}
	chinese, _ := v.translator.GetTranslator("zh")
	if err := zhLang.RegisterDefaultTranslations(v.validator, chinese); nil != err {
		panic(err)
	}

	if err := v.initValidation(); nil != err {
		panic(err)
	}

	return
}

func (v *Validate) Validate(data interface{}) error {
	return v.validator.Struct(data)
}

func (v *Validate) initValidation() (err error) {
	if err = v.validator.RegisterValidation("mobile", v.checkMobile); nil != err {
		return
	}

	if err = v.validator.RegisterValidation("password", v.checkPassword); nil != err {
		return
	}

	if err = v.validator.RegisterValidation("start_with_alpha", v.checkStartWithAlpha); nil != err {
		return
	}

	return
}
