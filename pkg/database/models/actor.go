package models

type Actor struct {
	Id 			string		`json:"id,omitempty"`
	Name 		string		`json:"name:"`
	YearOfDebut string		`json:"yearOfDebut"`
	CreatedAt 	int64		`json:"createdAt"`
}