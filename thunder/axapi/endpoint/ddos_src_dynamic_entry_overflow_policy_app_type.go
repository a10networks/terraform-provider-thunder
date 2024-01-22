

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosSrcDynamicEntryOverflowPolicyAppType struct {
	Inst struct {

    Protocol string `json:"protocol"`
    Template DdosSrcDynamicEntryOverflowPolicyAppTypeTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    DefaultAddressType string 

	} `json:"app-type"`
}


type DdosSrcDynamicEntryOverflowPolicyAppTypeTemplate struct {
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
}

func (p *DdosSrcDynamicEntryOverflowPolicyAppType) GetId() string{
    return p.Inst.Protocol
}

func (p *DdosSrcDynamicEntryOverflowPolicyAppType) getPath() string{
    return "ddos/src/dynamic-entry-overflow-policy/" +p.Inst.DefaultAddressType + "/app-type"
}

func (p *DdosSrcDynamicEntryOverflowPolicyAppType) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcDynamicEntryOverflowPolicyAppType::Post")
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

func (p *DdosSrcDynamicEntryOverflowPolicyAppType) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcDynamicEntryOverflowPolicyAppType::Get")
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
func (p *DdosSrcDynamicEntryOverflowPolicyAppType) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcDynamicEntryOverflowPolicyAppType::Put")
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

func (p *DdosSrcDynamicEntryOverflowPolicyAppType) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcDynamicEntryOverflowPolicyAppType::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
