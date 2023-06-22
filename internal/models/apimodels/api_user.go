package apimodels

type UserAPI struct {
	ID          int64           `json:"id,omitempty" validate:"required"`
	Username    string          `json:"username,omitempty" validate:"min=3,ascii,required"`
	FirstName   string          `json:"firstName,omitempty" validate:"alpha,required"`
	LastName    string          `json:"lastName,omitempty" validate:"alpha,required"`
	Email       string          `json:"email,omitempty" validate:"email,required"`
	Password    string          `json:"password,omitempty" validate:"ascii,required"`
	Phone       string          `json:"phone,omitempty"`
	UserStatus  UserStatus      `json:"userStatus,omitempty" validate:"required,user_status_custom_validation"`
	Description string          `json:"description,omitempty"`
	UsersSkills []UsersSkillAPI `json:"usersSkills,omitempty"`
}

type UsersSkillAPI struct {
	ID    int64     `json:"id,omitempty" validate:"required"`
	Skill *SkillAPI `json:"skill,omitempty" validate:"required"`
	Score int32     `json:"score,omitempty" validate:"required"`
}

type UserStatus int

const (
	UserSigned UserStatus = iota + 1
	UserUnsigned
)

func (s UserStatus) String() string {
	statuses := []string{"signed", "unsigned"}
	if s < UserSigned || s > UserUnsigned {
		return "undefined"
	}
	return statuses[s]
}
