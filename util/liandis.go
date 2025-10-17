// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package util

import (
    "crypto/tls"
    "net"
    "net/http"
    "os"

    "github.com/88250/gulu"
    "github.com/gin-gonic/gin"
)

// Logger
var logger = gulu.Log.NewLogger(os.Stdout)

// CommunityURL is the URL of LianDi community.
const CommunityURL = ""

// CommunityAPI is a reverse proxy.
func CommunityAPI() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{"error": "Community API is disabled"})
    }
}

func dialTLS(network, addr string) (net.Conn, error) {
    conn, err := net.Dial(network, addr)
    if err != nil {
        return nil, err
    }

    host, _, err := net.SplitHostPort(addr)
    if err != nil {
        return nil, err
    }
    cfg := &tls.Config{ServerName: host}

    tlsConn := tls.Client(conn, cfg)
    if err := tlsConn.Handshake(); err != nil {
        conn.Close()
        return nil, err
    }

    cs := tlsConn.ConnectionState()
    cert := cs.PeerCertificates[0]

    cert.VerifyHostname(host)

    return tlsConn, nil
}

// HacPaiUserInfo returns HacPai community user info specified by the given access token.
func HacPaiUserInfo(accessToken string) (ret map[string]interface{}) {
    return nil
}
