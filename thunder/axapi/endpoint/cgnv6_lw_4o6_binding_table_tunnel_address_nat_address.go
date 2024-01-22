

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Lw4o6BindingTableTunnelAddressNatAddress struct {
	Inst struct {

    Ipv4NatAddr string `json:"ipv4-nat-addr"`
    PortRangeList []Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeList `json:"port-range-list"`
    UserTag string `json:"user-tag"`
    Name string 
    Ipv6TunnelAddr string 

	} `json:"nat-address"`
}


type Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRangeList struct {
    PortStart int `json:"port-start"`
    PortEnd int `json:"port-end"`
    TunnelEndpointAddress string `json:"tunnel-endpoint-address"`
}

func (p *Cgnv6Lw4o6BindingTableTunnelAddressNatAddress) GetId() string{
    return p.Inst.Ipv4NatAddr
}

func (p *Cgnv6Lw4o6BindingTableTunnelAddressNatAddress) getPath() string{
    return "cgnv6/lw-4o6/binding-table/" +p.Inst.Name + "/tunnel-address/" +p.Inst.Ipv6TunnelAddr + "/nat-address"
}

func (p *Cgnv6Lw4o6BindingTableTunnelAddressNatAddress) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Lw4o6BindingTableTunnelAddressNatAddress::Post")
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

func (p *Cgnv6Lw4o6BindingTableTunnelAddressNatAddress) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Lw4o6BindingTableTunnelAddressNatAddress::Get")
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
func (p *Cgnv6Lw4o6BindingTableTunnelAddressNatAddress) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Lw4o6BindingTableTunnelAddressNatAddress::Put")
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

func (p *Cgnv6Lw4o6BindingTableTunnelAddressNatAddress) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Lw4o6BindingTableTunnelAddressNatAddress::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
