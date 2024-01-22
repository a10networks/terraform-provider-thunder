

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type IpAsPath struct {
	Inst struct {

    AccessList string `json:"access-list"`
    Action string `json:"action"`
    Uuid string `json:"uuid"`
    Value string `json:"value"`

	} `json:"as-path"`
}

func (p *IpAsPath) GetId() string{
    return p.Inst.AccessList+"+"+p.Inst.Action+"+"+url.QueryEscape(p.Inst.Value)
}

func (p *IpAsPath) getPath() string{
    return "ip/as-path"
}

func (p *IpAsPath) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAsPath::Post")
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

func (p *IpAsPath) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAsPath::Get")
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
func (p *IpAsPath) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAsPath::Put")
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

func (p *IpAsPath) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAsPath::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
