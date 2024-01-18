package entity

type User struct {
	ID        uint64  `gorm:"column:id;primaryKey"`
	Username  string  `gorm:"column:username"`
	Name      string  `gorm:"column:name"`
	Email     string  `gorm:"column:email"`
	Password  string  `gorm:"column:password"`
	CreatedAt uint64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt uint64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	MajorID   *uint64 `gorm:"column:major_id"`
	RoleID    uint64  `gorm:"column:role_id"`

	Major    *Major     `gorm:"foreignKey:major_id;references:id"`
	Role     *Role      `gorm:"foreignKey:role_id;references:id"`
	Posts    []*Post    `gorm:"foreignKey:user_id;references:id"`
	Appliers []*Applier `gorm:"foreignKey:user_id;references:id"`
	Contacts []*Contact `gorm:"foreignKey:user_id;references:id"`
}

func (u *User) TableName() string {
	return "users"
}
