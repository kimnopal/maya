package entity

type Faculty struct {
	ID        uint64 `gorm:"column:id;primaryKey"`
	Name      string `gorm:"column:username;unique"`
	CreatedAt string `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt string `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *Faculty) TableName() string {
	return "faculties"
}
