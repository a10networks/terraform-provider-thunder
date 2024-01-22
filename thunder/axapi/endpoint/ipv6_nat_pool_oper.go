

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Ipv6NatPoolOper struct {
    
    Oper Ipv6NatPoolOperOper `json:"oper"`
    PoolName string `json:"pool-name"`

}
type DataIpv6NatPoolOper struct {
    DtIpv6NatPoolOper Ipv6NatPoolOper `json:"pool"`
}


type Ipv6NatPoolOperOper struct {
    NatPoolAddrList []Ipv6NatPoolOperOperNatPoolAddrList `json:"nat-pool-addr-list"`
}


type Ipv6NatPoolOperOperNatPoolAddrList struct {
    Address string `json:"Address"`
    PortUsage int `json:"Port-Usage"`
    TotalUsed int `json:"Total-Used"`
    TotalFreed int `json:"Total-Freed"`
    Failed int `json:"Failed"`
}

func (p *Ipv6NatPoolOper) GetId() string{
    return "1"
}

func (p *Ipv6NatPoolOper) getPath() string{
    
    return "ipv6/nat/pool/"+p.PoolName+"/oper"
}

func (p *Ipv6NatPoolOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpv6NatPoolOper,error) {
logger.Println("Ipv6NatPoolOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpv6NatPoolOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
