

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateMonitor struct {
	Inst struct {

    ClearCfg []SlbTemplateMonitorClearCfg `json:"clear-cfg"`
    Id1 int `json:"id1"`
    LinkDisableCfg []SlbTemplateMonitorLinkDisableCfg `json:"link-disable-cfg"`
    LinkDownCfg []SlbTemplateMonitorLinkDownCfg `json:"link-down-cfg"`
    LinkEnableCfg []SlbTemplateMonitorLinkEnableCfg `json:"link-enable-cfg"`
    LinkUpCfg []SlbTemplateMonitorLinkUpCfg `json:"link-up-cfg"`
    MonitorRelation string `json:"monitor-relation" dval:"monitor-and"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"monitor"`
}


type SlbTemplateMonitorClearCfg struct {
    Sessions string `json:"sessions"`
    ClearAllSequence int `json:"clear-all-sequence"`
    ClearSequence int `json:"clear-sequence"`
}


type SlbTemplateMonitorLinkDisableCfg struct {
    Diseth int `json:"diseth"`
    DisSequence int `json:"dis-sequence"`
}


type SlbTemplateMonitorLinkDownCfg struct {
    LinkdownEthernet1 int `json:"linkdown-ethernet1"`
    LinkDownSequence1 int `json:"link-down-sequence1"`
    LinkdownEthernet2 int `json:"linkdown-ethernet2"`
    LinkDownSequence2 int `json:"link-down-sequence2"`
    LinkdownEthernet3 int `json:"linkdown-ethernet3"`
    LinkDownSequence3 int `json:"link-down-sequence3"`
}


type SlbTemplateMonitorLinkEnableCfg struct {
    Enaeth int `json:"enaeth"`
    EnaSequence int `json:"ena-sequence"`
}


type SlbTemplateMonitorLinkUpCfg struct {
    LinkupEthernet1 int `json:"linkup-ethernet1"`
    LinkUpSequence1 int `json:"link-up-sequence1"`
    LinkupEthernet2 int `json:"linkup-ethernet2"`
    LinkUpSequence2 int `json:"link-up-sequence2"`
    LinkupEthernet3 int `json:"linkup-ethernet3"`
    LinkUpSequence3 int `json:"link-up-sequence3"`
}

func (p *SlbTemplateMonitor) GetId() string{
    return strconv.Itoa(p.Inst.Id1)
}

func (p *SlbTemplateMonitor) getPath() string{
    return "slb/template/monitor"
}

func (p *SlbTemplateMonitor) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateMonitor::Post")
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

func (p *SlbTemplateMonitor) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateMonitor::Get")
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
func (p *SlbTemplateMonitor) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateMonitor::Put")
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

func (p *SlbTemplateMonitor) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateMonitor::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
