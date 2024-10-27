package models

import (
	"gorm.io/gorm"
	"time"
)

// 自增ID主键
type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

// 创建、更新时间
type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 软删除
type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}


// 多语言
type FieldLanguage struct {
	nameAr string `json:"nameAr"`
	nameEn string `json:"nameEn"`
	nameCn string `json:"nameCn"`
}

// 状态的不同颜色
type StatusColor struct {
	TextColor string `json:"textColor"`
	TextBgcColor string `json:"textBgcColor"`
}
