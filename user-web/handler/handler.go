package handler

import (
	"context"
	"github.com/google/uuid"

	"io/ioutil"
	"os"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/auth"
	"github.com/micro/micro/v3/service/auth/jwt"
	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"

	"regexp"

	userweb "github.com/wen-qu/xuesou-backend-service/user-web/proto"
	"github.com/wen-qu/xuesou-backend-service/user-web/utils"

	usersrv "github.com/wen-qu/xuesou-backend-service/user-srv/proto"
)

var (
	UserClient usersrv.UserSrvService
	JWTClient auth.Auth
)

func Init(){
	srv := service.New()
	UserClient = usersrv.NewUserSrvService("user-srv", srv.Client())
	JWTClient = jwt.NewAuth()

	JWTClient.Init(func(o *auth.Options) {
		privateFile, err := os.Open("/home/micro/.ssh/id_rsa_micro")
		if err != nil {
			log.Error(err)
			return
		}
		publicFile, err := os.Open("/home/micro/.ssh/id_rsa_micro.pub")
		if err != nil {
			log.Error(err)
			return
		}
		defer privateFile.Close()
		defer publicFile.Close()

		privateKeyContent, _ := ioutil.ReadAll(privateFile)
		publicKeyContent, _ := ioutil.ReadAll(publicFile)

		o.PrivateKey = string(privateKeyContent)
		o.PublicKey = string(publicKeyContent)
		o.LoginURL = "/login"
	})
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
	log.Info("Received UserWeb.Login request")

	// TODO: check the validation code [service: security]
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

		acc, err := JWTClient.Generate(loginRsp.User.Uid, func(o *auth.GenerateOptions) {
			o.Type = "user"
			o.Name = loginRsp.User.Uid
			o.Secret = uuid.New().String()
		})

		if err != nil {
			return errors.InternalServerError("user-web.UserWeb.Login:fatal:002", err.Error())
		}

		// if err := security.UpdateLoginInformation(loginRsp.User.Uid); err != nil {
		//     return errors.InternalServerError("fatal:001", "cannot update login information")
		// }

		rsp.Status = 200
		rsp.Uid = loginRsp.User.Uid
		rsp.Msg = "login success"
		rsp.Token = acc.Secret

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

	if len(req.Uid) == 0 && len(req.Tel) == 0 {
		return errors.BadRequest("para:001", "missing parameters")
	}

	switch req.InformationType {
	case 1: // general
		rsp.Profile, _ = utils.ReadGeneralProfile(ctx, &UserClient, req.Uid, req.Tel)
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
