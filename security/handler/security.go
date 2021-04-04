package handler

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/store"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/micro/micro/v3/service/auth"
	"github.com/micro/micro/v3/service/auth/jwt"

	security "github.com/wen-qu/xuesou-backend-service/security/proto"
)

var JWTClient auth.Auth

func Init(){
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
	})
}

type Security struct{}

func (sec *Security)GenerateValidation(ctx context.Context, req *security.GenerateValidationRequest, rsp *security.GenerateValidationResponse) error {
	if len(req.Tel) == 0 {
		return errors.BadRequest("security:001", "missing parameters: tel")
	}

	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	// TODO: send the code to the tel
	fmt.Println(code)

	if err := store.Write(&store.Record{Key: req.Tel, Value: []byte(code)}); err != nil {
		return errors.InternalServerError("security.GenerateValidation:fatal:001", err.Error())
	}

	rsp.Status = 200
	rsp.Msg = "success"

	return nil
}

func (sec *Security)CheckValidation(ctx context.Context, req *security.CheckValidationRequest, rsp *security.CheckValidationResponse) error {
	if len(req.Tel) == 0 {
		return errors.BadRequest("security:001", "missing parameters: tel")
	}

	records, err := store.Read(req.Tel)
	if err != nil {
		return errors.InternalServerError("security.CheckValidation:fatal:001", err.Error())
	}

	if req.Code == string(records[0].Value) {
		rsp.Status = 200
		rsp.Msg = "pass"
	} else {
		rsp.Status = 401
		rsp.Msg = "invalid"
	}
	return nil
}

func (sec *Security)GenerateToken(ctx context.Context, req *security.GenerateTokenRequest, rsp *security.GenerateTokenResponse) error {
	if len(req.Type) == 0 || len(req.Name) == 0 {
		return errors.BadRequest("security:001", "missing parameters: type or name")
	}

	if len(req.Secret) == 0 {
		req.Secret = uuid.New().String()
	}
	acc, err := JWTClient.Generate(req.Name, func(o *auth.GenerateOptions) {
		o.Type = req.Type
		o.Name = req.Name
		o.Secret = req.Secret
	})

	if err != nil {
		return errors.InternalServerError("security.GenerateToken:fatal:001", err.Error())
	}

	rsp.Token = acc.Secret
	rsp.Msg = "success"
	rsp.Status = 200

	return nil
}

func (sec *Security)CheckToken(ctx context.Context, req *security.CheckTokenRequest, rsp *security.CheckTokenResponse) error {
	if len(req.Token) == 0 {
		return errors.BadRequest("security:001", "missing parameters: token")
	}

	acc, err := JWTClient.Inspect(req.Token)
	if err == auth.ErrInvalidToken {
		rsp.Status = 401
		rsp.Msg = "invalid token"
	} else if err != nil {
		return errors.InternalServerError("security.CheckToken:fatal:001", err.Error())
	}

	if acc != nil {
		rsp.Status = 200
		rsp.Msg = "pass"
	} else {
		rsp.Status = 500
		rsp.Msg = "unknown error"
	}

	return nil
}
