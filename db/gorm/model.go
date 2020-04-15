package gorm

// Model 用来替代 gorm.Model，时间字段的类型均为 Time，而非 time.Time。
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt Time
	UpdatedAt Time
	DeletedAt *Time `sql:"index"`
}
