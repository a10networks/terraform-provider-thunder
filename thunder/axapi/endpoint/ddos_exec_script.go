

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosExecScript struct {
	Inst struct {

    AlertType int `json:"alert-type"`
    ExecScriptIpPortocol string `json:"exec-script-ip-portocol"`
    ExecScriptPortOtherProtocol string `json:"exec-script-port-other-protocol"`
    Level int `json:"level"`
    Mock int `json:"mock"`
    PortNum int `json:"port-num"`
    PortOther string `json:"port-other"`
    Protocol string `json:"protocol"`
    ProtocolNum int `json:"protocol-num"`
    Script string `json:"script"`
    SrcIp []DdosExecScriptSrcIp `json:"src-ip"`
    SrcIpv6 []DdosExecScriptSrcIpv6 `json:"src-ipv6"`
    Threshold int `json:"threshold"`
    Timeout int `json:"timeout"`
    Zone string `json:"zone"`

	} `json:"exec-script"`
}


type DdosExecScriptSrcIp struct {
    IpAddr string `json:"ip-addr"`
    SubnetIpAddr string `json:"subnet-ip-addr"`
}


type DdosExecScriptSrcIpv6 struct {
    Ip6Addr string `json:"ip6-addr"`
    SubnetIpv6Addr string `json:"subnet-ipv6-addr"`
}

func (p *DdosExecScript) GetId() string{
    return "1"
}

func (p *DdosExecScript) getPath() string{
    return "ddos/exec-script"
}

func (p *DdosExecScript) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosExecScript::Post")
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

func (p *DdosExecScript) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosExecScript::Get")
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
func (p *DdosExecScript) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosExecScript::Put")
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

func (p *DdosExecScript) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosExecScript::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
