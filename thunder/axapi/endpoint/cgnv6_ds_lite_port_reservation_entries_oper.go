

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6DsLitePortReservationEntriesOper struct {
    
    Oper Cgnv6DsLitePortReservationEntriesOperOper `json:"oper"`

}
type DataCgnv6DsLitePortReservationEntriesOper struct {
    DtCgnv6DsLitePortReservationEntriesOper Cgnv6DsLitePortReservationEntriesOper `json:"port-reservation-entries"`
}


type Cgnv6DsLitePortReservationEntriesOperOper struct {
    EntryList []Cgnv6DsLitePortReservationEntriesOperOperEntryList `json:"entry-list"`
    EntryCount int `json:"entry-count"`
}


type Cgnv6DsLitePortReservationEntriesOperOperEntryList struct {
    TunnelSrcAddress string `json:"tunnel-src-address"`
    TunnelDestAddress string `json:"tunnel-dest-address"`
    InsideAddress string `json:"inside-address"`
    InsideStartPort int `json:"inside-start-port"`
    InsideEndPort int `json:"inside-end-port"`
    NatAddress string `json:"nat-address"`
    NatStartPort int `json:"nat-start-port"`
    NatEndPort int `json:"nat-end-port"`
}

func (p *Cgnv6DsLitePortReservationEntriesOper) GetId() string{
    return "1"
}

func (p *Cgnv6DsLitePortReservationEntriesOper) getPath() string{
    return "cgnv6/ds-lite/port-reservation-entries/oper"
}

func (p *Cgnv6DsLitePortReservationEntriesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6DsLitePortReservationEntriesOper,error) {
logger.Println("Cgnv6DsLitePortReservationEntriesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6DsLitePortReservationEntriesOper
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
