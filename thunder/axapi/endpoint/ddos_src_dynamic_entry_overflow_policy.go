

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosSrcDynamicEntryOverflowPolicy struct {
	Inst struct {

    AppTypeList []DdosSrcDynamicEntryOverflowPolicyAppTypeList `json:"app-type-list"`
    DefaultAddressType string `json:"default-address-type"`
    ExceedLogCfg DdosSrcDynamicEntryOverflowPolicyExceedLogCfg `json:"exceed-log-cfg"`
    Glid string `json:"glid"`
    L4TypeList []DdosSrcDynamicEntryOverflowPolicyL4TypeList `json:"l4-type-list"`
    LogPeriodic int `json:"log-periodic"`
    Template DdosSrcDynamicEntryOverflowPolicyTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"dynamic-entry-overflow-policy"`
}


type DdosSrcDynamicEntryOverflowPolicyAppTypeList struct {
    Protocol string `json:"protocol"`
    Template DdosSrcDynamicEntryOverflowPolicyAppTypeListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosSrcDynamicEntryOverflowPolicyAppTypeListTemplate struct {
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
}


type DdosSrcDynamicEntryOverflowPolicyExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
    WithSflowSample int `json:"with-sflow-sample"`
}


type DdosSrcDynamicEntryOverflowPolicyL4TypeList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosSrcDynamicEntryOverflowPolicyL4TypeListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosSrcDynamicEntryOverflowPolicyL4TypeListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosSrcDynamicEntryOverflowPolicyTemplate struct {
    Logging string `json:"logging"`
}

func (p *DdosSrcDynamicEntryOverflowPolicy) GetId() string{
    return p.Inst.DefaultAddressType
}

func (p *DdosSrcDynamicEntryOverflowPolicy) getPath() string{
    return "ddos/src/dynamic-entry-overflow-policy"
}

func (p *DdosSrcDynamicEntryOverflowPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcDynamicEntryOverflowPolicy::Post")
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

func (p *DdosSrcDynamicEntryOverflowPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcDynamicEntryOverflowPolicy::Get")
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
func (p *DdosSrcDynamicEntryOverflowPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcDynamicEntryOverflowPolicy::Put")
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

func (p *DdosSrcDynamicEntryOverflowPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcDynamicEntryOverflowPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
