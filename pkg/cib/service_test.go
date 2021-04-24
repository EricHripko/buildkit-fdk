package cib

import (
	"context"
	"errors"
	"testing"

	cib_mock "github.com/EricHripko/buildkit-fdk/pkg/cib/mock"

	"github.com/golang/mock/gomock"
	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/frontend/gateway/client"
	"github.com/moby/buildkit/solver/pb"
	"github.com/moby/buildkit/util/apicaps"
	"github.com/opencontainers/go-digest"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestNewServiceSucceeds(t *testing.T) {
	// Act
	build := NewService(context.Background(), nil)

	// Arrange
	require.NotNil(t, build)
}

type serviceTestSuite struct {
	suite.Suite
	ctrl   *gomock.Controller
	ctx    context.Context
	client *cib_mock.MockClient
	build  Service
}

func (suite *serviceTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.ctx = context.Background()
	suite.client = cib_mock.NewMockClient(suite.ctrl)
	suite.build = NewService(suite.ctx, suite.client)

	suite.client.EXPECT().
		BuildOpts().
		Return(client.BuildOpts{
			Opts: map[string]string{
				"key":               "value",
				keyImageResolveMode: pb.AttrImageResolveModePreferLocal,
			},
			LLBCaps: apicaps.CapSet{},
			Workers: []client.WorkerInfo{
				{Platforms: []specs.Platform{
					{
						Architecture: "amd64",
						OS:           "linux",
						OSVersion:    "0.0.0",
					},
				}},
			},
		}).
		AnyTimes()
}

func (suite *serviceTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *serviceTestSuite) TestSrcSolveFails() {
	// Arrange
	actual := errors.New("something went wrong")
	suite.client.EXPECT().
		Solve(suite.ctx, gomock.Any()).
		Return(nil, actual)

	// Act
	_, expected := suite.build.Src()

	// Assert
	require.Same(suite.T(), expected, actual)
}

func (suite *serviceTestSuite) TestSrcSucceeds() {
	// Arrange
	expected := cib_mock.NewMockReference(suite.ctrl)
	res := client.NewResult()
	res.SetRef(expected)
	suite.client.EXPECT().
		Solve(suite.ctx, gomock.Any()).
		Return(res, nil)

	// Act
	actual, err := suite.build.Src()

	// Assert
	require.Nil(suite.T(), err)
	require.Same(suite.T(), expected, actual)
}

func (suite *serviceTestSuite) TestSrcState() {
	// Arrange
	ref := cib_mock.NewMockReference(suite.ctrl)
	res := client.NewResult()
	res.SetRef(ref)
	suite.client.EXPECT().
		Solve(suite.ctx, gomock.Any()).
		Return(res, nil)

	expected := llb.Local("test")
	ref.EXPECT().
		ToState().
		Return(expected, nil)

	// Act
	_, err := suite.build.SrcState()

	// Assert
	require.Nil(suite.T(), err)
}

func (suite *serviceTestSuite) TestGetOpts() {
	// Act
	opts := suite.build.GetOpts()

	// Assert
	require.Contains(suite.T(), opts, "key")
	require.Equal(suite.T(), opts["key"], "value")
}

func (suite *serviceTestSuite) TestGetCaps() {
	// Act
	caps := suite.build.GetCaps()

	// Assert
	require.Equal(suite.T(), apicaps.CapSet{}, caps)
}

func (suite *serviceTestSuite) TestGetResolveMode() {
	// Act
	resolveMode, err := suite.build.GetResolveMode()

	// Assert
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), llb.ResolveModePreferLocal, resolveMode)
}

func (suite *serviceTestSuite) TestGetMetadataFileName() {
	// Act
	name := suite.build.GetMetadataFileName()

	// Assert
	require.Equal(suite.T(), "Dockerfile", name)
}

func (suite *serviceTestSuite) TestGetMetadataSolveFails() {
	// Arrange
	suite.client.EXPECT().
		Solve(suite.ctx, gomock.Any()).
		Return(nil, errors.New("something went wrong"))

	// Act
	_, err := suite.build.GetMetadata()

	// Assert
	require.NotNil(suite.T(), err)
	require.Contains(suite.T(), err.Error(), "something went wrong")
}

func (suite *serviceTestSuite) TestGetMetadataRefFails() {
	// Arrange
	res := client.NewResult()
	res.AddRef("test", cib_mock.NewMockReference(suite.ctrl))
	suite.client.EXPECT().
		Solve(suite.ctx, gomock.Any()).
		Return(res, nil)

	// Act
	_, err := suite.build.GetMetadata()

	// Assert
	require.NotNil(suite.T(), err)
}

