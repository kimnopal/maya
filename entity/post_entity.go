package entity

type Post struct {
	ID          uint64 `gorm:"column:id;primaryKey"`
	Title       string `gorm:"column:title"`
	Code        string `gorm:"column:code"`
	Description string `gorm:"column:description"`
	CreatedAt   uint64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   uint64 `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (m *Post) TableName() string {
	return "majors"
}
