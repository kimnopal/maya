package entity

type ContactType struct {
	ID   uint64 `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}

func (c *ContactType) TableName() string {
	return "contact_types"
}
