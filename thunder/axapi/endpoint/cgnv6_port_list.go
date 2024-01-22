

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6PortList struct {
	Inst struct {

    Name string `json:"name"`
    PortConfig []Cgnv6PortListPortConfig `json:"port-config"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"port-list"`
}


type Cgnv6PortListPortConfig struct {
    OriginalPort int `json:"original-port"`
    TranslatedPort int `json:"translated-port"`
}

func (p *Cgnv6PortList) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *Cgnv6PortList) getPath() string{
    return "cgnv6/port-list"
}

func (p *Cgnv6PortList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6PortList::Post")
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

func (p *Cgnv6PortList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6PortList::Get")
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
func (p *Cgnv6PortList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6PortList::Put")
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

func (p *Cgnv6PortList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6PortList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
