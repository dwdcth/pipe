// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package cron

import (
    "time"

    "github.com/88250/gulu"
)

// BlacklistIPs saves all banned IPs.
var BlacklistIPs map[string]bool

func refreshBlacklistIPsPeriodically() {
    go refreshBlacklistIPs()

    go func() {
        for range time.Tick(time.Minute * 30) {
            refreshBlacklistIPs()
        }
    }()
}

func refreshBlacklistIPs() {
    defer gulu.Panic.Recover(nil)

    BlacklistIPs = map[string]bool{}
}
