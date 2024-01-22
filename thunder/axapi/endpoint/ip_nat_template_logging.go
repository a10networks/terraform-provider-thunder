

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type IpNatTemplateLogging struct {
	Inst struct {

    Facility string `json:"facility" dval:"local0"`
    IncludeDestination int `json:"include-destination"`
    IncludeRipRport int `json:"include-rip-rport"`
    Log IpNatTemplateLoggingLog `json:"log"`
    Name string `json:"name"`
    ServiceGroup string `json:"service-group"`
    Severity IpNatTemplateLoggingSeverity `json:"severity"`
    SourcePort IpNatTemplateLoggingSourcePort `json:"source-port"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"logging"`
}


type IpNatTemplateLoggingLog struct {
    PortMappings string `json:"port-mappings"`
}


type IpNatTemplateLoggingSeverity struct {
    SeverityString string `json:"severity-string" dval:"debug"`
    SeverityVal int `json:"severity-val" dval:"7"`
}


type IpNatTemplateLoggingSourcePort struct {
    SourcePortNum int `json:"source-port-num" dval:"514"`
    Any int `json:"any"`
}

func (p *IpNatTemplateLogging) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *IpNatTemplateLogging) getPath() string{
    return "ip/nat/template/logging"
}

func (p *IpNatTemplateLogging) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatTemplateLogging::Post")
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

func (p *IpNatTemplateLogging) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatTemplateLogging::Get")
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
func (p *IpNatTemplateLogging) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatTemplateLogging::Put")
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

func (p *IpNatTemplateLogging) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatTemplateLogging::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
