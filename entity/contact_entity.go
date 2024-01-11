package entity

type Contact struct {
	ID            uint64 `gorm:"column:id;primaryKey"`
	Value         string `gorm:"column:value"`
	CreatedAt     uint64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt     uint64 `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	ContactTypeID uint64 `gorm:"column:contact_type_id"`
	UserID        uint64 `gorm:"column:user_id"`

	ContactType ContactType `gorm:"foreignKey:contact_type_id;references:id"`
	User        User        `gorm:"foreignKey:user_id;references:id"`
}

func (c *Contact) TableName() string {
	return "contacts"
}
