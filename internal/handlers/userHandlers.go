package handlers

import (
	userService2 "Poehali/internal/userService"
	"Poehali/internal/web/users"
	"context"
)

type UserHandler struct {
	Service *userService2.UserService
}

func NewUserHandler(service *userService2.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (u *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.Service.GetUser()
	if err != nil {
		return nil, err
	}
	response := users.GetUsers200JSONResponse{}

	for _, user := range allUsers {
		user := users.User{
			Id:       &user.ID,
			Email:    &user.Email,
			Password: &user.Password,
		}
		response = append(response, user)
	}
	return response, nil
}

func (u *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {

	userRequest := request.Body

	userToCreate := userService2.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := u.Service.CreateUser(userToCreate)

	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	return response, nil
}

func (u *UserHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {

	id := request.Id

	updatedUser := userService2.User{}

	user, err := u.Service.UpdateUserByID(uint(id), updatedUser)
	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		Id:       &user.ID,
		Email:    &user.Email,
		Password: &user.Password,
	}
	return response, nil
}

func (u *UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id

	if err := u.Service.DeleteUserByID(uint(id)); err != nil {
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil
}
