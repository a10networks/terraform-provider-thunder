

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type NetworkArpStatic struct {
	Inst struct {

    Ethernet int `json:"ethernet"`
    IpAddr string `json:"ip-addr"`
    MacAddr string `json:"mac-addr"`
    Trunk int `json:"trunk"`
    Uuid string `json:"uuid"`
    Vlan int `json:"vlan"`

	} `json:"static"`
}

func (p *NetworkArpStatic) GetId() string{
    return p.Inst.IpAddr+"+"+strconv.Itoa(p.Inst.Vlan)
}

func (p *NetworkArpStatic) getPath() string{
    return "network/arp/static"
}

func (p *NetworkArpStatic) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkArpStatic::Post")
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

func (p *NetworkArpStatic) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkArpStatic::Get")
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
func (p *NetworkArpStatic) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkArpStatic::Put")
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

func (p *NetworkArpStatic) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkArpStatic::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
