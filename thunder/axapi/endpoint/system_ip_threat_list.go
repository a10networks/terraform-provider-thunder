package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type SystemIpThreatList struct {
	Inst struct {
		Ipv4DestList SystemIpThreatListIpv4DestList1682 `json:"ipv4-dest-list"`

		Ipv4InternetHostList SystemIpThreatListIpv4InternetHostList1684 `json:"ipv4-internet-host-list"`

		Ipv4SourceList SystemIpThreatListIpv4SourceList1686 `json:"ipv4-source-list"`

		Ipv6DestList SystemIpThreatListIpv6DestList1688 `json:"ipv6-dest-list"`

		Ipv6InternetHostList SystemIpThreatListIpv6InternetHostList1690 `json:"ipv6-internet-host-list"`

		Ipv6SourceList SystemIpThreatListIpv6SourceList1692 `json:"ipv6-source-list"`

		SamplingEnable []SystemIpThreatListSamplingEnable `json:"sampling-enable"`

		Uuid string `json:"uuid"`
	} `json:"ip-threat-list"`
}

type SystemIpThreatListIpv4DestList1682 struct {
	ClassListCfg []SystemIpThreatListIpv4DestListClassListCfg1683 `json:"class-list-cfg"`
	Uuid         string                                           `json:"uuid"`
}

type SystemIpThreatListIpv4DestListClassListCfg1683 struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIpThreatListIpv4InternetHostList1684 struct {
	WhiteList    string                                                   `json:"white-list"`
	ClassListCfg []SystemIpThreatListIpv4InternetHostListClassListCfg1685 `json:"class-list-cfg"`
	Uuid         string                                                   `json:"uuid"`
}

type SystemIpThreatListIpv4InternetHostListClassListCfg1685 struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIpThreatListIpv4SourceList1686 struct {
	ClassListCfg []SystemIpThreatListIpv4SourceListClassListCfg1687 `json:"class-list-cfg"`
	Uuid         string                                             `json:"uuid"`
}

type SystemIpThreatListIpv4SourceListClassListCfg1687 struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIpThreatListIpv6DestList1688 struct {
	ClassListCfg []SystemIpThreatListIpv6DestListClassListCfg1689 `json:"class-list-cfg"`
	Uuid         string                                           `json:"uuid"`
}

type SystemIpThreatListIpv6DestListClassListCfg1689 struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIpThreatListIpv6InternetHostList1690 struct {
	WhiteList    string                                                   `json:"white-list"`
	ClassListCfg []SystemIpThreatListIpv6InternetHostListClassListCfg1691 `json:"class-list-cfg"`
	Uuid         string                                                   `json:"uuid"`
}

type SystemIpThreatListIpv6InternetHostListClassListCfg1691 struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIpThreatListIpv6SourceList1692 struct {
	ClassListCfg []SystemIpThreatListIpv6SourceListClassListCfg1693 `json:"class-list-cfg"`
	Uuid         string                                             `json:"uuid"`
}

type SystemIpThreatListIpv6SourceListClassListCfg1693 struct {
	ClassList          string `json:"class-list"`
	IpThreatActionTmpl int    `json:"ip-threat-action-tmpl"`
}

type SystemIpThreatListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}

func (p *SystemIpThreatList) GetId() string {
	return "1"
}

func (p *SystemIpThreatList) getPath() string {
	return "system/ip-threat-list"
}

func (p *SystemIpThreatList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemIpThreatList::Post")
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

func (p *SystemIpThreatList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemIpThreatList::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return err
}
func (p *SystemIpThreatList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SystemIpThreatList::Put")
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

func (p *SystemIpThreatList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SystemIpThreatList::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
	return err
}
