

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateDynamicService struct {
	Inst struct {

    ClassListList []SlbTemplateDynamicServiceClassListList `json:"class-list-list"`
    DnsServer []SlbTemplateDynamicServiceDnsServer `json:"dns-server"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"dynamic-service"`
}


type SlbTemplateDynamicServiceClassListList struct {
    DnsClassList string `json:"dns-class-list"`
    Priority int `json:"priority"`
    DnsServer []SlbTemplateDynamicServiceClassListListDnsServer `json:"dns-server"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type SlbTemplateDynamicServiceClassListListDnsServer struct {
    Ipv4DnsServer string `json:"ipv4-dns-server"`
    Ipv6DnsServer string `json:"ipv6-dns-server"`
}


type SlbTemplateDynamicServiceDnsServer struct {
    Ipv4DnsServer string `json:"ipv4-dns-server"`
    Ipv6DnsServer string `json:"ipv6-dns-server"`
}

func (p *SlbTemplateDynamicService) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateDynamicService) getPath() string{
    return "slb/template/dynamic-service"
}

func (p *SlbTemplateDynamicService) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDynamicService::Post")
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

func (p *SlbTemplateDynamicService) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDynamicService::Get")
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
func (p *SlbTemplateDynamicService) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDynamicService::Put")
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

func (p *SlbTemplateDynamicService) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDynamicService::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
