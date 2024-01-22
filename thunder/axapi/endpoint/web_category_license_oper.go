

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WebCategoryLicenseOper struct {
    
    Oper WebCategoryLicenseOperOper `json:"oper"`

}
type DataWebCategoryLicenseOper struct {
    DtWebCategoryLicenseOper WebCategoryLicenseOper `json:"license"`
}


type WebCategoryLicenseOperOper struct {
    ModuleStatus string `json:"module-status"`
    LicenseStatus string `json:"license-status"`
    LicenseType string `json:"license-type"`
    LicenseExpiry string `json:"license-expiry"`
    RemainingPeriod string `json:"remaining-period"`
    IsGrace string `json:"is-grace"`
    GracePeriod string `json:"grace-period"`
    SerialNumber string `json:"serial-number"`
}

func (p *WebCategoryLicenseOper) GetId() string{
    return "1"
}

func (p *WebCategoryLicenseOper) getPath() string{
    return "web-category/license/oper"
}

func (p *WebCategoryLicenseOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataWebCategoryLicenseOper,error) {
logger.Println("WebCategoryLicenseOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataWebCategoryLicenseOper
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
