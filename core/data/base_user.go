package data

// 基础用户数据
type BaseUser struct {
	BaseStruct `xorm:"extends"`

	// 用户名
	Username string `xorm:"varchar(32) notnull default('') unique(uidx_name)" json:"username" validate:"omitempty,min=1,max=32,email"`
	// 手机号
	// 类似于+86-13684001001
	Phone string `xorm:"varchar(15) notnull default('') unique(uidx_phone)" json:"phone" validate:"omitempty,mobile"`
	// 密码
	Password string `xorm:"varchar(512) notnull default('')" json:"-"`
}
