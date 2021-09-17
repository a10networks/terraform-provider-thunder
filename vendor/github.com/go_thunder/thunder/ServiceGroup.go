package go_thunder

import (
	"bytes"
	"fmt"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Priorities struct {
	Priority       int    `json:"priority,omitempty"`
	PriorityAction string `json:"priority-action,omitempty"`
}

type SamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

type MemberList struct {
	MemberPriority         int              `json:"member-priority,omitempty"`
	UUID                   string           `json:"uuid,omitempty"`
	FqdnName               string           `json:"fqdn-name,omitempty"`
	ResolveAs              string           `json:"resolve-as,omitempty"`
	Counters1              []SamplingEnable `json:"sampling-enable,omitempty"`
	MemberTemplate         string           `json:"member-template,omitempty"`
	Name                   string           `json:"name,omitempty"`
	Host                   string           `json:"host,omitempty"`
	UserTag                string           `json:"user-tag,omitempty"`
	MemberState            string           `json:"member-state,omitempty"`
	ServerIpv6Addr         string           `json:"server-ipv6-addr,omitempty"`
	Port                   int              `json:"port,omitempty,omitempty"`
	MemberStatsDataDisable int              `json:"member-stats-data-disable,omitempty"`
}

type Reset struct {
	AutoSwitch int `json:"auto-switch,omitempty"`
}

type ServiceGroupInstance struct {
	ConnRate                         int          `json:"conn-rate,omitempty"`
	ResetOnServerSelectionFail       int          `json:"reset-on-server-selection-fail,omitempty"`
	HealthCheckDisable               int          `json:"health-check-disable,omitempty"`
	Protocol                         string       `json:"protocol,omitempty"`
	TrafficReplicationMirrorIPRepl   int          `json:"traffic-replication-mirror-ip-repl,omitempty"`
	ResetPriorityAffinity            int          `json:"reset-priority-affinity,omitempty"`
	Priority                         []Priorities `json:"priorities,omitempty"`
	MinActiveMember                  int          `json:"min-active-member,omitempty"`
	Host                             []MemberList `json:"member-list,omitempty"`
	StatsDataAction                  string       `json:"stats-data-action,omitempty"`
	TrafficReplicationMirrorDaRepl   int          `json:"traffic-replication-mirror-da-repl,omitempty"`
	TemplatePolicyShared             string       `json:"template-policy-shared,omitempty"`
	RptExtServer                     int          `json:"rpt-ext-server,omitempty"`
	TemplatePort                     string       `json:"template-port,omitempty"`
	ConnRateGracePeriod              int          `json:"conn-rate-grace-period,omitempty"`
	L4SessionUsageDuration           int          `json:"l4-session-usage-duration,omitempty"`
	UUID                             string       `json:"uuid,omitempty"`
	BackupServerEventLog             int          `json:"backup-server-event-log,omitempty"`
	LcMethod                         string       `json:"lc-method,omitempty"`
	PseudoRoundRobin                 int          `json:"pseudo-round-robin,omitempty"`
	SharedPartitionPolicyTemplate    int          `json:"shared-partition-policy-template,omitempty"`
	L4SessionUsageRevertRate         int          `json:"l4-session-usage-revert-rate,omitempty"`
	SharedPartitionSvcgrpHealthCheck int          `json:"shared-partition-svcgrp-health-check,omitempty"`
	TemplateServer                   string       `json:"template-server,omitempty"`
	SvcgrpHealthCheckShared          string       `json:"svcgrp-health-check-shared,omitempty"`
	TrafficReplicationMirror         int          `json:"traffic-replication-mirror,omitempty"`
	L4SessionRevertDuration          int          `json:"l4-session-revert-duration,omitempty"`
	TrafficReplicationMirrorSaDaRepl int          `json:"traffic-replication-mirror-sa-da-repl,omitempty"`
	LbMethod                         string       `json:"lb-method,omitempty"`
	StatelessAutoSwitch              int          `json:"stateless-auto-switch,omitempty"`
	MinActiveMemberAction            string       `json:"min-active-member-action,omitempty"`
	L4SessionUsage                   int          `json:"l4-session-usage,omitempty"`
	ExtendedStats                    int          `json:"extended-stats,omitempty"`
	ConnRateRevertDuration           int          `json:"conn-rate-revert-duration,omitempty"`
	StrictSelect                     int          `json:"strict-select,omitempty"`
	Name                             string       `json:"name,omitempty"`
	//		AutoSwitch                       Reset 	`json:"reset,omitempty"`
	TrafficReplicationMirrorSaRepl int              `json:"traffic-replication-mirror-sa-repl,omitempty"`
	ReportDelay                    int              `json:"report-delay,omitempty"`
	ConnRateLog                    int              `json:"conn-rate-log,omitempty"`
	L4SessionUsageLog              int              `json:"l4-session-usage-log,omitempty"`
	ConnRateDuration               int              `json:"conn-rate-duration,omitempty"`
	StatelessLbMethod              string           `json:"stateless-lb-method,omitempty"`
	TemplatePolicy                 string           `json:"template-policy,omitempty"`
	StatelessLbMethod2             string           `json:"stateless-lb-method2,omitempty"`
	UserTag                        string           `json:"user-tag,omitempty"`
	SampleRspTime                  int              `json:"sample-rsp-time,omitempty"`
	Counters1                      []SamplingEnable `json:"sampling-enable,omitempty"`
	TopFastest                     int              `json:"top-fastest,omitempty"`
	ConnRevertRate                 int              `json:"conn-revert-rate,omitempty"`
	L4SessionUsageGracePeriod      int              `json:"l4-session-usage-grace-period,omitempty"`
	PriorityAffinity               int              `json:"priority-affinity,omitempty"`
	TopSlowest                     int              `json:"top-slowest,omitempty"`
	HealthCheck                    string           `json:"health-check,omitempty"`
}

type ServiceGroup struct {
	Name ServiceGroupInstance `json:"service-group,omitempty"`
}

func GetSG(id string, name string, host string) (*ServiceGroup, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/service-group/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ServiceGroup
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)

			err := check_api_status("GetSG", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}
}

func PostSG(id string, sg ServiceGroup, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(sg)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/service-group", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ServiceGroup
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))

			err := check_api_status("PostSG", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func PutSG(id string, name string, sg ServiceGroup, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	logger.Println("[INFO] received V  --" + sg.Name.Name)

	payloadBytes, err := json.Marshal(sg)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/service-group/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ServiceGroup
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))

			err := check_api_status("PutSG", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func DeleteSG(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/service-group/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ServiceGroup
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
		}
	}
	return nil
}
