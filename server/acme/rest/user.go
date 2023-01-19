package rest

import (
	"context"

	"github.com/cortezaproject/corteza/server/acme/rest/request"
	"github.com/cortezaproject/corteza/server/acme/service"
	"github.com/cortezaproject/corteza/server/acme/types"
)

type (
	userServicer interface {
		Create(ctx context.Context, new *types.Cuser) (*types.Cuser, error)
		Find(ctx context.Context, filter types.CuserFilter) (types.CuserSet, types.CuserFilter, error)
	}

	Users struct {
		svcUser userServicer
	}

	usersAccessController interface {
	}
)

func (Users) New() *Users {
	return &Users{
		svcUser: service.DefaultCuser,
	}
}

func (ctrl Users) List(ctx context.Context, r *request.UsersList) (l interface{}, err error) {
	l, _, err = ctrl.svcUser.Find(ctx, types.CuserFilter{Query: r.Query})
	return
}

func (ctrl Users) Create(ctx context.Context, r *request.UsersCreate) (interface{}, error) {
	return ctrl.svcUser.Create(ctx, &types.Cuser{})
}
