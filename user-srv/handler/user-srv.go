package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	usersrv "user-srv/proto"

)

// UserSrv the UserSrv struct
type UserSrv struct{}

// Login login service
func (e *UserSrv) Login(ctx context.Context, req *usersrv.UserRequest, rsp *usersrv.UserResponse) error {
	log.Info("Received UserSrv.Login request")
	rsp.Msg = "Hello Login, " + req.Tel
	return nil
}

// Register register service
func (e *UserSrv) Register(ctx context.Context, req *usersrv.UserRequest, rsp *usersrv.UserResponse) error {
	log.Info("Received UserSrv.Register request")
	rsp.Msg = "Hello Register, " + req.Tel
	return nil
}

// Validation validation service (e.g. get a validation code, etc.)
func (e *UserSrv) Validation(ctx context.Context, req *usersrv.UserRequest, rsp *usersrv.UserResponse) error {
	log.Info("Received UserSrv.Validation request")
	rsp.Msg = "Hello Validation, " + req.Tel
	return nil
}

// UpdateProfile update user's profile
func (e *UserSrv) UpdateProfile(ctx context.Context, req *usersrv.UserProfileRequest, rsp *usersrv.UserProfileResponse) error {
	log.Info("Received UserSrv.Register request")
	rsp.Msg = "Hello Register, " + req.Tel
	return nil
}

// ReadProfile get user's profile
func (e *UserSrv) ReadProfile(ctx context.Context, req *usersrv.UserProfileRequest, rsp *usersrv.UserProfileResponse) error {
	log.Info("Received UserSrv.Register request")
	rsp.Msg = "Hello Register, " + req.Tel
	return nil
}