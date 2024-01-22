

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDetectionSettingsStandaloneSettings struct {
	Inst struct {

    Action string `json:"action" dval:"disable"`
    DeEscalationQuietTime int `json:"de-escalation-quiet-time"`
    Netflow DdosDetectionSettingsStandaloneSettingsNetflow130 `json:"netflow"`
    Sflow DdosDetectionSettingsStandaloneSettingsSflow131 `json:"sflow"`
    Uuid string `json:"uuid"`

	} `json:"standalone-settings"`
}


type DdosDetectionSettingsStandaloneSettingsNetflow130 struct {
    ListeningPort int `json:"listening-port" dval:"9996"`
    TemplateActiveTimeout int `json:"template-active-timeout" dval:"30"`
    Uuid string `json:"uuid"`
}


type DdosDetectionSettingsStandaloneSettingsSflow131 struct {
    ListeningPort int `json:"listening-port" dval:"6343"`
    Uuid string `json:"uuid"`
}

func (p *DdosDetectionSettingsStandaloneSettings) GetId() string{
    return "1"
}

func (p *DdosDetectionSettingsStandaloneSettings) getPath() string{
    return "ddos/detection/settings/standalone-settings"
}

func (p *DdosDetectionSettingsStandaloneSettings) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionSettingsStandaloneSettings::Post")
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

func (p *DdosDetectionSettingsStandaloneSettings) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionSettingsStandaloneSettings::Get")
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
func (p *DdosDetectionSettingsStandaloneSettings) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionSettingsStandaloneSettings::Put")
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

func (p *DdosDetectionSettingsStandaloneSettings) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionSettingsStandaloneSettings::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
