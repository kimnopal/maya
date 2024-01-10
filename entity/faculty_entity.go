package entity

type Faculty struct {
	ID        uint64 `gorm:"column:id;primaryKey"`
	Name      string `gorm:"column:username;unique"`
	CreatedAt uint64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt uint64 `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (f *Faculty) TableName() string {
	return "faculties"
}
