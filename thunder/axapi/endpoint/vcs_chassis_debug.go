

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VcsChassisDebug struct {
	Inst struct {

    Daemon int `json:"daemon"`
    Daemon_msg int `json:"daemon_msg"`
    Election int `json:"election"`
    Election_pdu int `json:"election_pdu"`
    Encoder int `json:"encoder"`
    Handler int `json:"handler"`
    Info int `json:"info"`
    Lib int `json:"lib"`
    Net int `json:"net"`
    Ssl int `json:"ssl"`
    Util int `json:"util"`
    Vblade int `json:"vblade"`
    Vblade_msg int `json:"vblade_msg"`
    Vmaster int `json:"vmaster"`
    Vmaster_msg int `json:"vmaster_msg"`

	} `json:"debug"`
}

func (p *VcsChassisDebug) GetId() string{
    return "1"
}

func (p *VcsChassisDebug) getPath() string{
    return "vcs-chassis/debug"
}

func (p *VcsChassisDebug) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VcsChassisDebug::Post")
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

func (p *VcsChassisDebug) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VcsChassisDebug::Get")
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
func (p *VcsChassisDebug) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VcsChassisDebug::Put")
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

func (p *VcsChassisDebug) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VcsChassisDebug::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
