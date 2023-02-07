package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type RuleSet struct {
	RuleSetInstanceName RuleSetInstance `json:"rule-set,omitempty"`
}

type RuleSetInstance struct {
	RuleSetInstanceAppUUID                 RuleSetInstanceApp              `json:"app,omitempty"`
	RuleSetInstanceApplicationUUID         RuleSetInstanceApplication      `json:"application,omitempty"`
	RuleSetInstanceName                    string                          `json:"name,omitempty"`
	RuleSetInstancePacketCaptureTemplate   string                          `json:"packet-capture-template,omitempty"`
	RuleSetInstanceRemark                  string                          `json:"remark,omitempty"`
	RuleSetInstanceRuleListStatus          []RuleSetInstanceRuleList       `json:"rule-list,omitempty"`
	RuleSetInstanceRulesByZoneUUID         RuleSetInstanceRulesByZone      `json:"rules-by-zone,omitempty"`
	RuleSetInstanceSamplingEnableCounters1 []RuleSetInstanceSamplingEnable `json:"sampling-enable,omitempty"`
	RuleSetInstanceSessionStatistic        string                          `json:"session-statistic,omitempty"`
	RuleSetInstanceTagUUID                 RuleSetInstanceTag              `json:"tag,omitempty"`
	RuleSetInstanceTrackAppRuleListUUID    RuleSetInstanceTrackAppRuleList `json:"track-app-rule-list,omitempty"`
	RuleSetInstanceUUID                    string                          `json:"uuid,omitempty"`
	RuleSetInstanceUserTag                 string                          `json:"user-tag,omitempty"`
}

type RuleSetInstanceApp struct {
	RuleSetInstanceAppUUID string `json:"uuid,omitempty"`
}

type RuleSetInstanceApplication struct {
	RuleSetInstanceApplicationUUID string `json:"uuid,omitempty"`
}

