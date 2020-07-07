package dto

type User struct {
	Email      string `json:"email,omitempty" sql:",pk,unique"`
	Hash       []byte `json:"-"`
	Indenticon []byte `json:"indenticon,omitempty"`
}
