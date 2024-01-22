

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type AccessListStandard struct {
	Inst struct {

    Std int `json:"std"`
    Stdrules []AccessListStandardStdrules `json:"stdrules"`
    Uuid string `json:"uuid"`

	} `json:"standard"`
}


type AccessListStandardStdrules struct {
    SeqNum int `json:"seq-num"`
    StdRemark string `json:"std-remark"`
    Action string `json:"action"`
    Any int `json:"any"`
    Host string `json:"host"`
    Subnet string `json:"subnet"`
    RevSubnetMask string `json:"rev-subnet-mask"`
    Log int `json:"log"`
    TransparentSessionOnly int `json:"transparent-session-only"`
}

func (p *AccessListStandard) GetId() string{
    return strconv.Itoa(p.Inst.Std)
}

func (p *AccessListStandard) getPath() string{
    return "access-list/standard"
}

func (p *AccessListStandard) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AccessListStandard::Post")
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

func (p *AccessListStandard) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AccessListStandard::Get")
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
func (p *AccessListStandard) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AccessListStandard::Put")
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

func (p *AccessListStandard) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AccessListStandard::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
