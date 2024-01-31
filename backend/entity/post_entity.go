package entity

type Post struct {
	ID             uint64 `gorm:"column:id;primaryKey"`
	Title          string `gorm:"column:title"`
	Code           string `gorm:"column:code"`
	Description    string `gorm:"column:description"`
	CreatedAt      uint64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt      uint64 `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	UserID         uint64 `gorm:"column:user_id"`
	PostCategoryID uint64 `gorm:"column:post_category_id"`

	User         User         `gorm:"foreignKey:user_id;references:id"`
	PostCategory PostCategory `gorm:"foreignKey:post_category_id;references:id"`
	Tags         []Tag        `gorm:"many2many:post_tag;foreignKey:id;joinForeignKey:post_id;references:id;joinReferences:tag_id"`
	Appliers     []Applier    `gorm:"foreignKey:post_id;references:id"`
}

func (m *Post) TableName() string {
	return "posts"
}
