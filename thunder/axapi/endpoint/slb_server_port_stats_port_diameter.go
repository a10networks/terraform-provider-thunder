

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbServerPortStats struct {
	Inst struct {

    PortNumber int `json:"port-number"`
    Protocol string `json:"protocol"`
    Stats SlbServerPortStatsStats `json:"stats"`
    Name string 

	} `json:"stats"`
}


type SlbServerPortStatsStats struct {
    PortDiameter SlbServerPortStatsStatsPortDiameter `json:"port-diameter"`
}


type SlbServerPortStatsStatsPortDiameter struct {
    Num int `json:"num"`
    Curr int `json:"curr"`
    Total int `json:"total"`
    Svrsel_fail int `json:"svrsel_fail"`
    No_route int `json:"no_route"`
    Snat_fail int `json:"snat_fail"`
    Client_fail int `json:"client_fail"`
    Server_fail int `json:"server_fail"`
    No_sess int `json:"no_sess"`
    User_session int `json:"user_session"`
    Acr_out int `json:"acr_out"`
    Acr_in int `json:"acr_in"`
    Aca_out int `json:"aca_out"`
    Aca_in int `json:"aca_in"`
    Cea_out int `json:"cea_out"`
    Cea_in int `json:"cea_in"`
    Cer_out int `json:"cer_out"`
    Cer_in int `json:"cer_in"`
    Dwr_out int `json:"dwr_out"`
    Dwr_in int `json:"dwr_in"`
    Dwa_out int `json:"dwa_out"`
    Dwa_in int `json:"dwa_in"`
    Str_out int `json:"str_out"`
    Str_in int `json:"str_in"`
    Sta_out int `json:"sta_out"`
    Sta_in int `json:"sta_in"`
    Asr_out int `json:"asr_out"`
    Asr_in int `json:"asr_in"`
    Asa_out int `json:"asa_out"`
    Asa_in int `json:"asa_in"`
    Other_out int `json:"other_out"`
    Other_in int `json:"other_in"`
    Cca_out int `json:"cca_out"`
    Cca_in int `json:"cca_in"`
    Ccr_out int `json:"ccr_out"`
    Ccr_in int `json:"ccr_in"`
    Dpr_out int `json:"dpr_out"`
    Dpr_in int `json:"dpr_in"`
    Dpa_out int `json:"dpa_out"`
    Dpa_in int `json:"dpa_in"`
}

func (p *SlbServerPortStats) GetId() string{
    return "1"
}

func (p *SlbServerPortStats) getPath() string{
    return "slb/server/" +p.Inst.Name + "/port/" +strconv.Itoa(p.Inst.PortNumber)+"+"+p.Inst.Protocol+"/stats?port-diameter=true"
}

func (p *SlbServerPortStats) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServerPortStats::Post")
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

func (p *SlbServerPortStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServerPortStats::Get")
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
func (p *SlbServerPortStats) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServerPortStats::Put")
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

func (p *SlbServerPortStats) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbServerPortStats::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
