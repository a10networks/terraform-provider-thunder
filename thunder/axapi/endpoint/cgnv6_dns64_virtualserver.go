

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Dns64Virtualserver struct {
	Inst struct {

    EnableDisableAction string `json:"enable-disable-action" dval:"enable"`
    Ethernet int `json:"ethernet"`
    IpAddress string `json:"ip-address"`
    Ipv6Address string `json:"ipv6-address"`
    Name string `json:"name"`
    Netmask string `json:"netmask"`
    Policy int `json:"policy"`
    PortList []Cgnv6Dns64VirtualserverPortList `json:"port-list"`
    TemplatePolicy string `json:"template-policy"`
    UseIfIp int `json:"use-if-ip"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Vrid int `json:"vrid"`

	} `json:"dns64-virtualserver"`
}


type Cgnv6Dns64VirtualserverPortList struct {
    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Action string `json:"action" dval:"enable"`
    Pool string `json:"pool"`
    Auto int `json:"auto"`
    Precedence int `json:"precedence"`
    ServiceGroup string `json:"service-group"`
    TemplateDns string `json:"template-dns"`
    TemplatePolicy string `json:"template-policy"`
    AclIdList []Cgnv6Dns64VirtualserverPortListAclIdList `json:"acl-id-list"`
    AclNameList []Cgnv6Dns64VirtualserverPortListAclNameList `json:"acl-name-list"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []Cgnv6Dns64VirtualserverPortListSamplingEnable `json:"sampling-enable"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
}


type Cgnv6Dns64VirtualserverPortListAclIdList struct {
    AclId int `json:"acl-id"`
    AclIdSrcNatPool string `json:"acl-id-src-nat-pool"`
    AclIdSeqNum int `json:"acl-id-seq-num"`
}


type Cgnv6Dns64VirtualserverPortListAclNameList struct {
    AclName string `json:"acl-name"`
    AclNameSrcNatPool string `json:"acl-name-src-nat-pool"`
    AclNameSeqNum int `json:"acl-name-seq-num"`
}


type Cgnv6Dns64VirtualserverPortListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6Dns64Virtualserver) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *Cgnv6Dns64Virtualserver) getPath() string{
    return "cgnv6/dns64-virtualserver"
}

func (p *Cgnv6Dns64Virtualserver) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Dns64Virtualserver::Post")
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

func (p *Cgnv6Dns64Virtualserver) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Dns64Virtualserver::Get")
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
func (p *Cgnv6Dns64Virtualserver) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Dns64Virtualserver::Put")
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

func (p *Cgnv6Dns64Virtualserver) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Dns64Virtualserver::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
