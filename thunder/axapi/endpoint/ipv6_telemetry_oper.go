

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Ipv6TelemetryOper struct {
    
    Oper Ipv6TelemetryOperOper `json:"oper"`

}
type DataIpv6TelemetryOper struct {
    DtIpv6TelemetryOper Ipv6TelemetryOper `json:"telemetry"`
}


type Ipv6TelemetryOperOper struct {
    Total int `json:"Total"`
    Db []Ipv6TelemetryOperOperDb `json:"db"`
}


type Ipv6TelemetryOperOperDb struct {
    Ipv6_addr string `json:"ipv6_addr"`
    NetMaskLen int `json:"net-mask-len"`
    ClassListName string `json:"class-list-name"`
    ClassListId int `json:"class-list-id"`
    ZoneName string `json:"zone-name"`
    Origin_as int `json:"origin_as"`
    Peer_as int `json:"peer_as"`
    Ipv6Nexthop string `json:"ipv6-nexthop"`
}

func (p *Ipv6TelemetryOper) GetId() string{
    return "1"
}

func (p *Ipv6TelemetryOper) getPath() string{
    return "ipv6/telemetry/oper"
}

func (p *Ipv6TelemetryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpv6TelemetryOper,error) {
logger.Println("Ipv6TelemetryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpv6TelemetryOper
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