type RuleSetInstanceRuleList struct {
	RuleSetInstanceRuleListAction                   string                                  `json:"action,omitempty"`
	RuleSetInstanceRuleListActionGroupType          RuleSetInstanceRuleListActionGroup      `json:"action-group,omitempty"`
	RuleSetInstanceRuleListAppListObjGrpApplication []RuleSetInstanceRuleListAppList        `json:"app-list,omitempty"`
	RuleSetInstanceRuleListApplicationAny           string                                  `json:"application-any,omitempty"`
	RuleSetInstanceRuleListCgnv6FixedNatLog         int                                     `json:"cgnv6-fixed-nat-log,omitempty"`
	RuleSetInstanceRuleListCgnv6Log                 int                                     `json:"cgnv6-log,omitempty"`
	RuleSetInstanceRuleListCgnv6LsnLid              int                                     `json:"cgnv6-lsn-lid,omitempty"`
	RuleSetInstanceRuleListCgnv6LsnLog              int                                     `json:"cgnv6-lsn-log,omitempty"`
	RuleSetInstanceRuleListCgnv6Policy              string                                  `json:"cgnv6-policy,omitempty"`
	RuleSetInstanceRuleListDestListDstIPSubnet      []RuleSetInstanceRuleListDestList       `json:"dest-list,omitempty"`
	RuleSetInstanceRuleListDscpListDscpValue        []RuleSetInstanceRuleListDscpList       `json:"dscp-list,omitempty"`
	RuleSetInstanceRuleListDstClassList             string                                  `json:"dst-class-list,omitempty"`
	RuleSetInstanceRuleListDstDomainList            string                                  `json:"dst-domain-list,omitempty"`
	RuleSetInstanceRuleListDstGeolocList            string                                  `json:"dst-geoloc-list,omitempty"`
	RuleSetInstanceRuleListDstGeolocListShared      int                                     `json:"dst-geoloc-list-shared,omitempty"`
	RuleSetInstanceRuleListDstGeolocName            string                                  `json:"dst-geoloc-name,omitempty"`
	RuleSetInstanceRuleListDstIpv4Any               string                                  `json:"dst-ipv4-any,omitempty"`
	RuleSetInstanceRuleListDstIpv6Any               string                                  `json:"dst-ipv6-any,omitempty"`
	RuleSetInstanceRuleListDstThreatList            string                                  `json:"dst-threat-list,omitempty"`
	RuleSetInstanceRuleListDstZone                  string                                  `json:"dst-zone,omitempty"`
	RuleSetInstanceRuleListDstZoneAny               string                                  `json:"dst-zone-any,omitempty"`
	RuleSetInstanceRuleListForwardListenOnPort      int                                     `json:"forward-listen-on-port,omitempty"`
	RuleSetInstanceRuleListForwardLog               int                                     `json:"forward-log,omitempty"`
	RuleSetInstanceRuleListFwLog                    int                                     `json:"fw-log,omitempty"`
	RuleSetInstanceRuleListFwlog                    int                                     `json:"fwlog,omitempty"`
	RuleSetInstanceRuleListGtpTemplate              string                                  `json:"gtp-template,omitempty"`
	RuleSetInstanceRuleListIPVersion                string                                  `json:"ip-version,omitempty"`
	RuleSetInstanceRuleListIdleTimeout              int                                     `json:"idle-timeout,omitempty"`
	RuleSetInstanceRuleListLid                      int                                     `json:"lid,omitempty"`
	RuleSetInstanceRuleListLidlog                   int                                     `json:"lidlog,omitempty"`
	RuleSetInstanceRuleListListenOnPort             int                                     `json:"listen-on-port,omitempty"`
	RuleSetInstanceRuleListListenOnPortLid          int                                     `json:"listen-on-port-lid,omitempty"`
	RuleSetInstanceRuleListListenOnPortLidlog       int                                     `json:"listen-on-port-lidlog,omitempty"`
	RuleSetInstanceRuleListLog                      int                                     `json:"log,omitempty"`
	RuleSetInstanceRuleListMoveRuleLocation         RuleSetInstanceRuleListMoveRule         `json:"move-rule,omitempty"`
	RuleSetInstanceRuleListName                     string                                  `json:"name,omitempty"`
	RuleSetInstanceRuleListPacketCaptureTemplate    string                                  `json:"packet-capture-template,omitempty"`
	RuleSetInstanceRuleListPolicy                   string                                  `json:"policy,omitempty"`
	RuleSetInstanceRuleListRemark                   string                                  `json:"remark,omitempty"`
	RuleSetInstanceRuleListResetLid                 int                                     `json:"reset-lid,omitempty"`
	RuleSetInstanceRuleListResetLidlog              int                                     `json:"reset-lidlog,omitempty"`
	RuleSetInstanceRuleListSamplingEnableCounters1  []RuleSetInstanceRuleListSamplingEnable `json:"sampling-enable,omitempty"`
	RuleSetInstanceRuleListServiceAny               string                                  `json:"service-any,omitempty"`
	RuleSetInstanceRuleListServiceListProtocols     []RuleSetInstanceRuleListServiceList    `json:"service-list,omitempty"`
	RuleSetInstanceRuleListSourceListSrcIPSubnet    []RuleSetInstanceRuleListSourceList     `json:"source-list,omitempty"`
	RuleSetInstanceRuleListSrcClassList             string                                  `json:"src-class-list,omitempty"`
	RuleSetInstanceRuleListSrcGeolocList            string                                  `json:"src-geoloc-list,omitempty"`
	RuleSetInstanceRuleListSrcGeolocListShared      int                                     `json:"src-geoloc-list-shared,omitempty"`
	RuleSetInstanceRuleListSrcGeolocName            string                                  `json:"src-geoloc-name,omitempty"`
	RuleSetInstanceRuleListSrcIpv4Any               string                                  `json:"src-ipv4-any,omitempty"`
	RuleSetInstanceRuleListSrcIpv6Any               string                                  `json:"src-ipv6-any,omitempty"`
	RuleSetInstanceRuleListSrcThreatList            string                                  `json:"src-threat-list,omitempty"`
	RuleSetInstanceRuleListSrcZone                  string                                  `json:"src-zone,omitempty"`
	RuleSetInstanceRuleListSrcZoneAny               string                                  `json:"src-zone-any,omitempty"`
	RuleSetInstanceRuleListStatus                   string                                  `json:"status,omitempty"`
	RuleSetInstanceRuleListTrackApplication         int                                     `json:"track-application,omitempty"`
	RuleSetInstanceRuleListUUID                     string                                  `json:"uuid,omitempty"`
	RuleSetInstanceRuleListUserTag                  string                                  `json:"user-tag,omitempty"`
	RuleSetInstanceRuleListVpnIpsecName             string                                  `json:"vpn-ipsec-name,omitempty"`
}

