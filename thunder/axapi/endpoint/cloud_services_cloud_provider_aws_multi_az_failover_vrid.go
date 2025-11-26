package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"strconv"
)

// based on ACOS 7_0_2-102
type CloudServicesCloudProviderAwsMultiAzFailoverVrid struct {
	Inst struct {
		FipDest string `json:"fip-dest"`

		FipInterfaceId string `json:"fip-interface-id"`

		RouteTableId string `json:"route-table-id"`

		UserTag string `json:"user-tag"`

		Uuid string `json:"uuid"`

		VipDest string `json:"vip-dest"`

		VipInterfaceId string `json:"vip-interface-id"`

		VipList []CloudServicesCloudProviderAwsMultiAzFailoverVridVipList `json:"vip-list"`

		VridNumber int `json:"vrid-number"`
	} `json:"vrid"`
}

type CloudServicesCloudProviderAwsMultiAzFailoverVridVipList struct {
	VipNumber int    `json:"vip-number"`
	PrivateIp string `json:"private-ip"`
	ElasticIp string `json:"elastic-ip"`
	Uuid      string `json:"uuid"`
	UserTag   string `json:"user-tag"`
}

func (p *CloudServicesCloudProviderAwsMultiAzFailoverVrid) GetId() string {
	return strconv.Itoa(p.Inst.VridNumber)
}

func (p *CloudServicesCloudProviderAwsMultiAzFailoverVrid) getPath() string {
	return "cloud-services/cloud-provider/aws/multi-az-failover/vrid"
}

func (p *CloudServicesCloudProviderAwsMultiAzFailoverVrid) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("CloudServicesCloudProviderAwsMultiAzFailoverVrid::Post")
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

func (p *CloudServicesCloudProviderAwsMultiAzFailoverVrid) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("CloudServicesCloudProviderAwsMultiAzFailoverVrid::Get")
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
func (p *CloudServicesCloudProviderAwsMultiAzFailoverVrid) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("CloudServicesCloudProviderAwsMultiAzFailoverVrid::Put")
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

func (p *CloudServicesCloudProviderAwsMultiAzFailoverVrid) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("CloudServicesCloudProviderAwsMultiAzFailoverVrid::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
