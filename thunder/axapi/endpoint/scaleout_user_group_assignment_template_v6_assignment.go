package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 7_0_2-102
type ScaleoutUserGroupAssignmentTemplateV6Assignment struct {
	Inst struct {
		AssignmentPrefixLength int `json:"assignment-prefix-length" dval:"128"`

		Ipv6Prefix string `json:"ipv6-prefix"`

		ServiceConfigTemplate string `json:"service-config-template"`

		UserGroupRangeEnd int `json:"user-group-range-end"`

		UserGroupRangeStart int `json:"user-group-range-start"`

		Uuid string `json:"uuid"`

		User_group_assignment_template_name string
	} `json:"v6-assignment"`
}

func (p *ScaleoutUserGroupAssignmentTemplateV6Assignment) GetId() string {
	return url.QueryEscape(p.Inst.Ipv6Prefix)
}

func (p *ScaleoutUserGroupAssignmentTemplateV6Assignment) getPath() string {
	return "scaleout/user-group-assignment-template/" + p.Inst.User_group_assignment_template_name + "/v6-assignment"
}

func (p *ScaleoutUserGroupAssignmentTemplateV6Assignment) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutUserGroupAssignmentTemplateV6Assignment::Post")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload:", string(payloadBytes))
	_, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
	return err
}

func (p *ScaleoutUserGroupAssignmentTemplateV6Assignment) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutUserGroupAssignmentTemplateV6Assignment::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *ScaleoutUserGroupAssignmentTemplateV6Assignment) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutUserGroupAssignmentTemplateV6Assignment::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
	return err
}

func (p *ScaleoutUserGroupAssignmentTemplateV6Assignment) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutUserGroupAssignmentTemplateV6Assignment::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
