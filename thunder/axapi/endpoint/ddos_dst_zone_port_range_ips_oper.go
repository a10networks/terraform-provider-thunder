

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortRangeIpsOper struct {
    
    Oper DdosDstZonePortRangeIpsOperOper `json:"oper"`
    Protocol string 
    ZoneName string 
    PortRangeEnd string 
    PortRangeStart string 

}
type DataDdosDstZonePortRangeIpsOper struct {
    DtDdosDstZonePortRangeIpsOper DdosDstZonePortRangeIpsOper `json:"ips"`
}


type DdosDstZonePortRangeIpsOperOper struct {
    SignatureList []DdosDstZonePortRangeIpsOperOperSignatureList `json:"signature-list"`
}


type DdosDstZonePortRangeIpsOperOperSignatureList struct {
    Sid int `json:"sid"`
    MatchCount int `json:"match-count"`
}

func (p *DdosDstZonePortRangeIpsOper) GetId() string{
    return "1"
}

func (p *DdosDstZonePortRangeIpsOper) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port-range/" +p.PortRangeStart + "+" +p.PortRangeEnd + "+" +p.Protocol + "/ips/oper"
}

func (p *DdosDstZonePortRangeIpsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortRangeIpsOper,error) {
logger.Println("DdosDstZonePortRangeIpsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortRangeIpsOper
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
