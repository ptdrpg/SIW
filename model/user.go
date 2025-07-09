package model

type User struct {
	Id                       string         `gorm:"primaryKey" json:"id"`
	Name                     string         `json:"name"`
	Email                    string         `gorm:"unique;not null" json:"email"`
	Image                    string         `json:"image"`
	Password                 string         `json:"password"`
	PasswordResetAttempts    int            `json:"passwordResetAttempts"`
	LastPasswordResetRequest string         `json:"lastPasswordResetRequest"`
	CreatedAt                string         `json:"createdAt"`
	UpdatedAt                string         `json:"updatedAt"`
}

type UserInput struct {
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Password   string   `json:"password"`
	CreatedAt  string   `json:"createdAt"`
	UpdatedAt  string   `json:"updatedAt"`
}

type UpdateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
