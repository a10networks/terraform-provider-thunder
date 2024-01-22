

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateLinkProbeDestination struct {
	Inst struct {

    Hostname string `json:"hostname"`
    ResolveAs string `json:"resolve-as"`
    StaticIpv4Addr string `json:"static-ipv4-addr"`
    StaticIpv6Addr string `json:"static-ipv6-addr"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"destination"`
}

func (p *SlbTemplateLinkProbeDestination) GetId() string{
    return "1"
}

func (p *SlbTemplateLinkProbeDestination) getPath() string{
    return "slb/template/link-probe/" +p.Inst.Name + "/destination"
}

func (p *SlbTemplateLinkProbeDestination) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateLinkProbeDestination::Post")
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

func (p *SlbTemplateLinkProbeDestination) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateLinkProbeDestination::Get")
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
func (p *SlbTemplateLinkProbeDestination) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateLinkProbeDestination::Put")
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

func (p *SlbTemplateLinkProbeDestination) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateLinkProbeDestination::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
