package handler

import (
	"context"
	"github.com/google/uuid"
	security "github.com/wen-qu/xuesou-backend-service/security/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"

	"regexp"

	userweb "github.com/wen-qu/xuesou-backend-service/user-web/proto"
	"github.com/wen-qu/xuesou-backend-service/user-web/utils"

	usersrv "github.com/wen-qu/xuesou-backend-service/user-srv/proto"
)

var (
	UserClient usersrv.UserSrvService
	SecClient security.SecurityService
)

func Init(){
	srv := service.New()
	UserClient = usersrv.NewUserSrvService("user-srv", srv.Client())
	SecClient = security.NewSecurityService("security", srv.Client())
}

// UserWeb the UserWeb struct
type UserWeb struct{}

// Login login service
func (e *UserWeb) Login(ctx context.Context, req *userweb.UserRequest, rsp *userweb.UserResponse) error {
	if len(req.Tel) == 0 || len(req.ValidationCode) == 0 {
		return errors.BadRequest("para:001", "missing parameter")
	}
	if ok, _ := regexp.Match("^1[3-9]\\d{9}$", []byte(req.Tel)); !ok {
		return errors.BadRequest("para:002", "invalid parameter: tel")
	}

	// TODO: check the validation code [service: security]
	rspCheck, err := SecClient.CheckValidation(ctx, &security.CheckValidationRequest{
		Code: req.ValidationCode,
		Tel:  req.Tel,
	})
	if err != nil {
		return err
	}
	if rspCheck.Status == 401 {
		rsp.Status = 401
		rsp.Msg = "invalid validation code"
	}

	loginRsp, err := UserClient.InspectUser(ctx, &usersrv.InspectRequest{
		Tel: req.Tel,
	})

	if err != nil {
		return errors.InternalServerError("user-web.UserWeb.Login:fatal:001", err.Error())
	}

	if len(loginRsp.User.Uid) > 0 { // login success
		// TODO: check login device, generate token, update user_login_inf table and then return 200. [service: security]

		// if err := security.CheckLoginDevice(loginRsp.User.Uid); err != nil {
		//     return errors.Forbidden("auth:003", "unknown device")
		// }


		token, err := SecClient.GenerateToken(ctx, &security.GenerateTokenRequest{
			Type:   "user",
			Name:   loginRsp.User.Uid,
			Secret: uuid.New().String(),
		})

		if err != nil {
			return errors.InternalServerError("user-web.UserWeb.Login:fatal:002", err.Error())
		}

		// if err := security.UpdateLoginInformation(loginRsp.User.Uid); err != nil {
		//     return errors.InternalServerError("fatal:001", "cannot update login information")
		// }

		rsp.Status = 200
		rsp.Uid = loginRsp.User.Uid
		rsp.Msg = "success"
		rsp.Token = token.Token

		return nil
	}

	if err := e.Register(ctx, req, rsp); err != nil {
		return errors.InternalServerError("user-web.UserWeb.Register:fatal:001", err.Error())
	}

	return e.Login(ctx, req, rsp)
}

// Register register service
func (e *UserWeb) Register(ctx context.Context, req *userweb.UserRequest, rsp *userweb.UserResponse) error {
	if len(req.Tel) == 0 || len(req.ValidationCode) == 0 {
		return errors.BadRequest("para:001", "missing parameter")
	}
	if ok, _ := regexp.Match("s/1[3-9]\\d{9}/g", []byte(req.Tel)); !ok {
		return errors.BadRequest("para:002", "invalid parameter: tel")
	}
	log.Info("Received UserWeb.Register request")

	// TODO: check the validation code.
	rspCheck, err := SecClient.CheckValidation(ctx, &security.CheckValidationRequest{
		Code: req.ValidationCode,
		Tel:  req.Tel,
	})
	if err != nil {
		return err
	}
	if rspCheck.Status == 401 {
		rsp.Status = 401
		rsp.Msg = "invalid validation code"
	}

	regRsp, err := UserClient.AddUser(ctx, &usersrv.AddRequest{
		User: &usersrv.User{
			Tel: req.Tel,
		},
	})

	if err != nil {
		return errors.InternalServerError("user-web.UserWeb.Register:fatal:001", err.Error())
	}

	if regRsp.Status == 400 {
		return errors.Forbidden("user:001", "registered")
	}

	rsp = &userweb.UserResponse{
		Status: 200,
		Msg: "register success",
	}
	return nil
}

// UpdateProfile update user's profile
func (e *UserWeb) UpdateProfile(ctx context.Context, req *userweb.UpdateProfileRequest, rsp *userweb.UpdateProfileResponse) error {
	log.Info("Received UserWeb.Register request")
	if len(req.Uid) == 0 {
		return errors.BadRequest("para:001", "missing parameters")
	}

	switch req.InformationType {
	case 1: // general
		ok, err := utils.UpdateGeneralProfile(ctx, &UserClient, req.Uid, req.Profile)
		if err != nil {
			return err
		}
		if !ok {
			return errors.Forbidden("user:001", "user not existed")
		}
		rsp.Type = "general"
		rsp.Status = 200
		rsp.Msg = ""
	case 2: // order
	case 3: // discount
	case 4: // likes
	case 5: // order_review
	case 6: // classes
	case 7: // collections
	}
	return nil
}

// ReadProfile get user's profile
func (e *UserWeb) ReadProfile(ctx context.Context, req *userweb.ReadProfileRequest, rsp *userweb.ReadProfileResponse) error {
	log.Info("Received UserWeb.Register request")

	if len(req.Uid) == 0 && len(req.Tel) == 0 {
		return errors.BadRequest("para:001", "missing parameters")
	}

	switch req.InformationType {
	case 1: // general
		var err error
		rsp.Profile, err = utils.ReadGeneralProfile(ctx, &UserClient, req.Uid, req.Tel)
		if err != nil {
			return err
		}
		rsp.Uid = req.Uid
		rsp.Type = "general"
	case 2: // order
	case 3: // discount
	case 4: // likes
	case 5: // order_review
	case 6: // classes
	case 7: // collections
	}

	return nil
}
