

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Terminal struct {
	Inst struct {

    AutoSize int `json:"auto-size" dval:"1"`
    Editing int `json:"editing" dval:"1"`
    GslbCfg TerminalGslbCfg `json:"gslb-cfg"`
    HistoryCfg TerminalHistoryCfg `json:"history-cfg"`
    IdleTimeout int `json:"idle-timeout" dval:"15"`
    Length int `json:"length" dval:"24"`
    PromptCfg TerminalPromptCfg `json:"prompt-cfg"`
    Uuid string `json:"uuid"`
    Width int `json:"width" dval:"80"`

	} `json:"terminal"`
}


type TerminalGslbCfg struct {
    GslbPrompt int `json:"gslb-prompt"`
    Disable int `json:"disable"`
    GroupRole int `json:"group-role"`
    Symbol int `json:"symbol"`
}


type TerminalHistoryCfg struct {
    Enable int `json:"enable" dval:"1"`
    Size int `json:"size" dval:"256"`
}


type TerminalPromptCfg struct {
    Prompt int `json:"prompt"`
    HaStatus int `json:"ha-status"`
    Hostname int `json:"hostname"`
    VcsCfg TerminalPromptCfgVcsCfg `json:"vcs-cfg"`
}


type TerminalPromptCfgVcsCfg struct {
    VcsStatus int `json:"vcs-status"`
}

func (p *Terminal) GetId() string{
    return "1"
}

func (p *Terminal) getPath() string{
    return "terminal"
}

func (p *Terminal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Terminal::Post")
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

func (p *Terminal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Terminal::Get")
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
func (p *Terminal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Terminal::Put")
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

func (p *Terminal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Terminal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
