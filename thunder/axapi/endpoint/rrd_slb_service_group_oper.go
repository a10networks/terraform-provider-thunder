

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RrdSlbServiceGroupOper struct {
    
    Oper RrdSlbServiceGroupOperOper `json:"oper"`

}
type DataRrdSlbServiceGroupOper struct {
    DtRrdSlbServiceGroupOper RrdSlbServiceGroupOper `json:"slb-service-group"`
}


type RrdSlbServiceGroupOperOper struct {
    StartTime int `json:"start-time"`
    EndTime int `json:"end-time"`
    SlbServiceGroupName string `json:"slb-service-group-name"`
    SlbServerStatistics []RrdSlbServiceGroupOperOperSlbServerStatistics `json:"slb-server-statistics"`
}


type RrdSlbServiceGroupOperOperSlbServerStatistics struct {
    Time int `json:"time"`
    In_pkts int `json:"in_pkts"`
    In_bytes int `json:"in_bytes"`
    Out_pkts int `json:"out_pkts"`
    Out_bytes int `json:"out_bytes"`
    Cur_conn int `json:"cur_conn"`
    P_conn int `json:"p_conn"`
}

func (p *RrdSlbServiceGroupOper) GetId() string{
    return "1"
}

func (p *RrdSlbServiceGroupOper) getPath() string{
    return "rrd/slb-service-group/oper"
}

func (p *RrdSlbServiceGroupOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRrdSlbServiceGroupOper,error) {
logger.Println("RrdSlbServiceGroupOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRrdSlbServiceGroupOper
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
