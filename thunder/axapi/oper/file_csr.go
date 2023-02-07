package oper

import (
	"errors"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

type FileCsr struct {
	SslCsr struct {
		Oper FileCsrList `json:"oper"`
	} `json:"csr"`
}

type FileCsrList struct {
	Items []FileCsrListItem `json:"ssl-csr"`
}

type FileCsrListItem struct {
	Name string `json:"name"`
}

func FileCsrGet(authToken string, host string, targetName string, logger *axapi.ThunderLog) error {
	logger.Println("oper.GetFileCsr")
	headers := axapi.GenRequestHeader(authToken)
	_, respPayload, err := axapi.SendGet(host, "file/csr/oper", "", nil, headers, logger)
	if err != nil {
		var m FileCsr
		err = json.Unmarshal(respPayload, &m)
		for _, item := range m.SslCsr.Oper.Items {
			if targetName == item.Name {
				return nil
			}
		}
		return errors.New("Not found: " + targetName)
	}
	return err
}
