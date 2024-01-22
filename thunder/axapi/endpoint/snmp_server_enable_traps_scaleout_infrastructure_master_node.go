

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerEnableTrapsScaleoutInfrastructureMasterNode struct {
	Inst struct {

    TrafficMapDistribution int `json:"traffic-map-distribution"`
    Uuid string `json:"uuid"`
    VserverTrafficMapUpdate int `json:"vserver-traffic-map-update"`

	} `json:"master-node"`
}

func (p *SnmpServerEnableTrapsScaleoutInfrastructureMasterNode) GetId() string{
    return "1"
}

func (p *SnmpServerEnableTrapsScaleoutInfrastructureMasterNode) getPath() string{
    return "snmp-server/enable/traps/scaleout/infrastructure/master-node"
}

func (p *SnmpServerEnableTrapsScaleoutInfrastructureMasterNode) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsScaleoutInfrastructureMasterNode::Post")
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

func (p *SnmpServerEnableTrapsScaleoutInfrastructureMasterNode) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsScaleoutInfrastructureMasterNode::Get")
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
func (p *SnmpServerEnableTrapsScaleoutInfrastructureMasterNode) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsScaleoutInfrastructureMasterNode::Put")
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

func (p *SnmpServerEnableTrapsScaleoutInfrastructureMasterNode) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsScaleoutInfrastructureMasterNode::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
