

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnPortReservationEntriesOper struct {
    
    Oper Cgnv6LsnPortReservationEntriesOperOper `json:"oper"`

}
type DataCgnv6LsnPortReservationEntriesOper struct {
    DtCgnv6LsnPortReservationEntriesOper Cgnv6LsnPortReservationEntriesOper `json:"port-reservation-entries"`
}


type Cgnv6LsnPortReservationEntriesOperOper struct {
    EntryList []Cgnv6LsnPortReservationEntriesOperOperEntryList `json:"entry-list"`
    EntryCount int `json:"entry-count"`
}


type Cgnv6LsnPortReservationEntriesOperOperEntryList struct {
    InsideAddress string `json:"inside-address"`
    InsideStartPort int `json:"inside-start-port"`
    InsideEndPort int `json:"inside-end-port"`
    NatAddress string `json:"nat-address"`
    NatStartPort int `json:"nat-start-port"`
    NatEndPort int `json:"nat-end-port"`
}

func (p *Cgnv6LsnPortReservationEntriesOper) GetId() string{
    return "1"
}

func (p *Cgnv6LsnPortReservationEntriesOper) getPath() string{
    return "cgnv6/lsn/port-reservation-entries/oper"
}

func (p *Cgnv6LsnPortReservationEntriesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnPortReservationEntriesOper,error) {
logger.Println("Cgnv6LsnPortReservationEntriesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnPortReservationEntriesOper
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
