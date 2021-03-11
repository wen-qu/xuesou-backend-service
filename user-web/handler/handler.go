package handler

import (
	"context"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"
	"regexp"

	userweb "github.com/wen-qu/xuesou-backend-service/user-web/proto"

	usersrv "github.com/wen-qu/xuesou-backend-service/user-srv/proto"

)

var (
	userClient usersrv.UserSrvService
)

func Init(){
	srv := service.New()
	userClient = usersrv.NewUserSrvService("user-srv", srv.Client())
}

// UserWeb the UserWeb struct
type UserWeb struct{}

// Login login service
func (e *UserWeb) Login(ctx context.Context, req *userweb.UserRequest, rsp *userweb.UserResponse) error {
	if len(req.Tel) == 0 || len(req.ValidationCode) == 0 {
		return errors.BadRequest("para:001", "missing parameter")
	}
	if ok, _ := regexp.Match("s/1[3-9]\\d{9}/g", []byte(req.Tel)); !ok {
		return errors.BadRequest("para:002", "invalid parameter: tel")
	}
	log.Info("Received UserWeb.Login request")

	loginRsp, err := userClient.InspectUser(ctx, &usersrv.InspectRequest{
		Tel: req.Tel,
	})

	if err != nil {
		return errors.InternalServerError("fatal:001", err.Error())
	}

	if len(loginRsp.User.Uid) > 0 { // login success
		// TODO: check login device, generate token, update user_login_inf table and then return 200.
		rsp = &userweb.UserResponse{
			Code: 200,
			Uid: loginRsp.User.Uid,
			Msg: "login success",
		}
	} else {
		if err := e.Register(ctx, req, rsp); err != nil {
			return e.Login(ctx, req, rsp)
		}
	}

	return nil
}

// Register register service
func (e *UserWeb) Register(ctx context.Context, req *userweb.UserRequest, rsp *userweb.UserResponse) error {
	if len(req.Tel) == 0 {
		return errors.BadRequest("para:001", "missing parameter")
	}
	if ok, _ := regexp.Match("s/1[3-9]\\d{9}/g", []byte(req.Tel)); !ok {
		return errors.BadRequest("para:002", "invalid parameter: tel")
	}
	log.Info("Received UserWeb.Register request")

	regRsp, err := userClient.AddUser(ctx, &usersrv.AddRequest{
		User: &usersrv.User{
			Tel: req.Tel,

		},
	})

	if err != nil {
		return errors.InternalServerError("fatal:001", err.Error())
	}



	// TODO: create user's chat table, class table, then insert login_inf into user_login_inf table
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
