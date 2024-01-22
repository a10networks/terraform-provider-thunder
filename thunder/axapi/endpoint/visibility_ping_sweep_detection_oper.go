

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityPingSweepDetectionOper struct {
    
    Oper VisibilityPingSweepDetectionOperOper `json:"oper"`

}
type DataVisibilityPingSweepDetectionOper struct {
    DtVisibilityPingSweepDetectionOper VisibilityPingSweepDetectionOper `json:"ping-sweep-detection"`
}


type VisibilityPingSweepDetectionOperOper struct {
    SrcIpList []VisibilityPingSweepDetectionOperOperSrcIpList `json:"src-ip-list"`
}


type VisibilityPingSweepDetectionOperOperSrcIpList struct {
    SrcIpAddr string `json:"src-ip-addr"`
    ScannedIps int `json:"scanned-ips"`
    IpList []VisibilityPingSweepDetectionOperOperSrcIpListIpList `json:"ip-list"`
}


type VisibilityPingSweepDetectionOperOperSrcIpListIpList struct {
    Ip string `json:"ip"`
    Scanned_time string `json:"scanned_time"`
}

func (p *VisibilityPingSweepDetectionOper) GetId() string{
    return "1"
}

func (p *VisibilityPingSweepDetectionOper) getPath() string{
    return "visibility/ping-sweep-detection/oper"
}

func (p *VisibilityPingSweepDetectionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityPingSweepDetectionOper,error) {
logger.Println("VisibilityPingSweepDetectionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityPingSweepDetectionOper
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
