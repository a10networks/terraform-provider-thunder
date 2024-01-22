

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSslJa3Oper struct {
    
    Oper SlbSslJa3OperOper `json:"oper"`

}
type DataSlbSslJa3Oper struct {
    DtSlbSslJa3Oper SlbSslJa3Oper `json:"ssl-ja3"`
}


type SlbSslJa3OperOper struct {
    Record []SlbSslJa3OperOperRecord `json:"record"`
}


type SlbSslJa3OperOperRecord struct {
    AddrV4 string `json:"addr-v4"`
    AddrV6 string `json:"addr-v6"`
    Amount int `json:"amount"`
}

func (p *SlbSslJa3Oper) GetId() string{
    return "1"
}

func (p *SlbSslJa3Oper) getPath() string{
    return "slb/ssl-ja3/oper"
}

func (p *SlbSslJa3Oper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSslJa3Oper,error) {
logger.Println("SlbSslJa3Oper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSslJa3Oper
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
