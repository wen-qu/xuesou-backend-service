package handler

import (
	"context"
	"database/sql"
	"github.com/jinzhu/copier"
	"github.com/micro/micro/v3/service/errors"
	agencysrv "github.com/wen-qu/xuesou-backend-service/agency-srv/proto"
	"github.com/wen-qu/xuesou-backend-service/basic/db"
	"strings"
)

type AgencySrv struct{}

func (agency *AgencySrv)ReadAgencyDetails(ctx context.Context, req *agencysrv.ReadAgencyRequest, rsp *agencysrv.ReadAgencyResponse) error {
	// read by AgencyID, name, search case or filter items

	if len(req.AgencyID) == 0 && len(req.Name) == 0 && len(req.Tags) == 0 && len(req.FilterItems) == 0 && len(req.S) == 0 {
		return errors.BadRequest("para:001", "missing parameters")
	}

	// Cautions: temporarily ignored req.FilterItems and req.Tags

	// accurate query
	if len(req.AgencyID) > 0 {
		var agency = new(agencysrv.Agency)
		var tagString string
		var characteristics string
		var brandHistory string
		if err := db.GetDB().QueryRow("select agency_id, name, tel, rating, comments, order, tags, address, " +
				"address_detail, icon, photos, brand_history, characteristic from agency_profile_table where agency_id = ?",
				req.AgencyID).Scan(
				&agency.AgencyID, &agency.Name, &agency.Tel, &agency.Rating, &agency.Comments,
				&agency.Order, &tagString, &agency.Address, &agency.AddressDetail, &agency.Icon,
				&agency.Photos, &brandHistory, &characteristics); err != nil {
			if err == sql.ErrNoRows {
				return nil
			} else {
				return errors.InternalServerError("agency-srv.AgencySrv.ReadAgencyDetails:fatal:001", err.Error())
			}
		}

		agency.Tags = strings.Split(tagString, ",") // separate the tags string to array

		rsp.Status = 200
		rsp.Agencies = append(rsp.Agencies, agency)
		rsp.BrandHistory = brandHistory
		rsp.Characteristics = strings.Split(characteristics, ",")
		rsp.Msg = ""

		return nil
	}

	if len(req.S) > 0 { // search agencies from search page, so needn't to read brandHistory, characteristics.
		var agency = new(agencysrv.Agency)
		var tagString string

		rows, err := db.GetDB().Query("select agency_id, name, tel, rating, comments, order, tags, address, " +
			"address_detail, icon, photos from agency_profile_table where name regexp ? ", req.S)
		if err == sql.ErrNoRows {
			return nil
		} else if err != nil {
			return errors.InternalServerError("agency-srv.AgencySrv.ReadAgencyDetails:fatal:002", err.Error())
		}

		for rows.Next() {
			err := rows.Scan(&agency.AgencyID, &agency.Name, &agency.Tel, &agency.Rating, &agency.Comments,
				&agency.Order, &tagString, &agency.Address, &agency.AddressDetail, &agency.Icon,
				&agency.Photos)
			if err != nil {
				return errors.InternalServerError("agency-srv.AgencySrv.ReadAgencyDetails:fatal:003", err.Error())
			}

			agency.Tags = strings.Split(tagString, ",")
			rsp.Agencies = append(rsp.Agencies, agency)
		}

		rsp.Msg = ""
		rsp.Status = 200

		return nil
	}

	return nil
}

