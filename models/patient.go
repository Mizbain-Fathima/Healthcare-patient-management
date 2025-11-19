package models

type Patient struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Diagnosis   string `json:"diagnosis"`
	PhoneNumber string `json:"phone_number"`
}
