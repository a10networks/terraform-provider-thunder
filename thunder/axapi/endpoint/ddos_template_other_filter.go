

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateOtherFilter struct {
	Inst struct {

    ByteOffsetFilter string `json:"byte-offset-filter"`
    OtherFilterAction string `json:"other-filter-action"`
    OtherFilterRegex string `json:"other-filter-regex"`
    OtherFilterSeq int `json:"other-filter-seq"`
    OtherFilterUnmatched int `json:"other-filter-unmatched"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"filter"`
}

func (p *DdosTemplateOtherFilter) GetId() string{
    return strconv.Itoa(p.Inst.OtherFilterSeq)
}

func (p *DdosTemplateOtherFilter) getPath() string{
    return "ddos/template/other/" +p.Inst.Name + "/filter"
}

func (p *DdosTemplateOtherFilter) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateOtherFilter::Post")
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

func (p *DdosTemplateOtherFilter) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateOtherFilter::Get")
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
func (p *DdosTemplateOtherFilter) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateOtherFilter::Put")
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

func (p *DdosTemplateOtherFilter) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateOtherFilter::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
