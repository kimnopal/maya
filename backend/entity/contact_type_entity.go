package entity

type ContactType struct {
	ID   uint64 `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`

	Contacts []Contact `gorm:"foreignKey:contact_type_id;references:id"`
}

func (c *ContactType) TableName() string {
	return "contact_types"
}
