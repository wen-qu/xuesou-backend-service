package handler

import (
	"context"
	"net/http"

	"github.com/micro/micro/v3/service"

	"github.com/gin-gonic/gin"

	proto "github.com/wen-qu/xuesou-backend-service/user-srv/proto"
)

var (
	userClient proto.UserSrvService
)

// Init init service
func Init() {
	srv := service.New()
	userClient = proto.NewUserSrvService("go.micro.user.srv", srv.Client()) // create a new object of UserSrvService, not open up the service.
}


// InitRouter init all routers and handlers
func InitRouter() *gin.Engine {
	r := gin.Default()

	v1User := r.Group("/v1/user/")

	v1User.POST("/login", loginHandler)

	v1User.POST("/register", registerHandler)

	v1User.POST("/profile/update", updateProfileHandler)

	v1User.GET("/profile/{type}", getProfileHandler)

	v1User.GET("/message", getMessageHandler)
	return r
}

func loginHandler(c *gin.Context) {
	var user struct {
		Tel string `json:"tel"`
	}
	c.ShouldBind(&user)
	rsp, err := userClient.Login(context.Background(), &proto.UserRequest{
		Tel: user.Tel,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"uid": rsp.Uid,
			"code": rsp.Code,
			"message": rsp.Msg,
		})
	}
}

func registerHandler(c *gin.Context) {
	var user struct {
		Tel string `json:"tel"`
	}
	c.ShouldBind(&user)
	_, err := userClient.Register(context.Background(), &proto.UserRequest{
		Tel: user.Tel,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		rsp, err := userClient.Login(context.Background(), &proto.UserRequest{
			Tel: user.Tel,
		})
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"uid": rsp.Uid,
				"code": rsp.Code,
				"message": rsp.Msg,
			})
		}
	}
}
