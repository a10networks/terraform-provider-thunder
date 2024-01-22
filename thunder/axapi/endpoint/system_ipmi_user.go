

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemIpmiUser struct {
	Inst struct {

    Add string `json:"add"`
    Administrator int `json:"administrator"`
    Callback int `json:"callback"`
    Disable string `json:"disable"`
    Newname string `json:"newname"`
    Newpass string `json:"newpass"`
    Operator int `json:"operator"`
    Password string `json:"password"`
    Privilege string `json:"privilege"`
    Setname string `json:"setname"`
    Setpass string `json:"setpass"`
    User int `json:"user"`

	} `json:"user"`
}

func (p *SystemIpmiUser) GetId() string{
    return "1"
}

func (p *SystemIpmiUser) getPath() string{
    return "system/ipmi/user"
}

func (p *SystemIpmiUser) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpmiUser::Post")
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

func (p *SystemIpmiUser) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpmiUser::Get")
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
func (p *SystemIpmiUser) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpmiUser::Put")
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

func (p *SystemIpmiUser) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpmiUser::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
