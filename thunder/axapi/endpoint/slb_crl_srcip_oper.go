

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbCrlSrcipOper struct {
    
    Oper SlbCrlSrcipOperOper `json:"oper"`

}
type DataSlbCrlSrcipOper struct {
    DtSlbCrlSrcipOper SlbCrlSrcipOper `json:"crl-srcip"`
}


type SlbCrlSrcipOperOper struct {
    Crl_srcip_lockedout_ips []SlbCrlSrcipOperOperCrl_srcip_lockedout_ips `json:"crl_srcip_lockedout_ips"`
    Lockedout_ips_count int `json:"lockedout_ips_count"`
}


type SlbCrlSrcipOperOperCrl_srcip_lockedout_ips struct {
    Client_ip string `json:"client_ip"`
    Start string `json:"start"`
    End string `json:"end"`
    Drops int `json:"drops"`
    Active int `json:"active"`
}

func (p *SlbCrlSrcipOper) GetId() string{
    return "1"
}

func (p *SlbCrlSrcipOper) getPath() string{
    return "slb/crl-srcip/oper"
}

func (p *SlbCrlSrcipOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbCrlSrcipOper,error) {
logger.Println("SlbCrlSrcipOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbCrlSrcipOper
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
