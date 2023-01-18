package service

import (
	"context"

	"github.com/cortezaproject/corteza/server/pkg/actionlog"
	"github.com/cortezaproject/corteza/server/pkg/id"
	"github.com/cortezaproject/corteza/server/pkg/logger"
	"github.com/cortezaproject/corteza/server/store"
	"go.uber.org/zap"
)

var (
	DefaultStore store.Storer

	DefaultLogger *zap.Logger

	// DefaultAccessControl *accessControl

	DefaultActionlog actionlog.Recorder

	DefaultCuser *cuser

	// wrapper around id.Next() that will aid service testing
	nextID = func() uint64 {
		return id.Next()
	}
)

func Initialize(_ context.Context, log *zap.Logger, s store.Storer) (err error) {

	// we're doing conversion to avoid having
	// store interface exposed or generated inside app package
	DefaultStore = s

	DefaultLogger = log.Named("service")

	DefaultActionlog = actionlog.NewService(DefaultStore, log, logger.MakeDebugLogger(), actionlog.MakeDebugPolicy())

	DefaultCuser = Cuser()

	// DefaultAccessControl = AccessControl(s)

	return
}
