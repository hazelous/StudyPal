package req

type ProfileReq struct {
	ProfileName 	string `json:"profile_name" validate:"required"`
	ProfileImage 	string `json:"profile_image" validate:"required"`
}