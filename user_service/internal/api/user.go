package user

import (
	"context"
	"fmt"
	"regexp"

	"github.com/qara-qurt/telegrum_plus/user_service/internal/converter"
	"github.com/qara-qurt/telegrum_plus/user_service/internal/service"
	desk "github.com/qara-qurt/telegrum_plus/user_service/pkg/gen/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Implementation struct {
	desk.UnimplementedUserServiceServer
	userService service.IUser
}

func New(service service.IUser) *Implementation {
	return &Implementation{
		userService: service,
	}
}

// ### Create User
func (i *Implementation) CreateUser(ctx context.Context, in *desk.CreateUserRequest) (*desk.CreateUserResponse, error) {
	if in.Username == "" && len(in.Username) <= 2 {
		return nil, status.Errorf(codes.InvalidArgument, "Missing required field - username")
	}
	if in.Login == "" && len(in.Username) <= 2 {
		return nil, status.Errorf(codes.InvalidArgument, "Missing required field - login")
	}
	if in.BirthdayDate == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Missing required field - birthday_date")
	}
	if in.Email == "" || !isValidEmail(in.Email) {
		return nil, status.Errorf(codes.InvalidArgument, "Missing required field or invalid - email ")
	}
	if in.Password == "" && len(in.Password) <= 5 {
		return nil, status.Errorf(codes.InvalidArgument, "Missing required field - password")
	}

	user := converter.ToUserFromService(in)
	uuid, err := i.userService.Create(user)
	if err != nil {
		return nil, err
	}

	return &desk.CreateUserResponse{
		Uuid: uuid,
	}, nil
}

// ### Get User by ID
func (i *Implementation) GetUser(ctx context.Context, in *desk.GetUserRequest) (*desk.User, error) {
	if in.Uuid == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Missing required field - uuid")
	}

	user, err := i.userService.GetUser(in.Uuid)
	if err != nil {
		return nil, err
	}

	userResp := converter.ToServiceFromUser(user)

	return userResp, nil
}

// ### Get users with arguments page and limit
func (i *Implementation) GetUsers(ctx context.Context, in *desk.GetUsersRequests) (*desk.GetUsersResponse, error) {
	limit := in.Limit
	page := in.Page

	if limit == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Missing required field - limit")
	}
	if page == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Missing required field - page")
	}

	users, err := i.userService.GetUsers(limit, page)
	if err != nil {
		return nil, err
	}

	userResp := converter.ToServiceFromUsers(users)

	return userResp, nil
}

// ### Get User by credentials like login, email, username
func (i *Implementation) GetUsersByCredential(ctx context.Context, in *desk.GetUsersByCredentialRequests) (*desk.GetUsersByCredentialResponse, error) {

	if in.Text == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing argiment - text")
	}
	users, err := i.userService.GetUsersByCredentials(in.Text)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to fetch users")
	}

	response := converter.ToServiceFromUsers(users)
	return (*desk.GetUsersByCredentialResponse)(response), nil
}

func (i *Implementation) DeleteUser(ctx context.Context, in *desk.DeleteUserRequest) (*desk.Empty, error) {
	fmt.Println("DEL")
	return nil, nil
}

func (i *Implementation) UpdateUser(ctx context.Context, in *desk.UpdateUserRequest) (*desk.User, error) {
	fmt.Println("Update")
	return nil, nil
}

// func (i *Implementation) mustEmbedUnimplementedUserServiceServer() {}

// utils
func isValidEmail(email string) bool {
	// Regular expression pattern for validating email addresses
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return emailRegex.MatchString(email)
}
