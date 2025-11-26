package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type IpRouteSourceLif struct {
	Inst struct {
		Ifname string `json:"ifname"`

		NexthopIp string `json:"nexthop-ip"`

		Uuid string `json:"uuid"`
	} `json:"lif"`
}

func (p *IpRouteSourceLif) GetId() string {
	return p.Inst.Ifname + "+" + p.Inst.NexthopIp
}

func (p *IpRouteSourceLif) getPath() string {
	return "ip/route/source/lif"
}

func (p *IpRouteSourceLif) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("IpRouteSourceLif::Post")
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

func (p *IpRouteSourceLif) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("IpRouteSourceLif::Get")
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
func (p *IpRouteSourceLif) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("IpRouteSourceLif::Put")
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

func (p *IpRouteSourceLif) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("IpRouteSourceLif::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
