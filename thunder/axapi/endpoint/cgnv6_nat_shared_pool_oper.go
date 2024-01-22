

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6NatSharedPoolOper struct {
    
    Oper Cgnv6NatSharedPoolOperOper `json:"oper"`

}
type DataCgnv6NatSharedPoolOper struct {
    DtCgnv6NatSharedPoolOper Cgnv6NatSharedPoolOper `json:"shared-pool"`
}


type Cgnv6NatSharedPoolOperOper struct {
    SharedPoolList []Cgnv6NatSharedPoolOperOperSharedPoolList `json:"shared-pool-list"`
}


type Cgnv6NatSharedPoolOperOperSharedPoolList struct {
    PoolName string `json:"pool-name"`
    StartAddress string `json:"start-address"`
    EndAddress string `json:"end-address"`
    Netmask string `json:"netmask"`
    Vird int `json:"vird"`
}

func (p *Cgnv6NatSharedPoolOper) GetId() string{
    return "1"
}

func (p *Cgnv6NatSharedPoolOper) getPath() string{
    return "cgnv6/nat/shared-pool/oper"
}

func (p *Cgnv6NatSharedPoolOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6NatSharedPoolOper,error) {
logger.Println("Cgnv6NatSharedPoolOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6NatSharedPoolOper
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
