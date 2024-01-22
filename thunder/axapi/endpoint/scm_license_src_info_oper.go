

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScmLicenseSrcInfoOper struct {
    
    Oper ScmLicenseSrcInfoOperOper `json:"oper"`

}
type DataScmLicenseSrcInfoOper struct {
    DtScmLicenseSrcInfoOper ScmLicenseSrcInfoOper `json:"license-src-info"`
}


type ScmLicenseSrcInfoOperOper struct {
    Uuid string `json:"uuid"`
    UsbUuid string `json:"usb-uuid"`
    BillingSerial string `json:"billing-serial"`
    Token string `json:"token"`
    Product string `json:"product"`
    Platform string `json:"platform"`
    Burst string `json:"burst"`
    Version string `json:"version"`
    GlmPingInterval int `json:"glm-ping-interval"`
    Source1 string `json:"source1"`
    Source1ModuleList []ScmLicenseSrcInfoOperOperSource1ModuleList `json:"source1-module-list"`
    Source2 string `json:"source2"`
    Source2ModuleList []ScmLicenseSrcInfoOperOperSource2ModuleList `json:"source2-module-list"`
    Source3 string `json:"source3"`
    Source3ModuleList []ScmLicenseSrcInfoOperOperSource3ModuleList `json:"source3-module-list"`
    HwSerialno string `json:"hw-serialno"`
    ProductDesc string `json:"product-desc"`
}


type ScmLicenseSrcInfoOperOperSource1ModuleList struct {
    Source1Module string `json:"source1-module"`
    Source1Expiry string `json:"source1-expiry"`
    Source1Notes string `json:"source1-notes"`
}


type ScmLicenseSrcInfoOperOperSource2ModuleList struct {
    Source2Module string `json:"source2-module"`
    Source2Expiry string `json:"source2-expiry"`
    Source2Notes string `json:"source2-notes"`
}


type ScmLicenseSrcInfoOperOperSource3ModuleList struct {
    Source3Module string `json:"source3-module"`
    Source3Expiry string `json:"source3-expiry"`
    Source3Notes string `json:"source3-notes"`
}

func (p *ScmLicenseSrcInfoOper) GetId() string{
    return "1"
}

func (p *ScmLicenseSrcInfoOper) getPath() string{
    return "scm/license-src-info/oper"
}

func (p *ScmLicenseSrcInfoOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScmLicenseSrcInfoOper,error) {
logger.Println("ScmLicenseSrcInfoOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScmLicenseSrcInfoOper
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
