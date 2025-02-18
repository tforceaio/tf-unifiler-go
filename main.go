// Copyright (C) 2024 T-Force I/O
// This file is part of TF Unifiler
//
// TF Unifiler is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// TF Unifiler is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with TF Unifiler. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/rs/zerolog/log"
	"github.com/tforceaio/tf-unifiler-go/cmd"
	"github.com/tforceaio/tf-unifiler-go/cmd/config"
	"github.com/tforceaio/tf-unifiler-go/diag"
	"github.com/tforceaio/tf-unifiler-go/extension/generic"
	"github.com/tforceaio/tf-unifiler-go/filesystem"
	"github.com/tforceaio/tf-unifiler-go/filesystem/exec"
)

var invokeArgs cmd.Args
var majorVersion = 0
var minorVersion = 4
var patchVersion = 0
var gitCommit, gitDate, gitBranch string

func version() string {
	originDate := time.Date(2024, time.August, 13, 0, 0, 0, 0, time.UTC)
	gitDate2, _ := time.Parse("20060102", gitDate)
	buildDate := generic.TernaryAssign(gitDate == "", time.Now().UTC(), gitDate2)
	duration := buildDate.Sub(originDate)
	minor := minorVersion
	patch := strconv.Itoa(patchVersion)
	if gitBranch == "master" {
		// do nothing
	} else if gitBranch == "release" {
		minor += 1
		patch = patch + "-rc"
	} else if strings.Contains(gitBranch, "feat/") {
		minor += 1
		patch = patch + "-dev"
	} else {
		patch = strconv.Itoa(patchVersion+1) + "-dev"
	}
	if gitCommit != "" {
		return fmt.Sprintf("%d.%d.%s.%d-%s", majorVersion, minor, patch, duration.Milliseconds()/int64(86400000), gitCommit)
	}
	return fmt.Sprintf("%d.%d.%s.%d", majorVersion, minor, patch, duration.Milliseconds()/int64(86400000))
}

func main() {
	cfg, cfgErr := config.InitKoanf()
	logFile := diag.InitZerolog(cfg.ConfigDir)
	if logFile != nil {
		defer logFile.Close()
	}
	filesystem.SetLogger(diag.GetModuleLogger("filesystem"))
	exec.SetLogger(diag.GetModuleLogger("exec"))

	invokeArgs = cmd.Args{}
	arg.MustParse(&invokeArgs)
	pwd, _ := os.Getwd()
	pwd, _ = filesystem.GetAbsPath(pwd)
	exec, _ := os.Executable()
	exec, _ = filesystem.GetAbsPath(exec)

	log.Info().Msgf("TF UNIFILER v%s", version())
	gitDate2, _ := time.Parse("20060102", gitDate)
	buildDate := generic.TernaryAssign(gitDate == "", time.Now().UTC(), gitDate2)
	log.Info().Msgf("Copyright (C) %d T-Force I/O", buildDate.Year())
	log.Info().Msgf("Licensed under GPL-3.0 license. See COPYING file along with this program for more details.")
	log.Info().Msgf("Working directory %s", pwd)
	log.Info().Msgf("Config directory %s", cfg.ConfigDir)
	log.Info().Msgf("Executable file %s", exec)
	log.Info().Msgf("Portable mode %t", cfg.IsPortable)
	if cfgErr != nil {
		log.Err(cfgErr).Str("configFile", cfg.ConfigFile).Msg("Error initialize configuration file")
		return
	}

	if invokeArgs.File != nil {
		m := FileModule{
			logger: diag.GetModuleLogger("file"),
		}
		m.File(invokeArgs.File)
	}
	if invokeArgs.Hash != nil {
		m := HashModule{
			logger: diag.GetModuleLogger("hash"),
		}
		m.Hash(invokeArgs.Hash)
	}
	if invokeArgs.Mirror != nil {
		m := MirrorModule{
			logger: diag.GetModuleLogger("mirror"),
		}
		m.Mirror(invokeArgs.Mirror)
	}
	if invokeArgs.Video != nil {
		m := VideoModule{
			cfg:    cfg,
			logger: diag.GetModuleLogger("video"),
		}
		m.Video(invokeArgs.Video)
	}
}
