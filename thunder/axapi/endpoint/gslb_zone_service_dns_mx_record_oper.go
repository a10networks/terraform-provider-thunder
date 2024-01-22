

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneServiceDnsMxRecordOper struct {
    
    MxName string `json:"mx-name"`
    Oper GslbZoneServiceDnsMxRecordOperOper `json:"oper"`
    ServiceName string 
    ServicePort string 
    Name string 

}
type DataGslbZoneServiceDnsMxRecordOper struct {
    DtGslbZoneServiceDnsMxRecordOper GslbZoneServiceDnsMxRecordOper `json:"dns-mx-record"`
}


type GslbZoneServiceDnsMxRecordOperOper struct {
    LastServer string `json:"last-server"`
    Hits int `json:"hits"`
    Priority int `json:"priority"`
}

func (p *GslbZoneServiceDnsMxRecordOper) GetId() string{
    return "1"
}

func (p *GslbZoneServiceDnsMxRecordOper) getPath() string{
    
    return "gslb/zone/" +p.Name + "/service/" +p.ServicePort + "+" +p.ServiceName + "/dns-mx-record/"+p.MxName+"/oper"
}

func (p *GslbZoneServiceDnsMxRecordOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneServiceDnsMxRecordOper,error) {
logger.Println("GslbZoneServiceDnsMxRecordOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneServiceDnsMxRecordOper
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
