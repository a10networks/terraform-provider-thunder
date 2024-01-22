

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange struct {
	Inst struct {

    PortEnd int `json:"port-end"`
    PortStart int `json:"port-start"`
    TunnelEndpointAddress string `json:"tunnel-endpoint-address"`
    Ipv4NatAddr string 
    Ipv6TunnelAddr string 
    Name string 

	} `json:"port-range"`
}

func (p *Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange) GetId() string{
    return strconv.Itoa(p.Inst.PortStart)+"+"+strconv.Itoa(p.Inst.PortEnd)+"+"+p.Inst.TunnelEndpointAddress
}

func (p *Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange) getPath() string{
    return "cgnv6/lw-4o6/binding-table/" +p.Inst.Name + "/tunnel-address/" +p.Inst.Ipv6TunnelAddr + "/nat-address/" +p.Inst.Ipv4NatAddr + "/port-range"
}

func (p *Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange::Post")
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

func (p *Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange::Get")
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
func (p *Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange::Put")
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

func (p *Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Lw4o6BindingTableTunnelAddressNatAddressPortRange::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
