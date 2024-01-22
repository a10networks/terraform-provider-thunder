

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbHealthGatewayOper struct {
    
    Oper SlbHealthGatewayOperOper `json:"oper"`

}
type DataSlbHealthGatewayOper struct {
    DtSlbHealthGatewayOper SlbHealthGatewayOper `json:"health-gateway"`
}


type SlbHealthGatewayOperOper struct {
    Enabled int `json:"enabled"`
    Interval int `json:"interval"`
    Timeout int `json:"timeout"`
}

func (p *SlbHealthGatewayOper) GetId() string{
    return "1"
}

func (p *SlbHealthGatewayOper) getPath() string{
    return "slb/health-gateway/oper"
}

func (p *SlbHealthGatewayOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbHealthGatewayOper,error) {
logger.Println("SlbHealthGatewayOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbHealthGatewayOper
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
