package examples_test

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func build(t *testing.T, tag string, context string) {
	// Arrange
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd := exec.Command(
		"docker",
		"build",
		"-t",
		tag,
		context,
	)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	cmd.Env = append(os.Environ(), "DOCKER_BUILDKIT=1")

	// Act
	err := cmd.Run()

	// Assert
	if err != nil {
		t.Log("stdout: ", stdout.String())
		t.Log("stderr: ", stderr.String())
		t.Log(err)
		t.Fail()
	}
}

const (
	tagTestFrontend = "erichripko/ubuntu-dockerfile"
	tagTestImage    = "erichripko/ubuntu-dockerfile-test"
)

func TestSimpleFrontend(t *testing.T) {
	// Arrange
	build(t, tagTestFrontend, "frontend")

	// Act
	build(t, tagTestImage, ".")

	// Assert
	out, err := exec.Command( // Verify that header was injected
		"docker",
		"run",
		tagTestImage,
		"cat",
		"/etc/lsb-release",
	).Output()
	require.Nil(t, err)
	require.Contains(t, string(out), "Ubuntu")

	out, err = exec.Command( // Verify that footer was injected
		"docker",
		"inspect",
		tagTestImage,
		"--format",
		"{{index .Config.Labels \"org.opencontainers.image.licenses\"}}",
	).Output()
	require.Nil(t, err)
	require.Equal(t, "Apache-2.0", strings.TrimSpace(string(out)))
}
