

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SystemMonTemplateMonitor struct {
	Inst struct {

    ClearCfg []SystemMonTemplateMonitorClearCfg `json:"clear-cfg"`
    Id1 int `json:"id1"`
    LinkDisableCfg []SystemMonTemplateMonitorLinkDisableCfg `json:"link-disable-cfg"`
    LinkDownCfg []SystemMonTemplateMonitorLinkDownCfg `json:"link-down-cfg"`
    LinkEnableCfg []SystemMonTemplateMonitorLinkEnableCfg `json:"link-enable-cfg"`
    LinkUpCfg []SystemMonTemplateMonitorLinkUpCfg `json:"link-up-cfg"`
    MonitorRelation string `json:"monitor-relation" dval:"monitor-and"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"monitor"`
}


type SystemMonTemplateMonitorClearCfg struct {
    Sessions string `json:"sessions"`
    ClearAllSequence int `json:"clear-all-sequence"`
    ClearSequence int `json:"clear-sequence"`
}


type SystemMonTemplateMonitorLinkDisableCfg struct {
    Diseth int `json:"diseth"`
    DisSequence int `json:"dis-sequence"`
}


type SystemMonTemplateMonitorLinkDownCfg struct {
    LinkdownEthernet1 int `json:"linkdown-ethernet1"`
    LinkDownSequence1 int `json:"link-down-sequence1"`
    LinkdownEthernet2 int `json:"linkdown-ethernet2"`
    LinkDownSequence2 int `json:"link-down-sequence2"`
    LinkdownEthernet3 int `json:"linkdown-ethernet3"`
    LinkDownSequence3 int `json:"link-down-sequence3"`
}


type SystemMonTemplateMonitorLinkEnableCfg struct {
    Enaeth int `json:"enaeth"`
    EnaSequence int `json:"ena-sequence"`
}


type SystemMonTemplateMonitorLinkUpCfg struct {
    LinkupEthernet1 int `json:"linkup-ethernet1"`
    LinkUpSequence1 int `json:"link-up-sequence1"`
    LinkupEthernet2 int `json:"linkup-ethernet2"`
    LinkUpSequence2 int `json:"link-up-sequence2"`
    LinkupEthernet3 int `json:"linkup-ethernet3"`
    LinkUpSequence3 int `json:"link-up-sequence3"`
}

func (p *SystemMonTemplateMonitor) GetId() string{
    return strconv.Itoa(p.Inst.Id1)
}

func (p *SystemMonTemplateMonitor) getPath() string{
    return "system/mon-template/monitor"
}

func (p *SystemMonTemplateMonitor) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemMonTemplateMonitor::Post")
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

func (p *SystemMonTemplateMonitor) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemMonTemplateMonitor::Get")
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
func (p *SystemMonTemplateMonitor) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemMonTemplateMonitor::Put")
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

func (p *SystemMonTemplateMonitor) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemMonTemplateMonitor::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
