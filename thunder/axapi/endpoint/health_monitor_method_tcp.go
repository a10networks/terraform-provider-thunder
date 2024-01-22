

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HealthMonitorMethodTcp struct {
	Inst struct {

    Maintenance int `json:"maintenance"`
    MaintenanceText string `json:"maintenance-text"`
    MethodTcp int `json:"method-tcp"`
    PortHalfopen int `json:"port-halfopen"`
    PortResp HealthMonitorMethodTcpPortResp `json:"port-resp"`
    PortSend string `json:"port-send"`
    TcpPort int `json:"tcp-port"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"tcp"`
}


type HealthMonitorMethodTcpPortResp struct {
    PortContains string `json:"port-contains"`
}

func (p *HealthMonitorMethodTcp) GetId() string{
    return "1"
}

func (p *HealthMonitorMethodTcp) getPath() string{
    return "health/monitor/" +p.Inst.Name + "/method/tcp"
}

func (p *HealthMonitorMethodTcp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodTcp::Post")
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

func (p *HealthMonitorMethodTcp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodTcp::Get")
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
func (p *HealthMonitorMethodTcp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodTcp::Put")
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

func (p *HealthMonitorMethodTcp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodTcp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
