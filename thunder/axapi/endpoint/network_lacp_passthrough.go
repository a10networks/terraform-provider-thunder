

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type NetworkLacpPassthrough struct {
	Inst struct {

    PeerFrom int `json:"peer-from"`
    PeerTo int `json:"peer-to"`
    Uuid string `json:"uuid"`

	} `json:"lacp-passthrough"`
}

func (p *NetworkLacpPassthrough) GetId() string{
    return strconv.Itoa(p.Inst.PeerFrom)+"+"+strconv.Itoa(p.Inst.PeerTo)
}

func (p *NetworkLacpPassthrough) getPath() string{
    return "network/lacp-passthrough"
}

func (p *NetworkLacpPassthrough) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLacpPassthrough::Post")
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

func (p *NetworkLacpPassthrough) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLacpPassthrough::Get")
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
func (p *NetworkLacpPassthrough) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLacpPassthrough::Put")
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

func (p *NetworkLacpPassthrough) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLacpPassthrough::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
