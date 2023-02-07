package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"net/url"
	"util"
)

type ServiceGroup struct {
	ServiceGroupInstanceName ServiceGroupInstance `json:"service-group,omitempty"`
}

type ServiceGroupInstance struct {
	ServiceGroupInstanceBackupServerEventLog             int                                  `json:"backup-server-event-log,omitempty"`
	ServiceGroupInstanceConnRate                         int                                  `json:"conn-rate,omitempty"`
	ServiceGroupInstanceConnRateDuration                 int                                  `json:"conn-rate-duration,omitempty"`
	ServiceGroupInstanceConnRateGracePeriod              int                                  `json:"conn-rate-grace-period,omitempty"`
	ServiceGroupInstanceConnRateLog                      int                                  `json:"conn-rate-log,omitempty"`
	ServiceGroupInstanceConnRateRevertDuration           int                                  `json:"conn-rate-revert-duration,omitempty"`
	ServiceGroupInstanceConnRevertRate                   int                                  `json:"conn-revert-rate,omitempty"`
	ServiceGroupInstanceExtendedStats                    int                                  `json:"extended-stats,omitempty"`
	ServiceGroupInstanceHealthCheck                      string                               `json:"health-check,omitempty"`
	ServiceGroupInstanceHealthCheckDisable               int                                  `json:"health-check-disable,omitempty"`
	ServiceGroupInstanceL4SessionRevertDuration          int                                  `json:"l4-session-revert-duration,omitempty"`
	ServiceGroupInstanceL4SessionUsage                   int                                  `json:"l4-session-usage,omitempty"`
	ServiceGroupInstanceL4SessionUsageDuration           int                                  `json:"l4-session-usage-duration,omitempty"`
	ServiceGroupInstanceL4SessionUsageGracePeriod        int                                  `json:"l4-session-usage-grace-period,omitempty"`
	ServiceGroupInstanceL4SessionUsageLog                int                                  `json:"l4-session-usage-log,omitempty"`
	ServiceGroupInstanceL4SessionUsageRevertRate         int                                  `json:"l4-session-usage-revert-rate,omitempty"`
	ServiceGroupInstanceLbMethod                         string                               `json:"lb-method,omitempty"`
	ServiceGroupInstanceLcMethod                         string                               `json:"lc-method,omitempty"`
	ServiceGroupInstanceLclbMethod                       string                               `json:"lclb-method,omitempty"`
	ServiceGroupInstanceLinkProbeTemplate                string                               `json:"link-probe-template,omitempty"`
	ServiceGroupInstanceLlbMethod                        string                               `json:"llb-method,omitempty"`
	ServiceGroupInstanceMemberListName                   []ServiceGroupInstanceMemberList     `json:"member-list,omitempty"`
	ServiceGroupInstanceMinActiveMember                  int                                  `json:"min-active-member,omitempty"`
	ServiceGroupInstanceMinActiveMemberAction            string                               `json:"min-active-member-action,omitempty"`
	ServiceGroupInstanceName                             string                               `json:"name,omitempty"`
	ServiceGroupInstancePersistScoring                   string                               `json:"persist-scoring,omitempty"`
	ServiceGroupInstancePrioritiesPriority               []ServiceGroupInstancePriorities     `json:"priorities,omitempty"`
	ServiceGroupInstancePriorityAffinity                 int                                  `json:"priority-affinity,omitempty"`
	ServiceGroupInstanceProtocol                         string                               `json:"protocol,omitempty"`
	ServiceGroupInstancePseudoRoundRobin                 int                                  `json:"pseudo-round-robin,omitempty"`
	ServiceGroupInstanceReportDelay                      int                                  `json:"report-delay,omitempty"`
	ServiceGroupInstanceResetAutoSwitch                  ServiceGroupInstanceReset            `json:"reset,omitempty"`
	ServiceGroupInstanceResetOnServerSelectionFail       int                                  `json:"reset-on-server-selection-fail,omitempty"`
	ServiceGroupInstanceResetPriorityAffinity            int                                  `json:"reset-priority-affinity,omitempty"`
	ServiceGroupInstanceRptExtServer                     int                                  `json:"rpt-ext-server,omitempty"`
	ServiceGroupInstanceSampleRspTime                    int                                  `json:"sample-rsp-time,omitempty"`
	ServiceGroupInstanceSamplingEnableCounters1          []ServiceGroupInstanceSamplingEnable `json:"sampling-enable,omitempty"`
	ServiceGroupInstanceSharedPartitionPolicyTemplate    int                                  `json:"shared-partition-policy-template,omitempty"`
	ServiceGroupInstanceSharedPartitionSvcgrpHealthCheck int                                  `json:"shared-partition-svcgrp-health-check,omitempty"`
	ServiceGroupInstanceStatelessAutoSwitch              int                                  `json:"stateless-auto-switch,omitempty"`
	ServiceGroupInstanceStatelessLbMethod                string                               `json:"stateless-lb-method,omitempty"`
	ServiceGroupInstanceStatelessLbMethod2               string                               `json:"stateless-lb-method2,omitempty"`
	ServiceGroupInstanceStatsDataAction                  string                               `json:"stats-data-action,omitempty"`
	ServiceGroupInstanceStrictSelect                     int                                  `json:"strict-select,omitempty"`
	ServiceGroupInstanceSvcgrpHealthCheckShared          string                               `json:"svcgrp-health-check-shared,omitempty"`
	ServiceGroupInstanceTemplatePolicy                   string                               `json:"template-policy,omitempty"`
	ServiceGroupInstanceTemplatePolicyShared             string                               `json:"template-policy-shared,omitempty"`
	ServiceGroupInstanceTemplatePort                     string                               `json:"template-port,omitempty"`
	ServiceGroupInstanceTemplateServer                   string                               `json:"template-server,omitempty"`
	ServiceGroupInstanceTopFastest                       int                                  `json:"top-fastest,omitempty"`
	ServiceGroupInstanceTopSlowest                       int                                  `json:"top-slowest,omitempty"`
	ServiceGroupInstanceTrafficReplicationMirror         int                                  `json:"traffic-replication-mirror,omitempty"`
	ServiceGroupInstanceTrafficReplicationMirrorDaRepl   int                                  `json:"traffic-replication-mirror-da-repl,omitempty"`
	ServiceGroupInstanceTrafficReplicationMirrorIPRepl   int                                  `json:"traffic-replication-mirror-ip-repl,omitempty"`
	ServiceGroupInstanceTrafficReplicationMirrorSaDaRepl int                                  `json:"traffic-replication-mirror-sa-da-repl,omitempty"`
	ServiceGroupInstanceTrafficReplicationMirrorSaRepl   int                                  `json:"traffic-replication-mirror-sa-repl,omitempty"`
	ServiceGroupInstanceUUID                             string                               `json:"uuid,omitempty"`
	ServiceGroupInstanceUserTag                          string                               `json:"user-tag,omitempty"`
}

