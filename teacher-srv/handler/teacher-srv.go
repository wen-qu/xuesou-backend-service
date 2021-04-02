package handler

import (
	"context"
	teachersrv "teacher-srv/proto"
)

type TeacherSrv struct{}

func (teacher *TeacherSrv)GetTeachers(ctx context.Context, req *teachersrv.GetTeachersRequest, rsp *teachersrv.GetTeachersResponse) error {
	return nil
}

func (teacher *TeacherSrv)AddTeacher(ctx context.Context, req *teachersrv.AddTeacherRequest, rsp *teachersrv.AddTeacherResponse) error {
	return nil
}

func (teacher *TeacherSrv)UpdateTeacher(ctx context.Context, req *teachersrv.UpdateTeacherRequest, rsp *teachersrv.UpdateTeacherResponse) error {
	return nil
}

func (teacher *TeacherSrv)DeleteTeacher(ctx context.Context, req *teachersrv.DeleteTeacherRequest, rsp *teachersrv.DeleteTeacherResponse) error {
	return nil
}