

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RrdEthernetAllOper struct {
    
    Oper RrdEthernetAllOperOper `json:"oper"`

}
type DataRrdEthernetAllOper struct {
    DtRrdEthernetAllOper RrdEthernetAllOper `json:"ethernet-all"`
}


type RrdEthernetAllOperOper struct {
    StartTime int `json:"start-time"`
    EndTime int `json:"end-time"`
    EthernetStatistics []RrdEthernetAllOperOperEthernetStatistics `json:"ethernet-statistics"`
}


type RrdEthernetAllOperOperEthernetStatistics struct {
    EthernetIndex int `json:"ethernet-index"`
    Stat []RrdEthernetAllOperOperEthernetStatisticsStat `json:"stat"`
}


type RrdEthernetAllOperOperEthernetStatisticsStat struct {
    Time int `json:"time"`
    Tx_bits int `json:"tx_bits"`
}

func (p *RrdEthernetAllOper) GetId() string{
    return "1"
}

func (p *RrdEthernetAllOper) getPath() string{
    return "rrd/ethernet-all/oper"
}

func (p *RrdEthernetAllOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRrdEthernetAllOper,error) {
logger.Println("RrdEthernetAllOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRrdEthernetAllOper
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
