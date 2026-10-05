package v1_test

import (
	"context"
	"errors"
	"testing"

	"sms_backend/api/v1/sms"
	v1 "sms_backend/internal/controller/v1"
	"sms_backend/internal/service"
)

type projectDeleteServiceStub struct {
	service.IMainControllerProjectManagement
	deleteProject func(context.Context, *sms.ProjectDeleteReq) (*sms.ProjectDeleteRes, error)
}

func (s *projectDeleteServiceStub) DeleteProject(ctx context.Context, req *sms.ProjectDeleteReq) (*sms.ProjectDeleteRes, error) {
	return s.deleteProject(ctx, req)
}

func TestProjectDeleteDelegatesToProjectManagement(t *testing.T) {
	// The controller package does not import production service implementations.
	// Keep these cases serial because GoFrame's service registry is process-wide.
	t.Cleanup(func() { service.RegisterMainControllerProjectManagement(nil) })

	for _, tc := range []struct {
		name string
		res  *sms.ProjectDeleteRes
		err  error
	}{
		{name: "success", res: &sms.ProjectDeleteRes{}},
		{name: "service rejects deletion", err: errors.New("project cannot be deleted")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			req := &sms.ProjectDeleteReq{ProjectID: 42}
			calls := 0
			service.RegisterMainControllerProjectManagement(&projectDeleteServiceStub{
				deleteProject: func(gotCtx context.Context, gotReq *sms.ProjectDeleteReq) (*sms.ProjectDeleteRes, error) {
					calls++
					if gotCtx != ctx {
						t.Error("context was not forwarded unchanged")
					}
					if gotReq != req {
						t.Error("project-delete request was not forwarded unchanged")
					}
					return tc.res, tc.err
				},
			})

			res, err := v1.NewSms().ProjectDelete(ctx, req)
			if calls != 1 {
				t.Errorf("DeleteProject called %d times, want 1", calls)
			}
			if res != tc.res || err != tc.err {
				t.Errorf("got (%v, %v), want (%v, %v)", res, err, tc.res, tc.err)
			}
		})
	}
}
