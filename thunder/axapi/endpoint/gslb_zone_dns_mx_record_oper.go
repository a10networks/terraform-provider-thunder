

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneDnsMxRecordOper struct {
    
    MxName string `json:"mx-name"`
    Oper GslbZoneDnsMxRecordOperOper `json:"oper"`
    Name string 

}
type DataGslbZoneDnsMxRecordOper struct {
    DtGslbZoneDnsMxRecordOper GslbZoneDnsMxRecordOper `json:"dns-mx-record"`
}


type GslbZoneDnsMxRecordOperOper struct {
    LastServer string `json:"last-server"`
    Hits int `json:"hits"`
    Priority int `json:"priority"`
}

func (p *GslbZoneDnsMxRecordOper) GetId() string{
    return "1"
}

func (p *GslbZoneDnsMxRecordOper) getPath() string{
    
    return "gslb/zone/" +p.Name + "/dns-mx-record/"+p.MxName+"/oper"
}

func (p *GslbZoneDnsMxRecordOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneDnsMxRecordOper,error) {
logger.Println("GslbZoneDnsMxRecordOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneDnsMxRecordOper
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
