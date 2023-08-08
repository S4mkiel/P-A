package entity

type User struct {
	ID        int    `json:"id" gorm:"primaryKey;not null;autoIncrement"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Base      `gorm:"embedded"`
}
