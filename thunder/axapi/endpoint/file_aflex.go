package endpoint

//Operational

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

type FileAflex struct {
	Inst struct {
		Action string `json:"action,omitempty"`
		File   string `json:"file,omitempty"`
	} `json:"aflex,omitempty"`
}

func (p *FileAflex) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("FileAflex::Post")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := json.Marshal(p)
	if err != nil {
		logger.Println("json.Marshal() failed with error ", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	httpResp, axapiResp, err := axapi.SendPost(host, "file/aflex", payloadBytes, headers, logger)
	if httpResp == nil {
		logger.Println("HTTP request failed with error ", err)
	} else {
		logger.Println("HTTP status: " + httpResp.Status)
		logger.Println("axAPI response payload: " + string(axapiResp))
	}
	return err
}
