package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbTemplateDns struct {
	SlbTemplateDNSInstanceName SlbTemplateDNSInstance `json:"dns,omitempty"`
}

type SlbTemplateDNSInstance struct {
	SlbTemplateDNSInstanceAddPaddingToClient                string                                       `json:"add-padding-to-client,omitempty"`
	SlbTemplateDNSInstanceCacheRecordServingPolicy          string                                       `json:"cache-record-serving-policy,omitempty"`
	SlbTemplateDNSInstanceClassListName                     SlbTemplateDNSInstanceClassList              `json:"class-list,omitempty"`
	SlbTemplateDNSInstanceDNS64Enable                       SlbTemplateDNSInstanceDNS64                  `json:"dns64,omitempty"`
	SlbTemplateDNSInstanceDNSLogging                        string                                       `json:"dns-logging,omitempty"`
	SlbTemplateDNSInstanceDefaultPolicy                     string                                       `json:"default-policy,omitempty"`
	SlbTemplateDNSInstanceDisableDNSTemplate                int                                          `json:"disable-dns-template,omitempty"`
	SlbTemplateDNSInstanceDisableRpzAttachSoa               int                                          `json:"disable-rpz-attach-soa,omitempty"`
	SlbTemplateDNSInstanceDnssecServiceGroup                string                                       `json:"dnssec-service-group,omitempty"`
	SlbTemplateDNSInstanceDrop                              int                                          `json:"drop,omitempty"`
	SlbTemplateDNSInstanceEnableCacheSharing                int                                          `json:"enable-cache-sharing,omitempty"`
	SlbTemplateDNSInstanceForward                           string                                       `json:"forward,omitempty"`
	SlbTemplateDNSInstanceInsertIpv4                        int                                          `json:"insert-ipv4,omitempty"`
	SlbTemplateDNSInstanceInsertIpv6                        int                                          `json:"insert-ipv6,omitempty"`
	SlbTemplateDNSInstanceLocalDNSResolutionHostListCfg     SlbTemplateDNSInstanceLocalDNSResolution     `json:"local-dns-resolution,omitempty"`
	SlbTemplateDNSInstanceMaxCacheEntrySize                 int                                          `json:"max-cache-entry-size,omitempty"`
	SlbTemplateDNSInstanceMaxCacheSize                      int                                          `json:"max-cache-size,omitempty"`
	SlbTemplateDNSInstanceMaxQueryLength                    int                                          `json:"max-query-length,omitempty"`
	SlbTemplateDNSInstanceName                              string                                       `json:"name,omitempty"`
	SlbTemplateDNSInstancePeriod                            int                                          `json:"period,omitempty"`
	SlbTemplateDNSInstanceQueryClassFilterQueryClassAction  SlbTemplateDNSInstanceQueryClassFilter       `json:"query-class-filter,omitempty"`
	SlbTemplateDNSInstanceQueryIDSwitch                     int                                          `json:"query-id-switch,omitempty"`
	SlbTemplateDNSInstanceQueryTypeFilterQueryTypeAction    SlbTemplateDNSInstanceQueryTypeFilter        `json:"query-type-filter,omitempty"`
	SlbTemplateDNSInstanceRecursiveDNSResolutionHostListCfg SlbTemplateDNSInstanceRecursiveDNSResolution `json:"recursive-dns-resolution,omitempty"`
	SlbTemplateDNSInstanceRedirectToTCPPort                 int                                          `json:"redirect-to-tcp-port,omitempty"`
	SlbTemplateDNSInstanceRemoveAaFlag                      int                                          `json:"remove-aa-flag,omitempty"`
	SlbTemplateDNSInstanceRemoveEdnsCsubnetToServer         int                                          `json:"remove-edns-csubnet-to-server,omitempty"`
	SlbTemplateDNSInstanceRemovePaddingToServer             int                                          `json:"remove-padding-to-server,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingResponseRate  SlbTemplateDNSInstanceResponseRateLimiting   `json:"response-rate-limiting,omitempty"`
	SlbTemplateDNSInstanceRpzListSeqID                      []SlbTemplateDNSInstanceRpzList              `json:"rpz-list,omitempty"`
	SlbTemplateDNSInstanceUDPRetransmitRetryInterval        SlbTemplateDNSInstanceUDPRetransmit          `json:"udp-retransmit,omitempty"`
	SlbTemplateDNSInstanceUUID                              string                                       `json:"uuid,omitempty"`
	SlbTemplateDNSInstanceUserTag                           string                                       `json:"user-tag,omitempty"`
}

type SlbTemplateDNSInstanceClassList struct {
	SlbTemplateDNSInstanceClassListLidListLidnum []SlbTemplateDNSInstanceClassListLidList `json:"lid-list,omitempty"`
	SlbTemplateDNSInstanceClassListName          string                                   `json:"name,omitempty"`
	SlbTemplateDNSInstanceClassListUUID          string                                   `json:"uuid,omitempty"`
}

type SlbTemplateDNSInstanceDNS64 struct {
	SlbTemplateDNSInstanceDNS64Cache                 int    `json:"cache,omitempty"`
	SlbTemplateDNSInstanceDNS64ChangeQuery           int    `json:"change-query,omitempty"`
	SlbTemplateDNSInstanceDNS64Enable                int    `json:"enable,omitempty"`
	SlbTemplateDNSInstanceDNS64ParallelQuery         int    `json:"parallel-query,omitempty"`
	SlbTemplateDNSInstanceDNS64Retry                 int    `json:"retry,omitempty"`
	SlbTemplateDNSInstanceDNS64SingleResponseDisable int    `json:"single-response-disable,omitempty"`
	SlbTemplateDNSInstanceDNS64Timeout               int    `json:"timeout,omitempty"`
	SlbTemplateDNSInstanceDNS64UUID                  string `json:"uuid,omitempty"`
}

type SlbTemplateDNSInstanceLocalDNSResolution struct {
	SlbTemplateDNSInstanceLocalDNSResolutionHostListCfgHostnames          []SlbTemplateDNSInstanceLocalDNSResolutionHostListCfg      `json:"host-list-cfg,omitempty"`
	SlbTemplateDNSInstanceLocalDNSResolutionLocalResolverCfgLocalResolver []SlbTemplateDNSInstanceLocalDNSResolutionLocalResolverCfg `json:"local-resolver-cfg,omitempty"`
	SlbTemplateDNSInstanceLocalDNSResolutionUUID                          string                                                     `json:"uuid,omitempty"`
}

type SlbTemplateDNSInstanceQueryClassFilter struct {
	SlbTemplateDNSInstanceQueryClassFilterQueryClassStrQueryClass []SlbTemplateDNSInstanceQueryClassFilterQueryClass `json:"query-class,omitempty"`
	SlbTemplateDNSInstanceQueryClassFilterQueryClassAction        string                                             `json:"query-class-action,omitempty"`
	SlbTemplateDNSInstanceQueryClassFilterUUID                    string                                             `json:"uuid,omitempty"`
}

type SlbTemplateDNSInstanceQueryTypeFilter struct {
	SlbTemplateDNSInstanceQueryTypeFilterQueryTypeStrQueryType []SlbTemplateDNSInstanceQueryTypeFilterQueryType `json:"query-type,omitempty"`
	SlbTemplateDNSInstanceQueryTypeFilterQueryTypeAction       string                                           `json:"query-type-action,omitempty"`
	SlbTemplateDNSInstanceQueryTypeFilterUUID                  string                                           `json:"uuid,omitempty"`
}

type SlbTemplateDNSInstanceRecursiveDNSResolution struct {
	SlbTemplateDNSInstanceRecursiveDNSResolutionCsubnetRetry         int                                                       `json:"csubnet-retry,omitempty"`
	SlbTemplateDNSInstanceRecursiveDNSResolutionDefaultRecursive     int                                                       `json:"default-recursive,omitempty"`
	SlbTemplateDNSInstanceRecursiveDNSResolutionHostListCfgHostnames []SlbTemplateDNSInstanceRecursiveDNSResolutionHostListCfg `json:"host-list-cfg,omitempty"`
	SlbTemplateDNSInstanceRecursiveDNSResolutionIpv4NatPool          string                                                    `json:"ipv4-nat-pool,omitempty"`
	SlbTemplateDNSInstanceRecursiveDNSResolutionIpv6NatPool          string                                                    `json:"ipv6-nat-pool,omitempty"`
	SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryType SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrder   `json:"lookup-order,omitempty"`
	SlbTemplateDNSInstanceRecursiveDNSResolutionMaxTrials            int                                                       `json:"max-trials,omitempty"`
	SlbTemplateDNSInstanceRecursiveDNSResolutionMinimalResponse      int                                                       `json:"minimal-response,omitempty"`
	SlbTemplateDNSInstanceRecursiveDNSResolutionRetriesPerLevel      int                                                       `json:"retries-per-level,omitempty"`
	SlbTemplateDNSInstanceRecursiveDNSResolutionUUID                 string                                                    `json:"uuid,omitempty"`
}

type SlbTemplateDNSInstanceResponseRateLimiting struct {
	SlbTemplateDNSInstanceResponseRateLimitingAction               string                                                       `json:"action,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingEnableLog            int                                                          `json:"enable-log,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingFilterResponseRate   int                                                          `json:"filter-response-rate,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingResponseRate         int                                                          `json:"response-rate,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListName []SlbTemplateDNSInstanceResponseRateLimitingRrlClassListList `json:"rrl-class-list-list,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingSlipRate             int                                                          `json:"slip-rate,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingUUID                 string                                                       `json:"uuid,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingWindow               int                                                          `json:"window,omitempty"`
}

