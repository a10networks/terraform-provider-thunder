

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetworkAvailableTrunkListOper struct {
    
    Oper NetworkAvailableTrunkListOperOper `json:"oper"`

}
type DataNetworkAvailableTrunkListOper struct {
    DtNetworkAvailableTrunkListOper NetworkAvailableTrunkListOper `json:"available-trunk-list"`
}


type NetworkAvailableTrunkListOperOper struct {
    IfList []NetworkAvailableTrunkListOperOperIfList `json:"if-list"`
}


type NetworkAvailableTrunkListOperOperIfList struct {
    IfType string `json:"IF-Type"`
    IfNum int `json:"IF-Num"`
    IfStatus string `json:"IF-Status"`
}

func (p *NetworkAvailableTrunkListOper) GetId() string{
    return "1"
}

func (p *NetworkAvailableTrunkListOper) getPath() string{
    return "network/available-trunk-list/oper"
}

func (p *NetworkAvailableTrunkListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNetworkAvailableTrunkListOper,error) {
logger.Println("NetworkAvailableTrunkListOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNetworkAvailableTrunkListOper
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