func (suite *serviceTestSuite) TestGetMetadataReadFails() {
	// Arrange
	req := client.ReadRequest{Filename: "Dockerfile"}
	ref := cib_mock.NewMockReference(suite.ctrl)
	ref.EXPECT().
		ReadFile(suite.ctx, req).
		Return(nil, errors.New("something went wrong"))

	res := client.NewResult()
	res.SetRef(ref)
	suite.client.EXPECT().
		Solve(suite.ctx, gomock.Any()).
		Return(res, nil)

	// Act
	_, err := suite.build.GetMetadata()

	// Assert
	require.NotNil(suite.T(), err)
	require.Contains(suite.T(), err.Error(), "something went wrong")
}

func (suite *serviceTestSuite) TestGetMetadataSucceeds() {
	// Arrange
	expected := []byte("hello world")
	req := client.ReadRequest{Filename: "Dockerfile"}
	ref := cib_mock.NewMockReference(suite.ctrl)
	ref.EXPECT().
		ReadFile(suite.ctx, req).
		Return(expected, nil)

	res := client.NewResult()
	res.SetRef(ref)
	suite.client.EXPECT().
		Solve(suite.ctx, gomock.Any()).
		Return(res, nil)

	// Act
	actual, err := suite.build.GetMetadata()

	// Assert
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), expected, actual)
}

func (suite *serviceTestSuite) TestGetExcludesSolveFails() {
	// Arrange
	suite.client.EXPECT().
		Solve(suite.ctx, gomock.Any()).
		Return(nil, errors.New("something went wrong"))

	// Act
	_, err := suite.build.GetExcludes()

	// Assert
	require.NotNil(suite.T(), err)
	require.Contains(suite.T(), err.Error(), "something went wrong")
}

func (suite *serviceTestSuite) TestGetExcludesRefFails() {
	// Arrange
	res := client.NewResult()
	res.AddRef("test", cib_mock.NewMockReference(suite.ctrl))
	suite.client.EXPECT().
		Solve(suite.ctx, gomock.Any()).
		Return(res, nil)

	// Act
	_, err := suite.build.GetExcludes()

	// Assert
	require.NotNil(suite.T(), err)
}

func (suite *serviceTestSuite) TestGetExcludesReadFails() {
	// Arrange
	req := client.ReadRequest{Filename: ".dockerignore"}
	ref := cib_mock.NewMockReference(suite.ctrl)
	ref.EXPECT().
		ReadFile(suite.ctx, req).
		Return(nil, errors.New("something went wrong"))

	res := client.NewResult()
	res.SetRef(ref)
	suite.client.EXPECT().
		Solve(suite.ctx, gomock.Any()).
		Return(res, nil)

	// Act
	excludes, err := suite.build.GetExcludes()

	// Assert
	require.Nil(suite.T(), err)
	require.Empty(suite.T(), excludes)
}

func (suite *serviceTestSuite) TestGetExcludesSucceeds() {
	// Arrange
	data := []byte("vendor")
	req := client.ReadRequest{Filename: ".dockerignore"}
	ref := cib_mock.NewMockReference(suite.ctrl)
	ref.EXPECT().
		ReadFile(suite.ctx, req).
		Return(data, nil)

	res := client.NewResult()
	res.SetRef(ref)
	suite.client.EXPECT().
		Solve(suite.ctx, gomock.Any()).
		Return(res, nil)

	// Act
	excludes, err := suite.build.GetExcludes()

	// Assert
	require.Nil(suite.T(), err)
	require.Contains(suite.T(), excludes, "vendor")
}

func (suite *serviceTestSuite) TestFetchImageConfigResolveFails() {
	// Arrange
	platform := &specs.Platform{}
	opt := llb.ResolveImageConfigOpt{
		Platform:    platform,
		ResolveMode: "local",
		LogName:     "load metadata for fakeimage",
	}
	expected := errors.New("something went wrong")
	suite.client.EXPECT().
		ResolveImageConfig(suite.ctx, "fakeimage", opt).
		Return(digest.Digest("sha256:fake"), nil, expected)

	// Act
	_, actual := suite.build.FetchImageConfig("fakeimage", platform)

	// Assert
	require.Equal(suite.T(), expected, actual)
}

func (suite *serviceTestSuite) TestFetchImageConfigUnmarshalFails() {
	// Arrange
	platform := &specs.Platform{}
	opt := llb.ResolveImageConfigOpt{
		Platform:    platform,
		ResolveMode: "local",
		LogName:     "load metadata for fakeimage",
	}
	data := []byte("not json")
	suite.client.EXPECT().
		ResolveImageConfig(suite.ctx, "fakeimage", opt).
		Return(digest.Digest("sha256:fake"), data, nil)

	// Act
	_, err := suite.build.FetchImageConfig("fakeimage", platform)

	// Assert
	require.NotNil(suite.T(), err)
}

