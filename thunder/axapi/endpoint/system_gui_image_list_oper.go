

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemGuiImageListOper struct {
    
    Oper SystemGuiImageListOperOper `json:"oper"`

}
type DataSystemGuiImageListOper struct {
    DtSystemGuiImageListOper SystemGuiImageListOper `json:"gui-image-list"`
}


type SystemGuiImageListOperOper struct {
    PrePriGui string `json:"pre-pri-gui"`
    PreSecGui string `json:"pre-sec-gui"`
    GuiListPri []SystemGuiImageListOperOperGuiListPri `json:"gui-list-pri"`
    GuiListSec []SystemGuiImageListOperOperGuiListSec `json:"gui-list-sec"`
}


type SystemGuiImageListOperOperGuiListPri struct {
    GuiImage string `json:"gui-image"`
    Path string `json:"path"`
}


type SystemGuiImageListOperOperGuiListSec struct {
    GuiImage string `json:"gui-image"`
    Path string `json:"path"`
}

func (p *SystemGuiImageListOper) GetId() string{
    return "1"
}

func (p *SystemGuiImageListOper) getPath() string{
    return "system/gui-image-list/oper"
}

func (p *SystemGuiImageListOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemGuiImageListOper,error) {
logger.Println("SystemGuiImageListOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemGuiImageListOper
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
