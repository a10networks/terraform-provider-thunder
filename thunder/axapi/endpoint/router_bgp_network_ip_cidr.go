package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

//based on ACOS 5_2_1-P4_81
type RouterBgpNetworkIpCidr struct {
	Inst struct {
		AsNumber	string //an object key which is outside of instance
		Backdoor        int    `json:"backdoor"`
		CommValue       string `json:"comm-value"`
		Description     string `json:"description"`
		NetworkIpv4Cidr string `json:"network-ipv4-cidr"`
		RouteMap        string `json:"route-map"`
		Uuid            string `json:"uuid"`
	} `json:"ip-cidr"`
}

func (p *RouterBgpNetworkIpCidr) GetId() string {
	return url.QueryEscape(p.Inst.NetworkIpv4Cidr)
}

func (p *RouterBgpNetworkIpCidr) getPath() string {
	return "router/bgp/" + p.Inst.AsNumber + "/network/ip-cidr"
}

func (p *RouterBgpNetworkIpCidr) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("RouterBgpNetworkIpCidr::Post")
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

func (p *RouterBgpNetworkIpCidr) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("RouterBgpNetworkIpCidr::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
	if err == nil {
		err = json.Unmarshal(axResp, &p)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}

func (p *RouterBgpNetworkIpCidr) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("RouterBgpNetworkIpCidr::Put")
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

func (p *RouterBgpNetworkIpCidr) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("RouterBgpNetworkIpCidr::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
