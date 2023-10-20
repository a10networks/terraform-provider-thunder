package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_0-P1_10
type ClockShowOper struct {
	Oper ClockShowOperOper `json:"oper"`
}

type ClockShoww struct {
	DataClockShow ClockShowOper `json:"show"`
}

type ClockShowOperOper struct {
	SourceType int    `json:"source-type"`
	Time       string `json:"time"`
	Timezone   string `json:"timezone"`
	Offset     string `json:"offset"`
	Day        string `json:"day"`
	Date       string `json:"date"`
}

func (p *ClockShowOper) GetId() string {
	return "1"
}

func (p *ClockShowOper) getPath() string {
	return "clock/show/oper"
}

func (p *ClockShowOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (ClockShoww, error) {
	logger.Println("ClockShowOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var clockShow ClockShoww
	if err == nil {
		err = json.Unmarshal(axResp, &clockShow)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return clockShow, err
}
