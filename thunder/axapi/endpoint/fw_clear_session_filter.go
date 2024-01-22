

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwClearSessionFilter struct {
	Inst struct {

    Status string `json:"status" dval:"disable"`
    Uuid string `json:"uuid"`

	} `json:"clear-session-filter"`
}

func (p *FwClearSessionFilter) GetId() string{
    return "1"
}

func (p *FwClearSessionFilter) getPath() string{
    return "fw/clear-session-filter"
}

func (p *FwClearSessionFilter) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwClearSessionFilter::Post")
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

func (p *FwClearSessionFilter) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwClearSessionFilter::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *FwClearSessionFilter) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwClearSessionFilter::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *FwClearSessionFilter) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwClearSessionFilter::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
