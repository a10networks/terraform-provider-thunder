package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_0-P1_10
type SessionsOper struct {
	Ext      SessionsOperExt      `json:"ext"`
	Oper     SessionsOperOper     `json:"oper"`
	Smp      SessionsOperSmp      `json:"smp"`
	SmpTable SessionsOperSmpTable `json:"smp-table"`
}

type Sessionss struct {
	Datasession SessionsOper `json:"sessions"`
}

type SessionsOperExt struct {
	Oper SessionsOperExtOper `json:"oper"`
}

type SessionsOperExtOper struct {
	SessionExtList []SessionsOperExtOperSessionExtList `json:"session-ext-list"`
}

type SessionsOperExtOperSessionExtList struct {
	Type              string `json:"type"`
	Alloc             int    `json:"alloc"`
	Free              int    `json:"free"`
	Fail              int    `json:"fail"`
	CpuRoundRobinFail int    `json:"cpu-round-robin-fail"`
	AllocExceed       int    `json:"alloc-exceed"`
}

type SessionsOperOper struct {
	SessionList          []SessionsOperOperSessionList `json:"session-list"`
	TotalSessions        int                           `json:"total-sessions"`
	AppSessions          int                           `json:"app-sessions"`
	Filter_type          string                        `json:"filter_type"`
	SrcIpv4Addr          string                        `json:"src-ipv4-addr"`
	DstIpv4Addr          string                        `json:"dst-ipv4-addr"`
	NatIpv4Addr          string                        `json:"nat-ipv4-addr"`
	SrcIpv6Addr          string                        `json:"src-ipv6-addr"`
	DstIpv6Addr          string                        `json:"dst-ipv6-addr"`
	NameStr              string                        `json:"name-str"`
	DestPort             int                           `json:"dest-port"`
	SrcPort              int                           `json:"src-port"`
	NatPort              int                           `json:"nat-port"`
	Thread               int                           `json:"thread"`
	Bucket               int                           `json:"bucket"`
	AppCategory          string                        `json:"app-category"`
	App                  string                        `json:"app"`
	L4Protocol           string                        `json:"l4-protocol"`
	FwHelperSessions     int                           `json:"fw-helper-sessions"`
	FwIpType             string                        `json:"fw-ip-type"`
	FwRule               string                        `json:"fw-rule"`
	FwDestZone           string                        `json:"fw-dest-zone"`
	FwSrcZone            string                        `json:"fw-src-zone"`
	FwDestObj            string                        `json:"fw-dest-obj"`
	FwSrcObj             string                        `json:"fw-src-obj"`
	FwDestObjGrp         string                        `json:"fw-dest-obj-grp"`
	FwSrcObjGrp          string                        `json:"fw-src-obj-grp"`
	FwDestRserver        string                        `json:"fw-dest-rserver"`
	FwSrcRserver         string                        `json:"fw-src-rserver"`
	FwDestVserver        string                        `json:"fw-dest-vserver"`
	Application          string                        `json:"application"`
	SessionId            string                        `json:"session-id"`
	ZoneName             string                        `json:"zone-name"`
	SportRateLimitExceed int                           `json:"sport-rate-limit-exceed"`
	SportRateLimitCurr   int                           `json:"sport-rate-limit-curr"`
	SrcIpv6Prefix        string                        `json:"src-ipv6-prefix"`
	DstIpv6Prefix        string                        `json:"dst-ipv6-prefix"`
	CheckInsideUser      int                           `json:"check-inside-user"`
	RevDestTeid          int                           `json:"rev-dest-teid"`
	Msisdn               int                           `json:"msisdn"`
	MsisdnVal            string                        `json:"msisdn-val"`
	Imsi                 int                           `json:"imsi"`
	ImsiVal              string                        `json:"imsi-val"`
	GtpMsgType           string                        `json:"gtp-msg-type"`
	GtpVersion           string                        `json:"gtp-version"`
	FullWidth            int                           `json:"full-width"`
	ExtFilterName        string                        `json:"ext-filter-name"`
}

