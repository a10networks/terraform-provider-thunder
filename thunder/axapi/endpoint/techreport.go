

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Techreport struct {
	Inst struct {

    Disable int `json:"disable"`
    EnableFullHistory int `json:"enable-full-history"`
    Interval TechreportInterval1897 `json:"interval"`
    MaxLogfileSize TechreportMaxLogfileSize1898 `json:"max-logfile-size"`
    MaxPartitions TechreportMaxPartitions1899 `json:"max-partitions"`
    PriorityPartitionList []TechreportPriorityPartitionList `json:"priority-partition-list"`
    Uuid string `json:"uuid"`

	} `json:"techreport"`
}


type TechreportInterval1897 struct {
    Value int `json:"value" dval:"15"`
    Uuid string `json:"uuid"`
}


type TechreportMaxLogfileSize1898 struct {
    Value int `json:"value" dval:"1"`
    Uuid string `json:"uuid"`
}


type TechreportMaxPartitions1899 struct {
    Value int `json:"value" dval:"30"`
    Uuid string `json:"uuid"`
}


type TechreportPriorityPartitionList struct {
    PartName string `json:"part-name"`
    Uuid string `json:"uuid"`
}

func (p *Techreport) GetId() string{
    return "1"
}

func (p *Techreport) getPath() string{
    return "techreport"
}

func (p *Techreport) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Techreport::Post")
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

func (p *Techreport) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Techreport::Get")
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
func (p *Techreport) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Techreport::Put")
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

func (p *Techreport) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Techreport::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
