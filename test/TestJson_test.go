package test

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"
)

type SpSearch struct {
	Starttime string    `json:"starttime"`
	Endtime   string    `json:"endtime"`
	Cursor    int       `json:"cursor"`
	Size      int       `json:"size"`
	Filters   []Filters `json:"filters"`
}

type Filters struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func TestJson(t *testing.T) {
	const BugTemplateId = "C4NwUa9KdrQd2GxFxyivw8WErcJnuQZe8VoPhwg7q"
	const SrmTemplateId = "C4Nx4VnkEfJFcePk9mVJhzNRj5VShZXTmjr62Cgoh"

	now := time.Now()

	bugSpSearch := SpSearch{
		Starttime: strconv.Itoa(int(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Unix())),
		Endtime:   strconv.Itoa(int(time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.Local).Unix())),
		Filters:   []Filters{{Key: "sp_status", Value: "2"}, {Key: "template_id", Value: BugTemplateId}},
	}
	srmSpSearch := SpSearch{
		Starttime: strconv.Itoa(int(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Unix())),
		Endtime:   strconv.Itoa(int(time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.Local).Unix())),
		Filters:   []Filters{{Key: "sp_status", Value: "2"}, {Key: "template_id", Value: SrmTemplateId}},
	}
	bugJson, _ := json.Marshal(bugSpSearch)
	srmJson, _ := json.Marshal(srmSpSearch)
	fmt.Printf("bugJson is : %s \n", bugJson)
	fmt.Printf("srmJson is : %s \n", srmJson)
}
