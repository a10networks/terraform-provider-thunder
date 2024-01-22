

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type AxdebugFilterConfig struct {
	Inst struct {

    CompHex string `json:"comp-hex"`
    Dst int `json:"dst"`
    DstIp int `json:"dst-ip"`
    DstIpv4Address string `json:"dst-ipv4-address"`
    DstMac int `json:"dst-mac"`
    DstMacAddr string `json:"dst-mac-addr"`
    DstPort int `json:"dst-port"`
    DstPortNum int `json:"dst-port-num"`
    Hex int `json:"hex"`
    Integer int `json:"integer"`
    IntegerComp int `json:"integer-comp"`
    IntegerMax int `json:"integer-max"`
    IntegerMin int `json:"integer-min"`
    Ip int `json:"ip"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`
    Ipv6 int `json:"ipv6"`
    Ipv6Address string `json:"ipv6-address"`
    L3Proto string `json:"l3-proto"`
    Length int `json:"length"`
    Mac int `json:"mac"`
    MacAddr string `json:"mac-addr"`
    MaxHex string `json:"max-hex"`
    MinHex string `json:"min-hex"`
    Number int `json:"number"`
    Offset int `json:"offset"`
    OperRange string `json:"oper-range"`
    Port int `json:"port"`
    PortNumMax int `json:"port-num-max"`
    PortNumMin int `json:"port-num-min"`
    ProtNum int `json:"prot-num"`
    Proto int `json:"proto"`
    ProtoVal string `json:"proto-val"`
    Src int `json:"src"`
    SrcIp int `json:"src-ip"`
    SrcIpv4Address string `json:"src-ipv4-address"`
    SrcMac int `json:"src-mac"`
    SrcMacAddr string `json:"src-mac-addr"`
    SrcPort int `json:"src-port"`
    SrcPortNum int `json:"src-port-num"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Word int `json:"word"`
    Word0 string `json:"WORD0"`
    Word1 string `json:"WORD1"`
    Word2 string `json:"WORD2"`

	} `json:"filter-config"`
}

func (p *AxdebugFilterConfig) GetId() string{
    return strconv.Itoa(p.Inst.Number)
}

func (p *AxdebugFilterConfig) getPath() string{
    return "axdebug/filter-config"
}

func (p *AxdebugFilterConfig) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AxdebugFilterConfig::Post")
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

func (p *AxdebugFilterConfig) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AxdebugFilterConfig::Get")
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
func (p *AxdebugFilterConfig) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AxdebugFilterConfig::Put")
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

func (p *AxdebugFilterConfig) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AxdebugFilterConfig::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
