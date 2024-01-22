

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Axdebug struct {
	Inst struct {

    ApplyConfig AxdebugApplyConfig75 `json:"apply-config"`
    Capture AxdebugCapture76 `json:"capture"`
    Count1 int `json:"count1" dval:"3000"`
    Delete AxdebugDelete77 `json:"delete"`
    Exit AxdebugExit78 `json:"exit"`
    FilterConfigList []AxdebugFilterConfigList `json:"filter-config-list"`
    IncPortNum string `json:"inc-port-num"`
    Incoming int `json:"incoming"`
    Length int `json:"length" dval:"1518"`
    Maxfile int `json:"maxfile" dval:"100"`
    OutPortNum string `json:"out-port-num"`
    Outgoing int `json:"outgoing"`
    PcapngConfig AxdebugPcapngConfig79 `json:"pcapng-config"`
    SaveConfig AxdebugSaveConfig80 `json:"save-config"`
    SessFilterDis int `json:"sess-filter-dis"`
    Timeout int `json:"timeout" dval:"5"`
    Uuid string `json:"uuid"`

	} `json:"axdebug"`
}


type AxdebugApplyConfig75 struct {
    ConfigFile string `json:"config-file"`
}


type AxdebugCapture76 struct {
    Brief int `json:"brief"`
    Detail int `json:"detail"`
    Save string `json:"save"`
    CurrentSlot int `json:"current-slot"`
    NoStop int `json:"no-stop"`
}


type AxdebugDelete77 struct {
    CaptureFile string `json:"capture-file"`
    ConfigFile string `json:"config-file"`
}


type AxdebugExit78 struct {
    StopCapture int `json:"stop-capture"`
}


type AxdebugFilterConfigList struct {
    Number int `json:"number"`
    L3Proto string `json:"l3-proto"`
    Dst int `json:"dst"`
    Src int `json:"src"`
    Ip int `json:"ip"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
    Ipv6 int `json:"ipv6"`
    Ipv6Address string `json:"ipv6-address"`
    Mac int `json:"mac"`
    MacAddr string `json:"mac-addr"`
    Port int `json:"port"`
    DstIp int `json:"dst-ip"`
    DstIpv4Address string `json:"dst-ipv4-address"`
    SrcIp int `json:"src-ip"`
    SrcIpv4Address string `json:"src-ipv4-address"`
    DstMac int `json:"dst-mac"`
    DstMacAddr string `json:"dst-mac-addr"`
    SrcMac int `json:"src-mac"`
    SrcMacAddr string `json:"src-mac-addr"`
    DstPort int `json:"dst-port"`
    DstPortNum int `json:"dst-port-num"`
    SrcPort int `json:"src-port"`
    SrcPortNum int `json:"src-port-num"`
    PortNumMin int `json:"port-num-min"`
    PortNumMax int `json:"port-num-max"`
    Proto int `json:"proto"`
    ProtoVal string `json:"proto-val"`
    ProtNum int `json:"prot-num"`
    Offset int `json:"offset"`
    Length int `json:"length"`
    OperRange string `json:"oper-range"`
    Hex int `json:"hex"`
    MinHex string `json:"min-hex"`
    MaxHex string `json:"max-hex"`
    CompHex string `json:"comp-hex"`
    Integer int `json:"integer"`
    IntegerMin int `json:"integer-min"`
    IntegerMax int `json:"integer-max"`
    IntegerComp int `json:"integer-comp"`
    Word int `json:"word"`
    Word0 string `json:"WORD0"`
    Word1 string `json:"WORD1"`
    Word2 string `json:"WORD2"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type AxdebugPcapngConfig79 struct {
    PcapngEnable int `json:"pcapng-enable"`
    SslKeyEnable int `json:"ssl-key-enable"`
    Exit int `json:"exit"`
    Uuid string `json:"uuid"`
}


type AxdebugSaveConfig80 struct {
    ConfigFile string `json:"config-file"`
    Default int `json:"default"`
}

func (p *Axdebug) GetId() string{
    return "1"
}

func (p *Axdebug) getPath() string{
    return "axdebug"
}

func (p *Axdebug) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Axdebug::Post")
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

func (p *Axdebug) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Axdebug::Get")
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
func (p *Axdebug) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Axdebug::Put")
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

func (p *Axdebug) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Axdebug::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
