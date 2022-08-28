package ts_ad_cs_white

import "time"

// TsAdCsWhite 白名单表
//go:generate gormgen -structs TsAdCsWhite -input .
type TsAdCsWhite struct {
	Id          int32  `gorm:"id"`          // 主键
	Name        string `gorm:"name"`        //
	Os          string `gorm:"os"`          // 系统名称
	Idfa        string `gorm:"idfa"`        // idfa
	Oaid        string `gorm:"oaid"`        //
	DeviceId    string `gorm:"device_id"`   // device_id
	UserId      string `gorm:"user_id"`     // 用户ID
	Type        string `gorm:"type"`        //
	VerifySign  string `gorm:"verify_sign"` // 对应校验的标识cs就是计划ID、api就是api标识、SDK就是SDK标识
	Description string `gorm:"description"` //
	Remark      string `gorm:"remark"`      //
	IsDeleted   int32  `gorm:"is_deleted"`  //
	CreateTime  int64  `gorm:"create_time"` // 创建时间
	CreateBy    string `gorm:"create_by"`   //
	UpdateBy    string `gorm:"update_by"`   //
}
