package converter

import (
	"time"

	"github.com/qara-qurt/telegram_plus/user_service/internal/model"
	desk "github.com/qara-qurt/telegram_plus/user_service/pkg/gen/user"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromService(user *desk.CreateUserRequest) model.UserInfo {
	// Convert timestamp.Timestamp to time.Time
	birthday := time.Unix(user.BirthdayDate.GetSeconds(), int64(user.BirthdayDate.GetNanos()))

	return model.UserInfo{
		Username:     user.GetUsername(),
		Login:        user.GetLogin(),
		BirthDayDate: birthday,
		Email:        user.GetEmail(),
		Password:     user.GetPassword(),
	}
}

func ToServiceFromUser(user model.User) *desk.User {
	return &desk.User{
		Uuid:         user.ID,
		Username:     user.Username,
		Login:        user.Login,
		BirthdayDate: timestamppb.New(user.BirthDayDate),
		Email:        user.Email,
		Img:          user.Img,
		CreatedAt:    timestamppb.New(user.CreatedAt),
		UpdatedAt:    timestamppb.New(user.UpdatedAt),
	}
}

func ToServiceFromUsers(users []model.User) *desk.GetUsersResponse {
	var usersConverted []*desk.User
	for _, user := range users {
		user := ToServiceFromUser(user)
		usersConverted = append(usersConverted, user)
	}

	return &desk.GetUsersResponse{
		Users: usersConverted,
	}
}
