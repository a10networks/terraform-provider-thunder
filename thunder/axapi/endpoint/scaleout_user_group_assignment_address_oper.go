package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type ScaleoutUserGroupAssignmentAddressOper struct {
	Oper ScaleoutUserGroupAssignmentAddressOperOper `json:"oper"`
}
type DataScaleoutUserGroupAssignmentAddressOper struct {
	DtScaleoutUserGroupAssignmentAddressOper ScaleoutUserGroupAssignmentAddressOper `json:"address"`
}

type ScaleoutUserGroupAssignmentAddressOperOper struct {
	UserGroup        int    `json:"user-group"`
	ActiveNode       int    `json:"active-node"`
	StandbyNode      int    `json:"standby-node"`
	ServiceTemplate  string `json:"service-template"`
	Application_type string `json:"application_type"`
	Ip               string `json:"ip"`
	Ipv6             string `json:"ipv6"`
}

func (p *ScaleoutUserGroupAssignmentAddressOper) GetId() string {
	return "1"
}

func (p *ScaleoutUserGroupAssignmentAddressOper) getPath() string {
	return "scaleout/user-group-assignment/address/oper"
}

func (p *ScaleoutUserGroupAssignmentAddressOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutUserGroupAssignmentAddressOper, error) {
	logger.Println("ScaleoutUserGroupAssignmentAddressOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataScaleoutUserGroupAssignmentAddressOper
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
