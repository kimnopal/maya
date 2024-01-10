package entity

type User struct {
	ID        uint64 `gorm:"column:id;primaryKey"`
	Username  string `gorm:"column:username"`
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	CreatedAt uint64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt uint64 `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *User) TableName() string {
	return "users"
}
