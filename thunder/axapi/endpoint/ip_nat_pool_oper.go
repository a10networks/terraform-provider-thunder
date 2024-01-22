

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpNatPoolOper struct {
    
    Oper IpNatPoolOperOper `json:"oper"`
    PoolName string `json:"pool-name"`

}
type DataIpNatPoolOper struct {
    DtIpNatPoolOper IpNatPoolOper `json:"pool"`
}


type IpNatPoolOperOper struct {
    NatPoolAddrList []IpNatPoolOperOperNatPoolAddrList `json:"nat-pool-addr-list"`
}


type IpNatPoolOperOperNatPoolAddrList struct {
    Address string `json:"Address"`
    PortUsage int `json:"Port-Usage"`
    TotalUsed int `json:"Total-Used"`
    TotalFreed int `json:"Total-Freed"`
    Failed int `json:"Failed"`
}

func (p *IpNatPoolOper) GetId() string{
    return "1"
}

func (p *IpNatPoolOper) getPath() string{
    
    return "ip/nat/pool/"+p.PoolName+"/oper"
}

func (p *IpNatPoolOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpNatPoolOper,error) {
logger.Println("IpNatPoolOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpNatPoolOper
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
