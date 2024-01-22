

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VersionOper struct {
    
    Oper VersionOperOper `json:"oper"`

}
type DataVersionOper struct {
    DtVersionOper VersionOper `json:"version"`
}


type VersionOperOper struct {
    HwPlatform string `json:"hw-platform"`
    Copyright string `json:"copyright"`
    SwVersion string `json:"sw-version"`
    PlatFeatures string `json:"plat-features"`
    BootFrom string `json:"boot-from"`
    SerialNumber string `json:"serial-number"`
    AflexVersion string `json:"aflex-version"`
    AxapiVersion string `json:"axapi-version"`
    PriGuiVersion string `json:"pri-gui-version"`
    SecGuiVersion string `json:"sec-gui-version"`
    FirmwareVersion string `json:"firmware-version"`
    HdPri string `json:"hd-pri"`
    HdSec string `json:"hd-sec"`
    CfPri string `json:"cf-pri"`
    CfSec string `json:"cf-sec"`
    LastConfigSavedTime string `json:"last-config-saved-time"`
    VirtualizationType string `json:"virtualization-type"`
    SysPollMode string `json:"sys-poll-mode"`
    Product string `json:"product"`
    HwCode string `json:"hw-code"`
    CurrentTime string `json:"current-time"`
    UpTime string `json:"up-time"`
    NunCtrlCpus int `json:"nun-ctrl-cpus"`
    BuffSize int `json:"buff-size"`
    IoBuffEnabled string `json:"io-buff-enabled"`
    BuildType string `json:"build-type"`
    CotsSysMfg string `json:"cots-sys-mfg"`
    CotsSysName string `json:"cots-sys-name"`
    CotsSysVer string `json:"cots-sys-ver"`
    SeriesName string `json:"series-name"`
    Hostname string `json:"hostname"`
    Alldynamic int `json:"alldynamic"`
}

func (p *VersionOper) GetId() string{
    return "1"
}

func (p *VersionOper) getPath() string{
    return "version/oper"
}

func (p *VersionOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVersionOper,error) {
logger.Println("VersionOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVersionOper
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
