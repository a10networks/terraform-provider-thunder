package oper

import (
	"errors"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

type FileSslCrl struct {
	SslCrl struct {
		Oper FileSslCrlList `json:"oper"`
	} `json:"ssl-crl"`
}

type FileSslCrlList struct {
	Items []FileSslCrlListItem `json:"file-list"`
}

type FileSslCrlListItem struct {
	Name string `json:"file"`
}

func FileSslCrlGet(authToken string, host string, targetName string, logger *axapi.ThunderLog) error {
	logger.Println("oper.GetFileSslCrl")
	headers := axapi.GenRequestHeader(authToken)
	_, respPayload, err := axapi.SendGet(host, "file/ssl-crl/oper", "", nil, headers, logger)
	if err != nil {
		var m FileSslCrl
		err = json.Unmarshal(respPayload, &m)
		for _, item := range m.SslCrl.Oper.Items {
			if targetName == item.Name {
				return nil
			}
		}
		return errors.New("Not found: " + targetName)
	}
	return err
}
