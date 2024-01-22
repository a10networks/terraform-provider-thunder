

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Dns64VirtualserverPort struct {
	Inst struct {

    AclIdList []Cgnv6Dns64VirtualserverPortAclIdList `json:"acl-id-list"`
    AclNameList []Cgnv6Dns64VirtualserverPortAclNameList `json:"acl-name-list"`
    Action string `json:"action" dval:"enable"`
    Auto int `json:"auto"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    Pool string `json:"pool"`
    PortNumber int `json:"port-number"`
    Precedence int `json:"precedence"`
    Protocol string `json:"protocol"`
    SamplingEnable []Cgnv6Dns64VirtualserverPortSamplingEnable `json:"sampling-enable"`
    ServiceGroup string `json:"service-group"`
    TemplateDns string `json:"template-dns"`
    TemplatePolicy string `json:"template-policy"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"port"`
}


type Cgnv6Dns64VirtualserverPortAclIdList struct {
    AclId int `json:"acl-id"`
    AclIdSrcNatPool string `json:"acl-id-src-nat-pool"`
    AclIdSeqNum int `json:"acl-id-seq-num"`
}


type Cgnv6Dns64VirtualserverPortAclNameList struct {
    AclName string `json:"acl-name"`
    AclNameSrcNatPool string `json:"acl-name-src-nat-pool"`
    AclNameSeqNum int `json:"acl-name-seq-num"`
}


type Cgnv6Dns64VirtualserverPortSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6Dns64VirtualserverPort) GetId() string{
    return strconv.Itoa(p.Inst.PortNumber)+"+"+p.Inst.Protocol
}

func (p *Cgnv6Dns64VirtualserverPort) getPath() string{
    return "cgnv6/dns64-virtualserver/" +p.Inst.Name + "/port"
}

func (p *Cgnv6Dns64VirtualserverPort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Dns64VirtualserverPort::Post")
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

func (p *Cgnv6Dns64VirtualserverPort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Dns64VirtualserverPort::Get")
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
func (p *Cgnv6Dns64VirtualserverPort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Dns64VirtualserverPort::Put")
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

func (p *Cgnv6Dns64VirtualserverPort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Dns64VirtualserverPort::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
