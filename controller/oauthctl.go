// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package controller

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

var states = map[string]string{}

// redirectLoginAction redirects to HacPai auth page.
func redirectLoginAction(c *gin.Context) {
    c.JSON(http.StatusNotImplemented, gin.H{"error": "External login is disabled"})
}

func loginCallbackAction(c *gin.Context) {
    c.JSON(http.StatusNotImplemented, gin.H{"error": "External login callback is disabled"})
}
