

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type TupleFilterFilterRule struct {
	Inst struct {

    DstAddr string `json:"dst-addr"`
    DstPort int `json:"dst-port"`
    DstV6Addr string `json:"dst-v6-addr"`
    DstV6Port int `json:"dst-v6-port"`
    Id1 int `json:"id"`
    SrcAddr string `json:"src-addr"`
    SrcPort int `json:"src-port"`
    SrcV6Addr string `json:"src-v6-addr"`
    SrcV6Port int `json:"src-v6-port"`
    Uuid string `json:"uuid"`
    Name string

	} `json:"filter-rule"`
}

func (p *TupleFilterFilterRule) GetId() string{
    return strconv.Itoa(p.Inst.Id1)
}

func (p *TupleFilterFilterRule) getPath() string{
    return "tuple-filter/" +p.Inst.Name + "/filter-rule"
}

func (p *TupleFilterFilterRule) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TupleFilterFilterRule::Post")
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

func (p *TupleFilterFilterRule) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TupleFilterFilterRule::Get")
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
func (p *TupleFilterFilterRule) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TupleFilterFilterRule::Put")
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

func (p *TupleFilterFilterRule) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TupleFilterFilterRule::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
