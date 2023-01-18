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
)

type (
	cuser struct {
		// ac        exposedModuleAccessController
		module    cs.ModuleService
		namespace cs.NamespaceService
		role      ss.RoleService
		store     store.Storer
		actionlog actionlog.Recorder
	}

	// userAccessController interface {
	// 	CanCreateModuleOnNode(ctx context.Context, r *types.Node) bool
	// 	CanManageExposedModule(ctx context.Context, r *types.ExposedModule) bool
	// }

	userService interface {
		Create(ctx context.Context, new *types.Cuser) (*types.Cuser, error)
		// generate filter
		Find(ctx context.Context, filter types.CuserFilter) (types.CuserSet, types.CuserFilter, error)
	}

	// moduleUpdateHandler func(ctx context.Context, ns *types.Node, c *types.ExposedModule) (bool, bool, error)
)

func Cuser() *cuser {
	return &cuser{
		// ac:        DefaultAccessControl,
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
	return types.CuserSet{}, types.CuserFilter{}, nil
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
