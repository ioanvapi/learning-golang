package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	API  = "http://www.tuling123.com/openapi/api"
	TYPE = "application/json;charset=utf-8"
)

type RobotRequest struct {
	Info   string `json:"info"`
	Key    string `json:"key"`
	Loc    string `json:"loc"`
	UserID string `json:"userid"`
}

type RobotResponse struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func SendTuling(message, userId string) (*RobotResponse, error) {
	robotReq := &RobotRequest{
		Info:   message,
		Key:    Opts.TulingKey,
		UserID: userId,
	}
	b, _ := json.Marshal(robotReq)
	res, err := http.Post(API, TYPE, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	robotResp := &RobotResponse{}
	if err := json.Unmarshal(result, robotResp); err != nil {
		return nil, err
	}
	return robotResp, nil
}
