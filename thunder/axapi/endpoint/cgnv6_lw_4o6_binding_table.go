

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Lw4o6BindingTable struct {
	Inst struct {

    Name string `json:"name"`
    TunnelAddressList []Cgnv6Lw4o6BindingTableTunnelAddressList `json:"tunnel-address-list"`
    UserTag string `json:"user-tag"`

	} `json:"binding-table"`
}


type Cgnv6Lw4o6BindingTableTunnelAddressList struct {
    Ipv6TunnelAddr string `json:"ipv6-tunnel-addr"`
    UserTag string `json:"user-tag"`
    NatAddressList []Cgnv6Lw4o6BindingTableTunnelAddressListNatAddressList `json:"nat-address-list"`
}


type Cgnv6Lw4o6BindingTableTunnelAddressListNatAddressList struct {
    Ipv4NatAddr string `json:"ipv4-nat-addr"`
    UserTag string `json:"user-tag"`
    PortRangeList []Cgnv6Lw4o6BindingTableTunnelAddressListNatAddressListPortRangeList `json:"port-range-list"`
}


type Cgnv6Lw4o6BindingTableTunnelAddressListNatAddressListPortRangeList struct {
    PortStart int `json:"port-start"`
    PortEnd int `json:"port-end"`
    TunnelEndpointAddress string `json:"tunnel-endpoint-address"`
}

func (p *Cgnv6Lw4o6BindingTable) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *Cgnv6Lw4o6BindingTable) getPath() string{
    return "cgnv6/lw-4o6/binding-table"
}

func (p *Cgnv6Lw4o6BindingTable) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Lw4o6BindingTable::Post")
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

func (p *Cgnv6Lw4o6BindingTable) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Lw4o6BindingTable::Get")
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
func (p *Cgnv6Lw4o6BindingTable) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Lw4o6BindingTable::Put")
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

func (p *Cgnv6Lw4o6BindingTable) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Lw4o6BindingTable::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
