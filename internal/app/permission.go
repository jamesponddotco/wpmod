package app

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"

	"git.sr.ht/~jamesponddotco/wpmod/internal/find"
	"git.sr.ht/~jamesponddotco/wpmod/internal/xos"
	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"github.com/urfave/cli/v2"
)

const (
	// ErrInvalidPath is returned when the path given by the user is not the
	// path to a valid WordPress install.
	ErrInvalidPath xerrors.Error = "path doesn't contain a valid WordPress install"

	// ErrInvalidUser is returned when the user given by the user is invalid.
	ErrInvalidUser xerrors.Error = "failed to look up user"

	// ErrInvalidGroup is returned when the group given by the user is invalid.
	ErrInvalidGroup xerrors.Error = "failed to look up group"
)

const (
	_wpConfigFilename       string = "wp-config.php"
	_wpConfigSampleFilename string = "wp-config-sample.php"
)

// PermissionAction is the main action for the application.
func PermissionAction(ctx *cli.Context) error {
	var (
		path         = ctx.String("path")
		wpConfigPath = filepath.Join(path, _wpConfigFilename)
		owner        = ctx.String("user")
		group        = ctx.String("group")
		strict       = ctx.Bool("strict")
	)

	if _, err := os.Stat(wpConfigPath); err != nil {
		wpConfigSamplePath := filepath.Join(path, _wpConfigSampleFilename)

		if _, err = os.Stat(wpConfigSamplePath); err != nil {
			return fmt.Errorf("%w %q: %w", ErrInvalidPath, path, err)
		}

		wpConfigPath = ""
	}

	u, err := user.Lookup(owner)
	if err != nil {
		return fmt.Errorf("%w %q: %w", ErrInvalidUser, owner, err)
	}

	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		return fmt.Errorf("%w %q: %w", ErrInvalidUser, owner, err)
	}

	g, err := user.LookupGroup(group)
	if err != nil {
		return fmt.Errorf("%w %q: %w", ErrInvalidGroup, group, err)
	}

	gid, err := strconv.Atoi(g.Gid)
	if err != nil {
		return fmt.Errorf("%w %q: %w", ErrInvalidGroup, group, err)
	}

	if err := xos.Rchown(path, uid, gid); err != nil {
		return fmt.Errorf("%w", err)
	}

	var (
		fileMode = os.FileMode(0o644)
		dirMode  = os.FileMode(0o755)
	)

	if err := xos.Rchmod(path, fileMode, dirMode); err != nil {
		return fmt.Errorf("%w", err)
	}

	if strict { //nolint:nestif // what other way is there?
		strictMode := os.FileMode(0o400)

		if wpConfigPath != "" {
			if err := os.Chmod(wpConfigPath, strictMode); err != nil {
				return fmt.Errorf("%w", err)
			}
		}

		// We need to find the path to the "mu-plugins" directory, as the user
		// may have changed the path to the "wp-content" directory.
		muPluginsPath, err := find.MuPluginsDirectory(path)
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		if muPluginsPath != "" {
			if err := os.Chmod(muPluginsPath, strictMode); err != nil {
				return fmt.Errorf("%w", err)
			}
		}
	}

	return nil
}
