package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 5_2_1-P4_81
type ClassList struct {
	Inst struct {
		AcList   []ClassListAcList   `json:"ac-list"`
		Dns      []ClassListDns      `json:"dns"`
		File     int                 `json:"file"`
		Ipv4List []ClassListIpv4List `json:"ipv4-list"`
		Ipv6List []ClassListIpv6List `json:"ipv6-list"`
		Name     string              `json:"name"`
		StrList  []ClassListStrList  `json:"str-list"`
		Type     string              `json:"type"`
		UserTag  string              `json:"user-tag"`
		Uuid     string              `json:"uuid"`
	} `json:"class-list"`
}

type ClassListAcList struct {
	AcMatchType           string `json:"ac-match-type"`
	AcKeyString           string `json:"ac-key-string"`
	AcValue               string `json:"ac-value"`
	GtpRateLimitPolicyStr string `json:"gtp-rate-limit-policy-str"`
}

type ClassListDns struct {
	DnsMatchType           string `json:"dns-match-type"`
	DnsMatchString         string `json:"dns-match-string"`
	DnsLid                 int    `json:"dns-lid"`
	DnsGlid                int    `json:"dns-glid"`
	SharedPartitionDnsGlid int    `json:"shared-partition-dns-glid"`
	DnsGlidShared          int    `json:"dns-glid-shared"`
}

type ClassListIpv4List struct {
	Ipv4addr             string `json:"ipv4addr"`
	Lid                  int    `json:"lid"`
	Glid                 int    `json:"glid"`
	SharedPartitionGlid  int    `json:"shared-partition-glid"`
	GlidShared           int    `json:"glid-shared"`
	LsnLid               int    `json:"lsn-lid"`
	LsnRadiusProfile     int    `json:"lsn-radius-profile"`
	GtpRateLimitPolicyV4 string `json:"gtp-rate-limit-policy-v4"`
	Age                  int    `json:"age"`
}

type ClassListIpv6List struct {
	Ipv6Addr              string `json:"ipv6-addr"`
	V6Lid                 int    `json:"v6-lid"`
	V6Glid                int    `json:"v6-glid"`
	SharedPartitionV6Glid int    `json:"shared-partition-v6-glid"`
	V6GlidShared          int    `json:"v6-glid-shared"`
	V6LsnLid              int    `json:"v6-lsn-lid"`
	V6LsnRadiusProfile    int    `json:"v6-lsn-radius-profile"`
	GtpRateLimitPolicyV6  string `json:"gtp-rate-limit-policy-v6"`
	V6Age                 int    `json:"v6-age"`
}

type ClassListStrList struct {
	Str                    string `json:"str"`
	StrLidDummy            int    `json:"str-lid-dummy"`
	StrLid                 int    `json:"str-lid"`
	StrGlidDummy           int    `json:"str-glid-dummy"`
	StrGlid                int    `json:"str-glid"`
	SharedPartitionStrGlid int    `json:"shared-partition-str-glid"`
	StrGlidShared          int    `json:"str-glid-shared"`
	ValueStr               string `json:"value-str"`
}

func (p *ClassList) GetId() string {
	return url.QueryEscape(p.Inst.Name)
}

func (p *ClassList) getPath() string {
	return "class-list"
}

func (p *ClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ClassList::Post")
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

func (p *ClassList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ClassList::Get")
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

func (p *ClassList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("ClassList::Put")
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

func (p *ClassList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("ClassList::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
