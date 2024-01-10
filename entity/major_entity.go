package entity

type Major struct {
	ID        uint64 `gorm:"column:id;primaryKey"`
	Name      string `gorm:"column:name;unique"`
	CreatedAt uint64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt uint64 `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (m *Major) TableName() string {
	return "majors"
}
