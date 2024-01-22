

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Lw4o6AllBindingTablesOper struct {
    
    Oper Cgnv6Lw4o6AllBindingTablesOperOper `json:"oper"`

}
type DataCgnv6Lw4o6AllBindingTablesOper struct {
    DtCgnv6Lw4o6AllBindingTablesOper Cgnv6Lw4o6AllBindingTablesOper `json:"all-binding-tables"`
}


type Cgnv6Lw4o6AllBindingTablesOperOper struct {
    EntryList []Cgnv6Lw4o6AllBindingTablesOperOperEntryList `json:"entry-list"`
}


type Cgnv6Lw4o6AllBindingTablesOperOperEntryList struct {
    Name string `json:"name"`
}

func (p *Cgnv6Lw4o6AllBindingTablesOper) GetId() string{
    return "1"
}

func (p *Cgnv6Lw4o6AllBindingTablesOper) getPath() string{
    return "cgnv6/lw-4o6/all-binding-tables/oper"
}

func (p *Cgnv6Lw4o6AllBindingTablesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6Lw4o6AllBindingTablesOper,error) {
logger.Println("Cgnv6Lw4o6AllBindingTablesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6Lw4o6AllBindingTablesOper
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
