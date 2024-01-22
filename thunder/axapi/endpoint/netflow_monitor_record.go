

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetflowMonitorRecord struct {
	Inst struct {

    DdosGeneralStat int `json:"ddos-general-stat"`
    DdosHttpStat int `json:"ddos-http-stat"`
    Dslite int `json:"dslite"`
    Nat44 int `json:"nat44"`
    Nat64 int `json:"nat64"`
    NetflowV5 int `json:"netflow-v5"`
    NetflowV5Ext int `json:"netflow-v5-ext"`
    PortBatchDslite string `json:"port-batch-dslite"`
    PortBatchNat44 string `json:"port-batch-nat44"`
    PortBatchNat64 string `json:"port-batch-nat64"`
    PortBatchV2Dslite string `json:"port-batch-v2-dslite"`
    PortBatchV2Nat44 string `json:"port-batch-v2-nat44"`
    PortBatchV2Nat64 string `json:"port-batch-v2-nat64"`
    PortMappingDslite string `json:"port-mapping-dslite"`
    PortMappingNat44 string `json:"port-mapping-nat44"`
    PortMappingNat64 string `json:"port-mapping-nat64"`
    SesnEventDslite string `json:"sesn-event-dslite"`
    SesnEventFw4 string `json:"sesn-event-fw4"`
    SesnEventFw6 string `json:"sesn-event-fw6"`
    SesnEventNat44 string `json:"sesn-event-nat44"`
    SesnEventNat64 string `json:"sesn-event-nat64"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"record"`
}

func (p *NetflowMonitorRecord) GetId() string{
    return "1"
}

func (p *NetflowMonitorRecord) getPath() string{
    return "netflow/monitor/" +p.Inst.Name + "/record"
}

func (p *NetflowMonitorRecord) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorRecord::Post")
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

func (p *NetflowMonitorRecord) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorRecord::Get")
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
func (p *NetflowMonitorRecord) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorRecord::Put")
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

func (p *NetflowMonitorRecord) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowMonitorRecord::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
