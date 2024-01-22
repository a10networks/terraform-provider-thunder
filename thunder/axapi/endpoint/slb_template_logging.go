

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateLogging struct {
	Inst struct {

    Auto string `json:"auto" dval:"auto"`
    Format string `json:"format"`
    KeepEnd int `json:"keep-end"`
    KeepStart int `json:"keep-start"`
    LocalLogging int `json:"local-logging"`
    Mask string `json:"mask" dval:"X"`
    Name string `json:"name"`
    PcreMask string `json:"pcre-mask"`
    Pool string `json:"pool"`
    PoolShared string `json:"pool-shared"`
    ServiceGroup string `json:"service-group"`
    SharedPartitionPool int `json:"shared-partition-pool"`
    SharedPartitionTcpProxyTemplate int `json:"shared-partition-tcp-proxy-template"`
    TcpProxy string `json:"tcp-proxy"`
    TemplateTcpProxyShared string `json:"template-tcp-proxy-shared"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"logging"`
}

func (p *SlbTemplateLogging) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateLogging) getPath() string{
    return "slb/template/logging"
}

func (p *SlbTemplateLogging) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateLogging::Post")
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

func (p *SlbTemplateLogging) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateLogging::Get")
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
func (p *SlbTemplateLogging) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateLogging::Put")
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

func (p *SlbTemplateLogging) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateLogging::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
