package version

import "runtime/debug"

func Get() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "dev"
	}

	if info.Main.Version != "" {
		return info.Main.Version
	}

	return "dev"
}