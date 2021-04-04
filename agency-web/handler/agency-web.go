package handler

import (
	"context"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/errors"
	agencysrv "github.com/wen-qu/xuesou-backend-service/agency-srv/proto"
	classsrv "github.com/wen-qu/xuesou-backend-service/class-srv/proto"

	"github.com/jinzhu/copier"
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
		if err := copier.Copy(rsp.Agencies[i], rspAgencies.Agencies[i]); err != nil {
			return errors.InternalServerError("agency-web.AgencyWeb.GetAgencies:fatal:002", err.Error())
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

	if err := copier.Copy(rsp.General, rspAgency.Agencies[0]); err != nil {
		return errors.InternalServerError("agency-web.AgencyWeb.GetAgencyDetail:fatal:002", err.Error())
	}

	rsp.BrandStory = rspAgency.BrandHistory
	rsp.Characteristics = rspAgency.Characteristics

	// get classes information
	rspClass, err := ClassClient.ReadClassesByAgencyID(ctx, &classsrv.ReadClassRequest{
		AgencyID: rsp.General.AgencyID,
	})
	if err != nil {
		return errors.InternalServerError("agency-web.AgencyWeb.GetAgencyDetail:fatal:003", err.Error())
	}

	if err := copier.Copy(rsp.General.Classes, rspClass.Classes); err != nil {
		return errors.InternalServerError("agency-web.AgencyWeb.GetAgencyDetail:fatal:004", err.Error())
	}

	// TODO: read teachers, evaluations and nearby agencies information.
	return nil
}

func (agency *AgencyWeb)UpdateAgencyProfile(ctx context.Context, req *agencyweb.UpdateAgencyRequest, rsp *agencyweb.UpdateAgencyResponse) error {
	if len(req.General.AgencyID) == 0 {
		return errors.BadRequest("para:001", "missing parameters: General.agencyID")
	}
	var updateAgency *agencysrv.Agency
	if err := copier.Copy(&updateAgency, req.General); err != nil {
		return errors.InternalServerError("agency-web.AgencyWeb.UpdateAgencyProfile:fatal:001", err.Error())
	}
	rspAgency, err := AgencyClient.UpdateAgency(ctx, &agencysrv.UpdateAgencyRequest{
		AgencyID: req.General.AgencyID,
		Agency:  updateAgency,
	})

	if err != nil {
		return errors.InternalServerError("agency-web.AgencyWeb.UpdateAgencyProfile:fatal:002", err.Error())
	}

	if rspAgency.Status == 200 {
		rsp.Msg = "success"
		rsp.Status = 200
		return nil
	}

	return nil
}

func (agency *AgencyWeb)GetEvaluation(ctx context.Context, req *agencyweb.GetEvaluationRequest, rsp *agencyweb.GetEvaluationResponse) error {
	if len(req.AgencyID) == 0 {
		return errors.BadRequest("para:001", "missing parameters: agencyID")
	}


	return nil
}

func (agency *AgencyWeb)GetNearbyAgencies(ctx context.Context, req *agencyweb.GetNearbyAgenciesRequest, rsp *agencyweb.GetNearbyAgenciesResponse) error {
	return nil
}