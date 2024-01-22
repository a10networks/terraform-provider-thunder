

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceTrunkBfd struct {
	Inst struct {

    Authentication InterfaceTrunkBfdAuthentication `json:"authentication"`
    Demand int `json:"demand"`
    Echo int `json:"echo"`
    IntervalCfg InterfaceTrunkBfdIntervalCfg `json:"interval-cfg"`
    PerMemberPort InterfaceTrunkBfdPerMemberPort742 `json:"per-member-port"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"bfd"`
}


type InterfaceTrunkBfdAuthentication struct {
    KeyId int `json:"key-id"`
    Method string `json:"method"`
    Password string `json:"password"`
    Encrypted string `json:"encrypted"`
}


type InterfaceTrunkBfdIntervalCfg struct {
    Interval int `json:"interval"`
    MinRx int `json:"min-rx"`
    Multiplier int `json:"multiplier"`
}


type InterfaceTrunkBfdPerMemberPort742 struct {
    LocalAddress string `json:"local-address"`
    NeighborAddress string `json:"neighbor-address"`
    Ipv6Local string `json:"ipv6-local"`
    Ipv6Nbr string `json:"ipv6-nbr"`
    Uuid string `json:"uuid"`
}

func (p *InterfaceTrunkBfd) GetId() string{
    return "1"
}

func (p *InterfaceTrunkBfd) getPath() string{
    return "interface/trunk/" +p.Inst.Ifnum + "/bfd"
}

func (p *InterfaceTrunkBfd) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkBfd::Post")
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

func (p *InterfaceTrunkBfd) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkBfd::Get")
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
func (p *InterfaceTrunkBfd) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkBfd::Put")
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

func (p *InterfaceTrunkBfd) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTrunkBfd::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
