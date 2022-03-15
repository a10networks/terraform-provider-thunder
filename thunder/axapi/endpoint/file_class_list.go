package endpoint

//Operational

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

type FileClassList struct {
	Inst struct {
		Action string `json:"action,omitempty"`
		File   string `json:"file,omitempty"`
	} `json:"class-list,omitempty"`
}

func (p *FileClassList) getPath() string {
	return "file/class-list"
}

func (p *FileClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FileClassList::Post")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := json.Marshal(p)
	if err != nil {
		logger.Println("json.Marshal() failed with error ", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	httpResp, axapiResp, err := axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
	if httpResp == nil {
		logger.Println("HTTP request failed with error ", err)
	} else {
		logger.Println("HTTP status: " + httpResp.Status)
		logger.Println("axAPI response payload: " + string(axapiResp))
	}
	return err
}

func (p *FileClassList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("FileClassList::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
