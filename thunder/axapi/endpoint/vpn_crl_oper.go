

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnCrlOper struct {
    
    Oper VpnCrlOperOper `json:"oper"`

}
type DataVpnCrlOper struct {
    DtVpnCrlOper VpnCrlOper `json:"crl"`
}


type VpnCrlOperOper struct {
    CrlList []VpnCrlOperOperCrlList `json:"crl-list"`
    TotalCrls int `json:"total-crls"`
}


type VpnCrlOperOperCrlList struct {
    Subject string `json:"subject"`
    Issuer string `json:"issuer"`
    Updates string `json:"updates"`
    Serial string `json:"serial"`
    Revoked string `json:"revoked"`
    StorageType string `json:"storage-type"`
}

func (p *VpnCrlOper) GetId() string{
    return "1"
}

func (p *VpnCrlOper) getPath() string{
    return "vpn/crl/oper"
}

func (p *VpnCrlOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnCrlOper,error) {
logger.Println("VpnCrlOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnCrlOper
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
