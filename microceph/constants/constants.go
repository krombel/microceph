// Package common
package constants

import (
	"os"
	"path/filepath"
)

// Constants for Size Constraints
const MinOSDSize = uint64(2147483648) // 2GB i.e. 2*1024*1024*1024

const ClientConfigGlobalHostConst = "*"
const BootstrapPortConst = 7443

// string templates
const LoopSpecId = "loop,"
const DevicePathPrefix = "/dev/disk/by-id/"
const CliForcePrompt = "If you understand the *RISK* and you're *ABSOLUTELY CERTAIN* that is what you want, pass --yes-i-really-mean-it."

type PathConst struct {
	ConfPath string
	RunPath  string
	DataPath string
	LogPath  string
	ProcPath string
}

type PathFileMode map[string]os.FileMode

func GetPathConst() PathConst {
	return PathConst{
		ConfPath: filepath.Join(os.Getenv("SNAP_DATA"), "conf"),
		RunPath:  filepath.Join(os.Getenv("SNAP_DATA"), "run"),
		DataPath: filepath.Join(os.Getenv("SNAP_COMMON"), "data"),
		LogPath:  filepath.Join(os.Getenv("SNAP_COMMON"), "logs"),
		ProcPath: filepath.Join(os.Getenv("TEST_ROOT_PATH"), "/proc"),
	}
}

func GetPathFileMode() PathFileMode {
	pathConsts := GetPathConst()
	return PathFileMode{
		pathConsts.ConfPath: 0750,
		pathConsts.RunPath:  0700,
		pathConsts.DataPath: 0700,
		pathConsts.LogPath:  0700,
	}
}
