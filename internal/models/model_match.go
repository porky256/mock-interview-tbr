package models

type Match struct {
	Id int64 `json:"id,omitempty"`

	UserAsker *User `json:"userAsker,omitempty"`

	UserMatch *User `json:"userMatch,omitempty"`

	MatchScore int32 `json:"matchScore,omitempty"`
}
