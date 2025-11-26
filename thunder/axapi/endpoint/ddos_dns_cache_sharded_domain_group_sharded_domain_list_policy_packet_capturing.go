package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing struct {
	Inst struct {
		RootZoneList []DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRootZoneList `json:"root-zone-list"`

		Uuid string `json:"uuid"`

		Sharded_domain_list_policy_name string

		Sharded_domain_group_name string

		Dns_cache_name string
	} `json:"packet-capturing"`
}

type DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturingRootZoneList struct {
	RootZone      string `json:"root-zone"`
	CaptureConfig string `json:"capture-config"`
	CaptureMode   string `json:"capture-mode"`
}

func (p *DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing) GetId() string {
	return "1"
}

func (p *DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing) getPath() string {
	return "ddos/dns-cache/" + p.Inst.Dns_cache_name + "/sharded-domain-group/" + p.Inst.Sharded_domain_group_name + "/sharded-domain-list-policy/" + p.Inst.Sharded_domain_list_policy_name + "/packet-capturing"
}

func (p *DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing::Post")
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

func (p *DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
	return err
}

func (p *DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("DdosDnsCacheShardedDomainGroupShardedDomainListPolicyPacketCapturing::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
