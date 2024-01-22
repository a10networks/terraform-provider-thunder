

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceManagement struct {
	Inst struct {

    AccessList InterfaceManagementAccessList `json:"access-list"`
    Action string `json:"action" dval:"enable"`
    BroadcastRateLimit InterfaceManagementBroadcastRateLimit `json:"broadcast-rate-limit"`
    Duplexity string `json:"duplexity" dval:"auto"`
    FlowControl int `json:"flow-control"`
    Ip InterfaceManagementIp `json:"ip"`
    Ipv6 []InterfaceManagementIpv6 `json:"ipv6"`
    Lldp InterfaceManagementLldp737 `json:"lldp"`
    Mtu int `json:"mtu"`
    SamplingEnable []InterfaceManagementSamplingEnable `json:"sampling-enable"`
    SecondaryIp InterfaceManagementSecondaryIp `json:"secondary-ip"`
    Speed string `json:"speed" dval:"auto"`
    Uuid string `json:"uuid"`

	} `json:"management"`
}


type InterfaceManagementAccessList struct {
    AclId int `json:"acl-id"`
    AclName string `json:"acl-name"`
}


type InterfaceManagementBroadcastRateLimit struct {
    BcastRateLimitEnable int `json:"bcast-rate-limit-enable"`
    Rate int `json:"rate" dval:"500"`
}


type InterfaceManagementIp struct {
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
    Dhcp int `json:"dhcp"`
    ControlAppsUseMgmtPort int `json:"control-apps-use-mgmt-port"`
    DefaultGateway string `json:"default-gateway"`
}


type InterfaceManagementIpv6 struct {
    Ipv6Addr string `json:"ipv6-addr"`
    AddressType string `json:"address-type"`
    V6AclName string `json:"v6-acl-name"`
    Inbound int `json:"inbound"`
    DefaultIpv6Gateway string `json:"default-ipv6-gateway"`
}


type InterfaceManagementLldp737 struct {
    EnableCfg InterfaceManagementLldpEnableCfg738 `json:"enable-cfg"`
    NotificationCfg InterfaceManagementLldpNotificationCfg739 `json:"notification-cfg"`
    TxDot1Cfg InterfaceManagementLldpTxDot1Cfg740 `json:"tx-dot1-cfg"`
    TxTlvsCfg InterfaceManagementLldpTxTlvsCfg741 `json:"tx-tlvs-cfg"`
    Uuid string `json:"uuid"`
}


type InterfaceManagementLldpEnableCfg738 struct {
    RtEnable int `json:"rt-enable"`
    Rx int `json:"rx"`
    Tx int `json:"tx"`
}


type InterfaceManagementLldpNotificationCfg739 struct {
    Notification int `json:"notification"`
    NotifEnable int `json:"notif-enable"`
}


type InterfaceManagementLldpTxDot1Cfg740 struct {
    TxDot1Tlvs int `json:"tx-dot1-tlvs"`
    LinkAggregation int `json:"link-aggregation"`
    Vlan int `json:"vlan"`
}


type InterfaceManagementLldpTxTlvsCfg741 struct {
    TxTlvs int `json:"tx-tlvs"`
    Exclude int `json:"exclude"`
    ManagementAddress int `json:"management-address"`
    PortDescription int `json:"port-description"`
    SystemCapabilities int `json:"system-capabilities"`
    SystemDescription int `json:"system-description"`
    SystemName int `json:"system-name"`
}


type InterfaceManagementSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type InterfaceManagementSecondaryIp struct {
    SecondaryIp int `json:"secondary-ip"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
    Dhcp int `json:"dhcp"`
    ControlAppsUseMgmtPort int `json:"control-apps-use-mgmt-port"`
    DefaultGateway string `json:"default-gateway"`
}

func (p *InterfaceManagement) GetId() string{
    return "1"
}

func (p *InterfaceManagement) getPath() string{
    return "interface/management"
}

func (p *InterfaceManagement) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceManagement::Post")
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

func (p *InterfaceManagement) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceManagement::Get")
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
func (p *InterfaceManagement) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceManagement::Put")
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

func (p *InterfaceManagement) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceManagement::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
