

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScmLicensestatusOper struct {
    
    Oper ScmLicensestatusOperOper `json:"oper"`

}
type DataScmLicensestatusOper struct {
    DtScmLicensestatusOper ScmLicensestatusOper `json:"licensestatus"`
}


type ScmLicensestatusOperOper struct {
    PrimaryOrgId string `json:"primary-org-id"`
    PrimaryOrgEmail string `json:"primary-org-email"`
    LicenseList []ScmLicensestatusOperOperLicenseList `json:"license-list"`
}


type ScmLicensestatusOperOperLicenseList struct {
    LicenseType string `json:"license-type"`
    LicenseName string `json:"license-name"`
    OrgName string `json:"org-name"`
    Status string `json:"status"`
}

func (p *ScmLicensestatusOper) GetId() string{
    return "1"
}

func (p *ScmLicensestatusOper) getPath() string{
    return "scm/licensestatus/oper"
}

func (p *ScmLicensestatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScmLicensestatusOper,error) {
logger.Println("ScmLicensestatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScmLicensestatusOper
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
