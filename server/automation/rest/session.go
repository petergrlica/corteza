package rest

import (
	"context"
	"github.com/cortezaproject/corteza/server/automation/rest/request"
	"github.com/cortezaproject/corteza/server/automation/service"
	"github.com/cortezaproject/corteza/server/automation/types"
	"github.com/cortezaproject/corteza/server/pkg/api"
	"github.com/cortezaproject/corteza/server/pkg/auth"
	"github.com/cortezaproject/corteza/server/pkg/expr"
	"github.com/cortezaproject/corteza/server/pkg/filter"
	"github.com/cortezaproject/corteza/server/pkg/payload"
	"github.com/cortezaproject/corteza/server/pkg/wfexec"
)

type (
	Session struct {
		svc sessionService
	}

	sessionService interface {
		Search(ctx context.Context, filter types.SessionFilter) (types.SessionSet, types.SessionFilter, error)
		LookupByID(ctx context.Context, sessionID uint64) (*types.Session, error)
		Resume(sessionID, stateID uint64, i auth.Identifiable, input *expr.Vars) error
		PendingPrompts(context.Context) []*wfexec.PendingPrompt
		Cancel(context.Context, uint64) error
	}

	sessionSetPayload struct {
		Filter types.SessionFilter `json:"filter"`
		Set    []*sessionSetItem   `json:"set"`
	}

	sessionSetItem struct {
		*types.Session

		// Make sure stacktrace is not included in the list
		Stacktrace *struct{} `json:"stacktrace,omitempty"`
		// Make sure output values are not included in the list
		Output *struct{} `json:"output,omitempty"`
		// Make sure input values are not included in the list
		Input *struct{} `json:"input,omitempty"`
	}
)

func (Session) New() *Session {
	ctrl := &Session{}
	ctrl.svc = service.DefaultSession
	return ctrl
}

func (ctrl Session) List(ctx context.Context, r *request.SessionList) (interface{}, error) {
	var (
		err error
		f   = types.SessionFilter{
			WorkflowID:   payload.ParseUint64s(r.WorkflowID),
			SessionID:    payload.ParseUint64s(r.SessionID),
			CreatedBy:    payload.ParseUint64s(r.CreatedBy),
			EventType:    r.EventType,
			ResourceType: r.ResourceType,
			Completed:    filter.State(r.Completed),
			Status:       r.Status,
		}
	)

	if f.Paging, err = filter.NewPaging(r.Limit, r.PageCursor); err != nil {
		return nil, err
	}

	f.IncTotal = r.IncTotal

	if f.Sorting, err = filter.NewSorting(r.Sort); err != nil {
		return nil, err
	}

	set, filter, err := ctrl.svc.Search(ctx, f)
	return ctrl.makeFilterPayload(ctx, set, filter, err)
}

func (ctrl Session) Read(ctx context.Context, r *request.SessionRead) (interface{}, error) {
	return ctrl.svc.LookupByID(ctx, r.SessionID)
}

func (ctrl Session) Cancel(ctx context.Context, r *request.SessionCancel) (interface{}, error) {
	return true, ctrl.svc.Cancel(ctx, r.SessionID)
}

func (ctrl Session) ListPrompts(ctx context.Context, r *request.SessionListPrompts) (interface{}, error) {
	return struct {
		Set []*wfexec.PendingPrompt `json:"set"`
	}{
		Set: ctrl.svc.PendingPrompts(ctx),
	}, nil
}

func (ctrl Session) ResumeState(ctx context.Context, r *request.SessionResumeState) (interface{}, error) {
	return api.OK(), ctrl.svc.Resume(r.SessionID, r.StateID, auth.GetIdentityFromContext(ctx), r.Input)
}

func (ctrl Session) makeFilterPayload(ctx context.Context, ss types.SessionSet, f types.SessionFilter, err error) (*sessionSetPayload, error) {
	if err != nil {
		return nil, err
	}

	out := &sessionSetPayload{Filter: f, Set: make([]*sessionSetItem, len(ss))}

	for i, s := range ss {
		s.Stacktrace = nil
		out.Set[i] = &sessionSetItem{Session: s}
	}

	return out, nil
}
