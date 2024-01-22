

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Authentication struct {
	Inst struct {

    Console AuthenticationConsole67 `json:"console"`
    EnableCfg AuthenticationEnableCfg `json:"enable-cfg"`
    LoginCfg AuthenticationLoginCfg `json:"login-cfg"`
    ModeCfg AuthenticationModeCfg `json:"mode-cfg"`
    MultipleAuthReject int `json:"multiple-auth-reject"`
    TypeCfg AuthenticationTypeCfg `json:"type-cfg"`
    Uuid string `json:"uuid"`

	} `json:"authentication"`
}


type AuthenticationConsole67 struct {
    TypeCfg AuthenticationConsoleTypeCfg68 `json:"type-cfg"`
    Uuid string `json:"uuid"`
}


type AuthenticationConsoleTypeCfg68 struct {
    Type int `json:"type"`
    ConsoleType string `json:"console-type"`
}


type AuthenticationEnableCfg struct {
    EnableAuthType string `json:"enable-auth-type" dval:"local"`
}


type AuthenticationLoginCfg struct {
    PrivilegeMode int `json:"privilege-mode"`
    Local int `json:"local"`
}


type AuthenticationModeCfg struct {
    Mode string `json:"mode" dval:"single"`
}


type AuthenticationTypeCfg struct {
    Type string `json:"type" dval:"local"`
}

func (p *Authentication) GetId() string{
    return "1"
}

func (p *Authentication) getPath() string{
    return "authentication"
}

func (p *Authentication) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Authentication::Post")
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

func (p *Authentication) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Authentication::Get")
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
func (p *Authentication) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Authentication::Put")
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

func (p *Authentication) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Authentication::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
