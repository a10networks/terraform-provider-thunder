

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RrdSlbServerOper struct {
    
    Oper RrdSlbServerOperOper `json:"oper"`

}
type DataRrdSlbServerOper struct {
    DtRrdSlbServerOper RrdSlbServerOper `json:"slb-server"`
}


type RrdSlbServerOperOper struct {
    StartTime int `json:"start-time"`
    EndTime int `json:"end-time"`
    SlbServerName string `json:"slb-server-name"`
    SlbServerStatistics []RrdSlbServerOperOperSlbServerStatistics `json:"slb-server-statistics"`
}


type RrdSlbServerOperOperSlbServerStatistics struct {
    Time int `json:"time"`
    In_pkts int `json:"in_pkts"`
    In_bytes int `json:"in_bytes"`
    Out_pkts int `json:"out_pkts"`
    Out_bytes int `json:"out_bytes"`
    Cur_conn int `json:"cur_conn"`
    P_conn int `json:"p_conn"`
}

func (p *RrdSlbServerOper) GetId() string{
    return "1"
}

func (p *RrdSlbServerOper) getPath() string{
    return "rrd/slb-server/oper"
}

func (p *RrdSlbServerOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRrdSlbServerOper,error) {
logger.Println("RrdSlbServerOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRrdSlbServerOper
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
