

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetflowMonitorDestination struct {
	Inst struct {

    IpCfg NetflowMonitorDestinationIpCfg `json:"ip-cfg"`
    Ipv6Cfg NetflowMonitorDestinationIpv6Cfg `json:"ipv6-cfg"`
    ServiceGroup string `json:"service-group"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"destination"`
}


type NetflowMonitorDestinationIpCfg struct {
    Ip string `json:"ip"`
    Port4 int `json:"port4" dval:"9996"`
}


type NetflowMonitorDestinationIpv6Cfg struct {
    Ipv6 string `json:"ipv6"`
    Port6 int `json:"port6" dval:"9996"`
}

func (p *NetflowMonitorDestination) GetId() string{
    return "1"
}

func (p *NetflowMonitorDestination) getPath() string{
    return "netflow/monitor/" +p.Inst.Name + "/destination"
}

func (p *NetflowMonitorDestination) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorDestination::Post")
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

func (p *NetflowMonitorDestination) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorDestination::Get")
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
func (p *NetflowMonitorDestination) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorDestination::Put")
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

func (p *NetflowMonitorDestination) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorDestination::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
