package service

import (
	"context"
	"fmt"

	"github.com/cortezaproject/corteza/server/acme/types"
	cs "github.com/cortezaproject/corteza/server/compose/service"
	"github.com/cortezaproject/corteza/server/pkg/actionlog"
	"github.com/cortezaproject/corteza/server/pkg/errors"
	"github.com/cortezaproject/corteza/server/store"
	ss "github.com/cortezaproject/corteza/server/system/service"
	"github.com/davecgh/go-spew/spew"
)

type (
	cuser struct {
		ac        userAccessController
		module    cs.ModuleService
		namespace cs.NamespaceService
		role      ss.RoleService
		store     store.Storer
		actionlog actionlog.Recorder
	}

	userAccessController interface {
		CanSearchUser(ctx context.Context) bool
	}

	userService interface {
		Create(ctx context.Context, new *types.Cuser) (*types.Cuser, error)
		Find(ctx context.Context, filter types.CuserFilter) (types.CuserSet, types.CuserFilter, error)
	}
)

func Cuser() *cuser {
	return &cuser{
		ac:        DefaultAccessControl,
		module:    cs.DefaultModule,
		role:      ss.DefaultRole,
		namespace: cs.DefaultNamespace,
		store:     DefaultStore,
		actionlog: DefaultActionlog,
	}
}

func (svc cuser) Create(ctx context.Context, new *types.Cuser) (*types.Cuser, error) {
	err := svc.store.CreateAcmeCuser(ctx, &types.Cuser{ID: nextID(), Name: "foo"})
	return new, err
}

func (svc cuser) Find(ctx context.Context, filter types.CuserFilter) (types.CuserSet, types.CuserFilter, error) {
	spew.Dump("can search user", svc.ac.CanSearchUser(ctx))
	return svc.store.SearchAcmeCusers(ctx, types.CuserFilter{})
}

func loadCuser(ctx context.Context, s store.AcmeCusers, ID uint64) (res *types.Cuser, err error) {
	if ID == 0 {
		return nil, fmt.Errorf("not found")
	}

	if res, err = store.LookupAcmeCuserByID(ctx, s, ID); errors.IsNotFound(err) {
		err = fmt.Errorf("not found")
	}

	return
}
