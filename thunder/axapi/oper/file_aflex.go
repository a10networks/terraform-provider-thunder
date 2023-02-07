package oper

import (
	"errors"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

type FileAflex struct {
	Aflex struct {
		Oper FileAflexList `json:"oper"`
	} `json:"aflex"`
}

type FileAflexList struct {
	Items []FileAflexListItem `json:"file-list"`
}

type FileAflexListItem struct {
	Name string `json:"file"`
}

func FileAflexGet(authToken string, host string, aflexName string, logger *axapi.ThunderLog) error {
	logger.Println("oper.GetFileAflex")
	headers := axapi.GenRequestHeader(authToken)
	_, respPayload, err := axapi.SendGet(host, "file/aflex/oper", "", nil, headers, logger)
	if err != nil {
		var m FileAflex
		err = json.Unmarshal(respPayload, &m)
		for _, item := range m.Aflex.Oper.Items {
			if aflexName == item.Name {
				return nil
			}
		}
		return errors.New("Not found: " + aflexName)
	}
	return err
}