type SlbTemplateDNSInstanceRpzList struct {
	SlbTemplateDNSInstanceRpzListLoggingEnable SlbTemplateDNSInstanceRpzListLogging `json:"logging,omitempty"`
	SlbTemplateDNSInstanceRpzListName          string                               `json:"name,omitempty"`
	SlbTemplateDNSInstanceRpzListSeqID         int                                  `json:"seq-id,omitempty"`
	SlbTemplateDNSInstanceRpzListUUID          string                               `json:"uuid,omitempty"`
	SlbTemplateDNSInstanceRpzListUserTag       string                               `json:"user-tag,omitempty"`
}

type SlbTemplateDNSInstanceUDPRetransmit struct {
	SlbTemplateDNSInstanceUDPRetransmitMaxTrials     int    `json:"max-trials,omitempty"`
	SlbTemplateDNSInstanceUDPRetransmitRetryInterval int    `json:"retry-interval,omitempty"`
	SlbTemplateDNSInstanceUDPRetransmitUUID          string `json:"uuid,omitempty"`
}

type SlbTemplateDNSInstanceClassListLidList struct {
	SlbTemplateDNSInstanceClassListLidListLidAction       string `json:"lid-action,omitempty"`
	SlbTemplateDNSInstanceClassListLidListLidEnableLog    int    `json:"lid-enable-log,omitempty"`
	SlbTemplateDNSInstanceClassListLidListLidResponseRate int    `json:"lid-response-rate,omitempty"`
	SlbTemplateDNSInstanceClassListLidListLidSlipRate     int    `json:"lid-slip-rate,omitempty"`
	SlbTemplateDNSInstanceClassListLidListLidWindow       int    `json:"lid-window,omitempty"`
	SlbTemplateDNSInstanceClassListLidListLidnum          int    `json:"lidnum,omitempty"`
	SlbTemplateDNSInstanceClassListLidListUUID            string `json:"uuid,omitempty"`
	SlbTemplateDNSInstanceClassListLidListUserTag         string `json:"user-tag,omitempty"`
}

