

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateTcpFilter struct {
	Inst struct {

    ByteOffsetFilter string `json:"byte-offset-filter"`
    TcpFilterAction string `json:"tcp-filter-action"`
    TcpFilterRegex string `json:"tcp-filter-regex"`
    TcpFilterSeq int `json:"tcp-filter-seq"`
    TcpFilterUnmatched int `json:"tcp-filter-unmatched"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"filter"`
}

func (p *DdosTemplateTcpFilter) GetId() string{
    return strconv.Itoa(p.Inst.TcpFilterSeq)
}

func (p *DdosTemplateTcpFilter) getPath() string{
    return "ddos/template/tcp/" +p.Inst.Name + "/filter"
}

func (p *DdosTemplateTcpFilter) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateTcpFilter::Post")
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

func (p *DdosTemplateTcpFilter) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateTcpFilter::Get")
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
func (p *DdosTemplateTcpFilter) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateTcpFilter::Put")
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

func (p *DdosTemplateTcpFilter) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateTcpFilter::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
