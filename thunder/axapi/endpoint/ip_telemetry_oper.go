

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpTelemetryOper struct {
    
    Oper IpTelemetryOperOper `json:"oper"`

}
type DataIpTelemetryOper struct {
    DtIpTelemetryOper IpTelemetryOper `json:"telemetry"`
}


type IpTelemetryOperOper struct {
    Total int `json:"Total"`
    Db []IpTelemetryOperOperDb `json:"db"`
}


type IpTelemetryOperOperDb struct {
    Ipv4_addr string `json:"ipv4_addr"`
    NetMaskLen int `json:"net-mask-len"`
    ClassListName string `json:"class-list-name"`
    ClassListId int `json:"class-list-id"`
    ZoneName string `json:"zone-name"`
    Origin_as int `json:"origin_as"`
    Peer_as int `json:"peer_as"`
    Ipv4Nexthop string `json:"ipv4-nexthop"`
}

func (p *IpTelemetryOper) GetId() string{
    return "1"
}

func (p *IpTelemetryOper) getPath() string{
    return "ip/telemetry/oper"
}

func (p *IpTelemetryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpTelemetryOper,error) {
logger.Println("IpTelemetryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpTelemetryOper
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
