package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 7_0_2-102
type SlbRpzOper struct {
	Oper SlbRpzOperOper `json:"oper"`
}
type DataSlbRpzOper struct {
	DtSlbRpzOper SlbRpzOper `json:"rpz"`
}

type SlbRpzOperOper struct {
	Filter_entry            string                      `json:"filter_entry"`
	FileList                []SlbRpzOperOperFileList    `json:"file-list"`
	RpzFileSizeMax          int                         `json:"rpz-file-size-max"`
	RpzCount                int                         `json:"rpz-count"`
	RpzRuleCount            int                         `json:"rpz-rule-count"`
	ClassList               string                      `json:"class-list"`
	Type                    string                      `json:"type"`
	FileOrString            string                      `json:"file-or-string"`
	UserTag                 string                      `json:"user-tag"`
	Ipv4TotalSingleIp       int                         `json:"ipv4-total-single-ip"`
	Ipv4TotalSubnet         int                         `json:"ipv4-total-subnet"`
	Ipv6TotalSingleIp       int                         `json:"ipv6-total-single-ip"`
	Ipv6TotalSubnet         int                         `json:"ipv6-total-subnet"`
	DnsTotalEntries         int                         `json:"dns-total-entries"`
	StringTotalEntries      int                         `json:"string-total-entries"`
	AcTotalEntries          int                         `json:"ac-total-entries"`
	GeoLocationTotalEntries int                         `json:"geo-location-total-entries"`
	Ipv4Entries             []SlbRpzOperOperIpv4Entries `json:"ipv4-entries"`
	Ipv6Entries             []SlbRpzOperOperIpv6Entries `json:"ipv6-entries"`
	DnsEntries              []SlbRpzOperOperDnsEntries  `json:"dns-entries"`
}

type SlbRpzOperOperFileList struct {
	File               string `json:"file"`
	Dns_template_bound string `json:"dns_template_bound"`
}

type SlbRpzOperOperIpv4Entries struct {
	Ipv4Addr             string `json:"ipv4-addr"`
	Ipv4Lid              int    `json:"ipv4-lid"`
	Ipv4Glid             int    `json:"ipv4-glid"`
	Ipv4LsnLid           int    `json:"ipv4-lsn-lid"`
	Ipv4LsnRadiusProfile int    `json:"ipv4-lsn-radius-profile"`
	Ipv4GtpPolicy        string `json:"ipv4-gtp-policy"`
	Ipv4HitCount         int    `json:"ipv4-hit-count"`
	Ipv4Age              int    `json:"ipv4-age"`
	Ipv4RpzType          int    `json:"ipv4-rpz-type"`
}

type SlbRpzOperOperIpv6Entries struct {
	Ipv6addr             string `json:"ipv6addr"`
	Ipv6Lid              int    `json:"ipv6-lid"`
	Ipv6Glid             int    `json:"ipv6-glid"`
	Ipv6LsnLid           int    `json:"ipv6-lsn-lid"`
	Ipv6LsnRadiusProfile int    `json:"ipv6-lsn-radius-profile"`
	Ipv6GtpPolicy        string `json:"ipv6-gtp-policy"`
	Ipv6HitCount         int    `json:"ipv6-hit-count"`
	Ipv6Age              int    `json:"ipv6-age"`
	Ipv6RpzType          int    `json:"ipv6-rpz-type"`
}

type SlbRpzOperOperDnsEntries struct {
	DnsMatchType   string `json:"dns-match-type"`
	DnsMatchString string `json:"dns-match-string"`
	DnsLid         int    `json:"dns-lid"`
	DnsGlid        int    `json:"dns-glid"`
	DnsHitCount    int    `json:"dns-hit-count"`
	DnsRpzType     int    `json:"dns-rpz-type"`
}

func (p *SlbRpzOper) GetId() string {
	return "1"
}

func (p *SlbRpzOper) getPath() string {
	return "slb/rpz/oper"
}

func (p *SlbRpzOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbRpzOper, error) {
	logger.Println("SlbRpzOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataSlbRpzOper
	if err == nil {
		if len(axResp) > 0 {
			err = json.Unmarshal(axResp, &p)
		}
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return payload, err
}
