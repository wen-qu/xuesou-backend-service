package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	"github.com/wen-qu/xuesou-backend-service/basic/db"

	usersrv "github.com/wen-qu/xuesou-backend-service/user-srv/proto"

)

// UserSrv the UserSrv struct
type UserSrv struct{}

// Login login service
func (e *UserSrv) AddUser(ctx context.Context, req *usersrv.AddRequest, rsp *usersrv.AddResponse) error {
	db.GetDB()
	log.Info("Received UserSrv.Login request")
	rsp.Msg = "Hello AddUser, " + req.User.Tel
	return nil
}

// Register register service
func (e *UserSrv) InspectUser(ctx context.Context, req *usersrv.InspectRequest, rsp *usersrv.InspectResponse) error {
	log.Info("Received UserSrv.Register request")
	rsp.Msg = "Hello InspectUser, " + req.Tel
	return nil
}

// Validation validation service (e.g. get a validation code, etc.)
func (e *UserSrv) UpdateUser(ctx context.Context, req *usersrv.UpdateRequest, rsp *usersrv.UpdateResponse) error {
	log.Info("Received UserSrv.Validation request")
	rsp.Msg = "Hello UpdateUser, " + req.User.Tel
	return nil
}

// UpdateProfile update user's profile
func (e *UserSrv) DeleteUser(ctx context.Context, req *usersrv.DeleteRequest, rsp *usersrv.DeleteResponse) error {
	log.Info("Received UserSrv.Register request")
	rsp.Msg = "Hello DeleteUser, " + req.Tel
	return nil
}
