/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Deprecated: use options to configure logging
package logging

import (
	"github.com/rs/zerolog"
	globalLog "github.com/rs/zerolog/log"
)

var oldLogger zerolog.Logger

// Deprecated: use config.WithCustomLogger
// init is used for _ imports for easy globalLog config
func init() {
	oldLogger = globalLog.Logger
	globalLog.Logger = globalLog.Logger.Level(zerolog.ErrorLevel)
}

// Deprecated: use config.WithCustomLogger
// ErrorLevel configures zerolog to WarnLevel
func ErrorLevel() {
	globalLog.Logger = globalLog.Logger.Level(zerolog.ErrorLevel)
}

// Deprecated: use config.WithCustomLogger
// WarnLevel configures zerolog to WarnLevel
func WarnLevel() {
	globalLog.Logger = globalLog.Logger.Level(zerolog.WarnLevel)
}

// Deprecated: use config.WithCustomLogger
// InfoLevel configures zerolog to InfoLevel
func InfoLevel() {
	globalLog.Logger = globalLog.Logger.Level(zerolog.InfoLevel)
}

// Deprecated: use config.WithCustomLogger
// DebugLevel configures zerolog to DebugLevel
func DebugLevel() {
	globalLog.Logger = globalLog.Logger.Level(zerolog.DebugLevel)
}

// Deprecated: use config.WithCustomLogger
// TraceLevel configures zerolog to TraceLevel
func TraceLevel() {
	globalLog.Logger = globalLog.Logger.Level(zerolog.TraceLevel)
}

// Deprecated: use config.WithCustomLogger
// ResetLogging can be used to reset to the old globalLog settings
func ResetLogging() {
	globalLog.Logger = oldLogger
}
