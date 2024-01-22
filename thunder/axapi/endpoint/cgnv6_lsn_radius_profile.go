

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnRadiusProfile struct {
	Inst struct {

    LidProfileIndex int `json:"lid-profile-index"`
    Radius []Cgnv6LsnRadiusProfileRadius `json:"radius"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"lsn-radius-profile"`
}


type Cgnv6LsnRadiusProfileRadius struct {
    Attribute string `json:"attribute"`
    StartsWith string `json:"starts-with"`
    StartsWithLsnLid int `json:"starts-with-lsn-lid"`
    ExactValue string `json:"exact-value"`
    ExactValueLsnLid int `json:"exact-value-lsn-lid"`
    DefaultLsnLid int `json:"default-lsn-lid"`
}

func (p *Cgnv6LsnRadiusProfile) GetId() string{
    return strconv.Itoa(p.Inst.LidProfileIndex)
}

func (p *Cgnv6LsnRadiusProfile) getPath() string{
    return "cgnv6/lsn-radius-profile"
}

func (p *Cgnv6LsnRadiusProfile) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRadiusProfile::Post")
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

func (p *Cgnv6LsnRadiusProfile) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRadiusProfile::Get")
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
func (p *Cgnv6LsnRadiusProfile) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRadiusProfile::Put")
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

func (p *Cgnv6LsnRadiusProfile) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnRadiusProfile::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
