// Copyright (C) 2024 T-Force I/O
// This file is part of TFunifiler Build library.
//
// TFunifiler Build library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// TFunifiler Build library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with TFunifiler Build library. If not, see <http://www.gnu.org/licenses/>.
//
// This file is a modified version of The go-ethereum library, Copyright 2016 The go-ethereum Authors

package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
)

var (
	// These flags override values in build env.
	GitCommitFlag   = flag.String("git-commit", "", `Overrides git commit hash embedded into executables`)
	GitBranchFlag   = flag.String("git-branch", "", `Overrides git branch being built`)
	GitTagFlag      = flag.String("git-tag", "", `Overrides git tag being built`)
	BuildNumFlag    = flag.String("build-num", "", `Overrides CI build number`)
	PullRequestFlag = flag.Bool("pull-request", false, `Overrides pull request status of the build`)
)

type Environment struct {
	Name          string // name of the environment
	Repo          string // name of GitHub repo
	Commit        string // git info
	Date          string // git info
	Branch        string // git info
	Tag           string // git info
	Platform      string // os info (GOOS)
	Architecture  string // os info (GOARCH)
	BuildNum      string // ci info
	IsPullRequest bool   // ci info
}

func (env Environment) String() string {
	return fmt.Sprintf("%s env (repo: %s commit:%s date:%s branch:%s tag:%s build-num:%s pull-request:%t)",
		env.Repo, env.Name, env.Commit, env.Date, env.Branch, env.Tag, env.BuildNum, env.IsPullRequest)
}

func Env() Environment {
	env := applyEnvFlags(Environment{Name: "local", Repo: "tforceaio/tf-unifiler-go"})
	env.Platform = runtime.GOOS
	env.Architecture = runtime.GOARCH

	head := readGitFile("HEAD")
	if fields := strings.Fields(head); len(fields) == 2 {
		head = fields[1]
	} else {
		// In this case we are in "detached head" state
		// see: https://git-scm.com/docs/git-checkout#_detached_head
		// Additional check required to verify, that file contains commit hash
		commitRe, _ := regexp.Compile("^([0-9a-f]{40})$")
		if commit := commitRe.FindString(head); commit != "" && env.Commit == "" {
			env.Commit = commit
			env.Date = getDate(env.Commit)
		}
		return env
	}
	if env.Commit == "" {
		env.Commit = readGitFile(head)
	}
	env.Date = getDate(env.Commit)
	if env.Branch == "" {
		if head != "HEAD" {
			env.Branch = strings.TrimPrefix(head, "refs/heads/")
		}
	}
	if info, err := os.Stat(".git/objects"); err == nil && info.IsDir() && env.Tag == "" {
		env.Tag = firstLine(RunGit("tag", "-l", "--points-at", "HEAD"))
	}
	return env
}

func applyEnvFlags(env Environment) Environment {
	if !flag.Parsed() {
		panic("you need to call flag.Parse before Env or LocalEnv")
	}
	if *GitCommitFlag != "" {
		env.Commit = *GitCommitFlag
	}
	if *GitBranchFlag != "" {
		env.Branch = *GitBranchFlag
	}
	if *GitTagFlag != "" {
		env.Tag = *GitTagFlag
	}
	if *BuildNumFlag != "" {
		env.BuildNum = *BuildNumFlag
	}
	if *PullRequestFlag {
		env.IsPullRequest = true
	}
	return env
}

func firstLine(s string) string {
	return strings.Split(s, "\n")[0]
}
