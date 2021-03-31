package handler

import (
	"context"
	"database/sql"
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

	if len(req.S) > 0 {
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
		"agencyID, name, tel, rating, comments, order, tags, address, address_detail, " +
		"icon, photos, brand_history, characteristics) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		req.Agency.AgencyID, req.Agency.Name, req.Agency.Tel, req.Agency.Rating,
		req.Agency.Comments, req.Agency.Order, strings.Join(req.Agency.Tags, ","), req.Agency.Address,
		req.Agency.AddressDetail, req.Agency.Icon, strings.Join(req.Agency.Photos, ","), req.BrandHistory,
		strings.Join(req.Characteristics, ",")); err != nil {
			return errors.InternalServerError("agency-srv.AgencySrv.AddAgency:fatal:001", err.Error())
	}

	rsp.AgencyID = req.Agency.AgencyID
	rsp.Status = 200
	rsp.Msg = "success"

	return nil
}

func (agency *AgencySrv)UpdateAgency(ctx context.Context, req *agencysrv.UpdateAgencyRequest, rsp *agencysrv.UpdateAgencyResponse) error {
	return nil
}

func (agency *AgencySrv)DeleteAgency(ctx context.Context, req *agencysrv.DeleteAgencyRequest, rsp *agencysrv.DeleteAgencyResponse) error {
	return nil
}