func (agency *AgencySrv)AddAgency(ctx context.Context, req *agencysrv.AddAgencyRequest, rsp *agencysrv.AddAgencyResponse) error {
	if len(req.Agency.Name) == 0 {
		return errors.BadRequest("para:001", "missing parameters")
	}


	if _, err := db.GetDB().Exec("insert into agency_profile_table (" +
		"agencyID, name, tel, comments, order, tags, address, address_detail, " +
		"icon, photos, brand_history, characteristics) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		req.Agency.AgencyID, req.Agency.Name, req.Agency.Tel, req.Agency.Comments,
		req.Agency.Order, strings.Join(req.Agency.Tags, ","), req.Agency.Address,
		req.Agency.AddressDetail, req.Agency.Icon, strings.Join(req.Agency.Photos, ","),
		req.BrandHistory, strings.Join(req.Characteristics, ",")); err != nil {
			return errors.InternalServerError("agency-srv.AgencySrv.AddAgency:fatal:001", err.Error())
	}

	rsp.AgencyID = req.Agency.AgencyID
	rsp.Status = 200
	rsp.Msg = "success"

	return nil
}

func (agency *AgencySrv)UpdateAgency(ctx context.Context, req *agencysrv.UpdateAgencyRequest, rsp *agencysrv.UpdateAgencyResponse) error {
	if len(req.Agency.AgencyID) == 0 {
		return errors.BadRequest("para:001", "missing parameters")
	}

	var currentAgency agencysrv.ReadAgencyResponse
	if err := agency.ReadAgencyDetails(ctx, &agencysrv.ReadAgencyRequest{
		AgencyID:    req.Agency.AgencyID,
	}, &currentAgency); err != nil {
		return errors.InternalServerError("agency-srv.AgencySrv.UpdateAgency:fatal:001", err.Error())
	}

	if len(currentAgency.Agencies) == 0 {
		return errors.Forbidden("agency:001", "agency not existed")
	}

	if err := copier.Copy(&currentAgency.Agencies[0], req.Agency); err != nil {
		return errors.InternalServerError("agency-srv.AgencySrv.UpdateAgency:fatal:002", err.Error())
	}
	if len(req.BrandHistory) != 0 {
		currentAgency.BrandHistory = req.BrandHistory
	}
	if len(req.Characteristics) != 0 {
		currentAgency.Characteristics = req.Characteristics
	}
	// 1. update agency's profile
	if _, err := db.GetDB().Exec("update agency_profile_table set (" +
		"agencyID, name, tel, comments, order, tags, address, address_detail, " +
		"icon, photos, brand_history, characteristics) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		currentAgency.Agencies[0].AgencyID, currentAgency.Agencies[0].Name,
		currentAgency.Agencies[0].Tel, currentAgency.Agencies[0].Comments,
		currentAgency.Agencies[0].Order, strings.Join(currentAgency.Agencies[0].Tags, ","),
		currentAgency.Agencies[0].Address, currentAgency.Agencies[0].AddressDetail,
		currentAgency.Agencies[0].Icon, strings.Join(currentAgency.Agencies[0].Photos, ","),
		currentAgency.BrandHistory, strings.Join(currentAgency.Characteristics, ",")); err != nil {
		return errors.InternalServerError("agency-srv.AgencySrv.AddAgency:fatal:001", err.Error())
	}
	// TODO: 2. update teacher's profile

	rsp.Status = 200
	rsp.Msg = "success"

	return nil
}

func (agency *AgencySrv)DeleteAgency(ctx context.Context, req *agencysrv.DeleteAgencyRequest, rsp *agencysrv.DeleteAgencyResponse) error {
	return nil
}

