

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Lw4o6ValidationLogFilesOper struct {
    
    Oper Cgnv6Lw4o6ValidationLogFilesOperOper `json:"oper"`

}
type DataCgnv6Lw4o6ValidationLogFilesOper struct {
    DtCgnv6Lw4o6ValidationLogFilesOper Cgnv6Lw4o6ValidationLogFilesOper `json:"validation-log-files"`
}


type Cgnv6Lw4o6ValidationLogFilesOperOper struct {
    EntryList []Cgnv6Lw4o6ValidationLogFilesOperOperEntryList `json:"entry-list"`
    EntryCount int `json:"entry-count"`
}


type Cgnv6Lw4o6ValidationLogFilesOperOperEntryList struct {
    Name string `json:"name"`
}

func (p *Cgnv6Lw4o6ValidationLogFilesOper) GetId() string{
    return "1"
}

func (p *Cgnv6Lw4o6ValidationLogFilesOper) getPath() string{
    return "cgnv6/lw-4o6/validation-log-files/oper"
}

func (p *Cgnv6Lw4o6ValidationLogFilesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6Lw4o6ValidationLogFilesOper,error) {
logger.Println("Cgnv6Lw4o6ValidationLogFilesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6Lw4o6ValidationLogFilesOper
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
