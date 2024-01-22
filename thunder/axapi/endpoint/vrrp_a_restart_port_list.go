

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpARestartPortList struct {
	Inst struct {

    EthernetCfg []VrrpARestartPortListEthernetCfg `json:"ethernet-cfg"`
    Uuid string `json:"uuid"`
    VridList []VrrpARestartPortListVridList `json:"vrid-list"`

	} `json:"restart-port-list"`
}


type VrrpARestartPortListEthernetCfg struct {
    FlapEthernetStart int `json:"flap-ethernet-start"`
    FlapEthernetEnd int `json:"flap-ethernet-end"`
}


type VrrpARestartPortListVridList struct {
    VridVal int `json:"vrid-val"`
    EthernetCfg []VrrpARestartPortListVridListEthernetCfg `json:"ethernet-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type VrrpARestartPortListVridListEthernetCfg struct {
    FlapEthernetStart int `json:"flap-ethernet-start"`
    FlapEthernetEnd int `json:"flap-ethernet-end"`
}

func (p *VrrpARestartPortList) GetId() string{
    return "1"
}

func (p *VrrpARestartPortList) getPath() string{
    return "vrrp-a/restart-port-list"
}

func (p *VrrpARestartPortList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpARestartPortList::Post")
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

func (p *VrrpARestartPortList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpARestartPortList::Get")
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
func (p *VrrpARestartPortList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpARestartPortList::Put")
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

func (p *VrrpARestartPortList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpARestartPortList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
