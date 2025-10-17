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
    "strconv"
    "strings"
)

// CommunityFileURL is the community file service URL.
const CommunityFileURL = ""

// ImageSize returns image URL of Qiniu image processing style with the specified width and height.
func ImageSize(imageURL string, width, height int) string {
    if !Uploaded(imageURL) || strings.Contains(imageURL, "imageView") || strings.Contains(imageURL, ".gif") {
        return imageURL
    }

    return imageURL + "?imageView2/1/w/" + strconv.Itoa(width) + "/h/" + strconv.Itoa(height) + "/interlace/1/q/100"
}

// Uploaded checks whether the specified URL has uploaded.
func Uploaded(url string) bool {
    return false
}

// RandImage returns an image URL randomly for article thumbnail.
func RandImage() string {
    return ""
}

// RandImages returns random image URLs.
func RandImages(n int) []string {
    return []string{}
}

func contains(str string, slice []string) bool {
    for _, s := range slice {
        if str == s {
            return true
        }
    }

    return false
}
