package datadto

type CreateDataRequest struct {
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
}

type DataFile struct {
	Image    string `json:"image"`
	Sizefile int64  `json:"Sizefile"`
	Format   string `json:"format"`
	FileName string `json:"filename"`
}
