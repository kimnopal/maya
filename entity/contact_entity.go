package entity

type Contact struct {
	ID        uint64 `gorm:"column:id;primaryKey"`
	Value     string `gorm:"column:value"`
	CreatedAt uint64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt uint64 `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}