func (suite *serviceTestSuite) TestFetchImageConfigSucceeds() {
	// Arrange
	platform := &specs.Platform{}
	opt := llb.ResolveImageConfigOpt{
		Platform:    platform,
		ResolveMode: "local",
		LogName:     "load metadata for fakeimage",
	}
	data := []byte(`{"Architecture": "amd64"}`)
	suite.client.EXPECT().
		ResolveImageConfig(suite.ctx, "fakeimage", opt).
		Return(digest.Digest("sha256:fake"), data, nil)

	// Act
	img, err := suite.build.FetchImageConfig("fakeimage", platform)

	// Assert
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), "amd64", img.Architecture)
}

func (suite *serviceTestSuite) TestGetBuildPlatform() {
	// Act
	platform := suite.build.GetBuildPlatform()

	// Assert
	require.Contains(suite.T(), platform.Architecture, "amd64")
	require.Equal(suite.T(), platform.OS, "linux")
	require.Equal(suite.T(), platform.OSVersion, "0.0.0")
}

func (suite *serviceTestSuite) TestGetTargetPlatformsImplicit() {
	// Arrange
	build := suite.build.GetBuildPlatform()

	// Act
	targets, err := suite.build.GetTargetPlatforms()

	// Assert
	require.Nil(suite.T(), err)
	require.Len(suite.T(), targets, 1)
	require.Equal(suite.T(), build, targets[0])
}

func (suite *serviceTestSuite) TestGetIgnoreCache() {
	// Act
	ignoreCache := suite.build.GetIgnoreCache()

	// Assert
	require.False(suite.T(), ignoreCache)
}

func (suite *serviceTestSuite) TestFromInvalidReference() {
	// Arrange
	platform := suite.build.GetBuildPlatform()

	// Act
	_, _, err := suite.build.From("nope::nope", platform, "")

	// Assert
	require.NotNil(suite.T(), err)
}

func (suite *serviceTestSuite) TestFromSucceeds() {
	// Arrange
	expected := suite.build.GetBuildPlatform()
	opt := llb.ResolveImageConfigOpt{
		Platform:    expected,
		ResolveMode: "local",
		LogName:     "load metadata for docker.io/library/fakeimage:latest",
	}
	data := []byte(`{"Config":{"User": "somebody","Env":["Key=Val"]}}`)
	suite.client.EXPECT().
		ResolveImageConfig(suite.ctx, "docker.io/library/fakeimage:latest", opt).
		Return(digest.Digest("sha256:fake"), data, nil)

	// Act
	state, img, err := suite.build.From("fakeimage", expected, "")

	// Assert
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), "somebody", img.Config.User)
	require.Contains(suite.T(), img.Config.Env, "Key=Val")

	actual, err := state.GetPlatform(suite.ctx)
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), expected.OS, actual.OS)
	require.Equal(suite.T(), expected.Architecture, actual.Architecture)

	value, exists, err := state.GetEnv(suite.ctx, "Key")
	require.True(suite.T(), exists)
	require.Nil(suite.T(), err)
	require.Equal(suite.T(), "Val", value)
}

func TestService(t *testing.T) {
	suite.Run(t, new(serviceTestSuite))
}

func TestGetMetadataFileNameCustom(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cli := cib_mock.NewMockClient(ctrl)
	cli.EXPECT().
		BuildOpts().
		Return(client.BuildOpts{
			Opts: map[string]string{
				keyFilename: "custom.yaml",
			},
		})
	build := NewService(context.Background(), cli)

	// Act
	name := build.GetMetadataFileName()

	// Assert
	require.Equal(t, "custom.yaml", name)
}

func TestGetTargetPlatformsExplicitFails(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cli := cib_mock.NewMockClient(ctrl)
	cli.EXPECT().
		BuildOpts().
		Return(client.BuildOpts{
			Opts: map[string]string{
				keyTargetPlatform: "vtx:cpu32",
			},
		})
	build := NewService(context.Background(), cli)

	// Act
	_, err := build.GetTargetPlatforms()

	// Assert
	require.NotNil(t, err)
}

func TestGetTargetPlatformsExplicitSucceeds(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cli := cib_mock.NewMockClient(ctrl)
	cli.EXPECT().
		BuildOpts().
		Return(client.BuildOpts{
			Opts: map[string]string{
				keyTargetPlatform: "linux/amd64",
			},
		})
	build := NewService(context.Background(), cli)

	// Act
	platforms, err := build.GetTargetPlatforms()

	// Assert
	require.Nil(t, err)
	require.Len(t, platforms, 1)
	require.Equal(t, platforms[0].Architecture, "amd64")
	require.Equal(t, platforms[0].OS, "linux")
}
