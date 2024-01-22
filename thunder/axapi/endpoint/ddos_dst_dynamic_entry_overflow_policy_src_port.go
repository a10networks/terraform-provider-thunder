

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type DdosDstDynamicEntryOverflowPolicySrcPort struct {
	Inst struct {

    Deny int `json:"deny"`
    Glid string `json:"glid"`
    PortNum int `json:"port-num"`
    Protocol string `json:"protocol"`
    Template DdosDstDynamicEntryOverflowPolicySrcPortTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    DefaultAddressType string 

	} `json:"src-port"`
}


type DdosDstDynamicEntryOverflowPolicySrcPortTemplate struct {
    SrcUdp string `json:"src-udp"`
    SrcTcp string `json:"src-tcp"`
}

func (p *DdosDstDynamicEntryOverflowPolicySrcPort) GetId() string{
    return strconv.Itoa(p.Inst.PortNum)+"+"+p.Inst.Protocol
}

func (p *DdosDstDynamicEntryOverflowPolicySrcPort) getPath() string{
    return "ddos/dst/dynamic-entry-overflow-policy/" +p.Inst.DefaultAddressType + "/src-port"
}

func (p *DdosDstDynamicEntryOverflowPolicySrcPort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDynamicEntryOverflowPolicySrcPort::Post")
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

func (p *DdosDstDynamicEntryOverflowPolicySrcPort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDynamicEntryOverflowPolicySrcPort::Get")
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
func (p *DdosDstDynamicEntryOverflowPolicySrcPort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDynamicEntryOverflowPolicySrcPort::Put")
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

func (p *DdosDstDynamicEntryOverflowPolicySrcPort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstDynamicEntryOverflowPolicySrcPort::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
