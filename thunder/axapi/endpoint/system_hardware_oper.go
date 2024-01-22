

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemHardwareOper struct {
    
    Oper SystemHardwareOperOper `json:"oper"`

}
type DataSystemHardwareOper struct {
    DtSystemHardwareOper SystemHardwareOper `json:"hardware"`
}


type SystemHardwareOperOper struct {
    PlatformDescription string `json:"platform-description"`
    Serial string `json:"serial"`
    Cpu string `json:"cpu"`
    CpuCores int `json:"cpu-cores"`
    CpuStepping int `json:"cpu-stepping"`
    Storage string `json:"storage"`
    Memory string `json:"memory"`
    SslCards SystemHardwareOperOperSslCards `json:"ssl-cards"`
    Octeon int `json:"octeon"`
    CompressionCards SystemHardwareOperOperCompressionCards `json:"compression-cards"`
    L23Asic string `json:"l23-asic"`
    Ipmi string `json:"ipmi"`
    Ports string `json:"ports"`
    PlatFlag string `json:"plat-flag"`
    BiosVersion string `json:"bios-version"`
    BiosReleaseDate string `json:"bios-release-date"`
    NvmFirmwareVersion string `json:"nvm-firmware-version"`
    FpgaSummary string `json:"fpga-summary"`
    FpgaDate string `json:"fpga-date"`
    DiskTotal int `json:"disk-total"`
    DiskUsed int `json:"disk-used"`
    DiskFree int `json:"disk-free"`
    DiskPercentage int `json:"disk-percentage"`
    Disk1Status string `json:"disk1-status"`
    Disk2Status string `json:"disk2-status"`
    NumDisks int `json:"num-disks"`
    Raid_present int `json:"raid_present"`
    RaidList []SystemHardwareOperOperRaidList `json:"raid-list"`
    Psu1_np15 string `json:"psu1_np15"`
    Psu2_np15 string `json:"psu2_np15"`
    Spe_present string `json:"spe_present"`
    BypassPr int `json:"bypass-pr"`
    BypassList []SystemHardwareOperOperBypassList `json:"bypass-list"`
    Alldynamic int `json:"alldynamic"`
    McpldType int `json:"mcpld-type"`
    McpldDate string `json:"mcpld-date"`
}


type SystemHardwareOperOperSslCards struct {
    SslDevices int `json:"ssl-devices"`
    Nitroxpx int `json:"nitroxpx"`
    Nitrox3 int `json:"nitrox3"`
    Nitrox3Cores int `json:"nitrox3-cores"`
    Nitrox5 int `json:"nitrox5"`
    Nitrox5Cores int `json:"nitrox5-cores"`
    Nitrox2 int `json:"nitrox2"`
    Nitrox1 int `json:"nitrox1"`
    Hsm int `json:"hsm"`
    UnknownSslCards int `json:"unknown-ssl-cards"`
    ColetoSslCards int `json:"coleto-ssl-cards"`
}


type SystemHardwareOperOperCompressionCards struct {
    GzipDevices int `json:"gzip-devices"`
    Aha363 int `json:"aha363"`
    UnknownCompression int `json:"unknown-compression"`
}


type SystemHardwareOperOperRaidList struct {
    MdName string `json:"md-name"`
    MdPri string `json:"md-pri"`
    MdSec string `json:"md-sec"`
}


type SystemHardwareOperOperBypassList struct {
    BypassName string `json:"bypass-name"`
    BypassInfo string `json:"bypass-info"`
}

func (p *SystemHardwareOper) GetId() string{
    return "1"
}

func (p *SystemHardwareOper) getPath() string{
    return "system/hardware/oper"
}

func (p *SystemHardwareOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemHardwareOper,error) {
logger.Println("SystemHardwareOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemHardwareOper
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
