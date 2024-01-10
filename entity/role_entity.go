package entity

type Role struct {
	ID   uint64 `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}

func (m *Role) TableName() string {
	return "roles"
}

// migrate -database "mysql://root@tcp(localhost:3306)/maya" -path db/migrations up
