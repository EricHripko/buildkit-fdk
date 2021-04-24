package cib

import (
	"context"
	"errors"
	"os"
	"testing"

	cib_mock "github.com/EricHripko/buildkit-fdk/pkg/cib/mock"

	"github.com/golang/mock/gomock"
	"github.com/moby/buildkit/frontend/gateway/client"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	fsutil "github.com/tonistiigi/fsutil/types"
)

func walkFnNoop(*fsutil.Stat) error {
	return nil
}

type walkRecursiveTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller
	ctx  context.Context
	ref  *cib_mock.MockReference
}

func (suite *walkRecursiveTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.ctx = context.Background()
	suite.ref = cib_mock.NewMockReference(suite.ctrl)
}

func (suite *walkRecursiveTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *walkRecursiveTestSuite) TestReadDirFails() {
	// Arrange
	actual := errors.New("something went wrong")
	req := client.ReadDirRequest{Path: "."}
	suite.ref.EXPECT().
		ReadDir(suite.ctx, req).
		Return(nil, actual)

	// Act
	expected := WalkRecursive(suite.ctx, suite.ref, walkFnNoop)

	// Assert
	require.Same(suite.T(), expected, actual)
}

func (suite *walkRecursiveTestSuite) TestWalkFnFails() {
	// Arrange
	req := client.ReadDirRequest{Path: "."}
	files := []*fsutil.Stat{
		{Path: "README.md"},
	}
	suite.ref.EXPECT().
		ReadDir(suite.ctx, req).
		Return(files, nil)
	actual := errors.New("something went wrong")
	walkFn := func(*fsutil.Stat) error { return actual }

	// Act
	expected := WalkRecursive(suite.ctx, suite.ref, walkFn)

	// Assert
	require.Same(suite.T(), expected, actual)
}

func (suite *walkRecursiveTestSuite) TestNestedFails() {
	// Arrange
	req := client.ReadDirRequest{Path: "."}
	files := []*fsutil.Stat{
		{Path: "vendor", Mode: uint32(os.ModeDir)},
	}
	suite.ref.EXPECT().
		ReadDir(suite.ctx, req).
		Return(files, nil)
	req = client.ReadDirRequest{Path: "vendor"}
	actual := errors.New("something went wrong")
	suite.ref.EXPECT().
		ReadDir(suite.ctx, req).
		Return(nil, actual)

	// Act
	expected := WalkRecursive(suite.ctx, suite.ref, walkFnNoop)

	// Assert
	require.Same(suite.T(), expected, actual)
}

func (suite *walkRecursiveTestSuite) TestSucceeds() {
	// Arrange
	req := client.ReadDirRequest{Path: "."}
	files := []*fsutil.Stat{}
	suite.ref.EXPECT().
		ReadDir(suite.ctx, req).
		Return(files, nil)

	// Act
	err := WalkRecursive(suite.ctx, suite.ref, walkFnNoop)

	// Assert
	require.Nil(suite.T(), err)
}

func TestWalkRecursive(t *testing.T) {
	suite.Run(t, new(walkRecursiveTestSuite))
}
