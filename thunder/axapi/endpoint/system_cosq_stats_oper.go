

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemCosqStatsOper struct {
    
    Oper SystemCosqStatsOperOper `json:"oper"`

}
type DataSystemCosqStatsOper struct {
    DtSystemCosqStatsOper SystemCosqStatsOper `json:"cosq-stats"`
}


type SystemCosqStatsOperOper struct {
    CosqIndex int `json:"cosq-index"`
    CosqList []SystemCosqStatsOperOperCosqList `json:"cosq-list"`
}


type SystemCosqStatsOperOperCosqList struct {
    EthernetIndex int `json:"ethernet-index"`
    XauiIndex int `json:"xaui-index"`
    InterfaceName string `json:"interface-name"`
    CosqDrop int `json:"cosq-drop"`
    CosqInpkt int `json:"cosq-inpkt"`
    CosqOutpkt int `json:"cosq-outpkt"`
    CosqSharedpkt int `json:"cosq-sharedpkt"`
    CosqUcpeakpkt int `json:"cosq-ucpeakpkt"`
    CosqUccongdrop int `json:"cosq-uccongdrop"`
    CosqCongdrop int `json:"cosq-congdrop"`
}

func (p *SystemCosqStatsOper) GetId() string{
    return "1"
}

func (p *SystemCosqStatsOper) getPath() string{
    return "system/cosq-stats/oper"
}

func (p *SystemCosqStatsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemCosqStatsOper,error) {
logger.Println("SystemCosqStatsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemCosqStatsOper
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
