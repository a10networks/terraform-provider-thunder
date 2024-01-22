

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwAppOper struct {
    
    Oper FwAppOperOper `json:"oper"`

}
type DataFwAppOper struct {
    DtFwAppOper FwAppOper `json:"app"`
}


type FwAppOperOper struct {
    Category string `json:"category"`
    Contains string `json:"contains"`
    Related string `json:"related"`
    GroupList []FwAppOperOperGroupList `json:"group-list"`
}


type FwAppOperOperGroupList struct {
    CategoryName string `json:"category-name"`
    AppName string `json:"app-name"`
    AppDesc string `json:"app-desc"`
}

func (p *FwAppOper) GetId() string{
    return "1"
}

func (p *FwAppOper) getPath() string{
    return "fw/app/oper"
}

func (p *FwAppOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwAppOper,error) {
logger.Println("FwAppOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataFwAppOper
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
