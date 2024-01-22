

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ThreatIntelWebrootIpCategoryOper struct {
    
    Oper ThreatIntelWebrootIpCategoryOperOper `json:"oper"`

}
type DataThreatIntelWebrootIpCategoryOper struct {
    DtThreatIntelWebrootIpCategoryOper ThreatIntelWebrootIpCategoryOper `json:"webroot-ip-category"`
}


type ThreatIntelWebrootIpCategoryOperOper struct {
    CategoryList []ThreatIntelWebrootIpCategoryOperOperCategoryList `json:"category-list"`
    Ip string `json:"ip"`
}


type ThreatIntelWebrootIpCategoryOperOperCategoryList struct {
    Category string `json:"category"`
}

func (p *ThreatIntelWebrootIpCategoryOper) GetId() string{
    return "1"
}

func (p *ThreatIntelWebrootIpCategoryOper) getPath() string{
    return "threat-intel/webroot-ip-category/oper"
}

func (p *ThreatIntelWebrootIpCategoryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataThreatIntelWebrootIpCategoryOper,error) {
logger.Println("ThreatIntelWebrootIpCategoryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataThreatIntelWebrootIpCategoryOper
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
