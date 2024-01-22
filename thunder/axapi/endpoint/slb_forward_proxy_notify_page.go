

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbForwardProxyNotifyPage struct {
	Inst struct {

    Message string `json:"message"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"notify-page"`
}

func (p *SlbForwardProxyNotifyPage) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbForwardProxyNotifyPage) getPath() string{
    return "slb/forward-proxy/notify-page"
}

func (p *SlbForwardProxyNotifyPage) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbForwardProxyNotifyPage::Post")
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

func (p *SlbForwardProxyNotifyPage) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbForwardProxyNotifyPage::Get")
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
func (p *SlbForwardProxyNotifyPage) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbForwardProxyNotifyPage::Put")
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

func (p *SlbForwardProxyNotifyPage) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbForwardProxyNotifyPage::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
