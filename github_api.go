// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "fmt"
    "code.google.com/p/goauth2/oauth"
    "github.com/google/go-github/github"
)

func main() {
    fmt.Println("Recently updated repositories owned by user ramtiga:")

    t := &oauth.Transport{
        Token: &oauth.Token{AccessToken: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx"},
    }

    client := github.NewClient(t.Client())
    opt := &github.RepositoryListOptions{Type: "owner", Sort: "updated", Direction: "desc"}

    repos, _, err := client.Repositories.List("", opt)
    if err != nil {
        fmt.Printf("error: %v\n\n", err)
    } else {
        fmt.Printf("%v\n\n", github.Stringify(repos))
    }
}
