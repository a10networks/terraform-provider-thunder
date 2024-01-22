

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HealthMonitorMethodSnmp struct {
	Inst struct {

    Community string `json:"community" dval:"public"`
    Oid HealthMonitorMethodSnmpOid `json:"oid"`
    Operation HealthMonitorMethodSnmpOperation `json:"operation"`
    Snmp int `json:"snmp"`
    SnmpPort int `json:"snmp-port" dval:"161"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"snmp"`
}


type HealthMonitorMethodSnmpOid struct {
    Mib string `json:"mib"`
    Asn string `json:"asn"`
}


type HealthMonitorMethodSnmpOperation struct {
    OperType string `json:"oper-type"`
}

func (p *HealthMonitorMethodSnmp) GetId() string{
    return "1"
}

func (p *HealthMonitorMethodSnmp) getPath() string{
    return "health/monitor/" +p.Inst.Name + "/method/snmp"
}

func (p *HealthMonitorMethodSnmp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodSnmp::Post")
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

func (p *HealthMonitorMethodSnmp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodSnmp::Get")
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
func (p *HealthMonitorMethodSnmp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodSnmp::Put")
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

func (p *HealthMonitorMethodSnmp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodSnmp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
