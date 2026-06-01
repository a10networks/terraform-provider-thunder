package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SystemClThreatCategoryOper struct {
	Oper SystemClThreatCategoryOperOper `json:"oper"`
}
type DataSystemClThreatCategoryOper struct {
	DtSystemClThreatCategoryOper SystemClThreatCategoryOper `json:"cl-threat-category"`
}

type SystemClThreatCategoryOperOper struct {
	CategoryList []SystemClThreatCategoryOperOperCategoryList `json:"category-list"`
}

type SystemClThreatCategoryOperOperCategoryList struct {
	Category string `json:"category"`
}

func (p *SystemClThreatCategoryOper) GetId() string {
	return "1"
}

func (p *SystemClThreatCategoryOper) getPath() string {
	return "system/cl-threat-category/oper"
}

func (p *SystemClThreatCategoryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemClThreatCategoryOper, error) {
	logger.Println("SystemClThreatCategoryOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataSystemClThreatCategoryOper
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
