

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ThreatIntelThreatList struct {
	Inst struct {

    AllCategories int `json:"all-categories"`
    Botnets int `json:"botnets"`
    DosAttacks int `json:"dos-attacks"`
    MobileThreats int `json:"mobile-threats"`
    Name string `json:"name"`
    Phishing int `json:"phishing"`
    Proxy int `json:"proxy"`
    Reputation int `json:"reputation"`
    SamplingEnable []ThreatIntelThreatListSamplingEnable `json:"sampling-enable"`
    Scanners int `json:"scanners"`
    SpamSources int `json:"spam-sources"`
    TorProxy int `json:"tor-proxy"`
    Type string `json:"type"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    WebAttacks int `json:"web-attacks"`
    WindowsExploits int `json:"windows-exploits"`

	} `json:"threat-list"`
}


type ThreatIntelThreatListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *ThreatIntelThreatList) GetId() string{
    return p.Inst.Name
}

func (p *ThreatIntelThreatList) getPath() string{
    return "threat-intel/threat-list"
}

func (p *ThreatIntelThreatList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ThreatIntelThreatList::Post")
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

func (p *ThreatIntelThreatList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ThreatIntelThreatList::Get")
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
func (p *ThreatIntelThreatList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ThreatIntelThreatList::Put")
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

func (p *ThreatIntelThreatList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ThreatIntelThreatList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
