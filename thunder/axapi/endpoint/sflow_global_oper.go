

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SflowGlobalOper struct {
    
    Oper SflowGlobalOperOper `json:"oper"`

}
type DataSflowGlobalOper struct {
    DtSflowGlobalOper SflowGlobalOper `json:"global"`
}


type SflowGlobalOperOper struct {
    IfStatsList []SflowGlobalOperOperIfStatsList `json:"if-stats-list"`
}


type SflowGlobalOperOperIfStatsList struct {
    IfType string `json:"if-type"`
    IfNum int `json:"if-num"`
    PacketSampleRecords int `json:"packet-sample-records"`
    CounterSampleRecords int `json:"counter-sample-records"`
}

func (p *SflowGlobalOper) GetId() string{
    return "1"
}

func (p *SflowGlobalOper) getPath() string{
    return "sflow/global/oper"
}

func (p *SflowGlobalOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSflowGlobalOper,error) {
logger.Println("SflowGlobalOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSflowGlobalOper
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
