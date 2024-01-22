

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerEnableTrapsSystem struct {
	Inst struct {

    All int `json:"all"`
    AppsGlobal SnmpServerEnableTrapsSystemAppsGlobal1482 `json:"apps-global"`
    ControlCpuHigh int `json:"control-cpu-high"`
    DataCpuHigh int `json:"data-cpu-high"`
    Fan int `json:"fan"`
    FileSysReadOnly int `json:"file-sys-read-only"`
    HighDiskUse int `json:"high-disk-use"`
    HighMemoryUse int `json:"high-memory-use"`
    HighTemp int `json:"high-temp"`
    LicenseManagement int `json:"license-management"`
    LowTemp int `json:"low-temp"`
    PacketDrop int `json:"packet-drop"`
    Power int `json:"power"`
    PriDisk int `json:"pri-disk"`
    Restart int `json:"restart"`
    SecDisk int `json:"sec-disk"`
    Shutdown int `json:"shutdown"`
    SmpResourceEvent int `json:"smp-resource-event"`
    Start int `json:"start"`
    SyslogSeverityOne int `json:"syslog-severity-one"`
    TacacsServerUpDown int `json:"tacacs-server-up-down"`
    Uuid string `json:"uuid"`

	} `json:"system"`
}


type SnmpServerEnableTrapsSystemAppsGlobal1482 struct {
    SessionsThreshold int `json:"sessions-threshold"`
    CpsThreshold int `json:"cps-threshold"`
    Uuid string `json:"uuid"`
}

func (p *SnmpServerEnableTrapsSystem) GetId() string{
    return "1"
}

func (p *SnmpServerEnableTrapsSystem) getPath() string{
    return "snmp-server/enable/traps/system"
}

func (p *SnmpServerEnableTrapsSystem) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSystem::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *SnmpServerEnableTrapsSystem) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSystem::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *SnmpServerEnableTrapsSystem) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSystem::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *SnmpServerEnableTrapsSystem) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerEnableTrapsSystem::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
