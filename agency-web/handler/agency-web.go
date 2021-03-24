package handler

import (
	"context"
	"github.com/micro/micro/v3/service"
	agencysrv "github.com/wen-qu/xuesou-backend-service/agency-srv/proto"

	// log "github.com/micro/micro/v3/service/logger"

	agencyweb "github.com/wen-qu/xuesou-backend-service/agency-web/proto"
)

type AgencyWeb struct{}

var AgencyClient agencysrv.AgencySrvService

func Init(){
	srv := service.New()
	AgencyClient = agencysrv.NewAgencySrvService("agency-srv", srv.Client())
}

func (agency *AgencyWeb) Login(ctx context.Context, req *agencyweb.LoginRequest, rsp *agencyweb.LoginResponse) error {
	return nil
}

func (agency *AgencyWeb) GetAgencies(ctx context.Context, req *agencyweb.GetAgenciesRequest, rsp *agencyweb.GetAgenciesResponse) error {
	return nil
}

func (agency *AgencyWeb) Search(ctx context.Context, req *agencyweb.SearchRequest, rsp *agencyweb.SearchResponse) error {
	return nil
}

func (agency *AgencyWeb) GetAgencyDetail(ctx context.Context, req *agencyweb.GetAgencyDetailRequest, rsp *agencyweb.GetAgencyDetailResponse) error {
	return nil
}