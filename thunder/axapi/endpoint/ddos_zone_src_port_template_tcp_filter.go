

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneSrcPortTemplateTcpFilter struct {
	Inst struct {

    ByteOffsetFilter string `json:"byte-offset-filter"`
    TcpFilterAction string `json:"tcp-filter-action"`
    TcpFilterActionListName string `json:"tcp-filter-action-list-name"`
    TcpFilterInverseMatch int `json:"tcp-filter-inverse-match"`
    TcpFilterName string `json:"tcp-filter-name"`
    TcpFilterRegex string `json:"tcp-filter-regex"`
    TcpFilterSeq int `json:"tcp-filter-seq"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"filter"`
}

func (p *DdosZoneSrcPortTemplateTcpFilter) GetId() string{
    return url.QueryEscape(p.Inst.TcpFilterName)
}

func (p *DdosZoneSrcPortTemplateTcpFilter) getPath() string{
    return "ddos/zone-src-port-template/tcp/" +p.Inst.Name + "/filter"
}

func (p *DdosZoneSrcPortTemplateTcpFilter) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateTcpFilter::Post")
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

func (p *DdosZoneSrcPortTemplateTcpFilter) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateTcpFilter::Get")
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
func (p *DdosZoneSrcPortTemplateTcpFilter) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateTcpFilter::Put")
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

func (p *DdosZoneSrcPortTemplateTcpFilter) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateTcpFilter::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