type SlbTemplateDNSInstanceLocalDNSResolutionHostListCfg struct {
	SlbTemplateDNSInstanceLocalDNSResolutionHostListCfgHostnames string `json:"hostnames,omitempty"`
}

type SlbTemplateDNSInstanceLocalDNSResolutionLocalResolverCfg struct {
	SlbTemplateDNSInstanceLocalDNSResolutionLocalResolverCfgLocalResolver string `json:"local-resolver,omitempty"`
}

type SlbTemplateDNSInstanceQueryClassFilterQueryClass struct {
	SlbTemplateDNSInstanceQueryClassFilterQueryClassNumQueryClass int    `json:"num-query-class,omitempty"`
	SlbTemplateDNSInstanceQueryClassFilterQueryClassStrQueryClass string `json:"str-query-class,omitempty"`
}

type SlbTemplateDNSInstanceQueryTypeFilterQueryType struct {
	SlbTemplateDNSInstanceQueryTypeFilterQueryTypeNumQueryType int    `json:"num-query-type,omitempty"`
	SlbTemplateDNSInstanceQueryTypeFilterQueryTypeOrder        string `json:"order,omitempty"`
	SlbTemplateDNSInstanceQueryTypeFilterQueryTypeStrQueryType string `json:"str-query-type,omitempty"`
}

