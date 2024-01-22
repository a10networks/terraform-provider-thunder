

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwAlg struct {
	Inst struct {

    Dns FwAlgDns355 `json:"dns"`
    Esp FwAlgEsp356 `json:"esp"`
    Ftp FwAlgFtp358 `json:"ftp"`
    Icmp FwAlgIcmp360 `json:"icmp"`
    Pptp FwAlgPptp361 `json:"pptp"`
    Rtsp FwAlgRtsp363 `json:"rtsp"`
    Sip FwAlgSip365 `json:"sip"`
    Tftp FwAlgTftp367 `json:"tftp"`
    Uuid string `json:"uuid"`

	} `json:"alg"`
}


type FwAlgDns355 struct {
    DefaultPortDisable string `json:"default-port-disable"`
    Uuid string `json:"uuid"`
}


type FwAlgEsp356 struct {
    DefaultPortDisable string `json:"default-port-disable"`
    Uuid string `json:"uuid"`
    SamplingEnable []FwAlgEspSamplingEnable357 `json:"sampling-enable"`
}


type FwAlgEspSamplingEnable357 struct {
    Counters1 string `json:"counters1"`
}


type FwAlgFtp358 struct {
    DefaultPortDisable string `json:"default-port-disable"`
    Uuid string `json:"uuid"`
    SamplingEnable []FwAlgFtpSamplingEnable359 `json:"sampling-enable"`
}


type FwAlgFtpSamplingEnable359 struct {
    Counters1 string `json:"counters1"`
}


type FwAlgIcmp360 struct {
    Disable string `json:"disable"`
    Uuid string `json:"uuid"`
}


type FwAlgPptp361 struct {
    DefaultPortDisable string `json:"default-port-disable"`
    Uuid string `json:"uuid"`
    SamplingEnable []FwAlgPptpSamplingEnable362 `json:"sampling-enable"`
}


type FwAlgPptpSamplingEnable362 struct {
    Counters1 string `json:"counters1"`
}


type FwAlgRtsp363 struct {
    DefaultPortDisable string `json:"default-port-disable"`
    Uuid string `json:"uuid"`
    SamplingEnable []FwAlgRtspSamplingEnable364 `json:"sampling-enable"`
}


type FwAlgRtspSamplingEnable364 struct {
    Counters1 string `json:"counters1"`
}


type FwAlgSip365 struct {
    DefaultPortDisable string `json:"default-port-disable"`
    Uuid string `json:"uuid"`
    SamplingEnable []FwAlgSipSamplingEnable366 `json:"sampling-enable"`
}


type FwAlgSipSamplingEnable366 struct {
    Counters1 string `json:"counters1"`
}


type FwAlgTftp367 struct {
    DefaultPortDisable string `json:"default-port-disable"`
    Uuid string `json:"uuid"`
    SamplingEnable []FwAlgTftpSamplingEnable368 `json:"sampling-enable"`
}


type FwAlgTftpSamplingEnable368 struct {
    Counters1 string `json:"counters1"`
}

func (p *FwAlg) GetId() string{
    return "1"
}

func (p *FwAlg) getPath() string{
    return "fw/alg"
}

func (p *FwAlg) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwAlg::Post")
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

func (p *FwAlg) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwAlg::Get")
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
func (p *FwAlg) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwAlg::Put")
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

func (p *FwAlg) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwAlg::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
