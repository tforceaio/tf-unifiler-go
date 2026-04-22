// Copyright (C) 2025 T-Force I/O
// This file is part of TFunifiler
//
// TFunifiler is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// TFunifiler is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with TFunifiler. If not, see <https://www.gnu.org/licenses/>.

package engine

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/tforceaio/tf-unifiler/config"
)

// Controller is the entrypoint for working with application configurations and
// loggings.
type Controller struct {
	Root   *config.RootConfig
	Logger zerolog.Logger

	logFile *os.File
}

// Entrypoint for creating new instance of Controller.
// useFS will instruct this function to read configurations and create log file.
func NewController(useFS bool) *Controller {
	cfg, err := config.InitKoanf(useFS)
	logger, logFile, err2 := config.InitZerolog(cfg.ConfigDir, useFS)
	if err != nil {
		logger.Err(err).Msg("error initializing config")
	}
	if err2 != nil {
		logger.Err(err2).Msg("error initializing log file")
	}
	return &Controller{
		Root:   cfg,
		Logger: logger,

		logFile: logFile,
	}
}

// Execute additional clean up when terminate the app.
func (c *Controller) Close() {
	if c.logFile != nil {
		c.logFile.Close()
		c.logFile = nil
	}
}

// Get a ZeroLog logger instance for command handler from root instance.
func (c *Controller) CommandLogger(module, command string) zerolog.Logger {
	return c.Logger.With().Str("module", module).Str("command", command).Logger()
}

// Get a ZeroLog logger instance for module from root instance.
func (c *Controller) ModuleLogger(module string) zerolog.Logger {
	return c.Logger.With().Str("module", module).Logger()
}