type SlbTemplateDNSInstanceRecursiveDNSResolutionHostListCfg struct {
	SlbTemplateDNSInstanceRecursiveDNSResolutionHostListCfgHostnames string `json:"hostnames,omitempty"`
}

type SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrder struct {
	SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryTypeStrQueryType []SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryType `json:"query-type,omitempty"`
	SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderUUID                  string                                                             `json:"uuid,omitempty"`
}

type SlbTemplateDNSInstanceResponseRateLimitingRrlClassListList struct {
	SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListLidnum []SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidList `json:"lid-list,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListName          string                                                              `json:"name,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListUUID          string                                                              `json:"uuid,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListUserTag       string                                                              `json:"user-tag,omitempty"`
}

type SlbTemplateDNSInstanceRpzListLogging struct {
	SlbTemplateDNSInstanceRpzListLoggingEnable                int                                             `json:"enable,omitempty"`
	SlbTemplateDNSInstanceRpzListLoggingRpzActionStrRpzAction []SlbTemplateDNSInstanceRpzListLoggingRpzAction `json:"rpz-action,omitempty"`
	SlbTemplateDNSInstanceRpzListLoggingUUID                  string                                          `json:"uuid,omitempty"`
}

type SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryType struct {
	SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryTypeNumQueryType int    `json:"num-query-type,omitempty"`
	SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryTypeOrder        string `json:"order,omitempty"`
	SlbTemplateDNSInstanceRecursiveDNSResolutionLookupOrderQueryTypeStrQueryType string `json:"str-query-type,omitempty"`
}

type SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidList struct {
	SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListLidAction       string `json:"lid-action,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListLidEnableLog    int    `json:"lid-enable-log,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListLidResponseRate int    `json:"lid-response-rate,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListLidSlipRate     int    `json:"lid-slip-rate,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListLidWindow       int    `json:"lid-window,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListLidnum          int    `json:"lidnum,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListUUID            string `json:"uuid,omitempty"`
	SlbTemplateDNSInstanceResponseRateLimitingRrlClassListListLidListUserTag         string `json:"user-tag,omitempty"`
}

type SlbTemplateDNSInstanceRpzListLoggingRpzAction struct {
	SlbTemplateDNSInstanceRpzListLoggingRpzActionStrRpzAction string `json:"str-rpz-action,omitempty"`
}

func PostSlbTemplateDns(id string, inst SlbTemplateDns, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateDns")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/dns", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateDns
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateDns", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateDns(id string, name1 string, host string) (*SlbTemplateDns, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateDns")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/dns/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateDns
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateDns", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateDns(id string, name1 string, inst SlbTemplateDns, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateDns")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/dns/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateDns
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateDns", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateDns(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateDns")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/dns/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateDns
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteSlbTemplateDns", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
