package entity

type Role struct {
	ID   uint64 `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name;unique"`
}

func (m *Role) TableName() string {
	return "roles"
}
