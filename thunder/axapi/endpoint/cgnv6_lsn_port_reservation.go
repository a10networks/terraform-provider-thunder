

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnPortReservation struct {
	Inst struct {

    Inside string `json:"inside"`
    InsidePortEnd int `json:"inside-port-end"`
    InsidePortStart int `json:"inside-port-start"`
    Nat string `json:"nat"`
    NatPortEnd int `json:"nat-port-end"`
    NatPortStart int `json:"nat-port-start"`
    Uuid string `json:"uuid"`

	} `json:"port-reservation"`
}

func (p *Cgnv6LsnPortReservation) GetId() string{
    return p.Inst.Inside+"+"+strconv.Itoa(p.Inst.InsidePortStart)+"+"+strconv.Itoa(p.Inst.InsidePortEnd)+"+"+p.Inst.Nat+"+"+strconv.Itoa(p.Inst.NatPortStart)+"+"+strconv.Itoa(p.Inst.NatPortEnd)
}

func (p *Cgnv6LsnPortReservation) getPath() string{
    return "cgnv6/lsn/port-reservation"
}

func (p *Cgnv6LsnPortReservation) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnPortReservation::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *Cgnv6LsnPortReservation) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnPortReservation::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *Cgnv6LsnPortReservation) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnPortReservation::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *Cgnv6LsnPortReservation) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnPortReservation::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
