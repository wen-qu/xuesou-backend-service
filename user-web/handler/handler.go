package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	userweb "github.com/wen-qu/xuesou-backend-service/user-web/proto"

)

// UserWeb the UserWeb struct
type UserWeb struct{}

// Login login service
func (e *UserWeb) Login(ctx context.Context, req *userweb.UserRequest, rsp *userweb.UserResponse) error {

	log.Info("Received UserWeb.Login request")
	rsp.Msg = "Hello Login, " + req.Tel
	return nil
}

// Register register service
func (e *UserWeb) Register(ctx context.Context, req *userweb.UserRequest, rsp *userweb.UserResponse) error {
	log.Info("Received UserWeb.Register request")
	rsp.Msg = "Hello Register, " + req.Tel
	return nil
}

// Validation validation service (e.g. get a validation code, etc.)
func (e *UserWeb) Validation(ctx context.Context, req *userweb.UserRequest, rsp *userweb.UserResponse) error {
	log.Info("Received UserWeb.Validation request")
	rsp.Msg = "Hello Validation, " + req.Tel
	return nil
}

// UpdateProfile update user's profile
func (e *UserWeb) UpdateProfile(ctx context.Context, req *userweb.UserProfileRequest, rsp *userweb.UserProfileResponse) error {
	log.Info("Received UserWeb.Register request")
	rsp.Msg = "Hello Register, " + req.Tel
	return nil
}

// ReadProfile get user's profile
func (e *UserWeb) ReadProfile(ctx context.Context, req *userweb.UserProfileRequest, rsp *userweb.UserProfileResponse) error {
	log.Info("Received UserWeb.Register request")
	rsp.Msg = "Hello Register, " + req.Tel
	return nil
}
