package models

type User struct {
	ID      int         `json:"id"`
	Name    string      `json:"name"`
	Age     int         `json:"age"`
	Contact ContactInfo `json:"contactInfo"`
	Address []Address   `json:"address"`
}

type ContactInfo struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Address struct {
	Street  string `json:"street"`
	Zipcode string `json:"zipcode"`
	State   string `json:"state"`
}
