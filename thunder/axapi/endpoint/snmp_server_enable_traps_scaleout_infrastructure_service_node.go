

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerEnableTrapsScaleoutInfrastructureServiceNode struct {
	Inst struct {

    LocalDeviceDisabled int `json:"local-device-disabled"`
    ServiceMaster int `json:"service-master"`
    TrafficMapUpdate int `json:"traffic-map-update"`
    Uuid string `json:"uuid"`

	} `json:"service-node"`
}

func (p *SnmpServerEnableTrapsScaleoutInfrastructureServiceNode) GetId() string{
    return "1"
}

func (p *SnmpServerEnableTrapsScaleoutInfrastructureServiceNode) getPath() string{
    return "snmp-server/enable/traps/scaleout/infrastructure/service-node"
}

func (p *SnmpServerEnableTrapsScaleoutInfrastructureServiceNode) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsScaleoutInfrastructureServiceNode::Post")
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

func (p *SnmpServerEnableTrapsScaleoutInfrastructureServiceNode) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsScaleoutInfrastructureServiceNode::Get")
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
func (p *SnmpServerEnableTrapsScaleoutInfrastructureServiceNode) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsScaleoutInfrastructureServiceNode::Put")
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

func (p *SnmpServerEnableTrapsScaleoutInfrastructureServiceNode) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsScaleoutInfrastructureServiceNode::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
