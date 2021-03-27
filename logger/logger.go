// Copyright 2021 Herman Slatman
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	"fmt"
	"time"
)

type Logger interface {
	Log(v ...interface{})
	Logf(format string, v ...interface{})
	Info(message string)
	Debug(message string)
	Warning(message string)
	Error(message string)
}

type NullLogger struct {
}

type ExampleLogger struct {
}

func (nl NullLogger) Log(v ...interface{}) {}

func (nl NullLogger) Logf(format string, v ...interface{}) {}

func (nl NullLogger) Info(message string) {}

func (nl NullLogger) Debug(message string) {}

func (nl NullLogger) Warning(message string) {}

func (nl NullLogger) Error(message string) {}

func (l ExampleLogger) Log(v ...interface{}) {
	fmt.Println(v[0])
}

func (l ExampleLogger) Logf(format string, v ...interface{}) {
	fmt.Println(fmt.Sprintf(format, v...))
}

func (l ExampleLogger) Info(message string) {
	l.Log(fmt.Sprint(time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + "  [INFO] " + message))
}

func (l ExampleLogger) Debug(message string) {
	l.Log(fmt.Sprint(time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + " [DEBUG] " + message))
}

func (l ExampleLogger) Warning(message string) {
	l.Log(fmt.Sprint(time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + " [WARN] " + message))
}

func (l ExampleLogger) Error(message string) {
	l.Log(fmt.Sprint(time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + " [ERROR] " + message))
}
