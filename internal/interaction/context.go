// Copyright © 2023 OpenIM SDK. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package interaction

import (
	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/constant"
	"time"
)

type ConnContext struct {
	RemoteAddr string
}

func (c *ConnContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c *ConnContext) Done() <-chan struct{} {
	return nil
}

func (c *ConnContext) Err() error {
	return nil
}

func (c *ConnContext) Value(key any) any {
	switch key {
	case constant.RemoteAddr:
		return c.RemoteAddr
	default:
		return ""
	}
}

func newContext(remoteAddr string) *ConnContext {
	return &ConnContext{
		RemoteAddr: remoteAddr,
	}
}
