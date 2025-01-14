package server

import (
	"context"
	"testing"

	"github.com/go-logr/logr/testr"
	v1 "github.com/llmariner/model-manager/api/v1"
	"github.com/llmariner/model-manager/server/internal/store"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestHFModelRepo(t *testing.T) {
	st, tearDown := store.NewTest(t)
	defer tearDown()

	const (
		repoName = "r0"
	)

	wsrv := NewWorkerServiceServer(st, testr.New(t))
	ctx := fakeAuthInto(context.Background())
	_, err := wsrv.GetHFModelRepo(ctx, &v1.GetHFModelRepoRequest{
		Name: repoName,
	})
	assert.Error(t, err)
	assert.Equal(t, codes.NotFound, status.Code(err))

	_, err = wsrv.CreateHFModelRepo(ctx, &v1.CreateHFModelRepoRequest{
		Name: repoName,
	})
	assert.NoError(t, err)

	got, err := wsrv.GetHFModelRepo(ctx, &v1.GetHFModelRepoRequest{
		Name: repoName,
	})
	assert.NoError(t, err)
	assert.Equal(t, got.Name, repoName)
}
