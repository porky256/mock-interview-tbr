package models

type UsersSkill struct {
	Id int64 `json:"id,omitempty"`

	Skill *Skill `json:"skill,omitempty"`

	Score int32 `json:"score,omitempty"`
}
