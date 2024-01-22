

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LicenseManagerOverage struct {
	Inst struct {

    Bytes int `json:"bytes"`
    Days int `json:"days"`
    Gb int `json:"gb"`
    Hours int `json:"hours"`
    Kb int `json:"kb"`
    Mb int `json:"mb"`
    Minutes int `json:"minutes"`
    Seconds int `json:"seconds"`
    Uuid string `json:"uuid"`

	} `json:"overage"`
}

func (p *LicenseManagerOverage) GetId() string{
    return "1"
}

func (p *LicenseManagerOverage) getPath() string{
    return "license-manager/overage"
}

func (p *LicenseManagerOverage) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LicenseManagerOverage::Post")
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

func (p *LicenseManagerOverage) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LicenseManagerOverage::Get")
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
func (p *LicenseManagerOverage) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LicenseManagerOverage::Put")
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

func (p *LicenseManagerOverage) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LicenseManagerOverage::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
