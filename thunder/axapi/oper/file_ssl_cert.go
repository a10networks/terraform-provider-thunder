package oper

import (
	"errors"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

type FileSslCert struct {
	SslCert struct {
		Oper FileSslCertList `json:"oper"`
	} `json:"ssl-cert"`
}

type FileSslCertList struct {
	Items []FileSslCertListItem `json:"file-list"`
}

type FileSslCertListItem struct {
	Name string `json:"file"`
}

func FileSslCertGet(authToken string, host string, targetName string, logger *axapi.ThunderLog) error {
	logger.Println("oper.GetFileSslCert")
	headers := axapi.GenRequestHeader(authToken)
	_, respPayload, err := axapi.SendGet(host, "file/ssl-cert/oper", "", nil, headers, logger)
	if err != nil {
		var m FileSslCert
		err = json.Unmarshal(respPayload, &m)
		for _, item := range m.SslCert.Oper.Items {
			if targetName == item.Name {
				return nil
			}
		}
		return errors.New("Not found: " + targetName)
	}
	return err
}
