

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateDnsLocalDnsResolution struct {
	Inst struct {

    HostListCfg []SlbTemplateDnsLocalDnsResolutionHostListCfg `json:"host-list-cfg"`
    LocalResolverCfg []SlbTemplateDnsLocalDnsResolutionLocalResolverCfg `json:"local-resolver-cfg"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"local-dns-resolution"`
}


type SlbTemplateDnsLocalDnsResolutionHostListCfg struct {
    Hostnames string `json:"hostnames"`
}


type SlbTemplateDnsLocalDnsResolutionLocalResolverCfg struct {
    LocalResolver string `json:"local-resolver"`
}

func (p *SlbTemplateDnsLocalDnsResolution) GetId() string{
    return "1"
}

func (p *SlbTemplateDnsLocalDnsResolution) getPath() string{
    return "slb/template/dns/" +p.Inst.Name + "/local-dns-resolution"
}

func (p *SlbTemplateDnsLocalDnsResolution) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsLocalDnsResolution::Post")
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

func (p *SlbTemplateDnsLocalDnsResolution) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsLocalDnsResolution::Get")
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
func (p *SlbTemplateDnsLocalDnsResolution) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsLocalDnsResolution::Put")
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

func (p *SlbTemplateDnsLocalDnsResolution) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsLocalDnsResolution::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
