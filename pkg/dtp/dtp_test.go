package dtp

import (
	"context"
	"testing"

	cib_mock "github.com/EricHripko/buildkit-fdk/pkg/cib/mock"
	dtp_mock "github.com/EricHripko/buildkit-fdk/pkg/dtp/mock"

	"github.com/golang/mock/gomock"
	"github.com/moby/buildkit/frontend/gateway/client"
	"github.com/moby/buildkit/frontend/gateway/grpcclient"
	pb "github.com/moby/buildkit/frontend/gateway/pb"
	"github.com/stretchr/testify/require"
)

func noOpTransform(dockerfile []byte) ([]byte, error) {
	return dockerfile, nil
}

func TestHandlesPanics(t *testing.T) {
	// Arrange
	client := cib_mock.NewMockClient(nil)

	// Act
	err := InjectDockerfileTransform(noOpTransform, client)

	// Assert
	require.NotNil(t, err)
}

type invalidTypeClient struct {
	cib_mock.MockClient
	client string
}

func TestUnexpectedType(t *testing.T) {
	// Arrange
	client := &invalidTypeClient{client: "hello"}

	// Act
	err := InjectDockerfileTransform(noOpTransform, client)

	// Assert
	require.NotNil(t, err)
}

func TestInjectsSuccess(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	llb := dtp_mock.NewMockLLBBridgeClient(ctrl)
	llb.EXPECT().
		Ping(gomock.Any(), gomock.Any()).
		Return(&pb.PongResponse{}, nil)

	client, err := grpcclient.New(
		context.Background(),
		map[string]string{},
		"fakesession",
		"fakeproduct",
		llb,
		[]client.WorkerInfo{},
	)
	if err != nil {
		t.Fatal(err)
	}

	// Act
	err = InjectDockerfileTransform(noOpTransform, client)

	// Assert
	require.Nil(t, err)
}