type ServiceGroupInstanceMemberList struct {
	ServiceGroupInstanceMemberListFqdnName                string                                         `json:"fqdn-name,omitempty"`
	ServiceGroupInstanceMemberListHost                    string                                         `json:"host,omitempty"`
	ServiceGroupInstanceMemberListMemberPriority          int                                            `json:"member-priority,omitempty"`
	ServiceGroupInstanceMemberListMemberState             string                                         `json:"member-state,omitempty"`
	ServiceGroupInstanceMemberListMemberStatsDataDisable  int                                            `json:"member-stats-data-disable,omitempty"`
	ServiceGroupInstanceMemberListMemberTemplate          string                                         `json:"member-template,omitempty"`
	ServiceGroupInstanceMemberListName                    string                                         `json:"name,omitempty"`
	ServiceGroupInstanceMemberListPort                    int                                            `json:"port,omitempty"`
	ServiceGroupInstanceMemberListResolveAs               string                                         `json:"resolve-as,omitempty"`
	ServiceGroupInstanceMemberListSamplingEnableCounters1 []ServiceGroupInstanceMemberListSamplingEnable `json:"sampling-enable,omitempty"`
	ServiceGroupInstanceMemberListServerIpv6Addr          string                                         `json:"server-ipv6-addr,omitempty"`
	ServiceGroupInstanceMemberListUUID                    string                                         `json:"uuid,omitempty"`
	ServiceGroupInstanceMemberListUserTag                 string                                         `json:"user-tag,omitempty"`
}

type ServiceGroupInstancePriorities struct {
	ServiceGroupInstancePrioritiesPriority       int    `json:"priority,omitempty"`
	ServiceGroupInstancePrioritiesPriorityAction string `json:"priority-action,omitempty"`
}

type ServiceGroupInstanceReset struct {
	ServiceGroupInstanceResetAutoSwitch int `json:"auto-switch,omitempty"`
}

type ServiceGroupInstanceSamplingEnable struct {
	ServiceGroupInstanceSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type ServiceGroupInstanceMemberListSamplingEnable struct {
	ServiceGroupInstanceMemberListSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

func PostServiceGroup(id string, inst ServiceGroup, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostServiceGroup")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/service-group", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ServiceGroup
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostServiceGroup", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetServiceGroup(id string, name1 string, host string) (*ServiceGroup, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetServiceGroup")
	nameEncode := url.QueryEscape(name1)
	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/service-group/"+nameEncode, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ServiceGroup
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetServiceGroup", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutServiceGroup(id string, name1 string, inst ServiceGroup, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutServiceGroup")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}
	nameEncode := url.QueryEscape(name1)
	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/service-group/"+nameEncode, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ServiceGroup
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutServiceGroup", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteServiceGroup(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteServiceGroup")
	nameEncode := url.QueryEscape(name1)
	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/service-group/"+nameEncode, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ServiceGroup
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteServiceGroup", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
