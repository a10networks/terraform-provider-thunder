

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceOtherDynamicEntryOverflowPolicy struct {
	Inst struct {

    Action string `json:"action"`
    DummyName string `json:"dummy-name"`
    Glid string `json:"glid"`
    LogEnable int `json:"log-enable"`
    LogPeriodic int `json:"log-periodic"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    ZoneTemplate DdosDstZonePortZoneServiceOtherDynamicEntryOverflowPolicyZoneTemplate `json:"zone-template"`
    PortOther string 
    ZoneName string 
    Protocol string 

	} `json:"dynamic-entry-overflow-policy"`
}


type DdosDstZonePortZoneServiceOtherDynamicEntryOverflowPolicyZoneTemplate struct {
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Encap string `json:"encap"`
    Logging string `json:"logging"`
}

func (p *DdosDstZonePortZoneServiceOtherDynamicEntryOverflowPolicy) GetId() string{
    return p.Inst.DummyName
}

func (p *DdosDstZonePortZoneServiceOtherDynamicEntryOverflowPolicy) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/port/zone-service-other/" +p.Inst.PortOther + "+" +p.Inst.Protocol + "/dynamic-entry-overflow-policy"
}

func (p *DdosDstZonePortZoneServiceOtherDynamicEntryOverflowPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceOtherDynamicEntryOverflowPolicy::Post")
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

func (p *DdosDstZonePortZoneServiceOtherDynamicEntryOverflowPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceOtherDynamicEntryOverflowPolicy::Get")
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
func (p *DdosDstZonePortZoneServiceOtherDynamicEntryOverflowPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceOtherDynamicEntryOverflowPolicy::Put")
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

func (p *DdosDstZonePortZoneServiceOtherDynamicEntryOverflowPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZonePortZoneServiceOtherDynamicEntryOverflowPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
