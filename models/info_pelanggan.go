package models

type Info_pelanggan struct {
	Id_plg     int64  `gorm:"primaryKey" json:"id_plg"`
	Nama_plg   string `gorm:"type:varchar(50)" json:"nama_plg" binding:"required"`
	Alamat_plg string `gorm:"type:varchar(50)" json:"alamat_plg" binding:"required"`
}

type Tabler interface {
	TableName() string
}

func (Info_pelanggan) TableName() string {
	return "Info_pelanggan"
}
