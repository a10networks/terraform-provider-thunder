package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_0-P1_10
type CloudServicesCloudProviderAzureLog struct {
	Inst struct {
		Action     string `json:"action" dval:"disable"`
		CustomerId string `json:"customer-id"`
		Encrypted  string `json:"encrypted"`
		ResourceId string `json:"resource-id"`
		SharedKey  string `json:"shared-key"`
		Uuid       string `json:"uuid"`
	} `json:"log"`
}

func (p *CloudServicesCloudProviderAzureLog) GetId() string {
	return "1"
}

func (p *CloudServicesCloudProviderAzureLog) getPath() string {
	return "cloud-services/cloud-provider/azure/log"
}

func (p *CloudServicesCloudProviderAzureLog) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("CloudServicesCloudProviderAzureLog::Post")
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

func (p *CloudServicesCloudProviderAzureLog) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("CloudServicesCloudProviderAzureLog::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	if err == nil {
		err = json.Unmarshal(axResp, &p)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}

func (p *CloudServicesCloudProviderAzureLog) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("CloudServicesCloudProviderAzureLog::Put")
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

func (p *CloudServicesCloudProviderAzureLog) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("CloudServicesCloudProviderAzureLog::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
