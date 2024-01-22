

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDnsCacheDomainGroupDomainListPolicyPacketCapturing struct {
	Inst struct {

    RootZoneList []DdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRootZoneList `json:"root-zone-list"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"packet-capturing"`
}


type DdosDnsCacheDomainGroupDomainListPolicyPacketCapturingRootZoneList struct {
    RootZone string `json:"root-zone"`
    CaptureConfig string `json:"capture-config"`
    CaptureMode string `json:"capture-mode"`
}

func (p *DdosDnsCacheDomainGroupDomainListPolicyPacketCapturing) GetId() string{
    return "1"
}

func (p *DdosDnsCacheDomainGroupDomainListPolicyPacketCapturing) getPath() string{
    return "ddos/dns-cache/" +p.Inst.Name + "/domain-group/domain-list-policy/" +p.Inst.Name + "/packet-capturing"
}

func (p *DdosDnsCacheDomainGroupDomainListPolicyPacketCapturing) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheDomainGroupDomainListPolicyPacketCapturing::Post")
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

func (p *DdosDnsCacheDomainGroupDomainListPolicyPacketCapturing) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheDomainGroupDomainListPolicyPacketCapturing::Get")
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
func (p *DdosDnsCacheDomainGroupDomainListPolicyPacketCapturing) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheDomainGroupDomainListPolicyPacketCapturing::Put")
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

func (p *DdosDnsCacheDomainGroupDomainListPolicyPacketCapturing) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheDomainGroupDomainListPolicyPacketCapturing::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
