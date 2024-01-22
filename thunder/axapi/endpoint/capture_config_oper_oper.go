

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type CaptureConfigOperOper struct {
    
    Oper CaptureConfigOperOperOper `json:"oper"`

}
type DataCaptureConfigOperOper struct {
    DtCaptureConfigOperOper CaptureConfigOperOper `json:"capture-config-oper"`
}


type CaptureConfigOperOperOper struct {
    Name string `json:"name"`
    Action string `json:"action"`
    Status string `json:"status"`
    Filesize_kbyte int `json:"filesize_kbyte"`
    Pkt_count int `json:"pkt_count"`
    Max_pkt_count int `json:"max_pkt_count"`
    Has_file_history int `json:"has_file_history"`
    Max_filesize_kbyte int `json:"max_filesize_kbyte"`
    Snaplen int `json:"snaplen"`
    Filter string `json:"filter"`
}

func (p *CaptureConfigOperOper) GetId() string{
    return "1"
}

func (p *CaptureConfigOperOper) getPath() string{
    return "capture-config-oper/oper"
}

func (p *CaptureConfigOperOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCaptureConfigOperOper,error) {
logger.Println("CaptureConfigOperOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCaptureConfigOperOper
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
