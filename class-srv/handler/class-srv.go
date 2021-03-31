package handler

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/micro/micro/v3/service/errors"
	"github.com/wen-qu/xuesou-backend-service/basic/db"
	classsrv "github.com/wen-qu/xuesou-backend-service/class-srv/proto"
	"reflect"
	"regexp"
	"time"
)

type ClassSrv struct{}

func (class *ClassSrv) ReadClassesByAgencyID(ctx context.Context, req *classsrv.ReadClassRequest, rsp *classsrv.ReadClassResponse) error {
	if len(req.AgencyID) == 0 {
		return errors.BadRequest("para:001", "missing parameters")
	}

	if matched, _ := regexp.Match("/^agency_[0-9]{13}$/", []byte(req.AgencyID)); !matched {
		return errors.BadRequest("para:002", "invalid parameters")
	}
	rows, err := db.GetDB().Query("select agency_id, class_id, price, name, stu_number, age, level, sales from " +
		req.AgencyID + "_agency_class_table")

	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return errors.InternalServerError("class-srv.ClassSrv.ReadClassesByAgencyID:fatal:001", err.Error())
	}

	for rows.Next() {
		cl := new(classsrv.Class)
		err := rows.Scan(&cl.AgencyID, &cl.ClassID, &cl.Price, &cl.Name, &cl.StuNumber, &cl.Age, &cl.Level, &cl.Sales)
		if err != nil {
			return errors.InternalServerError("class-srv.ClassSrv.ReadClassesByAgencyID:fatal:002", err.Error())
		}
		rsp.Classes = append(rsp.Classes, cl)
	}

	rsp.Msg = ""
	rsp.Status = 200

	return nil
}

func (class *ClassSrv) AddClasses(ctx context.Context, req *classsrv.AddClassRequest, rsp *classsrv.AddClassResponse) error {
	if matched, _ := regexp.Match("/^agency_[0-9]{13}$/", []byte(req.Class.AgencyID)); !matched {
		return errors.BadRequest("para:002", "invalid parameters: agencyID")
	}

	if len(req.Class.Name) == 0 {
		return errors.BadRequest("para:001", "missing parameters: name")
	}
	tableName := req.Class.AgencyID + "_agency_class_table"
	createTime := time.Now().String()[:19] // e.g. "2006-01-02 15:04:05"
	lastUpdateTime := createTime
	classID := "class_" + uuid.New().String()

	_, err := db.GetDB().Exec("insert into " + tableName + "(agency_id, class_id, price, name, " +
		"stu_number, age, level, sales, create_time, last_update_time) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		req.Class.AgencyID, classID, req.Class.Price, req.Class.Name, req.Class.StuNumber,
		req.Class.Age, req.Class.Level, req.Class.Sales, createTime, lastUpdateTime)

	if err != nil {
		return errors.InternalServerError("class-srv.ClassSrv.AddClasses:fatal:001", err.Error())
	}

	rsp.ClassID = classID
	rsp.Msg = ""
	rsp.Status = 200

	return nil
}

func (class *ClassSrv) UpdateClass(ctx context.Context, req *classsrv.UpdateClassRequest, rsp *classsrv.UpdateClassResponse) error {
	if matched, _ := regexp.Match("/^agency_[0-9]{13}$/", []byte(req.Class.AgencyID)); !matched {
		return errors.BadRequest("para:002", "invalid parameters: agencyID")
	}

	if matched, _ := regexp.Match("/^class_[0-9]{13}$/", []byte(req.Class.ClassID)); !matched {
		return errors.BadRequest("para:002", "invalid parameters: classID")
	}

	if len(req.Class.Name) == 0 {
		return errors.BadRequest("para:002", "invalid parameters: name")
	}

	var currentClass classsrv.ReadClassResponse

	if err := class.ReadClassesByAgencyID(ctx, &classsrv.ReadClassRequest{
		AgencyID: req.Class.AgencyID,
	}, &currentClass); err != nil {
		return errors.InternalServerError("class-srv.ClassSrv.UpdateClass:fatal:001", err.Error())
	}
	if len(currentClass.Classes) == 0 {
		return errors.Forbidden("class:001", "class not existed")
	}

	vReqClass := reflect.ValueOf(req.Class)
	vCurClass := reflect.ValueOf(currentClass.Classes[0])

	for i := 0; i < vCurClass.NumField(); i++ {
		switch vCurClass.Field(i).Kind() {
		case reflect.String:
			if len(vReqClass.Field(i).Interface().(string)) == 0 {
				vField := reflect.ValueOf(vReqClass.Field(i).Pointer())
				vField.Elem().SetString(vCurClass.Field(i).Interface().(string))
			}
		case reflect.Int32:
			if vReqClass.Field(i).Interface().(int32) == 0 {
				vField := reflect.ValueOf(vReqClass.Field(i).Pointer())
				vField.Elem().SetInt(vCurClass.Field(i).Interface().(int64))
			}
		}
	}

	tableName := req.Class.AgencyID + "_agency_class_name"
	lastUpdateTime := time.Now().String()[:19]
	_, err := db.GetDB().Exec("update " + tableName + "set " +
		"name = ?, price = ?, stu_number = ?, age = ?, level = ?, " +
		"sales = ?, last_update_time = ? where agency_id = ? and class_id = ?",
		req.Class.Name, req.Class.Price, req.Class.StuNumber, req.Class.Age,
		req.Class.Level, req.Class.Sales, lastUpdateTime, req.Class.AgencyID, req.Class.ClassID)
	if err != nil {
		return errors.InternalServerError("class-srv.ClassSrv.UpdateClass:fatal:002", err.Error())
	}

	rsp.Status = 200
	rsp.Msg = "success"

	return nil
}

func (class *ClassSrv) DeleteClass(ctx context.Context, req *classsrv.DeleteClassRequest, rsp *classsrv.DeleteClassResponse) error {
	if len(req.ClassID) == 0 || len(req.AgencyID) == 0 {
		return errors.BadRequest("para:001", "missing parameters")
	}


	var currentClass classsrv.ReadClassResponse

	if err := class.ReadClassesByAgencyID(ctx, &classsrv.ReadClassRequest{
		AgencyID: req.AgencyID,
	}, &currentClass); err != nil {
		return errors.Forbidden("class:001", "class not existed")
	}

	if req.AgencyID != currentClass.Classes[0].ClassID {
		return errors.Forbidden("class:003", "classID does not match agencyID")
	}
	tableName := req.AgencyID + "_agency_class_name"
	_, err := db.GetDB().Exec("delete from " + tableName + "where class_id = ?", currentClass.Classes[0].ClassID)

	if err != nil {
		return errors.InternalServerError("class-srv.ClassSrv.DeleteClass:fatal:001", err.Error())
	}

	rsp.Msg = ""
	rsp.Status = 200

	return nil
}