type RuleSetInstanceRulesByZone struct {
	RuleSetInstanceRulesByZoneSamplingEnableCounters1 []RuleSetInstanceRulesByZoneSamplingEnable `json:"sampling-enable,omitempty"`
	RuleSetInstanceRulesByZoneUUID                    string                                     `json:"uuid,omitempty"`
}

type RuleSetInstanceSamplingEnable struct {
	RuleSetInstanceSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type RuleSetInstanceTag struct {
	RuleSetInstanceTagUUID string `json:"uuid,omitempty"`
}

type RuleSetInstanceTrackAppRuleList struct {
	RuleSetInstanceTrackAppRuleListUUID string `json:"uuid,omitempty"`
}

type RuleSetInstanceRuleListActionGroup struct {
	RuleSetInstanceRuleListActionGroupCgnv6                  int    `json:"cgnv6,omitempty"`
	RuleSetInstanceRuleListActionGroupCgnv6LsnLid            int    `json:"cgnv6-lsn-lid,omitempty"`
	RuleSetInstanceRuleListActionGroupCgnv6Policy            string `json:"cgnv6-policy,omitempty"`
	RuleSetInstanceRuleListActionGroupDenyLog                int    `json:"deny-log,omitempty"`
	RuleSetInstanceRuleListActionGroupDscpNumber             int    `json:"dscp-number,omitempty"`
	RuleSetInstanceRuleListActionGroupDscpValue              string `json:"dscp-value,omitempty"`
	RuleSetInstanceRuleListActionGroupForward                int    `json:"forward,omitempty"`
	RuleSetInstanceRuleListActionGroupIpsec                  int    `json:"ipsec,omitempty"`
	RuleSetInstanceRuleListActionGroupListenOnPort           int    `json:"listen-on-port,omitempty"`
	RuleSetInstanceRuleListActionGroupPermitLimitPolicy      int    `json:"permit-limit-policy,omitempty"`
	RuleSetInstanceRuleListActionGroupPermitLog              int    `json:"permit-log,omitempty"`
	RuleSetInstanceRuleListActionGroupPermitRespondToUserMac int    `json:"permit-respond-to-user-mac,omitempty"`
	RuleSetInstanceRuleListActionGroupResetLog               int    `json:"reset-log,omitempty"`
	RuleSetInstanceRuleListActionGroupResetRespondToUserMac  int    `json:"reset-respond-to-user-mac,omitempty"`
	RuleSetInstanceRuleListActionGroupSetDscp                int    `json:"set-dscp,omitempty"`
	RuleSetInstanceRuleListActionGroupType                   string `json:"type,omitempty"`
	RuleSetInstanceRuleListActionGroupUUID                   string `json:"uuid,omitempty"`
	RuleSetInstanceRuleListActionGroupVpnIpsecName           string `json:"vpn-ipsec-name,omitempty"`
}

type RuleSetInstanceRuleListAppList struct {
	RuleSetInstanceRuleListAppListObjGrpApplication string `json:"obj-grp-application,omitempty"`
	RuleSetInstanceRuleListAppListProtocol          string `json:"protocol,omitempty"`
	RuleSetInstanceRuleListAppListProtocolTag       string `json:"protocol-tag,omitempty"`
}

type RuleSetInstanceRuleListDestList struct {
	RuleSetInstanceRuleListDestListDstIPSubnet      string `json:"dst-ip-subnet,omitempty"`
	RuleSetInstanceRuleListDestListDstIpv6Subnet    string `json:"dst-ipv6-subnet,omitempty"`
	RuleSetInstanceRuleListDestListDstObjGrpNetwork string `json:"dst-obj-grp-network,omitempty"`
	RuleSetInstanceRuleListDestListDstObjNetwork    string `json:"dst-obj-network,omitempty"`
	RuleSetInstanceRuleListDestListDstSlbServer     string `json:"dst-slb-server,omitempty"`
	RuleSetInstanceRuleListDestListDstSlbVserver    string `json:"dst-slb-vserver,omitempty"`
}

type RuleSetInstanceRuleListDscpList struct {
	RuleSetInstanceRuleListDscpListDscpRangeEnd   int    `json:"dscp-range-end,omitempty"`
	RuleSetInstanceRuleListDscpListDscpRangeStart int    `json:"dscp-range-start,omitempty"`
	RuleSetInstanceRuleListDscpListDscpValue      string `json:"dscp-value,omitempty"`
}

type RuleSetInstanceRuleListMoveRule struct {
	RuleSetInstanceRuleListMoveRuleLocation   string `json:"location,omitempty"`
	RuleSetInstanceRuleListMoveRuleTargetRule string `json:"target-rule,omitempty"`
}

type RuleSetInstanceRuleListSamplingEnable struct {
	RuleSetInstanceRuleListSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type RuleSetInstanceRuleListServiceList struct {
	RuleSetInstanceRuleListServiceListAlg           string `json:"alg,omitempty"`
	RuleSetInstanceRuleListServiceListEqDstPort     int    `json:"eq-dst-port,omitempty"`
	RuleSetInstanceRuleListServiceListEqSrcPort     int    `json:"eq-src-port,omitempty"`
	RuleSetInstanceRuleListServiceListGtDstPort     int    `json:"gt-dst-port,omitempty"`
	RuleSetInstanceRuleListServiceListGtSrcPort     int    `json:"gt-src-port,omitempty"`
	RuleSetInstanceRuleListServiceListIcmp          int    `json:"icmp,omitempty"`
	RuleSetInstanceRuleListServiceListIcmpCode      int    `json:"icmp-code,omitempty"`
	RuleSetInstanceRuleListServiceListIcmpType      int    `json:"icmp-type,omitempty"`
	RuleSetInstanceRuleListServiceListIcmpv6        int    `json:"icmpv6,omitempty"`
	RuleSetInstanceRuleListServiceListIcmpv6Code    int    `json:"icmpv6-code,omitempty"`
	RuleSetInstanceRuleListServiceListIcmpv6Type    int    `json:"icmpv6-type,omitempty"`
	RuleSetInstanceRuleListServiceListLtDstPort     int    `json:"lt-dst-port,omitempty"`
	RuleSetInstanceRuleListServiceListLtSrcPort     int    `json:"lt-src-port,omitempty"`
	RuleSetInstanceRuleListServiceListObjGrpService string `json:"obj-grp-service,omitempty"`
	RuleSetInstanceRuleListServiceListPortNumEndDst int    `json:"port-num-end-dst,omitempty"`
	RuleSetInstanceRuleListServiceListPortNumEndSrc int    `json:"port-num-end-src,omitempty"`
	RuleSetInstanceRuleListServiceListProtoID       int    `json:"proto-id,omitempty"`
	RuleSetInstanceRuleListServiceListProtocols     string `json:"protocols,omitempty"`
	RuleSetInstanceRuleListServiceListRangeDstPort  int    `json:"range-dst-port,omitempty"`
	RuleSetInstanceRuleListServiceListRangeSrcPort  int    `json:"range-src-port,omitempty"`
	RuleSetInstanceRuleListServiceListSctpTemplate  string `json:"sctp-template,omitempty"`
	RuleSetInstanceRuleListServiceListSpecialCode   string `json:"special-code,omitempty"`
	RuleSetInstanceRuleListServiceListSpecialType   string `json:"special-type,omitempty"`
	RuleSetInstanceRuleListServiceListSpecialV6Code string `json:"special-v6-code,omitempty"`
	RuleSetInstanceRuleListServiceListSpecialV6Type string `json:"special-v6-type,omitempty"`
}

type RuleSetInstanceRuleListSourceList struct {
	RuleSetInstanceRuleListSourceListSrcIPSubnet      string `json:"src-ip-subnet,omitempty"`
	RuleSetInstanceRuleListSourceListSrcIpv6Subnet    string `json:"src-ipv6-subnet,omitempty"`
	RuleSetInstanceRuleListSourceListSrcObjGrpNetwork string `json:"src-obj-grp-network,omitempty"`
	RuleSetInstanceRuleListSourceListSrcObjNetwork    string `json:"src-obj-network,omitempty"`
	RuleSetInstanceRuleListSourceListSrcSlbServer     string `json:"src-slb-server,omitempty"`
}

type RuleSetInstanceRulesByZoneSamplingEnable struct {
	RuleSetInstanceRulesByZoneSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

func PostRuleSet(id string, inst RuleSet, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRuleSet")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/rule-set", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RuleSet
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRuleSet", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetRuleSet(id string, name1 string, host string) (*RuleSet, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRuleSet")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/rule-set/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RuleSet
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRuleSet", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutRuleSet(id string, name1 string, inst RuleSet, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutRuleSet")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/rule-set/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RuleSet
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutRuleSet", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteRuleSet(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteRuleSet")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/rule-set/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RuleSet
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteRuleSet", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
