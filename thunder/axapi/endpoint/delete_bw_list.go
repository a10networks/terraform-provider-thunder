package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

//based on ACOS 5_2_1-P3_70
type DeleteBwList struct {
	Inst struct {
		FileName string `json:"file-name"`
	} `json:"bw-list"`
}

func (p *DeleteBwList) GetId() string {
	return "1"
}

func (p *DeleteBwList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DeleteBwList::Post")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := json.Marshal(p)
	if err != nil {
		logger.Println("json.Marshal() failed with error", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, "delete/bw-list", payloadBytes, headers, logger)
	return err
}
