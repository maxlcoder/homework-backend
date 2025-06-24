package response

import "github.com/jinzhu/copier"

type UserResponse struct {
	BaseResponse
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   uint8  `json:"age"`
}

func ToUserResponse(T any) UserResponse {
	var userResponse UserResponse
	copier.Copy(&userResponse.BaseResponse, T)
	copier.Copy(&userResponse, T)
	return userResponse
}

func FlatToUserResponse(T any) UserResponse {
	var userResponse UserResponse
	copier.Copy(&userResponse.BaseResponse, T)
	copier.Copy(&userResponse, T)
	return userResponse
}
