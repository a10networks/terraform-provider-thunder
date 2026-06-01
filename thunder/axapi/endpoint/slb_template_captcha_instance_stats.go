package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SlbTemplateCaptchaInstanceStats struct {
	Name string `json:"name"`

	Stats SlbTemplateCaptchaInstanceStatsStats `json:"stats"`
}
type DataSlbTemplateCaptchaInstanceStats struct {
	DtSlbTemplateCaptchaInstanceStats SlbTemplateCaptchaInstanceStats `json:"instance"`
}

type SlbTemplateCaptchaInstanceStatsStats struct {
	Request      int `json:"request"`
	VerifySucc   int `json:"verify-succ"`
	ParseFail    int `json:"parse-fail"`
	JsonFail     int `json:"json-fail"`
	TimeoutError int `json:"timeout-error"`
	OtherError   int `json:"other-error"`
}

func (p *SlbTemplateCaptchaInstanceStats) GetId() string {
	return "1"
}

func (p *SlbTemplateCaptchaInstanceStats) getPath() string {
	return "slb/template/captcha/instance/" + p.Name + "/stats"
}

func (p *SlbTemplateCaptchaInstanceStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTemplateCaptchaInstanceStats, error) {
	logger.Println("SlbTemplateCaptchaInstanceStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataSlbTemplateCaptchaInstanceStats
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return payload, err
}
