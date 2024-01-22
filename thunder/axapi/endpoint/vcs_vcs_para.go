

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VcsVcsPara struct {
	Inst struct {

    ChassisId int `json:"chassis-id"`
    ConfigInfo string `json:"config-info"`
    ConfigSeq string `json:"config-seq"`
    DeadInterval int `json:"dead-interval" dval:"10"`
    DeadIntervalMseconds int `json:"dead-interval-mseconds"`
    FailureRetryCountValue int `json:"failure-retry-count-value" dval:"2"`
    FloatingIpCfg []VcsVcsParaFloatingIpCfg `json:"floating-ip-cfg"`
    FloatingIpv6Cfg []VcsVcsParaFloatingIpv6Cfg `json:"floating-ipv6-cfg"`
    ForceWaitInterval int `json:"force-wait-interval" dval:"5"`
    Forever int `json:"forever"`
    MemoryStatInterval int `json:"memory-stat-interval" dval:"30"`
    MulticastIp string `json:"multicast-ip" dval:"224.0.1.210"`
    MulticastIpv6 string `json:"multicast-ipv6"`
    MulticastPort int `json:"multicast-port" dval:"41217"`
    Size int `json:"size"`
    SlogLevel int `json:"slog-level" dval:"7"`
    SlogMethod int `json:"slog-method" dval:"1"`
    Speed_limit int `json:"speed_limit"`
    SslEnable int `json:"ssl-enable"`
    TcpChannelMonitor int `json:"tcp-channel-monitor"`
    TimeInterval int `json:"time-interval" dval:"3"`
    TimeIntervalMseconds int `json:"time-interval-mseconds"`
    TransmitFragmentSize int `json:"transmit-fragment-size" dval:"6000"`
    Uuid string `json:"uuid"`

	} `json:"vcs-para"`
}


type VcsVcsParaFloatingIpCfg struct {
    FloatingIp string `json:"floating-ip"`
    FloatingIpMask string `json:"floating-ip-mask"`
}


type VcsVcsParaFloatingIpv6Cfg struct {
    FloatingIpv6 string `json:"floating-ipv6"`
}

func (p *VcsVcsPara) GetId() string{
    return "1"
}

func (p *VcsVcsPara) getPath() string{
    return "vcs/vcs-para"
}

func (p *VcsVcsPara) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VcsVcsPara::Post")
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

func (p *VcsVcsPara) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VcsVcsPara::Get")
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
func (p *VcsVcsPara) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VcsVcsPara::Put")
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

func (p *VcsVcsPara) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VcsVcsPara::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
