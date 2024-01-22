

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemPathOper struct {
    
    L2hmPathName string `json:"l2hm-path-name"`
    Oper SystemPathOperOper `json:"oper"`

}
type DataSystemPathOper struct {
    DtSystemPathOper SystemPathOper `json:"path"`
}


type SystemPathOperOper struct {
    Path_list []SystemPathOperOperPath_list `json:"path_list"`
}


type SystemPathOperOperPath_list struct {
    L2hm_path string `json:"l2hm_path"`
    Endpoint_1 string `json:"endpoint_1"`
    Endpoint_2 string `json:"endpoint_2"`
    Health_check string `json:"health_check"`
    Path_state string `json:"path_state"`
    Apps string `json:"apps"`
}

func (p *SystemPathOper) GetId() string{
    return "1"
}

func (p *SystemPathOper) getPath() string{
    
    return "system/path/"+p.L2hmPathName+"/oper"
}

func (p *SystemPathOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemPathOper,error) {
logger.Println("SystemPathOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemPathOper
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
