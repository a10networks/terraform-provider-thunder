

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceTunnelIpOspfOspfGlobal struct {
	Inst struct {

    AuthenticationCfg InterfaceTunnelIpOspfOspfGlobalAuthenticationCfg `json:"authentication-cfg"`
    AuthenticationKey string `json:"authentication-key"`
    BfdCfg InterfaceTunnelIpOspfOspfGlobalBfdCfg `json:"bfd-cfg"`
    Cost int `json:"cost"`
    DatabaseFilterCfg InterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg `json:"database-filter-cfg"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    Disable string `json:"disable"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg `json:"message-digest-cfg"`
    Mtu int `json:"mtu"`
    MtuIgnore int `json:"mtu-ignore"`
    Network InterfaceTunnelIpOspfOspfGlobalNetwork `json:"network"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"ospf-global"`
}


type InterfaceTunnelIpOspfOspfGlobalAuthenticationCfg struct {
    Authentication int `json:"authentication"`
    Value string `json:"value"`
}


type InterfaceTunnelIpOspfOspfGlobalBfdCfg struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceTunnelIpOspfOspfGlobalDatabaseFilterCfg struct {
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
}


type InterfaceTunnelIpOspfOspfGlobalMessageDigestCfg struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5 InterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5 `json:"md5"`
}


type InterfaceTunnelIpOspfOspfGlobalMessageDigestCfgMd5 struct {
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceTunnelIpOspfOspfGlobalNetwork struct {
    Broadcast int `json:"broadcast"`
    NonBroadcast int `json:"non-broadcast"`
    PointToPoint int `json:"point-to-point"`
    PointToMultipoint int `json:"point-to-multipoint"`
    P2mpNbma int `json:"p2mp-nbma"`
}

func (p *InterfaceTunnelIpOspfOspfGlobal) GetId() string{
    return "1"
}

func (p *InterfaceTunnelIpOspfOspfGlobal) getPath() string{
    return "interface/tunnel/" +p.Inst.Ifnum + "/ip/ospf/ospf-global"
}

func (p *InterfaceTunnelIpOspfOspfGlobal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnelIpOspfOspfGlobal::Post")
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

func (p *InterfaceTunnelIpOspfOspfGlobal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnelIpOspfOspfGlobal::Get")
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
func (p *InterfaceTunnelIpOspfOspfGlobal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnelIpOspfOspfGlobal::Put")
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

func (p *InterfaceTunnelIpOspfOspfGlobal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceTunnelIpOspfOspfGlobal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
