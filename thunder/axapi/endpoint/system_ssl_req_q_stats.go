

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemSslReqQStats struct {
    
    Stats SystemSslReqQStatsStats `json:"stats"`

}
type DataSystemSslReqQStats struct {
    DtSystemSslReqQStats SystemSslReqQStats `json:"ssl-req-q"`
}


type SystemSslReqQStatsStats struct {
    NumSslQueues int `json:"num-ssl-queues"`
    SslReqQDepthTot int `json:"ssl-req-q-depth-tot"`
    SslReqQInuseTot int `json:"ssl-req-q-inuse-tot"`
    SslHwQDepthTot int `json:"ssl-hw-q-depth-tot"`
    SslHwQInuseTot int `json:"ssl-hw-q-inuse-tot"`
}

func (p *SystemSslReqQStats) GetId() string{
    return "1"
}

func (p *SystemSslReqQStats) getPath() string{
    return "system/ssl-req-q/stats"
}

func (p *SystemSslReqQStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemSslReqQStats,error) {
logger.Println("SystemSslReqQStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemSslReqQStats
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
