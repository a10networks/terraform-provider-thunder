

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceLifIpOspfOspfGlobal struct {
	Inst struct {

    AuthenticationCfg InterfaceLifIpOspfOspfGlobalAuthenticationCfg `json:"authentication-cfg"`
    AuthenticationKey string `json:"authentication-key"`
    BfdCfg InterfaceLifIpOspfOspfGlobalBfdCfg `json:"bfd-cfg"`
    Cost int `json:"cost"`
    DatabaseFilterCfg InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg `json:"database-filter-cfg"`
    DeadInterval int `json:"dead-interval" dval:"40"`
    Disable string `json:"disable"`
    HelloInterval int `json:"hello-interval" dval:"10"`
    MessageDigestCfg []InterfaceLifIpOspfOspfGlobalMessageDigestCfg `json:"message-digest-cfg"`
    Mtu int `json:"mtu"`
    MtuIgnore int `json:"mtu-ignore"`
    Network InterfaceLifIpOspfOspfGlobalNetwork `json:"network"`
    Priority int `json:"priority" dval:"1"`
    RetransmitInterval int `json:"retransmit-interval" dval:"5"`
    TransmitDelay int `json:"transmit-delay" dval:"1"`
    Uuid string `json:"uuid"`
    Ifname string 

	} `json:"ospf-global"`
}


type InterfaceLifIpOspfOspfGlobalAuthenticationCfg struct {
    Authentication int `json:"authentication"`
    Value string `json:"value"`
}


type InterfaceLifIpOspfOspfGlobalBfdCfg struct {
    Bfd int `json:"bfd"`
    Disable int `json:"disable"`
}


type InterfaceLifIpOspfOspfGlobalDatabaseFilterCfg struct {
    DatabaseFilter string `json:"database-filter"`
    Out int `json:"out"`
}


type InterfaceLifIpOspfOspfGlobalMessageDigestCfg struct {
    MessageDigestKey int `json:"message-digest-key"`
    Md5 InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5 `json:"md5"`
}


type InterfaceLifIpOspfOspfGlobalMessageDigestCfgMd5 struct {
    Md5Value string `json:"md5-value"`
    Encrypted string `json:"encrypted"`
}


type InterfaceLifIpOspfOspfGlobalNetwork struct {
    Broadcast int `json:"broadcast"`
    NonBroadcast int `json:"non-broadcast"`
    PointToPoint int `json:"point-to-point"`
    PointToMultipoint int `json:"point-to-multipoint"`
    P2mpNbma int `json:"p2mp-nbma"`
}

func (p *InterfaceLifIpOspfOspfGlobal) GetId() string{
    return "1"
}

func (p *InterfaceLifIpOspfOspfGlobal) getPath() string{
    return "interface/lif/" +p.Inst.Ifname + "/ip/ospf/ospf-global"
}

func (p *InterfaceLifIpOspfOspfGlobal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpOspfOspfGlobal::Post")
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

func (p *InterfaceLifIpOspfOspfGlobal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpOspfOspfGlobal::Get")
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
func (p *InterfaceLifIpOspfOspfGlobal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpOspfOspfGlobal::Put")
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

func (p *InterfaceLifIpOspfOspfGlobal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceLifIpOspfOspfGlobal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
