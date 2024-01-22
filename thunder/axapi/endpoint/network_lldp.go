

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkLldp struct {
	Inst struct {

    EnableCfg NetworkLldpEnableCfg `json:"enable-cfg"`
    ManagementAddress NetworkLldpManagementAddress1073 `json:"management-address"`
    NotificationCfg NetworkLldpNotificationCfg `json:"notification-cfg"`
    SystemDescription string `json:"system-description"`
    SystemName string `json:"system-name"`
    TxSet NetworkLldpTxSet `json:"tx-set"`
    Uuid string `json:"uuid"`

	} `json:"lldp"`
}


type NetworkLldpEnableCfg struct {
    Enable int `json:"enable"`
    Rx int `json:"rx"`
    Tx int `json:"tx"`
}


type NetworkLldpManagementAddress1073 struct {
    DnsList []NetworkLldpManagementAddressDnsList `json:"dns-list"`
    Ipv4AddrList []NetworkLldpManagementAddressIpv4AddrList `json:"ipv4-addr-list"`
    Ipv6AddrList []NetworkLldpManagementAddressIpv6AddrList `json:"ipv6-addr-list"`
}


type NetworkLldpManagementAddressDnsList struct {
    Dns string `json:"dns"`
    Interface NetworkLldpManagementAddressDnsListInterface `json:"interface"`
    Uuid string `json:"uuid"`
}


type NetworkLldpManagementAddressDnsListInterface struct {
    Ethernet int `json:"ethernet"`
    Ve int `json:"ve"`
    Management int `json:"management"`
}


type NetworkLldpManagementAddressIpv4AddrList struct {
    Ipv4 string `json:"ipv4"`
    InterfaceIpv4 NetworkLldpManagementAddressIpv4AddrListInterfaceIpv4 `json:"interface-ipv4"`
    Uuid string `json:"uuid"`
}


type NetworkLldpManagementAddressIpv4AddrListInterfaceIpv4 struct {
    Ipv4Eth int `json:"ipv4-eth"`
    Ipv4Ve int `json:"ipv4-ve"`
    Ipv4Mgmt int `json:"ipv4-mgmt"`
}


type NetworkLldpManagementAddressIpv6AddrList struct {
    Ipv6 string `json:"ipv6"`
    InterfaceIpv6 NetworkLldpManagementAddressIpv6AddrListInterfaceIpv6 `json:"interface-ipv6"`
    Uuid string `json:"uuid"`
}


type NetworkLldpManagementAddressIpv6AddrListInterfaceIpv6 struct {
    Ipv6Eth int `json:"ipv6-eth"`
    Ipv6Ve int `json:"ipv6-ve"`
    Ipv6Mgmt int `json:"ipv6-mgmt"`
}


type NetworkLldpNotificationCfg struct {
    Notification int `json:"notification"`
    Interval int `json:"interval" dval:"30"`
}


type NetworkLldpTxSet struct {
    FastCount int `json:"fast-count" dval:"4"`
    FastInterval int `json:"fast-interval" dval:"1"`
    Hold int `json:"hold" dval:"4"`
    TxInterval int `json:"tx-interval" dval:"30"`
    ReinitDelay int `json:"reinit-delay" dval:"2"`
}

func (p *NetworkLldp) GetId() string{
    return "1"
}

func (p *NetworkLldp) getPath() string{
    return "network/lldp"
}

func (p *NetworkLldp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLldp::Post")
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

func (p *NetworkLldp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLldp::Get")
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
func (p *NetworkLldp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLldp::Put")
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

func (p *NetworkLldp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetworkLldp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
