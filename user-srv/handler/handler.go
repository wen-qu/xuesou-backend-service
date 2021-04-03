package handler

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/micro/micro/v3/service/errors"

	log "github.com/micro/micro/v3/service/logger"

	"github.com/wen-qu/xuesou-backend-service/basic/db"

	usersrv "github.com/wen-qu/xuesou-backend-service/user-srv/proto"

)

// UserSrv the UserSrv struct
type UserSrv struct{}

// AddUser add a user
func (e *UserSrv) AddUser(ctx context.Context, req *usersrv.AddRequest, rsp *usersrv.AddResponse) error {
	log.Info("Received UserSrv.Login request")
	var user usersrv.InspectResponse

	if err := e.InspectUser(ctx, &usersrv.InspectRequest{
		Tel: req.User.Tel,
	}, &user); err != nil {
		return errors.InternalServerError("user-srv.UserSrv.AddUser:fatal:001", err.Error())
	}
	if user.User != nil {
		rsp.Status = 400
		rsp.Msg = "registered"
		return nil
	}

	uid := "user_" + uuid.New().String()
	if _, err := db.GetDB().Exec("insert into user " +
		"(uid, username, password, tel, email, sex, age, address, class_num, img) " +
		"values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", uid, req.User.Username, req.User.Password,
		req.User.Tel, req.User.Email, req.User.Sex, req.User.Age, req.User.Address,
		req.User.ClassNum, req.User.Img); err != nil {
		return errors.InternalServerError("user-srv.UserSrv.AddUser:fatal:002", err.Error())
	}

	// create table [uid]_user_class_table, [uid]_user_chatting_table, [uid]_user_evaluations_table
	tableName := uid + "_user_class_table"
	if _, err := db.GetDB().Exec("create table `" + tableName + "` (" +
		"`uid` varchar(18) not null," +
		"`class_id` varchar(19) not null," +
		"`bought_time` varchar(20) not null," +
		"`agency_id` varchar(20) not null" +
		") engine=innodb default charset=utf8"); err != nil {
		return errors.InternalServerError("user-srv.UserSrv.AddUser:fatal:003", err.Error())
	}
	tableName = uid + "_user_chatting_table"
	if _, err := db.GetDB().Exec("create table `" + tableName + "` (" +
		"`chat_id` varchar(18) primary key not null," +
		"`uid` varchar(18) not null," +
		"`msg_num` int not null," +
		"`agency_icon` varchar(60)," +
		"`agency_id` varchar(20) not null," +
		"`agency_name` varchar(50) not null" +
		") engine=innodb default charset=utf8"); err != nil {
		return errors.InternalServerError("user-srv.UserSrv.AddUser:fatal:004", err.Error())
	}
	tableName = uid + "_user_evaluations_table"
	if _, err := db.GetDB().Exec("create table `" + tableName + "` (" +
		"`evaluation_id` varchar(20) primary key not null," +
		"`favicon` varchar(60)," +
		"`rating` float not null," +
		"`username` varchar(50) not null," +
		"`agency_id` varchar(20) not null," +
		"`uid` varchar(18) not null," +
		"`class_id` varchar(19) not null," +
		"`detail` varchar(10000)," +
		"`pics` varchar(700)" +
		") engine=innodb default charset=utf8"); err != nil {
		return errors.InternalServerError("user-srv.UserSrv.AddUser:fatal:005", err.Error())
	}

	rsp.Status = 200
	rsp.Msg = "success"

	return nil
}

// InspectUser inspect a user by uid or tel
func (e *UserSrv) InspectUser(ctx context.Context, req *usersrv.InspectRequest, rsp *usersrv.InspectResponse) error {
	log.Info("Received UserSrv.Register request")
	var user usersrv.User
	var row *sql.Row

	if len(req.Uid) > 0 {
		row = db.GetDB().QueryRow("select uid, username, password, tel, age, sex, email, " +
			"address, class_num, img from user where uid = ?", req.Uid)
	} else if len(req.Tel) > 0 {
		row = db.GetDB().QueryRow("select uid, username, password, tel, age, sex, email, " +
			"address, class_num, img from user where uid = ?", req.Uid)
	} else {
		return errors.BadRequest("para:002", "missing uid or tel")
	}

	err := row.Scan(
		&user.Uid, &user.Username, &user.Password,
		&user.Tel, &user.Age, &user.Sex,
		&user.Email, &user.Address, &user.ClassNum, &user.Img)

	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return errors.InternalServerError("user-srv.UserSrv.InspectUser:fatal:001", err.Error())
	}

	rsp.User = &user

	return nil
}

// UpdateUser update user's profile
func (e *UserSrv) UpdateUser(ctx context.Context, req *usersrv.UpdateRequest, rsp *usersrv.UpdateResponse) error {
	log.Info("Received UserSrv.Validation request")

	var currentUser usersrv.InspectResponse

	if err := e.InspectUser(ctx, &usersrv.InspectRequest{
		Uid: req.User.Uid,
		Tel: req.User.Tel,
	}, &currentUser); err != nil {
		return errors.InternalServerError("user-srv.UserSrv.UpdateUser:fatal:001", err.Error())
	}
	if currentUser.User == nil {
		return errors.Forbidden("user:001", "user not existed")
	}

	if err := copier.Copy(req.User, currentUser.User); err != nil {
		return errors.InternalServerError("user-srv.UserSrv.UpdateUser:fatal:002", err.Error())
	}

	_, err := db.GetDB().Exec("update user set uid = ?, username = ?, password = ?, tel = ?, " +
		"age = ?, sex = ?, email = ?, address = ?, class_num = ?, img = ? where uid = ? ",
		req.User.Uid, req.User.Username, req.User.Password,
		req.User.Tel, req.User.Age, req.User.Sex,
		req.User.Email, req.User.Address, req.User.ClassNum,
		req.User.Img, req.User.Uid)

	if err != nil {
		return errors.InternalServerError("user-srv.UserSrv.UpdateUser:fatal:003", err.Error())
	}

	rsp.Status = 200
	rsp.Msg = "success"

	return nil
}

// DeleteUser delete a user
func (e *UserSrv) DeleteUser(ctx context.Context, req *usersrv.DeleteRequest, rsp *usersrv.DeleteResponse) error {
	log.Info("Received UserSrv.Register request")

	if len(req.Tel) == 0 && len(req.Uid) == 0 {
		return errors.BadRequest("para:002", "missing parameters")
	}

	var goalUser usersrv.InspectResponse
	var err error

	if err := e.InspectUser(ctx, &usersrv.InspectRequest{
		Uid: req.Tel,
		Tel: req.Uid,
	}, &goalUser); err != nil {
		return errors.InternalServerError("user-srv.UserSrv.DeleteUser:fatal:001", err.Error())
	}

	if goalUser.User == nil {
		return errors.Forbidden("user:001", "user not existed")
	}

	if len(req.Tel) > 0 {
		_, err = db.GetDB().Exec("delete from user where tel = ?", req.Tel)
	} else if len(req.Uid) > 0 {
		_, err = db.GetDB().Exec("delete from user where uid = ?", req.Uid)
	}

	if err != nil {
		return errors.InternalServerError("user-srv.UserSrv.DeleteUser:fatal:002", err.Error())
	}

	// TODO: delete user's chatting table, evaluations table, class table

	rsp.Status = 200
	rsp.Msg = "success"

	return nil
}
