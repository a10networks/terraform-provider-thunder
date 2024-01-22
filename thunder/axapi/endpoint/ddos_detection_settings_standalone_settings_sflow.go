

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDetectionSettingsStandaloneSettingsSflow struct {
	Inst struct {

    ListeningPort int `json:"listening-port" dval:"6343"`
    Uuid string `json:"uuid"`

	} `json:"sflow"`
}

func (p *DdosDetectionSettingsStandaloneSettingsSflow) GetId() string{
    return "1"
}

func (p *DdosDetectionSettingsStandaloneSettingsSflow) getPath() string{
    return "ddos/detection/settings/standalone-settings/sflow"
}

func (p *DdosDetectionSettingsStandaloneSettingsSflow) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionSettingsStandaloneSettingsSflow::Post")
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

func (p *DdosDetectionSettingsStandaloneSettingsSflow) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionSettingsStandaloneSettingsSflow::Get")
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
func (p *DdosDetectionSettingsStandaloneSettingsSflow) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionSettingsStandaloneSettingsSflow::Put")
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

func (p *DdosDetectionSettingsStandaloneSettingsSflow) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionSettingsStandaloneSettingsSflow::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
