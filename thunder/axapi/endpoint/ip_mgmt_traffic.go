

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpMgmtTraffic struct {
	Inst struct {

    SourceInterface IpMgmtTrafficSourceInterface `json:"source-interface"`
    TrafficType string `json:"traffic-type"`
    Uuid string `json:"uuid"`

	} `json:"mgmt-traffic"`
}


type IpMgmtTrafficSourceInterface struct {
    Loopback int `json:"loopback"`
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Trunk int `json:"trunk"`
    Lif int `json:"lif"`
    Tunnel int `json:"tunnel"`
}

func (p *IpMgmtTraffic) GetId() string{
    return p.Inst.TrafficType
}

func (p *IpMgmtTraffic) getPath() string{
    return "ip/mgmt-traffic"
}

func (p *IpMgmtTraffic) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpMgmtTraffic::Post")
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

func (p *IpMgmtTraffic) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpMgmtTraffic::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *IpMgmtTraffic) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpMgmtTraffic::Put")
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

func (p *IpMgmtTraffic) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpMgmtTraffic::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
