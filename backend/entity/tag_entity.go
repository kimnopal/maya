package entity

type Tag struct {
	ID    uint64 `gorm:"column:id;primaryKey"`
	Value string `gorm:"column:value"`

	Posts []Post `gorm:"many2many:post_tag;foreignKey:id;joinForeignKey:tag_id;references:id;joinReferences:post_id"`
}

func (m *Tag) TableName() string {
	return "tags"
}
