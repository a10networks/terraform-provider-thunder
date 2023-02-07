package oper

import (
	"errors"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

type FileClassList struct {
	ClassList struct {
		Oper FileClassListList `json:"oper"`
	} `json:"class-list"`
}

type FileClassListList struct {
	Items []FileClassListListItem `json:"file-list"`
}

type FileClassListListItem struct {
	Name string `json:"file"`
}

func FileClassListGet(authToken string, host string, aflexName string, logger *axapi.ThunderLog) error {
	logger.Println("oper.GetFileClassList")
	headers := axapi.GenRequestHeader(authToken)
	_, respPayload, err := axapi.SendGet(host, "file/class-list/oper", "", nil, headers, logger)
	if err != nil {
		var m FileClassList
		err = json.Unmarshal(respPayload, &m)
		for _, item := range m.ClassList.Oper.Items {
			if aflexName == item.Name {
				return nil
			}
		}
		return errors.New("Not found: " + aflexName)
	}
	return err
}
