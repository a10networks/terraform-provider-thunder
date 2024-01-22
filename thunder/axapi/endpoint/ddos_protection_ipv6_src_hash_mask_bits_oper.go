

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosProtectionIpv6SrcHashMaskBitsOper struct {
    
    Oper DdosProtectionIpv6SrcHashMaskBitsOperOper `json:"oper"`

}
type DataDdosProtectionIpv6SrcHashMaskBitsOper struct {
    DtDdosProtectionIpv6SrcHashMaskBitsOper DdosProtectionIpv6SrcHashMaskBitsOper `json:"ipv6-src-hash-mask-bits"`
}


type DdosProtectionIpv6SrcHashMaskBitsOperOper struct {
    Offsets []DdosProtectionIpv6SrcHashMaskBitsOperOperOffsets `json:"offsets"`
}


type DdosProtectionIpv6SrcHashMaskBitsOperOperOffsets struct {
    MaskBitOffset1 int `json:"mask-bit-offset-1"`
    MaskBitOffset2 int `json:"mask-bit-offset-2"`
    MaskBitOffset3 int `json:"mask-bit-offset-3"`
    MaskBitOffset4 int `json:"mask-bit-offset-4"`
    MaskBitOffset5 int `json:"mask-bit-offset-5"`
}

func (p *DdosProtectionIpv6SrcHashMaskBitsOper) GetId() string{
    return "1"
}

func (p *DdosProtectionIpv6SrcHashMaskBitsOper) getPath() string{
    return "ddos/protection/ipv6-src-hash-mask-bits/oper"
}

func (p *DdosProtectionIpv6SrcHashMaskBitsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosProtectionIpv6SrcHashMaskBitsOper,error) {
logger.Println("DdosProtectionIpv6SrcHashMaskBitsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosProtectionIpv6SrcHashMaskBitsOper
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
