

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkLldpManagementAddressIpv4Addr struct {
	Inst struct {

    InterfaceIpv4 NetworkLldpManagementAddressIpv4AddrInterfaceIpv4 `json:"interface-ipv4"`
    Ipv4 string `json:"ipv4"`
    Uuid string `json:"uuid"`

	} `json:"ipv4-addr"`
}


type NetworkLldpManagementAddressIpv4AddrInterfaceIpv4 struct {
    Ipv4Eth int `json:"ipv4-eth"`
    Ipv4Ve int `json:"ipv4-ve"`
    Ipv4Mgmt int `json:"ipv4-mgmt"`
}

func (p *NetworkLldpManagementAddressIpv4Addr) GetId() string{
    return p.Inst.Ipv4
}

func (p *NetworkLldpManagementAddressIpv4Addr) getPath() string{
    return "network/lldp/management-address/ipv4-addr"
}

func (p *NetworkLldpManagementAddressIpv4Addr) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLldpManagementAddressIpv4Addr::Post")
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

func (p *NetworkLldpManagementAddressIpv4Addr) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLldpManagementAddressIpv4Addr::Get")
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
func (p *NetworkLldpManagementAddressIpv4Addr) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLldpManagementAddressIpv4Addr::Put")
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

func (p *NetworkLldpManagementAddressIpv4Addr) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLldpManagementAddressIpv4Addr::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
