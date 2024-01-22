

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDetectionSettingsStandaloneSettingsNetflow struct {
	Inst struct {

    ListeningPort int `json:"listening-port" dval:"9996"`
    TemplateActiveTimeout int `json:"template-active-timeout" dval:"30"`
    Uuid string `json:"uuid"`

	} `json:"netflow"`
}

func (p *DdosDetectionSettingsStandaloneSettingsNetflow) GetId() string{
    return "1"
}

func (p *DdosDetectionSettingsStandaloneSettingsNetflow) getPath() string{
    return "ddos/detection/settings/standalone-settings/netflow"
}

func (p *DdosDetectionSettingsStandaloneSettingsNetflow) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionSettingsStandaloneSettingsNetflow::Post")
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

func (p *DdosDetectionSettingsStandaloneSettingsNetflow) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionSettingsStandaloneSettingsNetflow::Get")
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
func (p *DdosDetectionSettingsStandaloneSettingsNetflow) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionSettingsStandaloneSettingsNetflow::Put")
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

func (p *DdosDetectionSettingsStandaloneSettingsNetflow) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionSettingsStandaloneSettingsNetflow::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
