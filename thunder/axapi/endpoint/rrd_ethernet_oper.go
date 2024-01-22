

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RrdEthernetOper struct {
    
    Oper RrdEthernetOperOper `json:"oper"`

}
type DataRrdEthernetOper struct {
    DtRrdEthernetOper RrdEthernetOper `json:"ethernet"`
}


type RrdEthernetOperOper struct {
    StartTime int `json:"start-time"`
    EndTime int `json:"end-time"`
    EthernetIndex int `json:"ethernet-index"`
    EthernetStatistics []RrdEthernetOperOperEthernetStatistics `json:"ethernet-statistics"`
}


type RrdEthernetOperOperEthernetStatistics struct {
    Time int `json:"time"`
    Rx_packets int `json:"rx_packets"`
    Rx_bits int `json:"rx_bits"`
    Rx_error int `json:"rx_error"`
    Rx_drop int `json:"rx_drop"`
    Tx_packets int `json:"tx_packets"`
    Tx_bits int `json:"tx_bits"`
    Tx_error int `json:"tx_error"`
    Tx_drop int `json:"tx_drop"`
}

func (p *RrdEthernetOper) GetId() string{
    return "1"
}

func (p *RrdEthernetOper) getPath() string{
    return "rrd/ethernet/oper"
}

func (p *RrdEthernetOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRrdEthernetOper,error) {
logger.Println("RrdEthernetOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRrdEthernetOper
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
