package dtp

import (
	"context"
	"errors"
	"testing"

	dtp_mock "github.com/EricHripko/buildkit-fdk/pkg/dtp/mock"

	"github.com/golang/mock/gomock"
	pb "github.com/moby/buildkit/frontend/gateway/pb"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

func TestNewDefaultDockerfile(t *testing.T) {
	// Arrange
	opts := map[string]string{}

	// Act
	proxy := NewDockerfileTransformingLLBProxy(nil, opts, nil)

	// Assert
	impl, ok := proxy.(*dockerfileTransformingLLBProxy)
	if !ok {
		t.Fatal("wrong type")
	}
	require.Equal(t, impl.dockerfile, "Dockerfile")
}

func TestNewCustomDockerfile(t *testing.T) {
	// Arrange
	opts := map[string]string{
		keyFilename: "dev.Dockerfile",
	}

	// Act
	proxy := NewDockerfileTransformingLLBProxy(nil, opts, nil)

	// Assert
	impl, ok := proxy.(*dockerfileTransformingLLBProxy)
	if !ok {
		t.Fatal("wrong type")
	}
	require.Equal(t, impl.dockerfile, "dev.Dockerfile")
}

type passThroughTestSuite struct {
	suite.Suite
	ctrl   *gomock.Controller
	client *dtp_mock.MockLLBBridgeClient
	proxy  pb.LLBBridgeClient
}

func (suite *passThroughTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.client = dtp_mock.NewMockLLBBridgeClient(suite.ctrl)
	suite.proxy = NewDockerfileTransformingLLBProxy(suite.client, map[string]string{}, func(dockerfile []byte) ([]byte, error) {
		return []byte("hijacked"), nil
	})
}

func (suite *passThroughTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *passThroughTestSuite) TestResolveImageConfig() {
	// Arrange
	ctx := context.Background()
	in := &pb.ResolveImageConfigRequest{}
	opt := grpc.CallContentSubtype("json")
	expected := &pb.ResolveImageConfigResponse{}
	suite.client.EXPECT().
		ResolveImageConfig(ctx, in, opt).
		Return(expected, nil)

	// Act
	actual, err := suite.proxy.ResolveImageConfig(ctx, in, opt)

	// Assert
	require.Nil(suite.T(), err)
	require.Same(suite.T(), expected, actual)
}

func (suite *passThroughTestSuite) TestSolve() {
	// Arrange
	ctx := context.Background()
	in := &pb.SolveRequest{}
	opt := grpc.CallContentSubtype("json")
	expected := &pb.SolveResponse{}
	suite.client.EXPECT().
		Solve(ctx, in, opt).
		Return(expected, nil)

	// Act
	actual, err := suite.proxy.Solve(ctx, in, opt)

	// Assert
	require.Nil(suite.T(), err)
	require.Same(suite.T(), expected, actual)
}

func (suite *passThroughTestSuite) TestReadDir() {
	// Arrange
	ctx := context.Background()
	in := &pb.ReadDirRequest{}
	opt := grpc.CallContentSubtype("json")
	expected := &pb.ReadDirResponse{}
	suite.client.EXPECT().
		ReadDir(ctx, in, opt).
		Return(expected, nil)

	// Act
	actual, err := suite.proxy.ReadDir(ctx, in, opt)

	// Assert
	require.Nil(suite.T(), err)
	require.Same(suite.T(), expected, actual)
}

func (suite *passThroughTestSuite) TestReadFileNonDockerfile() {
	// Arrange
	ctx := context.Background()
	in := &pb.ReadFileRequest{FilePath: "go.mod"}
	opt := grpc.CallContentSubtype("json")
	expected := &pb.ReadFileResponse{Data: []byte("hello")}
	suite.client.EXPECT().
		ReadFile(ctx, in, opt).
		Return(expected, nil)

	// Act
	actual, err := suite.proxy.ReadFile(ctx, in, opt)

	// Assert
	require.Nil(suite.T(), err)
	require.Same(suite.T(), expected, actual)
}

func (suite *passThroughTestSuite) TestReadFileFails() {
	// Arrange
	ctx := context.Background()
	in := &pb.ReadFileRequest{}
	opt := grpc.CallContentSubtype("json")
	suite.client.EXPECT().
		ReadFile(ctx, in, opt).
		Return(nil, errors.New("failed"))

	// Act
	_, err := suite.proxy.ReadFile(ctx, in, opt)

	// Assert
	require.NotNil(suite.T(), err)
}

func (suite *passThroughTestSuite) TestStatFile() {
	// Arrange
	ctx := context.Background()
	in := &pb.StatFileRequest{}
	opt := grpc.CallContentSubtype("json")
	expected := &pb.StatFileResponse{}
	suite.client.EXPECT().
		StatFile(ctx, in, opt).
		Return(expected, nil)

	// Act
	actual, err := suite.proxy.StatFile(ctx, in, opt)

	// Assert
	require.Nil(suite.T(), err)
	require.Same(suite.T(), expected, actual)
}

func (suite *passThroughTestSuite) TestPing() {
	// Arrange
	ctx := context.Background()
	in := &pb.PingRequest{}
	opt := grpc.CallContentSubtype("json")
	expected := &pb.PongResponse{}
	suite.client.EXPECT().
		Ping(ctx, in, opt).
		Return(expected, nil)

	// Act
	actual, err := suite.proxy.Ping(ctx, in, opt)

	// Assert
	require.Nil(suite.T(), err)
	require.Same(suite.T(), expected, actual)
}

func (suite *passThroughTestSuite) TestReturn() {
	// Arrange
	ctx := context.Background()
	in := &pb.ReturnRequest{}
	opt := grpc.CallContentSubtype("json")
	expected := &pb.ReturnResponse{}
	suite.client.EXPECT().
		Return(ctx, in, opt).
		Return(expected, nil)

	// Act
	actual, err := suite.proxy.Return(ctx, in, opt)

	// Assert
	require.Nil(suite.T(), err)
	require.Same(suite.T(), expected, actual)
}

func (suite *passThroughTestSuite) TestInputs() {
	// Arrange
	ctx := context.Background()
	in := &pb.InputsRequest{}
	opt := grpc.CallContentSubtype("json")
	expected := &pb.InputsResponse{}
	suite.client.EXPECT().
		Inputs(ctx, in, opt).
		Return(expected, nil)

	// Act
	actual, err := suite.proxy.Inputs(ctx, in, opt)

	// Assert
	require.Nil(suite.T(), err)
	require.Same(suite.T(), expected, actual)
}

func (suite *passThroughTestSuite) TestNewContainer() {
	// Arrange
	ctx := context.Background()
	in := &pb.NewContainerRequest{}
	opt := grpc.CallContentSubtype("json")
	expected := &pb.NewContainerResponse{}
	suite.client.EXPECT().
		NewContainer(ctx, in, opt).
		Return(expected, nil)

	// Act
	actual, err := suite.proxy.NewContainer(ctx, in, opt)

	// Assert
	require.Nil(suite.T(), err)
	require.Same(suite.T(), expected, actual)
}

func (suite *passThroughTestSuite) TestReleaseContainer() {
	// Arrange
	ctx := context.Background()
	in := &pb.ReleaseContainerRequest{}
	opt := grpc.CallContentSubtype("json")
	expected := &pb.ReleaseContainerResponse{}
	suite.client.EXPECT().
		ReleaseContainer(ctx, in, opt).
		Return(expected, nil)

	// Act
	actual, err := suite.proxy.ReleaseContainer(ctx, in, opt)

	// Assert
	require.Nil(suite.T(), err)
	require.Same(suite.T(), expected, actual)
}

func (suite *passThroughTestSuite) TestExecProcess() {
	// Arrange
	ctx := context.Background()
	opt := grpc.CallContentSubtype("json")
	expected := dtp_mock.NewMockLLBBridge_ExecProcessClient(suite.ctrl)
	suite.client.EXPECT().
		ExecProcess(ctx, opt).
		Return(expected, nil)

	// Act
	actual, err := suite.proxy.ExecProcess(ctx, opt)

	// Assert
	require.Nil(suite.T(), err)
	require.Same(suite.T(), expected, actual)
}

func TestPassThrough(t *testing.T) {
	suite.Run(t, new(passThroughTestSuite))
}

func TestTransformDockerfile(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	client := dtp_mock.NewMockLLBBridgeClient(ctrl)

	opts := map[string]string{}
	transform := func(dockerfile []byte) ([]byte, error) {
		return append(dockerfile, ' ', 'w', 'o', 'r', 'l', 'd'), nil
	}
	proxy := NewDockerfileTransformingLLBProxy(client, opts, transform)

	ctx := context.Background()
	in := &pb.ReadFileRequest{FilePath: "Dockerfile"}
	client.EXPECT().
		ReadFile(ctx, in).
		Return(&pb.ReadFileResponse{Data: []byte("hello")}, nil)

	// Act
	out, err := proxy.ReadFile(ctx, in)

	// Assert
	require.Nil(t, err)
	require.Equal(t, "hello world", string(out.Data))
}
