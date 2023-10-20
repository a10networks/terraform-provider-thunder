package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
	"strconv"
)

// based on ACOS 5_2_1-P4_81
type SlbServerPort struct {
	Inst struct {
		ServerName                  string                        //outside key
		Action                      string                        `json:"action" dval:"enable"`
		AlternatePort               []SlbServerPortAlternatePort  `json:"alternate-port"`
		AuthCfg                     SlbServerPortAuthCfg          `json:"auth-cfg"`
		ConnLimit                   int                           `json:"conn-limit" dval:"64000000"`
		ConnResume                  int                           `json:"conn-resume"`
		ExtendedStats               int                           `json:"extended-stats"`
		FollowPortProtocol          string                        `json:"follow-port-protocol"`
		HealthCheck                 string                        `json:"health-check"`
		HealthCheckDisable          int                           `json:"health-check-disable"`
		HealthCheckFollowPort       int                           `json:"health-check-follow-port"`
		NoLogging                   int                           `json:"no-logging"`
		NoSsl                       int                           `json:"no-ssl"`
		PacketCaptureTemplate       string                        `json:"packet-capture-template"`
		PortNumber                  int                           `json:"port-number" dval:"-1"`
		Protocol                    string                        `json:"protocol"`
		Range                       int                           `json:"range"`
		RportHealthCheckShared      string                        `json:"rport-health-check-shared"`
		SamplingEnable              []SlbServerPortSamplingEnable `json:"sampling-enable"`
		SharedPartitionPortTemplate int                           `json:"shared-partition-port-template"`
		SharedRportHealthCheck      int                           `json:"shared-rport-health-check"`
		StatsDataAction             string                        `json:"stats-data-action" dval:"stats-data-enable"`
		SupportHttp2                int                           `json:"support-http2"`
		TemplatePort                string                        `json:"template-port" dval:"default"`
		TemplatePortShared          string                        `json:"template-port-shared"`
		TemplateServerSsl           string                        `json:"template-server-ssl"`
		UserTag                     string                        `json:"user-tag"`
		Uuid                        string                        `json:"uuid"`
		Weight                      int                           `json:"weight" dval:"1"`
	} `json:"port"`
}

type SlbServerPortAlternatePort struct {
	Alternate           int    `json:"alternate"`
	AlternateName       string `json:"alternate-name"`
	AlternateServerPort int    `json:"alternate-server-port"`
}

type SlbServerPortAuthCfg struct {
	ServicePrincipalName string `json:"service-principal-name"`
}

type SlbServerPortSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

func (p *SlbServerPort) GetId() string {
	return strconv.Itoa(p.Inst.PortNumber) + "+" + p.Inst.Protocol
}

func (p *SlbServerPort) getPath() string {
	return "slb/server/" + url.QueryEscape(p.Inst.ServerName) + "/port"
}

func (p *SlbServerPort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbServerPort::Post")
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

func (p *SlbServerPort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbServerPort::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
	if err == nil {
		err = json.Unmarshal(axResp, &p)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}

func (p *SlbServerPort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbServerPort::Put")
	headers := axapi.GenRequestHeader(authToken)
	payloadBytes, err := axapi.SerializeToJson(p)
	if err != nil {
		logger.Println("Failed to serialize struct as json", err)
		return err
	}
	logger.Println("payload: " + string(payloadBytes))
	_, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
	return err
}

func (p *SlbServerPort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbServerPort::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
