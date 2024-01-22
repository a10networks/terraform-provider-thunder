

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosSrcDefault struct {
	Inst struct {

    Age int `json:"age" dval:"5"`
    AppTypeList []DdosSrcDefaultAppTypeList `json:"app-type-list"`
    ApplyPolicyOnOverflow int `json:"apply-policy-on-overflow"`
    DefaultAddressType string `json:"default-address-type"`
    Disable int `json:"disable"`
    ExceedLogCfg DdosSrcDefaultExceedLogCfg `json:"exceed-log-cfg"`
    Glid string `json:"glid"`
    L4TypeList []DdosSrcDefaultL4TypeList `json:"l4-type-list"`
    LogPeriodic int `json:"log-periodic"`
    MaxDynamicEntryCount int `json:"max-dynamic-entry-count"`
    Template DdosSrcDefaultTemplate `json:"template"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"default"`
}


type DdosSrcDefaultAppTypeList struct {
    Protocol string `json:"protocol"`
    Template DdosSrcDefaultAppTypeListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosSrcDefaultAppTypeListTemplate struct {
    Dns string `json:"dns"`
    Http string `json:"http"`
    SslL4 string `json:"ssl-l4"`
    Sip string `json:"sip"`
}


type DdosSrcDefaultExceedLogCfg struct {
    LogEnable int `json:"log-enable"`
}


type DdosSrcDefaultL4TypeList struct {
    Protocol string `json:"protocol"`
    Deny int `json:"deny"`
    Glid string `json:"glid"`
    Template DdosSrcDefaultL4TypeListTemplate `json:"template"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type DdosSrcDefaultL4TypeListTemplate struct {
    Tcp string `json:"tcp"`
    Udp string `json:"udp"`
    Other string `json:"other"`
    TemplateIcmpV4 string `json:"template-icmp-v4"`
    TemplateIcmpV6 string `json:"template-icmp-v6"`
}


type DdosSrcDefaultTemplate struct {
    Logging string `json:"logging"`
}

func (p *DdosSrcDefault) GetId() string{
    return p.Inst.DefaultAddressType
}

func (p *DdosSrcDefault) getPath() string{
    return "ddos/src/default"
}

func (p *DdosSrcDefault) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcDefault::Post")
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

func (p *DdosSrcDefault) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcDefault::Get")
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
func (p *DdosSrcDefault) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcDefault::Put")
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

func (p *DdosSrcDefault) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcDefault::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
