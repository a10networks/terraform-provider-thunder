

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceIpsOper struct {
    
    Oper DdosDstZonePortZoneServiceIpsOperOper `json:"oper"`
    ZoneName string 
    PortNum string 
    Protocol string 

}
type DataDdosDstZonePortZoneServiceIpsOper struct {
    DtDdosDstZonePortZoneServiceIpsOper DdosDstZonePortZoneServiceIpsOper `json:"ips"`
}


type DdosDstZonePortZoneServiceIpsOperOper struct {
    SignatureList []DdosDstZonePortZoneServiceIpsOperOperSignatureList `json:"signature-list"`
}


type DdosDstZonePortZoneServiceIpsOperOperSignatureList struct {
    Sid int `json:"sid"`
    MatchCount int `json:"match-count"`
}

func (p *DdosDstZonePortZoneServiceIpsOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortZoneServiceIpsOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port/zone-service/" +p.PortNum + "+" +p.Protocol + "/ips/oper"
}

func (p *DdosDstZonePortZoneServiceIpsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceIpsOper,error) {
logger.Println("DdosDstZonePortZoneServiceIpsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortZoneServiceIpsOper
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
