package models

type Data struct {
	ID       int    `json:"id" gorm:"PRIMARY_KEY"`
	Image    string `json:"image"  form:"image" gorm:"type: varchar(255)"`
	FileName string `json:"filename"`
	Sizefile int64  `json:"Sizefile"`
	Format   string `json:"format"`
}
