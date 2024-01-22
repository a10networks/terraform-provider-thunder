

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type TemplateAppsOper struct {
    
    Oper TemplateAppsOperOper `json:"oper"`

}
type DataTemplateAppsOper struct {
    DtTemplateAppsOper TemplateAppsOper `json:"apps"`
}


type TemplateAppsOperOper struct {
    Installed []TemplateAppsOperOperInstalled `json:"installed"`
    Available []TemplateAppsOperOperAvailable `json:"available"`
}


type TemplateAppsOperOperInstalled struct {
    Version string `json:"version"`
    Category_slug string `json:"category_slug"`
    Description string `json:"description"`
    Type string `json:"type"`
    Release_notes string `json:"release_notes"`
    Video_url string `json:"video_url"`
    Guide_url string `json:"guide_url"`
    Developer string `json:"developer"`
    Icon string `json:"icon"`
    Supported_acos []TemplateAppsOperOperInstalledSupported_acos `json:"supported_acos"`
    File_name string `json:"file_name"`
    Deleted string `json:"deleted"`
    Color string `json:"color"`
    Category_name string `json:"category_name"`
}


type TemplateAppsOperOperInstalledSupported_acos struct {
    Version string `json:"version"`
}


type TemplateAppsOperOperAvailable struct {
    Version string `json:"version"`
    Category_slug string `json:"category_slug"`
    Description string `json:"description"`
    Type string `json:"type"`
    Release_notes string `json:"release_notes"`
    Video_url string `json:"video_url"`
    Guide_url string `json:"guide_url"`
    Developer string `json:"developer"`
    Icon string `json:"icon"`
    Supported_acos []TemplateAppsOperOperAvailableSupported_acos `json:"supported_acos"`
    File_name string `json:"file_name"`
    Deleted string `json:"deleted"`
    Color string `json:"color"`
    Category_name string `json:"category_name"`
}


type TemplateAppsOperOperAvailableSupported_acos struct {
    Version string `json:"version"`
}

func (p *TemplateAppsOper) GetId() string{
    return "1"
}

func (p *TemplateAppsOper) getPath() string{
    return "template/apps/oper"
}

func (p *TemplateAppsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataTemplateAppsOper,error) {
logger.Println("TemplateAppsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataTemplateAppsOper
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
