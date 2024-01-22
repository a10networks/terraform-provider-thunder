

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbVirtualServerMigrateVipOper struct {
    
    Oper SlbVirtualServerMigrateVipOperOper `json:"oper"`
    Name string 

}
type DataSlbVirtualServerMigrateVipOper struct {
    DtSlbVirtualServerMigrateVipOper SlbVirtualServerMigrateVipOper `json:"migrate-vip"`
}


type SlbVirtualServerMigrateVipOperOper struct {
    State string `json:"state"`
}

func (p *SlbVirtualServerMigrateVipOper) GetId() string{
    return "1"
}

func (p *SlbVirtualServerMigrateVipOper) getPath() string{
    
    return "slb/virtual-server/" +p.Name + "/migrate-vip/oper"
}

func (p *SlbVirtualServerMigrateVipOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbVirtualServerMigrateVipOper,error) {
logger.Println("SlbVirtualServerMigrateVipOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbVirtualServerMigrateVipOper
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
