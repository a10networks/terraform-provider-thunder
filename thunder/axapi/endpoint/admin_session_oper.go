

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AdminSessionOper struct {
    
    Oper AdminSessionOperOper `json:"oper"`

}
type DataAdminSessionOper struct {
    DtAdminSessionOper AdminSessionOper `json:"admin-session"`
}


type AdminSessionOperOper struct {
    SessionList []AdminSessionOperOperSessionList `json:"session-list"`
}


type AdminSessionOperOperSessionList struct {
    Sid string `json:"sid"`
    Name string `json:"name"`
    Start_time string `json:"start_time"`
    Src_ip string `json:"src_ip"`
    Type string `json:"type"`
    Partition string `json:"Partition"`
    Authen string `json:"Authen"`
    Role string `json:"Role"`
    Cfg_mode string `json:"cfg_mode"`
    Priv string `json:"priv"`
}

func (p *AdminSessionOper) GetId() string{
    return "1"
}

func (p *AdminSessionOper) getPath() string{
    return "admin-session/oper"
}

func (p *AdminSessionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAdminSessionOper,error) {
logger.Println("AdminSessionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAdminSessionOper
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
