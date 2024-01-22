

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosDnsCacheZoneManualOverrideAction struct {
	Inst struct {

    Action string `json:"action"`
    ZoneName string `json:"zone-name"`
    Name string 

	} `json:"zone-manual-override-action"`
}

func (p *DdosDnsCacheZoneManualOverrideAction) GetId() string{
    return url.QueryEscape(p.Inst.ZoneName)
}

func (p *DdosDnsCacheZoneManualOverrideAction) getPath() string{
    return "ddos/dns-cache/" +p.Inst.Name + "/zone-manual-override-action"
}

func (p *DdosDnsCacheZoneManualOverrideAction) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheZoneManualOverrideAction::Post")
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

func (p *DdosDnsCacheZoneManualOverrideAction) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheZoneManualOverrideAction::Get")
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
func (p *DdosDnsCacheZoneManualOverrideAction) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheZoneManualOverrideAction::Put")
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

func (p *DdosDnsCacheZoneManualOverrideAction) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDnsCacheZoneManualOverrideAction::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
