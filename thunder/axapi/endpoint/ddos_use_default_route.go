

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosUseDefaultRoute struct {
	Inst struct {

    EthernetStartCfg []DdosUseDefaultRouteEthernetStartCfg `json:"ethernet-start-cfg"`
    Uuid string `json:"uuid"`

	} `json:"use-default-route"`
}


type DdosUseDefaultRouteEthernetStartCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}

func (p *DdosUseDefaultRoute) GetId() string{
    return "1"
}

func (p *DdosUseDefaultRoute) getPath() string{
    return "ddos/use-default-route"
}

func (p *DdosUseDefaultRoute) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosUseDefaultRoute::Post")
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

func (p *DdosUseDefaultRoute) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosUseDefaultRoute::Get")
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
func (p *DdosUseDefaultRoute) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosUseDefaultRoute::Put")
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

func (p *DdosUseDefaultRoute) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosUseDefaultRoute::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
