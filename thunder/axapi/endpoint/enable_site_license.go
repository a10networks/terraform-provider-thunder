

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnableSiteLicense struct {
	Inst struct {

    HashKey string `json:"hash-key"`

	} `json:"enable-site-license"`
}

func (p *EnableSiteLicense) GetId() string{
    return p.Inst.HashKey
}

func (p *EnableSiteLicense) getPath() string{
    return "enable-site-license"
}

func (p *EnableSiteLicense) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableSiteLicense::Post")
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

func (p *EnableSiteLicense) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableSiteLicense::Get")
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
func (p *EnableSiteLicense) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableSiteLicense::Put")
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

func (p *EnableSiteLicense) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableSiteLicense::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
