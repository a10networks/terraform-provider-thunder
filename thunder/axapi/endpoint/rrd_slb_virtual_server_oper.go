

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RrdSlbVirtualServerOper struct {
    
    Oper RrdSlbVirtualServerOperOper `json:"oper"`

}
type DataRrdSlbVirtualServerOper struct {
    DtRrdSlbVirtualServerOper RrdSlbVirtualServerOper `json:"slb-virtual-server"`
}


type RrdSlbVirtualServerOperOper struct {
    StartTime int `json:"start-time"`
    EndTime int `json:"end-time"`
    SlbVirtualServerName string `json:"slb-virtual-server-name"`
    SlbServerStatistics []RrdSlbVirtualServerOperOperSlbServerStatistics `json:"slb-server-statistics"`
}


type RrdSlbVirtualServerOperOperSlbServerStatistics struct {
    Time int `json:"time"`
    In_pkts int `json:"in_pkts"`
    In_bytes int `json:"in_bytes"`
    Out_pkts int `json:"out_pkts"`
    Out_bytes int `json:"out_bytes"`
    Cur_conn int `json:"cur_conn"`
    P_conn int `json:"p_conn"`
}

func (p *RrdSlbVirtualServerOper) GetId() string{
    return "1"
}

func (p *RrdSlbVirtualServerOper) getPath() string{
    return "rrd/slb-virtual-server/oper"
}

func (p *RrdSlbVirtualServerOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRrdSlbVirtualServerOper,error) {
logger.Println("RrdSlbVirtualServerOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRrdSlbVirtualServerOper
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
