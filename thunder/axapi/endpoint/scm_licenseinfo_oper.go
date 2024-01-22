

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScmLicenseinfoOper struct {
    
    Oper ScmLicenseinfoOperOper `json:"oper"`

}
type DataScmLicenseinfoOper struct {
    DtScmLicenseinfoOper ScmLicenseinfoOper `json:"licenseinfo"`
}


type ScmLicenseinfoOperOper struct {
    Uuid string `json:"uuid"`
    UsbUuid string `json:"usb-uuid"`
    BillingSerial string `json:"billing-serial"`
    Token string `json:"token"`
    Product string `json:"product"`
    Platform string `json:"platform"`
    Burst string `json:"burst"`
    Version string `json:"version"`
    GlmPingInterval int `json:"glm-ping-interval"`
    ModuleList []ScmLicenseinfoOperOperModuleList `json:"module-list"`
    HwSerialno string `json:"hw-serialno"`
    ProductDesc string `json:"product-desc"`
}


type ScmLicenseinfoOperOperModuleList struct {
    Module string `json:"module"`
    Expiry string `json:"expiry"`
    Notes string `json:"notes"`
}

func (p *ScmLicenseinfoOper) GetId() string{
    return "1"
}

func (p *ScmLicenseinfoOper) getPath() string{
    return "scm/licenseinfo/oper"
}

func (p *ScmLicenseinfoOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScmLicenseinfoOper,error) {
logger.Println("ScmLicenseinfoOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScmLicenseinfoOper
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
