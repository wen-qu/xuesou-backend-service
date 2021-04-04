package utils

import (
	"context"
	"github.com/micro/micro/v3/service/errors"
	usersrv "github.com/wen-qu/xuesou-backend-service/user-srv/proto"
	userweb "github.com/wen-qu/xuesou-backend-service/user-web/proto"
)

func ReadGeneralProfile(ctx context.Context, userClient *usersrv.UserSrvService, uid string, tel string) (*userweb.Profile, error) {
	var profile *usersrv.InspectResponse
	var res = new(userweb.Profile)

	profile, err := (*userClient).InspectUser(ctx, &usersrv.InspectRequest{
		Uid: uid,
		Tel: tel,
	})

	if err != nil {
		return nil, errors.InternalServerError("user-web.UserWeb.ReadProfile.ReadGeneralProfile:fatal:001", err.Error())
	}

	if profile == nil {
		return nil, nil
	}

	res.Username = profile.User.Username
	res.Tel = profile.User.Tel
	res.Img = profile.User.Img
	res.Address = profile.User.Address
	res.Age = profile.User.Age
	res.Email = profile.User.Email
	res.Sex = profile.User.Sex
	res.ClassNum = profile.User.ClassNum

	return res, nil
}

func UpdateGeneralProfile(ctx context.Context, userClient *usersrv.UserSrvService, uid string, profile *userweb.Profile) (ok bool, err error) {
	var currProfile *usersrv.InspectResponse
	currProfile, err = (*userClient).InspectUser(ctx, &usersrv.InspectRequest{Uid: uid})

	if currProfile == nil {
		return false, nil
	}

	if err != nil {
		return false, errors.InternalServerError("user-web.UserWeb.UpdateProfile.UpdateGeneralProfile:fatal:001", err.Error())
	}

	if _, err := (*userClient).UpdateUser(ctx, &usersrv.UpdateRequest{
		User: &usersrv.User{
			Uid:      uid,
			Username: profile.Username,
			Tel:      profile.Tel,
			Email:    profile.Email,
			Sex:      profile.Sex,
			Age:      profile.Age,
			Address:  profile.Address,
			Img:      profile.Img,
		},
	}); err != nil {
		return false, errors.InternalServerError("user-web.UserWeb.UpdateProfile.UpdateGeneralProfile:fatal:002", err.Error())
	}

	return true, err
}

func ReadOrder(ctx context.Context, userClient *usersrv.UserSrvService, uid string, tel string) *userweb.Profile {
	return nil
}

func ReadDisCount(ctx context.Context, userClient *usersrv.UserSrvService, uid string, tel string) *userweb.Profile {
	return nil
}

func ReadLikes(ctx context.Context, userClient *usersrv.UserSrvService, uid string, tel string) *userweb.Profile {
	return nil
}

func ReadOrderReview(ctx context.Context, userClient *usersrv.UserSrvService, uid string, tel string) *userweb.Profile {
	return nil
}

func ReadClasses(ctx context.Context, userClient *usersrv.UserSrvService, uid string, tel string) *userweb.Profile {
	return nil
}

func ReadCollections(ctx context.Context, uid string, tel string) *userweb.Profile {
	return nil
}