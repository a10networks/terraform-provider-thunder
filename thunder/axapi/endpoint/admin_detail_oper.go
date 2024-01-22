

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AdminDetailOper struct {
    
    Oper AdminDetailOperOper `json:"oper"`

}
type DataAdminDetailOper struct {
    DtAdminDetailOper AdminDetailOper `json:"admin-detail"`
}


type AdminDetailOperOper struct {
    UserList []AdminDetailOperOperUserList `json:"user-list"`
}


type AdminDetailOperOperUserList struct {
    User_name string `json:"user_name"`
    Status string `json:"status"`
    Priviledge string `json:"priviledge"`
    Partition string `json:"partition"`
    Access_type string `json:"access_type"`
    Gui_role string `json:"gui_role"`
    Trusted_host string `json:"trusted_host"`
    Lock_status string `json:"lock_status"`
    Lock_time string `json:"lock_time"`
    Unlock_time string `json:"unlock_time"`
    Password_type string `json:"password_type"`
}

func (p *AdminDetailOper) GetId() string{
    return "1"
}

func (p *AdminDetailOper) getPath() string{
    return "admin-detail/oper"
}

func (p *AdminDetailOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAdminDetailOper,error) {
logger.Println("AdminDetailOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAdminDetailOper
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
