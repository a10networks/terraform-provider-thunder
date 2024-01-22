

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetflowMonitorOper struct {
    
    Name string `json:"name"`
    Oper NetflowMonitorOperOper `json:"oper"`

}
type DataNetflowMonitorOper struct {
    DtNetflowMonitorOper NetflowMonitorOper `json:"monitor"`
}


type NetflowMonitorOperOper struct {
    Protocol string `json:"protocol"`
    Status string `json:"status"`
    Filter string `json:"filter"`
    Destination string `json:"destination"`
    SourceIpUseMgmt string `json:"source-ip-use-mgmt"`
    SourceIpAddr string `json:"source-ip-addr"`
    FlowTimeout string `json:"flow-timeout"`
    CounterPollingInterval string `json:"counter-polling-interval"`
    ResendTemplatePerRecords string `json:"resend-template-per-records"`
    ResendTemplateTimeout string `json:"resend-template-timeout"`
}

func (p *NetflowMonitorOper) GetId() string{
    return "1"
}

func (p *NetflowMonitorOper) getPath() string{
    
    return "netflow/monitor/"+p.Name+"/oper"
}

func (p *NetflowMonitorOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetflowMonitorOper,error) {
logger.Println("NetflowMonitorOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetflowMonitorOper
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
