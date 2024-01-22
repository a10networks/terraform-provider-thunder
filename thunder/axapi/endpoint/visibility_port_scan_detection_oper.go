

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPortScanDetectionOper struct {
    
    Oper VisibilityPortScanDetectionOperOper `json:"oper"`

}
type DataVisibilityPortScanDetectionOper struct {
    DtVisibilityPortScanDetectionOper VisibilityPortScanDetectionOper `json:"port-scan-detection"`
}


type VisibilityPortScanDetectionOperOper struct {
    SrcIpList []VisibilityPortScanDetectionOperOperSrcIpList `json:"src-ip-list"`
}


type VisibilityPortScanDetectionOperOperSrcIpList struct {
    SrcIpAddr string `json:"src-ip-addr"`
    ScannedPorts int `json:"scanned-ports"`
    PortList []VisibilityPortScanDetectionOperOperSrcIpListPortList `json:"port-list"`
}


type VisibilityPortScanDetectionOperOperSrcIpListPortList struct {
    Port int `json:"port"`
    Protocol string `json:"protocol"`
    Scanned_time string `json:"scanned_time"`
}

func (p *VisibilityPortScanDetectionOper) GetId() string{
    return "1"
}

func (p *VisibilityPortScanDetectionOper) getPath() string{
    return "visibility/port-scan-detection/oper"
}

func (p *VisibilityPortScanDetectionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityPortScanDetectionOper,error) {
logger.Println("VisibilityPortScanDetectionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityPortScanDetectionOper
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
