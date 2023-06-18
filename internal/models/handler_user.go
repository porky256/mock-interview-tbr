package models

type UserHandler struct {
	ID          int64               `json:"id,omitempty"`
	Username    string              `json:"username,omitempty"`
	FirstName   string              `json:"firstName,omitempty"`
	LastName    string              `json:"lastName,omitempty"`
	Email       string              `json:"email,omitempty"`
	Password    string              `json:"password,omitempty"`
	Phone       string              `json:"phone,omitempty"`
	UserStatus  int32               `json:"userStatus,omitempty"`
	Description string              `json:"description,omitempty"`
	UsersSkills []UsersSkillHandler `json:"usersSkills,omitempty"`
}

type UsersSkillHandler struct {
	ID    int64         `json:"id,omitempty"`
	Skill *SkillHandler `json:"skill,omitempty"`
	Score int32         `json:"score,omitempty"`
}
