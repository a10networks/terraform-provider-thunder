

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type VrrpARestartPortListVrid struct {
	Inst struct {

    EthernetCfg []VrrpARestartPortListVridEthernetCfg `json:"ethernet-cfg"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VridVal int `json:"vrid-val"`

	} `json:"vrid"`
}


type VrrpARestartPortListVridEthernetCfg struct {
    FlapEthernetStart int `json:"flap-ethernet-start"`
    FlapEthernetEnd int `json:"flap-ethernet-end"`
}

func (p *VrrpARestartPortListVrid) GetId() string{
    return strconv.Itoa(p.Inst.VridVal)
}

func (p *VrrpARestartPortListVrid) getPath() string{
    return "vrrp-a/restart-port-list/vrid"
}

func (p *VrrpARestartPortListVrid) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpARestartPortListVrid::Post")
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

func (p *VrrpARestartPortListVrid) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpARestartPortListVrid::Get")
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
func (p *VrrpARestartPortListVrid) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpARestartPortListVrid::Put")
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

func (p *VrrpARestartPortListVrid) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VrrpARestartPortListVrid::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
