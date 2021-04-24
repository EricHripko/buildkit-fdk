package cib

import (
	"context"
	"os"
	"path"

	"github.com/moby/buildkit/frontend/gateway/client"
	fsutil "github.com/tonistiigi/fsutil/types"
)

// WalkFunc is the type of function called for each file or directory visited
// by WalkRecursive.
type WalkFunc func(file *fsutil.Stat) error

// WalkRecursive iterates all the files in the reference recursively.
func WalkRecursive(ctx context.Context, ref client.Reference, walkFn WalkFunc) error {
	return walkRecursive(ctx, ref, ".", walkFn)
}

func walkRecursive(ctx context.Context, ref client.Reference, root string, walkFn WalkFunc) error {
	files, err := ref.ReadDir(ctx, client.ReadDirRequest{Path: root})
	if err != nil {
		return err
	}

	for _, file := range files {
		// Make path absolute for easier integration
		file.Path = path.Join(root, file.Path)

		// Callback
		err = walkFn(file)
		if err != nil {
			return err
		}

		// Walk folders
		mode := os.FileMode(file.Mode)
		if mode.IsDir() {
			err = walkRecursive(ctx, ref, file.Path, walkFn)
			if err != nil {
				return err
			}
		}
	}
	return nil

}
