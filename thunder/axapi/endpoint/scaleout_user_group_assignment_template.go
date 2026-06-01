package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 6_0_8-219
type ScaleoutUserGroupAssignmentTemplate struct {
	Inst struct {
		Name string `json:"name"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		V4AssignmentList []ScaleoutUserGroupAssignmentTemplateV4AssignmentList `json:"v4-assignment-list"`

		V6AssignmentList []ScaleoutUserGroupAssignmentTemplateV6AssignmentList `json:"v6-assignment-list"`
	} `json:"user-group-assignment-template"`
}

type ScaleoutUserGroupAssignmentTemplateV4AssignmentList struct {
	Ipv4Prefix             string `json:"ipv4-prefix"`
	AssignmentPrefixLength int    `json:"assignment-prefix-length" dval:"32"`
	AssignmentPrefixAuto   int    `json:"assignment-prefix-auto"`
	UserGroupRangeStart    int    `json:"user-group-range-start"`
	UserGroupRangeEnd      int    `json:"user-group-range-end"`
	ServiceConfigTemplate  string `json:"service-config-template"`
	PrivateIp              int    `json:"private-ip"`
	Uuid                   string `json:"uuid"`
}

type ScaleoutUserGroupAssignmentTemplateV6AssignmentList struct {
	Ipv6Prefix             string `json:"ipv6-prefix"`
	AssignmentPrefixLength int    `json:"assignment-prefix-length" dval:"128"`
	AssignmentPrefixAuto   int    `json:"assignment-prefix-auto"`
	UserGroupRangeStart    int    `json:"user-group-range-start"`
	UserGroupRangeEnd      int    `json:"user-group-range-end"`
	ServiceConfigTemplate  string `json:"service-config-template"`
	PrivateIp              int    `json:"private-ip"`
	Uuid                   string `json:"uuid"`
}

func (p *ScaleoutUserGroupAssignmentTemplate) GetId() string {
	return url.QueryEscape(p.Inst.Name)
}

func (p *ScaleoutUserGroupAssignmentTemplate) getPath() string {
	return "scaleout/user-group-assignment-template"
}

func (p *ScaleoutUserGroupAssignmentTemplate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutUserGroupAssignmentTemplate::Post")
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

func (p *ScaleoutUserGroupAssignmentTemplate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutUserGroupAssignmentTemplate::Get")
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
func (p *ScaleoutUserGroupAssignmentTemplate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutUserGroupAssignmentTemplate::Put")
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

func (p *ScaleoutUserGroupAssignmentTemplate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ScaleoutUserGroupAssignmentTemplate::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
