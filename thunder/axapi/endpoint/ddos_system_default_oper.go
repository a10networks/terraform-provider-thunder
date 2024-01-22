

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosSystemDefaultOper struct {
    
    Oper DdosSystemDefaultOperOper `json:"oper"`

}
type DataDdosSystemDefaultOper struct {
    DtDdosSystemDefaultOper DdosSystemDefaultOper `json:"system-default"`
}


type DdosSystemDefaultOperOper struct {
    DstBitEntry string `json:"dst-bit-entry"`
    DstBitUdp string `json:"dst-bit-udp"`
    DstBitTcp string `json:"dst-bit-tcp"`
    DstBitIcmp string `json:"dst-bit-icmp"`
    DstBitOther string `json:"dst-bit-other"`
    DstPktEntry string `json:"dst-pkt-entry"`
    DstPktUdp string `json:"dst-pkt-udp"`
    DstPktTcp string `json:"dst-pkt-tcp"`
    DstPktIcmp string `json:"dst-pkt-icmp"`
    DstPktOther string `json:"dst-pkt-other"`
    DstConnEntry string `json:"dst-conn-entry"`
    DstConnUdp string `json:"dst-conn-udp"`
    DstConnTcp string `json:"dst-conn-tcp"`
    DstConnIcmp string `json:"dst-conn-icmp"`
    DstConnOther string `json:"dst-conn-other"`
    DstConnREntry string `json:"dst-conn-r-entry"`
    DstConnRUdp string `json:"dst-conn-r-udp"`
    DstConnRTcp string `json:"dst-conn-r-tcp"`
    DstConnRIcmp string `json:"dst-conn-r-icmp"`
    DstConnROther string `json:"dst-conn-r-other"`
    DstFragEntry string `json:"dst-frag-entry"`
    DstFragUdp string `json:"dst-frag-udp"`
    DstFragTcp string `json:"dst-frag-tcp"`
    DstFragIcmp string `json:"dst-frag-icmp"`
    DstFragOther string `json:"dst-frag-other"`
    SrcBitEntry string `json:"src-bit-entry"`
    SrcBitUdp string `json:"src-bit-udp"`
    SrcBitTcp string `json:"src-bit-tcp"`
    SrcBitIcmp string `json:"src-bit-icmp"`
    SrcBitOther string `json:"src-bit-other"`
    SrcPktEntry string `json:"src-pkt-entry"`
    SrcPktUdp string `json:"src-pkt-udp"`
    SrcPktTcp string `json:"src-pkt-tcp"`
    SrcPktIcmp string `json:"src-pkt-icmp"`
    SrcPktOther string `json:"src-pkt-other"`
    SrcConnEntry string `json:"src-conn-entry"`
    SrcConnUdp string `json:"src-conn-udp"`
    SrcConnTcp string `json:"src-conn-tcp"`
    SrcConnIcmp string `json:"src-conn-icmp"`
    SrcConnOther string `json:"src-conn-other"`
    SrcConnREntry string `json:"src-conn-r-entry"`
    SrcConnRUdp string `json:"src-conn-r-udp"`
    SrcConnRTcp string `json:"src-conn-r-tcp"`
    SrcConnRIcmp string `json:"src-conn-r-icmp"`
    SrcConnROther string `json:"src-conn-r-other"`
    SrcFragEntry string `json:"src-frag-entry"`
    SrcFragTcp string `json:"src-frag-tcp"`
    SrcFragUdp string `json:"src-frag-udp"`
    SrcFragIcmp string `json:"src-frag-icmp"`
    SrcFragOther string `json:"src-frag-other"`
}

func (p *DdosSystemDefaultOper) GetId() string{
    return "1"
}

func (p *DdosSystemDefaultOper) getPath() string{
    return "ddos/system-default/oper"
}

func (p *DdosSystemDefaultOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosSystemDefaultOper,error) {
logger.Println("DdosSystemDefaultOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosSystemDefaultOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
