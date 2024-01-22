

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetflowMonitorCustomRecord struct {
	Inst struct {

    CustomCfg []NetflowMonitorCustomRecordCustomCfg `json:"custom-cfg"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"custom-record"`
}


type NetflowMonitorCustomRecordCustomCfg struct {
    Event string `json:"event"`
    IpfixTemplate string `json:"ipfix-template"`
}

func (p *NetflowMonitorCustomRecord) GetId() string{
    return "1"
}

func (p *NetflowMonitorCustomRecord) getPath() string{
    return "netflow/monitor/" +p.Inst.Name + "/custom-record"
}

func (p *NetflowMonitorCustomRecord) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorCustomRecord::Post")
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

func (p *NetflowMonitorCustomRecord) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorCustomRecord::Get")
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
func (p *NetflowMonitorCustomRecord) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorCustomRecord::Put")
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

func (p *NetflowMonitorCustomRecord) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorCustomRecord::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
