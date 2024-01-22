

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemEnvironmentOper struct {
    
    Oper SystemEnvironmentOperOper `json:"oper"`

}
type DataSystemEnvironmentOper struct {
    DtSystemEnvironmentOper SystemEnvironmentOper `json:"environment"`
}


type SystemEnvironmentOperOper struct {
    PhysicalTemperature string `json:"physical-temperature"`
    PhysicalTemperature2 string `json:"physical-temperature2"`
    Cpu0Temperature string `json:"cpu0-temperature"`
    Cpu1Temperature string `json:"cpu1-temperature"`
    Fan1aReport string `json:"fan1a-report"`
    Fan1aValue int `json:"fan1a-value"`
    Fan1bReport string `json:"fan1b-report"`
    Fan1bValue int `json:"fan1b-value"`
    Fan2aReport string `json:"fan2a-report"`
    Fan2aValue int `json:"fan2a-value"`
    Fan2bReport string `json:"fan2b-report"`
    Fan2bValue int `json:"fan2b-value"`
    Fan3aReport string `json:"fan3a-report"`
    Fan3aValue int `json:"fan3a-value"`
    Fan3bReport string `json:"fan3b-report"`
    Fan3bValue int `json:"fan3b-value"`
    Fan4aReport string `json:"fan4a-report"`
    Fan4aValue int `json:"fan4a-value"`
    Fan4bReport string `json:"fan4b-report"`
    Fan4bValue int `json:"fan4b-value"`
    Fan5aReport string `json:"fan5a-report"`
    Fan5aValue int `json:"fan5a-value"`
    Fan5bReport string `json:"fan5b-report"`
    Fan5bValue int `json:"fan5b-value"`
    Fan6aReport string `json:"fan6a-report"`
    Fan6aValue int `json:"fan6a-value"`
    Fan6bReport string `json:"fan6b-report"`
    Fan6bValue int `json:"fan6b-value"`
    Fan7aReport string `json:"fan7a-report"`
    Fan7aValue int `json:"fan7a-value"`
    Fan7bReport string `json:"fan7b-report"`
    Fan7bValue int `json:"fan7b-value"`
    Fan8aReport string `json:"fan8a-report"`
    Fan8aValue int `json:"fan8a-value"`
    Fan8bReport string `json:"fan8b-report"`
    Fan8bValue int `json:"fan8b-value"`
    Fan9aReport string `json:"fan9a-report"`
    Fan9aValue int `json:"fan9a-value"`
    Fan9bReport string `json:"fan9b-report"`
    Fan9bValue int `json:"fan9b-value"`
    Fan10aReport string `json:"fan10a-report"`
    Fan10aValue int `json:"fan10a-value"`
    Fan10bReport string `json:"fan10b-report"`
    Fan10bValue int `json:"fan10b-value"`
    VoltageLabel1 string `json:"voltage-label-1"`
    VoltageLabel2 string `json:"voltage-label-2"`
    VoltageLabel3 string `json:"voltage-label-3"`
    VoltageLabel4 string `json:"voltage-label-4"`
    VoltageLabel5 string `json:"voltage-label-5"`
    VoltageLabel6 string `json:"voltage-label-6"`
    VoltageLabel7 string `json:"voltage-label-7"`
    VoltageLabel8 string `json:"voltage-label-8"`
    VoltageLabel9 string `json:"voltage-label-9"`
    VoltageLabel10 string `json:"voltage-label-10"`
    VoltageLabel11 string `json:"voltage-label-11"`
    VoltageLabel12 string `json:"voltage-label-12"`
    VoltageLabel13 string `json:"voltage-label-13"`
    VoltageLabel14 string `json:"voltage-label-14"`
    VoltageLabel15 string `json:"voltage-label-15"`
    VoltageLabel16 string `json:"voltage-label-16"`
    VoltageLabel17 string `json:"voltage-label-17"`
    PowerUnit1 string `json:"power-unit1"`
    PowerUnit2 string `json:"power-unit2"`
    PowerUnit3 string `json:"power-unit3"`
    PowerUnit4 string `json:"power-unit4"`
    Pu2PhysicalTemperature string `json:"PU2-physical-temperature"`
    Pu2PhysicalTemperature2 string `json:"PU2-physical-temperature2"`
    Pu2VoltageLabel7 string `json:"PU2-voltage-label-7"`
    Pu2VoltageLabel8 string `json:"PU2-voltage-label-8"`
    Pu2VoltageLabel9 string `json:"PU2-voltage-label-9"`
    Pu2VoltageLabel10 string `json:"PU2-voltage-label-10"`
}

func (p *SystemEnvironmentOper) GetId() string{
    return "1"
}

func (p *SystemEnvironmentOper) getPath() string{
    return "system/environment/oper"
}

func (p *SystemEnvironmentOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemEnvironmentOper,error) {
logger.Println("SystemEnvironmentOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemEnvironmentOper
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
