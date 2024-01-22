

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatPortMappingFilesOper struct {
    
    Oper Cgnv6FixedNatPortMappingFilesOperOper `json:"oper"`

}
type DataCgnv6FixedNatPortMappingFilesOper struct {
    DtCgnv6FixedNatPortMappingFilesOper Cgnv6FixedNatPortMappingFilesOper `json:"port-mapping-files"`
}


type Cgnv6FixedNatPortMappingFilesOperOper struct {
    FileList []Cgnv6FixedNatPortMappingFilesOperOperFileList `json:"file-list"`
    Type string `json:"type"`
    Contain string `json:"contain"`
    ContainCaseSensitive string `json:"contain-case-sensitive"`
}


type Cgnv6FixedNatPortMappingFilesOperOperFileList struct {
    FileName string `json:"file-name"`
    WriteStatus string `json:"write-status"`
}

func (p *Cgnv6FixedNatPortMappingFilesOper) GetId() string{
    return "1"
}

func (p *Cgnv6FixedNatPortMappingFilesOper) getPath() string{
    return "cgnv6/fixed-nat/port-mapping-files/oper"
}

func (p *Cgnv6FixedNatPortMappingFilesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6FixedNatPortMappingFilesOper,error) {
logger.Println("Cgnv6FixedNatPortMappingFilesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6FixedNatPortMappingFilesOper
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
