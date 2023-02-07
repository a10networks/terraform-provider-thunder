package oper

import (
	"errors"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

type FileCaCert struct {
	CaCert struct {
		Oper FileCaCertList `json:"oper"`
	} `json:"ca-cert"`
}

type FileCaCertList struct {
	Items []FileCaCertListItem `json:"file-list"`
}

type FileCaCertListItem struct {
	Name string `json:"file"`
}

func FileCaCertGet(authToken string, host string, targetName string, logger *axapi.ThunderLog) error {
	logger.Println("oper.GetFileCaCert")
	headers := axapi.GenRequestHeader(authToken)
	_, respPayload, err := axapi.SendGet(host, "file/ca-cert/oper", "", nil, headers, logger)
	if err != nil {
		var m FileCaCert
		err = json.Unmarshal(respPayload, &m)
		for _, item := range m.CaCert.Oper.Items {
			if targetName == item.Name {
				return nil
			}
		}
		return errors.New("Not found: " + targetName)
	}
	return err
}
