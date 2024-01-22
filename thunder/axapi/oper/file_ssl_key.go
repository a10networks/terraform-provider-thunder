package oper

import (
	"errors"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

type FileSslKey struct {
	SslKey struct {
		Oper FileSslKeyList `json:"oper"`
	} `json:"ssl-key"`
}

type FileSslKeyList struct {
	Items []FileSslKeyListItem `json:"file-list"`
}

type FileSslKeyListItem struct {
	Name string `json:"file"`
}

func FileSslKeyGet(authToken string, host string, targetName string, logger *axapi.ThunderLog) error {
	logger.Println("oper.GetFileSslKey")
	headers := axapi.GenRequestHeader(authToken)
	_, respPayload, err := axapi.SendGet(host, "file/ssl-key/oper", "", nil, headers, logger)
	if err != nil {
		var m FileSslKey
		err = json.Unmarshal(respPayload, &m)
		for _, item := range m.SslKey.Oper.Items {
			if targetName == item.Name {
				return nil
			}
		}
		return errors.New("Not found: " + targetName)
	}
	return err
}
