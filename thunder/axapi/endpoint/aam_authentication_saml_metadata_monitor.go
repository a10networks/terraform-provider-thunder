

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AamAuthenticationSamlMetadataMonitor struct {
	Inst struct {

    AcsContinuousFailThreshold int `json:"acs-continuous-fail-threshold"`
    AcsMissingPeriod int `json:"acs-missing-period"`
    AcsMissingThreshold int `json:"acs-missing-threshold"`
    Status string `json:"status" dval:"enable"`
    Uuid string `json:"uuid"`

	} `json:"metadata-monitor"`
}

func (p *AamAuthenticationSamlMetadataMonitor) GetId() string{
    return "1"
}

func (p *AamAuthenticationSamlMetadataMonitor) getPath() string{
    return "aam/authentication/saml/metadata-monitor"
}

func (p *AamAuthenticationSamlMetadataMonitor) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationSamlMetadataMonitor::Post")
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

func (p *AamAuthenticationSamlMetadataMonitor) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationSamlMetadataMonitor::Get")
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
func (p *AamAuthenticationSamlMetadataMonitor) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationSamlMetadataMonitor::Put")
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

func (p *AamAuthenticationSamlMetadataMonitor) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AamAuthenticationSamlMetadataMonitor::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
