package entity

type Applier struct {
	ID        uint64 `gorm:"column:id;primaryKey"`
	CreatedAt uint64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt uint64 `gorm:"column:updated_at;autoUpdateTime:milli"`
}

func (a *Applier) TableName() string {
	return "appliers"
}
