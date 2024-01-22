

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbFtpDataStats struct {
    
    Stats SlbFtpDataStatsStats `json:"stats"`

}
type DataSlbFtpDataStats struct {
    DtSlbFtpDataStats SlbFtpDataStats `json:"ftp-data"`
}


type SlbFtpDataStatsStats struct {
    Sessions_num int `json:"sessions_num"`
    Port_out_of_range int `json:"port_out_of_range"`
}

func (p *SlbFtpDataStats) GetId() string{
    return "1"
}

func (p *SlbFtpDataStats) getPath() string{
    return "slb/ftp-data/stats"
}

func (p *SlbFtpDataStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbFtpDataStats,error) {
logger.Println("SlbFtpDataStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbFtpDataStats
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