func (agency *AgencySrv)ReadEvaluations(ctx context.Context, req *agencysrv.ReadEvaluationsRequest, rsp *agencysrv.ReadEvaluationsResponse) error {
	if len(req.AgencyID) == 0 || len(req.Uid) == 0 {
		return errors.BadRequest("para:001", "missing parameters")
	}

	if len(req.AgencyID) > 0 {
		tableName := req.AgencyID + "_agency_evaluations_table"
		good := 0
		rows, err := db.GetDB().Query("select evaluation_id, favicon, rating, username, " +
			"class_id, detail, pics from " + tableName)
		if err == sql.ErrNoRows {
			return nil
		} else if err != nil {
			return errors.InternalServerError("agency-srv.AgencySrv.ReadEvaluations:fatal:001", err.Error())
		}
		for rows.Next() {
			evaluation := new(agencysrv.Evaluation)
			var pics string
			err := rows.Scan(&evaluation.EvaluationID, &evaluation.Favicon, &evaluation.Rating, &evaluation.Username,
				&evaluation.Class.ClassID, &evaluation.Detail, &pics)
			if err != nil {
				return errors.InternalServerError("agency-srv.AgencySrv.ReadEvaluations:fatal:002", err.Error())
			}
			evaluation.Pics = strings.Split(pics, ",")
			rsp.Evaluation = append(rsp.Evaluation, evaluation)
		}

		for _, evaluation := range rsp.Evaluation {
			if evaluation.Rating >= 8.0 {
				good++
			}
		}

		// GoodRate
		rsp.OverEvaluation.GoodRate = float32(good) / float32(len(rsp.Evaluation))
		// GeneralRate
		if err := db.GetDB().QueryRow("select rating from agency_profile_table " +
			"where agency_id = ?", req.AgencyID).Scan(&rsp.OverEvaluation.GeneralRate); err != nil {
			return errors.InternalServerError("agency-srv.AgencySrv.ReadEvaluations:fatal:003", err.Error())
		}
		// UpRate
		if err := db.GetDB().QueryRow("select concat(round(((@ranking - rank) / @ranking)*100, 2), '%') as percentileRank " +
			"from (select agency_id, rating, @ranking := if ( @previous = @curr, @ranking, @ranking+1) as rank " +
			"from agency_profile_table, (select @curr := null, @previous := null, @ranking := -1) q " +
			"order by rating desc) T where agency_id = ?", req.AgencyID).Scan(&rsp.OverEvaluation.UpRate); err != nil {
			return errors.InternalServerError("agency-srv.AgencySrv.ReadEvaluations:fatal:004", err.Error())
		}

	} else {
		// only need rsp.Evaluation
		tableName := req.Uid + "_user_evaluations_table"
		rows, err := db.GetDB().Query("select evaluation_id, favicon, rating, username, " +
			"class_id, detail, pics from " + tableName)
		if err == sql.ErrNoRows {
			return nil
		} else if err != nil {
			return errors.InternalServerError("agency-srv.AgencySrv.ReadEvaluations:fatal:005", err.Error())
		}

		for rows.Next() {
			evaluation := new(agencysrv.Evaluation)
			var pics string
			err := rows.Scan(&evaluation.EvaluationID, &evaluation.Favicon, &evaluation.Rating, &evaluation.Username,
				&evaluation.Class.ClassID, &evaluation.Detail, &pics)
			if err != nil {
				return errors.InternalServerError("agency-srv.AgencySrv.ReadEvaluations:fatal:002", err.Error())
			}
			evaluation.Pics = strings.Split(pics, ",")
			rsp.Evaluation = append(rsp.Evaluation, evaluation)
		}
	}

	rsp.Status = 200
	rsp.Msg = ""

	return nil
}

func (agency *AgencySrv)AddEvaluation(ctx context.Context, req *agencysrv.AddEvaluationRequest, rsp *agencysrv.AddEvaluationResponse) error {
	return nil
}

func (agency *AgencySrv)UpdateEvaluation(ctx context.Context, req *agencysrv.UpdateEvaluationRequest, rsp *agencysrv.UpdateAgencyResponse) error {
	return nil
}

func (agency *AgencySrv)DeleteEvaluation(ctx context.Context, req *agencysrv.DeleteEvaluationRequest, rsp *agencysrv.DeleteEvaluationResponse) error {
	return nil
}

func (agency *AgencySrv)GetNearbyAgencies(ctx context.Context, req *agencysrv.GetNearbyAgenciesRequest, rsp *agencysrv.GetNearbyAgenciesResponse) error {
	return nil
}