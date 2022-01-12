package valid

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

// 字符串必须以字母开头
func StartWithAlpha(str string) (checked bool) {
	if "" == str {
		return
	}
	checked = unicode.IsLetter(rune(str[0]))

	return
}

func (v *Validate) checkStartWithAlpha(fl validator.FieldLevel) bool {
	return StartWithAlpha(fl.Field().String())
}
