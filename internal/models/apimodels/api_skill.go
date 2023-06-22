package apimodels

type SkillAPI struct {
	ID          int64  `json:"id,omitempty" validate:"required"`
	Name        string `json:"name,omitempty" validate:"ascii,required"`
	Description string `json:"description,omitempty"`
}
