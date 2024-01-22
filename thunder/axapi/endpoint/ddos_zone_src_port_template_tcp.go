

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneSrcPortTemplateTcp struct {
	Inst struct {

    FilterList []DdosZoneSrcPortTemplateTcpFilterList `json:"filter-list"`
    FilterMatchType string `json:"filter-match-type" dval:"default"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"tcp"`
}


type DdosZoneSrcPortTemplateTcpFilterList struct {
    TcpFilterName string `json:"tcp-filter-name"`
    TcpFilterSeq int `json:"tcp-filter-seq"`
    TcpFilterRegex string `json:"tcp-filter-regex"`
    TcpFilterInverseMatch int `json:"tcp-filter-inverse-match"`
    ByteOffsetFilter string `json:"byte-offset-filter"`
    TcpFilterActionListName string `json:"tcp-filter-action-list-name"`
    TcpFilterAction string `json:"tcp-filter-action"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}

func (p *DdosZoneSrcPortTemplateTcp) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosZoneSrcPortTemplateTcp) getPath() string{
    return "ddos/zone-src-port-template/tcp"
}

func (p *DdosZoneSrcPortTemplateTcp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateTcp::Post")
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

func (p *DdosZoneSrcPortTemplateTcp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateTcp::Get")
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
func (p *DdosZoneSrcPortTemplateTcp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateTcp::Put")
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

func (p *DdosZoneSrcPortTemplateTcp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateTcp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
