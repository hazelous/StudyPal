package req

type EditProfileReq struct {
	ProfileID    int    `json:"profile_id" validate:"required"`
	ProfileName  string `json:"profile_name" validate:"required"`
	ProfileImage string `json:"profile_image" validate:"required"`
}
