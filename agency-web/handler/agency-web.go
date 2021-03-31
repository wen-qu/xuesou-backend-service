package handler

import (
	"context"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/errors"
	agencysrv "github.com/wen-qu/xuesou-backend-service/agency-srv/proto"
	classsrv "github.com/wen-qu/xuesou-backend-service/class-srv/proto"
	"reflect"

	// log "github.com/micro/micro/v3/service/logger"

	agencyweb "github.com/wen-qu/xuesou-backend-service/agency-web/proto"
)

type AgencyWeb struct{}

var AgencyClient agencysrv.AgencySrvService
var ClassClient classsrv.ClassSrvService

func Init(){
	srv := service.New()
	AgencyClient = agencysrv.NewAgencySrvService("agency-srv", srv.Client())
	ClassClient = classsrv.NewClassSrvService("class-srv", srv.Client())
}

func (agency *AgencyWeb) Login(ctx context.Context, req *agencyweb.LoginRequest, rsp *agencyweb.LoginResponse) error {

	return nil
}

func (agency *AgencyWeb) Register(ctx context.Context, req *agencyweb.RegisterRequest, rsp *agencyweb.LoginResponse) error {

	return nil
}

func (agency *AgencyWeb) GetAgencies(ctx context.Context, req *agencyweb.GetAgenciesRequest, rsp *agencyweb.GetAgenciesResponse) error {
	if len(req.S) == 0 {
		return errors.BadRequest("para:001", "missing parameters: s")
	}

	rspAgencies, err := AgencyClient.ReadAgencyDetails(ctx, &agencysrv.ReadAgencyRequest{
		S: req.S,
	})
	if err != nil {
		return errors.InternalServerError("agency-web.AgencyWeb.GetAgencies:fatal:001", err.Error())
	}

	if len(rspAgencies.Agencies) == 0 {
		return nil
	}

	for i := 0; i < len(rspAgencies.Agencies); i++ {
		rsp.Agencies = append(rsp.Agencies, new(agencyweb.Agency))
		vSrvAgency := reflect.ValueOf(rspAgencies.Agencies[i])
		vWebAgency := reflect.ValueOf(rsp.Agencies[i])

		for j := 0; j < vSrvAgency.NumField(); j++ {
			switch vSrvAgency.Field(j).Kind() {
			case reflect.String:
				vField := reflect.ValueOf(vWebAgency.Field(j).Pointer())
				vField.Elem().SetString(vSrvAgency.Field(j).Interface().(string))
			case reflect.Int32:
				vField := reflect.ValueOf(vWebAgency.Field(j).Pointer())
				vField.Elem().SetInt(vSrvAgency.Field(j).Interface().(int64))
			}
		}
	}

	rsp.Status = 200
	rsp.Msg = ""
	return nil
}

func (agency *AgencyWeb) Search(ctx context.Context, req *agencyweb.SearchRequest, rsp *agencyweb.SearchResponse) error {
	// TODO: read from users' historical search and read from hottest search
	return nil
}

func (agency *AgencyWeb) GetAgencyDetail(ctx context.Context, req *agencyweb.GetAgencyDetailRequest, rsp *agencyweb.GetAgencyDetailResponse) error {
	if len(req.AgencyID) == 0 {
		return errors.BadRequest("para:001", "missing parameters: agencyID")
	}

	// get general information
	rspAgency, err := AgencyClient.ReadAgencyDetails(ctx, &agencysrv.ReadAgencyRequest{
		AgencyID: req.AgencyID,
	})

	if err != nil {
		return errors.InternalServerError("agency-web.AgencyWeb.GetAgencyDetail:fatal:001", err.Error())
	}

	if len(rspAgency.Agencies) == 0 {
		return nil
	}

	rsp.General = new(agencyweb.Agency)
	vSrvAgency := reflect.ValueOf(rspAgency.Agencies[0])
	vWebAgency := reflect.ValueOf(rsp.General)

	// assign rspAgency.Agencies[0] to rsp.General
	for i := 0; i < vSrvAgency.NumField(); i++ {
		switch vSrvAgency.Field(i).Kind() {
		case reflect.String:
			vField := reflect.ValueOf(vWebAgency.Field(i).Pointer())
			vField.Elem().SetString(vSrvAgency.Field(i).Interface().(string))
		case reflect.Int32:
			vField := reflect.ValueOf(vWebAgency.Field(i).Pointer())
			vField.Elem().SetInt(vSrvAgency.Field(i).Interface().(int64))
		}
	}

	return nil
}