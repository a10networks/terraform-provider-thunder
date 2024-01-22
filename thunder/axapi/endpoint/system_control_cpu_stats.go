

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemControlCpuStats struct {
    
    Stats SystemControlCpuStatsStats `json:"stats"`

}
type DataSystemControlCpuStats struct {
    DtSystemControlCpuStats SystemControlCpuStats `json:"control-cpu"`
}


type SystemControlCpuStatsStats struct {
    CtrlCpuNumber int `json:"ctrl-cpu-number"`
    Cpu1 int `json:"cpu-1"`
    Cpu2 int `json:"cpu-2"`
    Cpu3 int `json:"cpu-3"`
    Cpu4 int `json:"cpu-4"`
    Cpu5 int `json:"cpu-5"`
    Cpu6 int `json:"cpu-6"`
    Cpu7 int `json:"cpu-7"`
    Cpu8 int `json:"cpu-8"`
    Cpu9 int `json:"cpu-9"`
    Cpu10 int `json:"cpu-10"`
    Cpu11 int `json:"cpu-11"`
    Cpu12 int `json:"cpu-12"`
    Cpu13 int `json:"cpu-13"`
    Cpu14 int `json:"cpu-14"`
    Cpu15 int `json:"cpu-15"`
    Cpu16 int `json:"cpu-16"`
    Cpu17 int `json:"cpu-17"`
    Cpu18 int `json:"cpu-18"`
    Cpu19 int `json:"cpu-19"`
    Cpu20 int `json:"cpu-20"`
    Cpu21 int `json:"cpu-21"`
    Cpu22 int `json:"cpu-22"`
    Cpu23 int `json:"cpu-23"`
    Cpu24 int `json:"cpu-24"`
    Cpu25 int `json:"cpu-25"`
    Cpu26 int `json:"cpu-26"`
    Cpu27 int `json:"cpu-27"`
    Cpu28 int `json:"cpu-28"`
    Cpu29 int `json:"cpu-29"`
    Cpu30 int `json:"cpu-30"`
    Cpu31 int `json:"cpu-31"`
    Cpu32 int `json:"cpu-32"`
    Cpu33 int `json:"cpu-33"`
    Cpu34 int `json:"cpu-34"`
    Cpu35 int `json:"cpu-35"`
    Cpu36 int `json:"cpu-36"`
    Cpu37 int `json:"cpu-37"`
    Cpu38 int `json:"cpu-38"`
    Cpu39 int `json:"cpu-39"`
    Cpu40 int `json:"cpu-40"`
    Cpu41 int `json:"cpu-41"`
    Cpu42 int `json:"cpu-42"`
    Cpu43 int `json:"cpu-43"`
    Cpu44 int `json:"cpu-44"`
    Cpu45 int `json:"cpu-45"`
    Cpu46 int `json:"cpu-46"`
    Cpu47 int `json:"cpu-47"`
    Cpu48 int `json:"cpu-48"`
    Cpu49 int `json:"cpu-49"`
    Cpu50 int `json:"cpu-50"`
    Cpu51 int `json:"cpu-51"`
    Cpu52 int `json:"cpu-52"`
    Cpu53 int `json:"cpu-53"`
    Cpu54 int `json:"cpu-54"`
    Cpu55 int `json:"cpu-55"`
    Cpu56 int `json:"cpu-56"`
    Cpu57 int `json:"cpu-57"`
    Cpu58 int `json:"cpu-58"`
    Cpu59 int `json:"cpu-59"`
    Cpu60 int `json:"cpu-60"`
    Cpu61 int `json:"cpu-61"`
    Cpu62 int `json:"cpu-62"`
    Cpu63 int `json:"cpu-63"`
    Cpu64 int `json:"cpu-64"`
}

func (p *SystemControlCpuStats) GetId() string{
    return "1"
}

func (p *SystemControlCpuStats) getPath() string{
    return "system/control-cpu/stats"
}

func (p *SystemControlCpuStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemControlCpuStats,error) {
logger.Println("SystemControlCpuStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemControlCpuStats
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
