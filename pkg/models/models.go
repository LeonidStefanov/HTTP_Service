package models

type User struct {
	ID      int      `json:"id,omitempty" bson:"id"`
	Name    string   `json:"name,omitempty" bson:"name"`
	Age     int      `json:"age,omitempty" bson:"age"`
	Friends []string `json:"friends" bson:"friends"`
}

type Users []User

type MakeFfriends struct {
	SourceID int `json:"source_id,omitempty"`
	TargetID int `json:"target_id,omitempty"`
}

type DeleteUser struct {
	TargetID int `json:"target_id"`
}
type ChangeAge struct {
	NewAge int `json:"new_age"`
}

type ErrorRecponse struct {
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}

type Response struct {
	Status string `json:"status,omitempty"`
	Info   string `json:"info,omitempty"`
}
