

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugAuth struct {
	Inst struct {

    Authd int `json:"authd"`
    ClientAddr string `json:"client-addr"`
    Level string `json:"level"`
    Saml int `json:"saml"`
    SamlSp string `json:"saml-sp"`
    Username string `json:"username"`
    Uuid string `json:"uuid"`
    VirtualServer string `json:"virtual-server"`

	} `json:"auth"`
}

func (p *DebugAuth) GetId() string{
    return "1"
}

func (p *DebugAuth) getPath() string{
    return "debug/auth"
}

func (p *DebugAuth) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugAuth::Post")
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

func (p *DebugAuth) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugAuth::Get")
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
func (p *DebugAuth) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugAuth::Put")
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

func (p *DebugAuth) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugAuth::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
