

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6OneToOneSharedPoolOper struct {
    
    Oper Cgnv6OneToOneSharedPoolOperOper `json:"oper"`

}
type DataCgnv6OneToOneSharedPoolOper struct {
    DtCgnv6OneToOneSharedPoolOper Cgnv6OneToOneSharedPoolOper `json:"shared-pool"`
}


type Cgnv6OneToOneSharedPoolOperOper struct {
    SharedPoolList []Cgnv6OneToOneSharedPoolOperOperSharedPoolList `json:"shared-pool-list"`
}


type Cgnv6OneToOneSharedPoolOperOperSharedPoolList struct {
    PoolName string `json:"pool-name"`
    StartAddress string `json:"start-address"`
    EndAddress string `json:"end-address"`
    Netmask string `json:"netmask"`
    Vird int `json:"vird"`
}

func (p *Cgnv6OneToOneSharedPoolOper) GetId() string{
    return "1"
}

func (p *Cgnv6OneToOneSharedPoolOper) getPath() string{
    return "cgnv6/one-to-one/shared-pool/oper"
}

func (p *Cgnv6OneToOneSharedPoolOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6OneToOneSharedPoolOper,error) {
logger.Println("Cgnv6OneToOneSharedPoolOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6OneToOneSharedPoolOper
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
