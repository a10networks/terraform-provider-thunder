

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DnssecDs struct {
	Inst struct {

    DsDelete int `json:"ds-delete"`
    ZoneName string `json:"zone-name"`

	} `json:"ds"`
}

func (p *DnssecDs) GetId() string{
    return "1"
}

func (p *DnssecDs) getPath() string{
    return "dnssec/ds"
}

func (p *DnssecDs) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DnssecDs::Post")
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

func (p *DnssecDs) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DnssecDs::Get")
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
func (p *DnssecDs) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DnssecDs::Put")
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

func (p *DnssecDs) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DnssecDs::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
