

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GlmCreateLicenseRequest struct {
	Inst struct {

    CreateLicenseRequest int `json:"create-license-request"`
    Uuid string `json:"uuid"`

	} `json:"create-license-request"`
}

func (p *GlmCreateLicenseRequest) GetId() string{
    return "1"
}

func (p *GlmCreateLicenseRequest) getPath() string{
    return "glm/create-license-request"
}

func (p *GlmCreateLicenseRequest) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GlmCreateLicenseRequest::Post")
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

func (p *GlmCreateLicenseRequest) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GlmCreateLicenseRequest::Get")
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
func (p *GlmCreateLicenseRequest) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GlmCreateLicenseRequest::Put")
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

func (p *GlmCreateLicenseRequest) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GlmCreateLicenseRequest::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
