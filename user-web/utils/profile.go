package utils

import (
	"context"
	"github.com/micro/micro/v3/service/errors"
	usersrv "github.com/wen-qu/xuesou-backend-service/user-srv/proto"
	"github.com/wen-qu/xuesou-backend-service/user-web/handler"
	userweb "github.com/wen-qu/xuesou-backend-service/user-web/proto"
)

func ReadGeneralProfile(ctx context.Context, uid string, tel string) (*userweb.Profile, error) {
	var profile *usersrv.InspectResponse
	var res = new(userweb.Profile)

	profile, err := handler.UserClient.InspectUser(ctx, &usersrv.InspectRequest{
		Uid: uid,
		Tel: tel,
	})

	if err != nil {
		return nil, errors.InternalServerError("user-web.UserWeb.ReadProfile:fatal:001", err.Error())
	}

	if profile == nil {
		return nil, errors.NotFound("user:002", "user not found")
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

func ReadOrder(ctx context.Context, uid string, tel string) *userweb.Profile {
	return nil
}

func ReadDisCount(ctx context.Context, uid string, tel string) *userweb.Profile {
	return nil
}

func ReadLikes(ctx context.Context, uid string, tel string) *userweb.Profile {
	return nil
}

func ReadOrderReview(ctx context.Context, uid string, tel string) *userweb.Profile {
	return nil
}

func ReadClasses(ctx context.Context, uid string, tel string) *userweb.Profile {
	return nil
}

func ReadCollections(ctx context.Context, uid string, tel string) *userweb.Profile {
	return nil
}