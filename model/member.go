package model

type Member struct {
	Id               string    `gorm:"primaryKey" json:"id"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	Sexe             string    `json:"sexe"`
	BirthDate        string `json:"birthDate"`
	Is_dead          bool      `json:"is_dead"`
	DeathDate        string `json:"deathDate"`
	Baptism          bool      `json:"Baptism"`
	BaptismDate      string`json:"BaptismDate"`
	BaptismRenew     bool      `json:"BaptismRenew"`
	BaptismRenewDate string `json:"BaptismRenewDate"`
	Confession       bool      `json:"Confession"`
	ConfessionDate   string `json:"ConfessionDate"`
	Communion        bool      `json:"Communion"`
	CommunionDate    string `json:"CommunionDate"`
	Confirmation     bool      `json:"Confirmation"`
	ConfirmationDate string `json:"ConfirmationDate"`
	Wedding          bool      `json:"Wedding"`
	Couple           string    `json:"Couple"`
	WeddingDate      string`json:"WeddingDate"`
	Faritra          int       `json:"faritra"`
}
