package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type ControllerPartitionTenantInfoOper struct {
	Oper ControllerPartitionTenantInfoOperOper `json:"oper"`
}
type DataControllerPartitionTenantInfoOper struct {
	DtControllerPartitionTenantInfoOper ControllerPartitionTenantInfoOper `json:"partition-tenant-info"`
}

type ControllerPartitionTenantInfoOperOper struct {
	PartitionName string `json:"partition-name"`
	TenantName    string `json:"tenant-name"`
	TenantId      string `json:"tenant-id"`
	ClusterName   string `json:"cluster-name"`
	ClusterId     string `json:"cluster-id"`
	LogRatePerSec int    `json:"log-rate-per-sec"`
}

func (p *ControllerPartitionTenantInfoOper) GetId() string {
	return "1"
}

func (p *ControllerPartitionTenantInfoOper) getPath() string {
	return "controller/partition-tenant-info/oper"
}

func (p *ControllerPartitionTenantInfoOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataControllerPartitionTenantInfoOper, error) {
	logger.Println("ControllerPartitionTenantInfoOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataControllerPartitionTenantInfoOper
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
