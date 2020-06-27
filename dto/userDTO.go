package dto

type User struct {
	Email      string `json:"email,omitempty"`
	Hash       []byte `json:"-"`
	Indenticon []byte `json:"indenticon,omitempty"`
}