type SessionsOperOperSessionList struct {
	Protocol                string                                           `json:"protocol"`
	ForwardSource           string                                           `json:"forward-source"`
	ForwardDest             string                                           `json:"forward-dest"`
	ReverseSource           string                                           `json:"reverse-source"`
	ReverseDest             string                                           `json:"reverse-dest"`
	Rate                    int                                              `json:"rate"`
	Limit                   int                                              `json:"limit"`
	Drop                    int                                              `json:"drop"`
	PeakRate                int                                              `json:"peak-rate"`
	Age                     int                                              `json:"age"`
	Hash                    int                                              `json:"hash"`
	Flags                   string                                           `json:"flags"`
	AppType                 string                                           `json:"app-type"`
	Ms100                   string                                           `json:"100ms"`
	Sip_call_id             string                                           `json:"sip_call_id"`
	AppName                 string                                           `json:"app-name"`
	ServiceName             string                                           `json:"service-name"`
	RserverName             string                                           `json:"rserver-name"`
	CategoryName            string                                           `json:"category-name"`
	Bytes                   int                                              `json:"bytes"`
	Duration                int                                              `json:"duration"`
	Conn_idx                int                                              `json:"conn_idx"`
	Hash_idx                int                                              `json:"hash_idx"`
	Ddos_total_fwd_bytes    int                                              `json:"ddos_total_fwd_bytes"`
	Ddos_total_rev_bytes    int                                              `json:"ddos_total_rev_bytes"`
	Ddos_total_out_of_order int                                              `json:"ddos_total_out_of_order"`
	Ddos_total_zero_window  int                                              `json:"ddos_total_zero_window"`
	Ddos_total_retrans      int                                              `json:"ddos_total_retrans"`
	Ddos_current_pkt_rate   int                                              `json:"ddos_current_pkt_rate"`
	Ddos_exceeded_pkt_rate  int                                              `json:"ddos_exceeded_pkt_rate"`
	ExtensionFieldsList     []SessionsOperOperSessionListExtensionFieldsList `json:"extension-fields-list"`
	Dns_id                  int                                              `json:"dns_id"`
}

type SessionsOperOperSessionListExtensionFieldsList struct {
	ExtFieldName string `json:"ext-field-name"`
	ExtFieldVal  string `json:"ext-field-val"`
}

type SessionsOperSmp struct {
	Oper SessionsOperSmpOper `json:"oper"`
}

type SessionsOperSmpOper struct {
	SessionSmpList []SessionsOperSmpOperSessionSmpList `json:"session-smp-list"`
}

type SessionsOperSmpOperSessionSmpList struct {
	Type      string `json:"type"`
	Alloc     int    `json:"alloc"`
	Free      int    `json:"free"`
	AllocFail int    `json:"alloc-fail"`
}

type SessionsOperSmpTable struct {
	Oper SessionsOperSmpTableOper `json:"oper"`
}

type SessionsOperSmpTableOper struct {
	EntryList []SessionsOperSmpTableOperEntryList `json:"entry-list"`
}

type SessionsOperSmpTableOperEntryList struct {
	Src4    string `json:"src4"`
	Src6    string `json:"src6"`
	Dst4    string `json:"dst4"`
	Dst6    string `json:"dst6"`
	Srcport int    `json:"srcport"`
	Dstport int    `json:"dstport"`
	Ttl     int    `json:"ttl"`
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

func (p *SessionsOper) GetId() string {
	return "1"
}

func (p *SessionsOper) getPath() string {
	return "sessions/oper"
}

func (p *SessionsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (Sessionss, error) {
	logger.Println("SessionsOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)

	var ss Sessionss
	if err == nil {
		err = json.Unmarshal(axResp, &ss)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return ss, err
}
