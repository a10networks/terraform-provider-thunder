

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosSrcPortTemplateTcpFilter struct {
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

func (p *DdosSrcPortTemplateTcpFilter) GetId() string{
    return strconv.Itoa(p.Inst.TcpFilterSeq)
}

func (p *DdosSrcPortTemplateTcpFilter) getPath() string{
    return "ddos/src-port-template/tcp/" +p.Inst.Name + "/filter"
}

func (p *DdosSrcPortTemplateTcpFilter) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcPortTemplateTcpFilter::Post")
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

func (p *DdosSrcPortTemplateTcpFilter) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcPortTemplateTcpFilter::Get")
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
func (p *DdosSrcPortTemplateTcpFilter) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcPortTemplateTcpFilter::Put")
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

func (p *DdosSrcPortTemplateTcpFilter) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcPortTemplateTcpFilter::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
