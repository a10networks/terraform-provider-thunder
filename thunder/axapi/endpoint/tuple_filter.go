

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type TupleFilter struct {
	Inst struct {

    FilterRuleList []TupleFilterFilterRuleList `json:"filter-rule-list"`
    FilterType string `json:"filter-type"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"tuple-filter"`
}


type TupleFilterFilterRuleList struct {
    Id1 int `json:"id"`
    SrcAddr string `json:"src-addr"`
    DstAddr string `json:"dst-addr"`
    SrcPort int `json:"src-port"`
    DstPort int `json:"dst-port"`
    SrcV6Addr string `json:"src-v6-addr"`
    DstV6Addr string `json:"dst-v6-addr"`
    SrcV6Port int `json:"src-v6-port"`
    DstV6Port int `json:"dst-v6-port"`
    Uuid string `json:"uuid"`
}

func (p *TupleFilter) GetId() string{
    return p.Inst.Name
}

func (p *TupleFilter) getPath() string{
    return "tuple-filter"
}

func (p *TupleFilter) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TupleFilter::Post")
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

func (p *TupleFilter) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TupleFilter::Get")
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
func (p *TupleFilter) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TupleFilter::Put")
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

func (p *TupleFilter) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TupleFilter::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
