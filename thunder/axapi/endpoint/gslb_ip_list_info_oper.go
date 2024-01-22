

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbIpListInfoOper struct {
    
    Oper GslbIpListInfoOperOper `json:"oper"`

}
type DataGslbIpListInfoOper struct {
    DtGslbIpListInfoOper GslbIpListInfoOper `json:"ip-list-info"`
}


type GslbIpListInfoOperOper struct {
    IpLists []GslbIpListInfoOperOperIpLists `json:"ip-lists"`
}


type GslbIpListInfoOperOperIpLists struct {
    Listname string `json:"listname"`
    TotalEntriesInList int `json:"total-entries-in-list"`
    Filename string `json:"filename"`
    FileLines int `json:"file-lines"`
    Errors int `json:"errors"`
    Ipaddr string `json:"ipaddr"`
    Ipmask string `json:"ipmask"`
    GroupId int `json:"group-id"`
    Last string `json:"last"`
    Hits int `json:"hits"`
    Type string `json:"type"`
    IpaddrFilter int `json:"ipaddr-filter"`
}

func (p *GslbIpListInfoOper) GetId() string{
    return "1"
}

func (p *GslbIpListInfoOper) getPath() string{
    return "gslb/ip-list-info/oper"
}

func (p *GslbIpListInfoOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbIpListInfoOper,error) {
logger.Println("GslbIpListInfoOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbIpListInfoOper
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
