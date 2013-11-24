package main

import (
    "fmt"
    "io/ioutil"
    "launchpad.net/goyaml"
)

type Config struct {
    ServerName string
    Port       string
}

var conf Config


func main() {

    file, err := ioutil.ReadFile("config.yml")
    if err != nil {
        panic(err)
    }
    err = goyaml.Unmarshal([]byte(file), &conf)
    if err != nil {
        panic(err)
    }

    fmt.Printf("ServerName:%s\nPort:%s", conf.ServerName, conf.Port)

}
