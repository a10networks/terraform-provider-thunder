

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkLldpManagementAddressIpv6Addr struct {
	Inst struct {

    InterfaceIpv6 NetworkLldpManagementAddressIpv6AddrInterfaceIpv6 `json:"interface-ipv6"`
    Ipv6 string `json:"ipv6"`
    Uuid string `json:"uuid"`

	} `json:"ipv6-addr"`
}


type NetworkLldpManagementAddressIpv6AddrInterfaceIpv6 struct {
    Ipv6Eth int `json:"ipv6-eth"`
    Ipv6Ve int `json:"ipv6-ve"`
    Ipv6Mgmt int `json:"ipv6-mgmt"`
}

func (p *NetworkLldpManagementAddressIpv6Addr) GetId() string{
    return p.Inst.Ipv6
}

func (p *NetworkLldpManagementAddressIpv6Addr) getPath() string{
    return "network/lldp/management-address/ipv6-addr"
}

func (p *NetworkLldpManagementAddressIpv6Addr) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLldpManagementAddressIpv6Addr::Post")
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

func (p *NetworkLldpManagementAddressIpv6Addr) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLldpManagementAddressIpv6Addr::Get")
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
func (p *NetworkLldpManagementAddressIpv6Addr) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLldpManagementAddressIpv6Addr::Put")
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

func (p *NetworkLldpManagementAddressIpv6Addr) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLldpManagementAddressIpv6Addr::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
