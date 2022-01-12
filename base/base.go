package base

import (
	"strconv"
)

type (
	// IdStruct 带序列号的模型
	IdStruct struct {
		// Id 编号
		Id int64 `xorm:"pk notnull unique index('idx_id') default(0)" json:"id,string"`
	}

	// CreatedStruct 带创建时间模型
	CreatedStruct struct {
		// CreatedAt 创建时间
		CreatedAt int64 `xorm:"created" json:"createdAt,string"`
	}

	// UpdatedStruct 带修改时间模型
	UpdatedStruct struct {
		// UpdatedAt 最后更新时间
		UpdatedAt int64 `xorm:"updated" json:"updatedAt,string"`
	}

	// DeletedStruct 软删除模型
	DeletedStruct struct {
		// DeletedAt 删除时间，用户软删除
		DeletedAt int64 `xorm:"deleted" json:"deletedAt,string"`
	}

	// BaseStruct 基础数据库模型
	BaseStruct struct {
		IdStruct      `xorm:"extends"`
		CreatedStruct `xorm:"extends"`
		UpdatedStruct `xorm:"extends"`
	}

	// SoftDeleteStruct 带软删除功能的数据库模型
	SoftDeleteStruct struct {
		BaseStruct    `xorm:"extends"`
		DeletedStruct `xorm:"extends"`
	}
)

// IdString Id的字符串形式
func (is *IdStruct) IdString() string {
	return strconv.FormatInt(is.Id, 10)
}

// Exists 对象是否存在
func (is *IdStruct) Exists() bool {
	return 0 != is.Id
}

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
