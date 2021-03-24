package handler

import (
	"context"
	classsrv "github.com/wen-qu/xuesou-backend-service/class-srv/proto"
)

type ClassSrv struct{}

func (class *ClassSrv) ReadClassesByAgencyID(ctx context.Context, req *classsrv.ReadClassRequest, rsp *classsrv.ReadClassResponse) error {
	return nil
}

func (class *ClassSrv) AddClasses(ctx context.Context, req *classsrv.AddClassRequest, rsp *classsrv.AddClassResponse) error {
	return nil
}

func (class *ClassSrv) UpdateClass(ctx context.Context, req *classsrv.UpdateClassRequest, rsp *classsrv.UpdateClassResponse) error {
	return nil
}

func (class *ClassSrv) DeleteClass(ctx context.Context, req *classsrv.DeleteClassRequest, rsp *classsrv.DeleteClassResponse) error {
	return nil
}