

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpAHostidOper struct {
    
    Oper VrrpAHostidOperOper `json:"oper"`

}
type DataVrrpAHostidOper struct {
    DtVrrpAHostidOper VrrpAHostidOper `json:"hostid"`
}


type VrrpAHostidOperOper struct {
    Set_id int `json:"set_ID"`
    HostidList []VrrpAHostidOperOperHostidList `json:"hostid-list"`
}


type VrrpAHostidOperOperHostidList struct {
    Device_id int `json:"device_ID"`
    Sn string `json:"SN"`
}

func (p *VrrpAHostidOper) GetId() string{
    return "1"
}

func (p *VrrpAHostidOper) getPath() string{
    return "vrrp-a/hostid/oper"
}

func (p *VrrpAHostidOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVrrpAHostidOper,error) {
logger.Println("VrrpAHostidOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVrrpAHostidOper
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
