package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const data = `
{
    "date": "2013-11-06", 
    "dateLabel": "明日", 
    "telop": "晴れ", 
    "temperature": {
        "max": {
            "celsius": "19", 
            "fahrenheit": "66.2"
        }, 
        "min": {
            "celsius": "13", 
            "fahrenheit": "55.4"
        }
    }
}
`

type JsonData struct {
	Date        string      `json:"date"`
	Datelabel   string      `json:"dateLabel"`
	Telop       string      `json:"telop"`
	Temperature Temperature `json:"temperature"`
}

type Temperature struct {
	Max Max `json:"max"`
	Min Min `json:"min"`
}

type Max struct {
	Celsius    string `json:"celsius"`
	Fahrenheit string `json:"fahrenheit"`
}

type Min struct {
	Celsius    string `json:"celsius"`
	Fahrenheit string `json:"fahrenheit"`
}

func main() {
	jd := new(JsonData)
	err := jd.JsonProc()
	if err != nil {
		log.Fatalf("Log: %v", err)
		return
	}
	fmt.Printf("%v %v: %v  最高気温：%v  最低気温：%v\n", jd.Date, jd.Datelabel, jd.Telop, jd.Temperature.Max.Celsius, jd.Temperature.Min.Celsius)
}

func (p *JsonData) JsonProc() (err error) {
	err = json.Unmarshal([]byte(data), &p)
	if err != nil {
		return err
	}
	return nil
}
