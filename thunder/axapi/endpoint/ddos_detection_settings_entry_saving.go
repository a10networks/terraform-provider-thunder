

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDetectionSettingsEntrySaving struct {
	Inst struct {

    Interval int `json:"interval"`
    ManualRestore int `json:"manual-restore"`
    ManualSave int `json:"manual-save"`
    Uuid string `json:"uuid"`

	} `json:"entry-saving"`
}

func (p *DdosDetectionSettingsEntrySaving) GetId() string{
    return "1"
}

func (p *DdosDetectionSettingsEntrySaving) getPath() string{
    return "ddos/detection/settings/entry-saving"
}

func (p *DdosDetectionSettingsEntrySaving) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionSettingsEntrySaving::Post")
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

func (p *DdosDetectionSettingsEntrySaving) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionSettingsEntrySaving::Get")
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
func (p *DdosDetectionSettingsEntrySaving) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionSettingsEntrySaving::Put")
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

func (p *DdosDetectionSettingsEntrySaving) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDetectionSettingsEntrySaving::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
