

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Monitor struct {
	Inst struct {

    BufferDrop int `json:"buffer-drop" dval:"4000"`
    BufferUsage int `json:"buffer-usage"`
    ConnType0 int `json:"conn-type0" dval:"32767"`
    ConnType1 int `json:"conn-type1" dval:"32767"`
    ConnType2 int `json:"conn-type2" dval:"32767"`
    ConnType3 int `json:"conn-type3" dval:"32767"`
    ConnType4 int `json:"conn-type4" dval:"32767"`
    CtrlCpu int `json:"ctrl-cpu" dval:"90"`
    DataCpu int `json:"data-cpu" dval:"90"`
    Disk int `json:"disk" dval:"85"`
    Memory int `json:"memory" dval:"95"`
    SmpType0 int `json:"smp-type0" dval:"32767"`
    SmpType1 int `json:"smp-type1" dval:"32767"`
    SmpType2 int `json:"smp-type2" dval:"32767"`
    SmpType3 int `json:"smp-type3" dval:"32767"`
    SmpType4 int `json:"smp-type4" dval:"32767"`
    Uuid string `json:"uuid"`
    WarnTemp int `json:"warn-temp"`

	} `json:"monitor"`
}

func (p *Monitor) GetId() string{
    return "1"
}

func (p *Monitor) getPath() string{
    return "monitor"
}

func (p *Monitor) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Monitor::Post")
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

func (p *Monitor) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Monitor::Get")
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
func (p *Monitor) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Monitor::Put")
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

func (p *Monitor) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Monitor::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
