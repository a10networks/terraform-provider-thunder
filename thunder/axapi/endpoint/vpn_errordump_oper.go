

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnErrordumpOper struct {
    
    Oper VpnErrordumpOperOper `json:"oper"`

}
type DataVpnErrordumpOper struct {
    DtVpnErrordumpOper VpnErrordumpOper `json:"errordump"`
}


type VpnErrordumpOperOper struct {
    IpsecErrorDumpPath string `json:"IPsec-error-dump-path"`
}

func (p *VpnErrordumpOper) GetId() string{
    return "1"
}

func (p *VpnErrordumpOper) getPath() string{
    return "vpn/errordump/oper"
}

func (p *VpnErrordumpOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnErrordumpOper,error) {
logger.Println("VpnErrordumpOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnErrordumpOper
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
