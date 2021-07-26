package valid

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// 检查密码
// 密码规则
// 支持数字、字母、特殊字符
// 至少两种类型的组合
// 密码长度至少8位 在外边验证
func CheckPassword(password string /*, minLength, maxLength int*/) bool {
	// regular := fmt.Sprintf(`/(?![\d]+$)(?![a-zA-Z]+$)(?![-=+_.,]+$)[\da-zA-Z-=+_.,]{%d,%d}/`, minLength, maxLength)
	regular := `/(?![\d]+$)(?![a-zA-Z]+$)(?![-=+_.,]+$)[\da-zA-Z-=+_.,]/`
	reg := regexp.MustCompile(regular)

	return reg.MatchString(password)
}

func (v *Validate) checkPassword(fl validator.FieldLevel) bool {
	return CheckPassword(fl.Field().String())
}
