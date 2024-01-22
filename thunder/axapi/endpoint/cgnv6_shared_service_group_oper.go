

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6SharedServiceGroupOper struct {
    
    Oper Cgnv6SharedServiceGroupOperOper `json:"oper"`

}
type DataCgnv6SharedServiceGroupOper struct {
    DtCgnv6SharedServiceGroupOper Cgnv6SharedServiceGroupOper `json:"shared-service-group"`
}


type Cgnv6SharedServiceGroupOperOper struct {
    SharedServiceGroupList []Cgnv6SharedServiceGroupOperOperSharedServiceGroupList `json:"shared-service-group-list"`
}


type Cgnv6SharedServiceGroupOperOperSharedServiceGroupList struct {
    ServiceGroupName string `json:"service-group-name"`
}

func (p *Cgnv6SharedServiceGroupOper) GetId() string{
    return "1"
}

func (p *Cgnv6SharedServiceGroupOper) getPath() string{
    return "cgnv6/shared-service-group/oper"
}

func (p *Cgnv6SharedServiceGroupOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6SharedServiceGroupOper,error) {
logger.Println("Cgnv6SharedServiceGroupOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6SharedServiceGroupOper
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
