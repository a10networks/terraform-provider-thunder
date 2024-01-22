

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemIpmi struct {
	Inst struct {

    Ip SystemIpmiIp1588 `json:"ip"`
    Ipsrc SystemIpmiIpsrc1589 `json:"ipsrc"`
    Reset int `json:"reset"`
    Tool SystemIpmiTool1590 `json:"tool"`
    User SystemIpmiUser1591 `json:"user"`

	} `json:"ipmi"`
}


type SystemIpmiIp1588 struct {
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
    DefaultGateway string `json:"default-gateway"`
}


type SystemIpmiIpsrc1589 struct {
    Dhcp int `json:"dhcp"`
    Static int `json:"static"`
}


type SystemIpmiTool1590 struct {
    Cmd string `json:"cmd"`
}


type SystemIpmiUser1591 struct {
    Add string `json:"add"`
    Password string `json:"password"`
    Administrator int `json:"administrator"`
    Callback int `json:"callback"`
    Operator int `json:"operator"`
    User int `json:"user"`
    Disable string `json:"disable"`
    Privilege string `json:"privilege"`
    Setname string `json:"setname"`
    Newname string `json:"newname"`
    Setpass string `json:"setpass"`
    Newpass string `json:"newpass"`
}

func (p *SystemIpmi) GetId() string{
    return "1"
}

func (p *SystemIpmi) getPath() string{
    return "system/ipmi"
}

func (p *SystemIpmi) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpmi::Post")
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

func (p *SystemIpmi) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpmi::Get")
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
func (p *SystemIpmi) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpmi::Put")
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

func (p *SystemIpmi) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpmi::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
