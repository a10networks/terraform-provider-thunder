

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbHealthStatOper struct {
    
    Oper SlbHealthStatOperOper `json:"oper"`

}
type DataSlbHealthStatOper struct {
    DtSlbHealthStatOper SlbHealthStatOper `json:"health-stat"`
}


type SlbHealthStatOperOper struct {
    HealthCheckList []SlbHealthStatOperOperHealthCheckList `json:"health-check-list"`
    Num_pins int `json:"num_pins"`
    Num_pins_stat_up int `json:"num_pins_stat_up"`
    Num_pins_stat_down int `json:"num_pins_stat_down"`
    Num_pins_stat_unkn int `json:"num_pins_stat_unkn"`
    Num_pins_stat_else int `json:"num_pins_stat_else"`
    Num_ssl_tickets int `json:"num_ssl_tickets"`
    Total_stat int `json:"total_stat"`
    Method string `json:"method"`
    Clear_ssl_ticket int `json:"clear_ssl_ticket"`
    Monitor string `json:"monitor"`
}


type SlbHealthStatOperOperHealthCheckList struct {
    IpAddress string `json:"ip-address"`
    Port string `json:"port"`
    HealthMonitor string `json:"health-monitor"`
    Status string `json:"status"`
    UpCause int `json:"up-cause"`
    DownCause int `json:"down-cause"`
    DownState int `json:"down-state"`
    Reason string `json:"reason"`
    TotalRetry int `json:"total-retry"`
    Retries int `json:"retries"`
    UpRetries int `json:"up-retries"`
    PartitionId int `json:"partition-id"`
    Server string `json:"server"`
    SslVersion string `json:"ssl-version"`
    SslCipher string `json:"ssl-cipher"`
    SslTicket int `json:"ssl-ticket"`
}

func (p *SlbHealthStatOper) GetId() string{
    return "1"
}

func (p *SlbHealthStatOper) getPath() string{
    return "slb/health-stat/oper"
}

func (p *SlbHealthStatOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbHealthStatOper,error) {
logger.Println("SlbHealthStatOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbHealthStatOper
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
