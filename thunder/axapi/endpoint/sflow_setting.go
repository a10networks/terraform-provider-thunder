

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SflowSetting struct {
	Inst struct {

    AppendMappingInfo int `json:"append-mapping-info"`
    CounterPollingInterval int `json:"counter-polling-interval" dval:"20"`
    DefaultCounterPollingMtu int `json:"default-counter-polling-mtu" dval:"1500"`
    LocalCollection string `json:"local-collection" dval:"enable"`
    LocalT1PollingInterval int `json:"local-t1-polling-interval"`
    LocalT2PollingInterval int `json:"local-t2-polling-interval"`
    ManagementLinkUtilization int `json:"management-link-utilization"`
    ManagementLinkUtilizationPercentage int `json:"management-link-utilization-percentage" dval:"30"`
    MaxHeader int `json:"max-header"`
    PacketSamplingRate int `json:"packet-sampling-rate" dval:"1000"`
    PortRangeEnd int `json:"port-range-end"`
    PortRangeStart int `json:"port-range-start"`
    RandomizeSourcePort string `json:"randomize-source-port" dval:"packet-sampling-only"`
    SourceIpUseMgmt int `json:"source-ip-use-mgmt"`
    Uuid string `json:"uuid"`

	} `json:"setting"`
}

func (p *SflowSetting) GetId() string{
    return "1"
}

func (p *SflowSetting) getPath() string{
    return "sflow/setting"
}

func (p *SflowSetting) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SflowSetting::Post")
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

func (p *SflowSetting) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SflowSetting::Get")
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
func (p *SflowSetting) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SflowSetting::Put")
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

func (p *SflowSetting) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SflowSetting::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
