

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Lw4o6BindingTableFilesStatusOper struct {
    
    Oper Cgnv6Lw4o6BindingTableFilesStatusOperOper `json:"oper"`

}
type DataCgnv6Lw4o6BindingTableFilesStatusOper struct {
    DtCgnv6Lw4o6BindingTableFilesStatusOper Cgnv6Lw4o6BindingTableFilesStatusOper `json:"binding-table-files-status"`
}


type Cgnv6Lw4o6BindingTableFilesStatusOperOper struct {
    EntryList []Cgnv6Lw4o6BindingTableFilesStatusOperOperEntryList `json:"entry-list"`
    EntryCount int `json:"entry-count"`
}


type Cgnv6Lw4o6BindingTableFilesStatusOperOperEntryList struct {
    Name string `json:"name"`
    Active string `json:"active"`
    Modified string `json:"modified"`
}

func (p *Cgnv6Lw4o6BindingTableFilesStatusOper) GetId() string{
    return "1"
}

func (p *Cgnv6Lw4o6BindingTableFilesStatusOper) getPath() string{
    return "cgnv6/lw-4o6/binding-table-files-status/oper"
}

func (p *Cgnv6Lw4o6BindingTableFilesStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6Lw4o6BindingTableFilesStatusOper,error) {
logger.Println("Cgnv6Lw4o6BindingTableFilesStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6Lw4o6BindingTableFilesStatusOper
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
