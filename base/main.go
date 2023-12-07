//go:generate goversioninfo

package main

import (
	"os"
	"time"
)

func main() {
	path, err := os.Executable()
	if err != nil {
		verboseFatal(err)
	}

	verbosef("Executable path: %s", path)

	closer, config, reader, err := openExeContent(path)
	if err != nil {
		verboseFatal(err)
	}

	defer closer.Close()

	destConfig, err := readDestConfig(config)
	if err != nil {
		verboseFatal(err)
	}

	mode := getInstallMode(config, destConfig)
	if mode == modOUTDATED {
		verbosef("SFX version (%s) is lower than installed version (%s). Running installed configuration (%s)...", config.Version, destConfig.Version, destConfig.Run)
		err = run(destConfig)
		if err != nil {
			verboseFatal(err)
		}

		return
	}

	exeModTime := time.Unix(0, 0)

	verbosef("Uncompressing resources to: %s", config.Dest)
	err = uncompress(reader, config, exeModTime)
	if err != nil {
		verboseFatal(err)
	}

	if mode >= modUPDATE {
		err = writeDestConfig(config)
		if err != nil {
			verboseFatal(err)
		}
	}

	verbosef("Running %s...", config.Run)
	err = run(config)
	if err != nil {
		verboseFatal(err)
	}
}

func getExeModTime(path string) (time.Time, error) {
	info, err := os.Stat(path)
	if err != nil {
		return time.Unix(0, 0), err
	}

	return info.ModTime(), nil
}
