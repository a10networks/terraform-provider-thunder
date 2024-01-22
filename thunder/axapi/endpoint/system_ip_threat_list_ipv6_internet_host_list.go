

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemIpThreatListIpv6InternetHostList struct {
	Inst struct {

    ClassListCfg []SystemIpThreatListIpv6InternetHostListClassListCfg `json:"class-list-cfg"`
    Uuid string `json:"uuid"`

	} `json:"ipv6-internet-host-list"`
}


type SystemIpThreatListIpv6InternetHostListClassListCfg struct {
    ClassList string `json:"class-list"`
    IpThreatActionTmpl int `json:"ip-threat-action-tmpl"`
}

func (p *SystemIpThreatListIpv6InternetHostList) GetId() string{
    return "1"
}

func (p *SystemIpThreatListIpv6InternetHostList) getPath() string{
    return "system/ip-threat-list/ipv6-internet-host-list"
}

func (p *SystemIpThreatListIpv6InternetHostList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpThreatListIpv6InternetHostList::Post")
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

func (p *SystemIpThreatListIpv6InternetHostList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpThreatListIpv6InternetHostList::Get")
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
func (p *SystemIpThreatListIpv6InternetHostList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpThreatListIpv6InternetHostList::Put")
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

func (p *SystemIpThreatListIpv6InternetHostList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpThreatListIpv6InternetHostList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
