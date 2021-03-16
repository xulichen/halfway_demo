package orm

type Demo struct {
	Name string `json:"name" gorm:"column:name"`
}

func (d *Demo) TableName() string {
	return "demo"
}
