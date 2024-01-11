package entity

type Applier struct {
	ID        uint64 `gorm:"column:id;primaryKey"`
	CreatedAt uint64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt uint64 `gorm:"column:updated_at;autoUpdateTime:milli"`
	UserID    uint64 `gorm:"column:user_id"`
	PostID    uint64 `gorm:"column:post_id"`
	User      User   `gorm:"foreignKey:user_id;references:id"`
	Post      Post   `gorm:"foreignKey:post_id;references:id"`
}

func (a *Applier) TableName() string {
	return "appliers"
}
