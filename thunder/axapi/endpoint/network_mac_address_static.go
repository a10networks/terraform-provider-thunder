

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type NetworkMacAddressStatic struct {
	Inst struct {

    Dest int `json:"dest"`
    Mac string `json:"mac"`
    Port int `json:"port"`
    RedirectDummyMac int `json:"redirect-dummy-mac"`
    Uuid string `json:"uuid"`
    Vlan int `json:"vlan"`

	} `json:"static"`
}

func (p *NetworkMacAddressStatic) GetId() string{
    return p.Inst.Mac+"+"+strconv.Itoa(p.Inst.Vlan)
}

func (p *NetworkMacAddressStatic) getPath() string{
    return "network/mac-address/static"
}

func (p *NetworkMacAddressStatic) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkMacAddressStatic::Post")
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

func (p *NetworkMacAddressStatic) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkMacAddressStatic::Get")
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
func (p *NetworkMacAddressStatic) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkMacAddressStatic::Put")
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

func (p *NetworkMacAddressStatic) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkMacAddressStatic::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
