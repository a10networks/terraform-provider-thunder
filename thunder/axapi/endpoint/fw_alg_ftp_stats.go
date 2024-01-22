

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwAlgFtpStats struct {
    
    Stats FwAlgFtpStatsStats `json:"stats"`

}
type DataFwAlgFtpStats struct {
    DtFwAlgFtpStats FwAlgFtpStats `json:"ftp"`
}


type FwAlgFtpStatsStats struct {
    ClientPortRequest int `json:"client-port-request"`
    ClientEprtRequest int `json:"client-eprt-request"`
    ServerPasvReply int `json:"server-pasv-reply"`
    ServerEpsvReply int `json:"server-epsv-reply"`
}

func (p *FwAlgFtpStats) GetId() string{
    return "1"
}

func (p *FwAlgFtpStats) getPath() string{
    return "fw/alg/ftp/stats"
}

func (p *FwAlgFtpStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwAlgFtpStats,error) {
logger.Println("FwAlgFtpStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwAlgFtpStats
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
