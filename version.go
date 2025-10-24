package version

import "runtime/debug"

func Version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "dev"
	}

	if info.Main.Version != "" {
		return info.Main.Version
	}

	return "dev"
}