

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6DsLitePortReservation struct {
	Inst struct {

    Inside string `json:"inside"`
    InsideAddr string `json:"inside-addr"`
    InsideEndPort int `json:"inside-end-port"`
    InsideStartPort int `json:"inside-start-port"`
    Nat string `json:"nat"`
    NatEndPort int `json:"nat-end-port"`
    NatStartPort int `json:"nat-start-port"`
    TunnelDestAddress string `json:"tunnel-dest-address"`
    Uuid string `json:"uuid"`

	} `json:"port-reservation"`
}

func (p *Cgnv6DsLitePortReservation) GetId() string{
    return p.Inst.Inside+"+"+p.Inst.TunnelDestAddress+"+"+p.Inst.InsideAddr+"+"+strconv.Itoa(p.Inst.InsideStartPort)+"+"+strconv.Itoa(p.Inst.InsideEndPort)+"+"+p.Inst.Nat+"+"+strconv.Itoa(p.Inst.NatStartPort)+"+"+strconv.Itoa(p.Inst.NatEndPort)
}

func (p *Cgnv6DsLitePortReservation) getPath() string{
    return "cgnv6/ds-lite/port-reservation"
}

func (p *Cgnv6DsLitePortReservation) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6DsLitePortReservation::Post")
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

func (p *Cgnv6DsLitePortReservation) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6DsLitePortReservation::Get")
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
func (p *Cgnv6DsLitePortReservation) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6DsLitePortReservation::Put")
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

func (p *Cgnv6DsLitePortReservation) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6DsLitePortReservation::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
