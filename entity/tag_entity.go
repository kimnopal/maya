package entity

type Tag struct {
	ID    uint64 `gorm:"column:id;primaryKey"`
	Value string `gorm:"column:value"`
}

func (m *Tag) TableName() string {
	return "tags"
}
