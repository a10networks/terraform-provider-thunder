

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RrdCgnOper struct {
    
    Oper RrdCgnOperOper `json:"oper"`

}
type DataRrdCgnOper struct {
    DtRrdCgnOper RrdCgnOper `json:"cgn"`
}


type RrdCgnOperOper struct {
    StartTime int `json:"start-time"`
    EndTime int `json:"end-time"`
    CgnData []RrdCgnOperOperCgnData `json:"cgn-data"`
}


type RrdCgnOperOperCgnData struct {
    Time int `json:"time"`
    Lsn_user_quota_create int `json:"lsn_user_quota_create"`
    Lsn_user_quota_delete int `json:"lsn_user_quota_delete"`
    Nat64_user_quota_create int `json:"nat64_user_quota_create"`
    Nat64_user_quota_delete int `json:"nat64_user_quota_delete"`
    Dslite_user_quota_create int `json:"dslite_user_quota_create"`
    Dslite_user_quota_delete int `json:"dslite_user_quota_delete"`
}

func (p *RrdCgnOper) GetId() string{
    return "1"
}

func (p *RrdCgnOper) getPath() string{
    return "rrd/cgn/oper"
}

func (p *RrdCgnOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRrdCgnOper,error) {
logger.Println("RrdCgnOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRrdCgnOper
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
