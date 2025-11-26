package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type ScaleoutUserGroupAssignmentTemplateInfoOper struct {
	Oper ScaleoutUserGroupAssignmentTemplateInfoOperOper `json:"oper"`
}
type DataScaleoutUserGroupAssignmentTemplateInfoOper struct {
	DtScaleoutUserGroupAssignmentTemplateInfoOper ScaleoutUserGroupAssignmentTemplateInfoOper `json:"template-info"`
}

type ScaleoutUserGroupAssignmentTemplateInfoOperOper struct {
	AssignmentList []ScaleoutUserGroupAssignmentTemplateInfoOperOperAssignmentList `json:"assignment-list"`
	Name           string                                                          `json:"name"`
}

type ScaleoutUserGroupAssignmentTemplateInfoOperOperAssignmentList struct {
	AddressPrefix        string `json:"address-prefix"`
	PrefixesCount        int    `json:"prefixes-count"`
	IpCountPerPrefix     int    `json:"ip-count-per-prefix"`
	PrefixesPerUserGroup int    `json:"prefixes-per-user-group"`
	UserGroupRangeStart  int    `json:"user-group-range-start"`
	UserGroupRangeEnd    int    `json:"user-group-range-end"`
}

func (p *ScaleoutUserGroupAssignmentTemplateInfoOper) GetId() string {
	return "1"
}

func (p *ScaleoutUserGroupAssignmentTemplateInfoOper) getPath() string {
	return "scaleout/user-group-assignment/template-info/oper"
}

func (p *ScaleoutUserGroupAssignmentTemplateInfoOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutUserGroupAssignmentTemplateInfoOper, error) {
	logger.Println("ScaleoutUserGroupAssignmentTemplateInfoOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataScaleoutUserGroupAssignmentTemplateInfoOper
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
