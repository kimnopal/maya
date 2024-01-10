package entity

type PostCategory struct {
	ID   uint64 `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`

	Posts []Post `gorm:"foreignKey:post_id;references:id"`
}

func (m *PostCategory) TableName() string {
	return "post_categories"
}
