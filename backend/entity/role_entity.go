package entity

type Role struct {
	ID   uint64 `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`

	Users []User `gorm:"foreignKey:role_id;references:id"`
}

func (m *Role) TableName() string {
	return "roles"
}
