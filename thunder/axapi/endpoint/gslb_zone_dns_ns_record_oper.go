

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneDnsNsRecordOper struct {
    
    NsName string `json:"ns-name"`
    Oper GslbZoneDnsNsRecordOperOper `json:"oper"`
    Name string 

}
type DataGslbZoneDnsNsRecordOper struct {
    DtGslbZoneDnsNsRecordOper GslbZoneDnsNsRecordOper `json:"dns-ns-record"`
}


type GslbZoneDnsNsRecordOperOper struct {
    LastServer string `json:"last-server"`
    Hits int `json:"hits"`
}

func (p *GslbZoneDnsNsRecordOper) GetId() string{
    return "1"
}

func (p *GslbZoneDnsNsRecordOper) getPath() string{
    
    return "gslb/zone/" +p.Name + "/dns-ns-record/"+p.NsName+"/oper"
}

func (p *GslbZoneDnsNsRecordOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneDnsNsRecordOper,error) {
logger.Println("GslbZoneDnsNsRecordOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneDnsNsRecordOper
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
