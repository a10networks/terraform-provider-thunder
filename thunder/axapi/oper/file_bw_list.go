package oper

import (
	"errors"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

type FileBwList struct {
	BwList struct {
		Oper FileBwListList `json:"oper"`
	} `json:"bw-list"`
}

type FileBwListList struct {
	Items []FileBwListListItem `json:"file-list"`
}

type FileBwListListItem struct {
	Name string `json:"file"`
}

func FileBwListGet(authToken string, host string, fileName string, logger *axapi.ThunderLog) error {
	logger.Println("oper.GetFileBwList")
	headers := axapi.GenRequestHeader(authToken)
	_, respPayload, err := axapi.SendGet(host, "file/bw-list/oper", "", nil, headers, logger)
	if err != nil {
		var m FileBwList
		err = json.Unmarshal(respPayload, &m)
		for _, item := range m.BwList.Oper.Items {
			if fileName == item.Name {
				return nil
			}
		}
		return errors.New("Not found: " + fileName)
	}
	return err
}
