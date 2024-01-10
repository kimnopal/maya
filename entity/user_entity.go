package entity

type User struct {
	ID        uint64 `gorm:"column:id;primaryKey"`
	Username  string `gorm:"column:username;unique"`
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email;unique"`
	Password  string `gorm:"column:password"`
	CreatedAt string `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt string `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *User) TableName() string {
	return "users"
}
