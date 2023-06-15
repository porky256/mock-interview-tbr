package models

type RequestToMatch struct {
	Id int64 `json:"id,omitempty"`

	Skill *User `json:"skill,omitempty"`

	RequiredSkills []Skill `json:"requiredSkills,omitempty"`
}
