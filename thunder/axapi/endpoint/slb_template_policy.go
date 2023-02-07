package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
	"net/url"
)

// based on ACOS 5_2_1-P6_74
type SlbTemplatePolicy struct {
	Inst struct {
		BwListId         []SlbTemplatePolicyBwListId       `json:"bw-list-id"`
		BwListName       string                            `json:"bw-list-name"`
		ClassList        SlbTemplatePolicyClassList        `json:"class-list"`
		ForwardPolicy    SlbTemplatePolicyForwardPolicy    `json:"forward-policy"`
		FullDomainTree   int                               `json:"full-domain-tree"`
		Interval         int                               `json:"interval"`
		Name             string                            `json:"name"`
		OverLimit        int                               `json:"over-limit"`
		OverLimitLockup  int                               `json:"over-limit-lockup"`
		OverLimitLogging int                               `json:"over-limit-logging"`
		OverLimitReset   int                               `json:"over-limit-reset"`
		Overlap          int                               `json:"overlap"`
		SamplingEnable   []SlbTemplatePolicySamplingEnable `json:"sampling-enable"`
		Share            int                               `json:"share"`
		Timeout          int                               `json:"timeout" dval:"5"`
		UseDestinationIp int                               `json:"use-destination-ip"`
		UserTag          string                            `json:"user-tag"`
		Uuid             string                            `json:"uuid"`
	} `json:"policy"`
}

type SlbTemplatePolicyBwListId struct {
	Id             int    `json:"id"`
	ServiceGroup   string `json:"service-group"`
	PbslbLogging   int    `json:"pbslb-logging"`
	PbslbInterval  int    `json:"pbslb-interval" dval:"3"`
	Fail           int    `json:"fail"`
	BwListAction   string `json:"bw-list-action"`
	LoggingDrpRst  int    `json:"logging-drp-rst"`
	ActionInterval int    `json:"action-interval" dval:"3"`
}
type SlbTemplatePolicyClassList struct {
	Name             string                              `json:"name"`
	ClientIpL3Dest   int                                 `json:"client-ip-l3-dest"`
	ClientIpL7Header int                                 `json:"client-ip-l7-header"`
	HeaderName       string                              `json:"header-name"`
	Uuid             string                              `json:"uuid"`
	LidList          []SlbTemplatePolicyClassListLidList `json:"lid-list"`
}
type SlbTemplatePolicyClassListLidList struct {
	Lidnum                int                                                      `json:"lidnum"`
	ConnLimit             int                                                      `json:"conn-limit"`
	ConnRateLimit         int                                                      `json:"conn-rate-limit"`
	ConnPer               int                                                      `json:"conn-per"`
	RequestLimit          int                                                      `json:"request-limit"`
	RequestRateLimit      int                                                      `json:"request-rate-limit"`
	RequestPer            int                                                      `json:"request-per"`
	BwRateLimit           int                                                      `json:"bw-rate-limit"`
	BwPer                 int                                                      `json:"bw-per"`
	OverLimitAction       int                                                      `json:"over-limit-action"`
	ActionValue           string                                                   `json:"action-value"`
	Lockout               int                                                      `json:"lockout"`
	Log                   int                                                      `json:"log"`
	Interval              int                                                      `json:"interval"`
	DirectAction          int                                                      `json:"direct-action"`
	DirectServiceGroup    string                                                   `json:"direct-service-group"`
	DirectPbslbLogging    int                                                      `json:"direct-pbslb-logging"`
	DirectPbslbInterval   int                                                      `json:"direct-pbslb-interval" dval:"3"`
	DirectFail            int                                                      `json:"direct-fail"`
	DirectActionValue     string                                                   `json:"direct-action-value"`
	DirectLoggingDrpRst   int                                                      `json:"direct-logging-drp-rst"`
	DirectActionInterval  int                                                      `json:"direct-action-interval" dval:"3"`
	ResponseCodeRateLimit []SlbTemplatePolicyClassListLidListResponseCodeRateLimit `json:"response-code-rate-limit"`
	Dns64                 SlbTemplatePolicyClassListLidListDns64                   `json:"dns64"`
	Uuid                  string                                                   `json:"uuid"`
	UserTag               string                                                   `json:"user-tag"`
}
type SlbTemplatePolicyClassListLidListResponseCodeRateLimit struct {
	CodeRangeStart int `json:"code-range-start"`
	CodeRangeEnd   int `json:"code-range-end"`
	Threshold      int `json:"threshold"`
	Period         int `json:"period"`
}
type SlbTemplatePolicyClassListLidListDns64 struct {
	Disable         int    `json:"disable"`
	ExclusiveAnswer int    `json:"exclusive-answer"`
	Prefix          string `json:"prefix"`
}
type SlbTemplatePolicyForwardPolicy struct {
	NoClientConnReuse  int                                          `json:"no-client-conn-reuse"`
	AcosEventLog       int                                          `json:"acos-event-log"`
	LocalLogging       int                                          `json:"local-logging"`
	RequireWebCategory int                                          `json:"require-web-category"`
	Filtering          []SlbTemplatePolicyForwardPolicyFiltering    `json:"filtering"`
	SanFiltering       []SlbTemplatePolicyForwardPolicySanFiltering `json:"san-filtering"`
	Uuid               string                                       `json:"uuid"`
	ActionList         []SlbTemplatePolicyForwardPolicyActionList   `json:"action-list"`
	SourceList         []SlbTemplatePolicyForwardPolicySourceList   `json:"source-list"`
}
type SlbTemplatePolicyForwardPolicyFiltering struct {
	SsliUrlFiltering string `json:"ssli-url-filtering"`
}
type SlbTemplatePolicyForwardPolicySanFiltering struct {
	SsliUrlFilteringSan string `json:"ssli-url-filtering-san"`
}
type SlbTemplatePolicyForwardPolicyActionList struct {
	Name                string                                                   `json:"name"`
	Action1             string                                                   `json:"action1"`
	FakeSg              string                                                   `json:"fake-sg"`
	RealSg              string                                                   `json:"real-sg"`
	ForwardSnat         string                                                   `json:"forward-snat"`
	FallBack            string                                                   `json:"fall-back"`
	FallBackSnat        string                                                   `json:"fall-back-snat"`
	ProxyChaining       int                                                      `json:"proxy-chaining"`
	ProxyChainingBypass int                                                      `json:"proxy-chaining-bypass"`
	SupportCertFetch    int                                                      `json:"support-cert-fetch"`
	Log                 int                                                      `json:"log"`
	DropResponseCode    int                                                      `json:"drop-response-code"`
	DropMessage         string                                                   `json:"drop-message"`
	DropRedirectUrl     string                                                   `json:"drop-redirect-url"`
	HttpStatusCode      string                                                   `json:"http-status-code" dval:"302"`
	Uuid                string                                                   `json:"uuid"`
	UserTag             string                                                   `json:"user-tag"`
	SamplingEnable      []SlbTemplatePolicyForwardPolicyActionListSamplingEnable `json:"sampling-enable"`
}
type SlbTemplatePolicyForwardPolicyActionListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}
type SlbTemplatePolicyForwardPolicySourceList struct {
	Name                 string                                                   `json:"name"`
	MatchClassList       string                                                   `json:"match-class-list"`
	MatchAny             int                                                      `json:"match-any"`
	MatchAuthorizePolicy string                                                   `json:"match-authorize-policy"`
	Priority             int                                                      `json:"priority"`
	Uuid                 string                                                   `json:"uuid"`
	UserTag              string                                                   `json:"user-tag"`
	SamplingEnable       []SlbTemplatePolicyForwardPolicySourceListSamplingEnable `json:"sampling-enable"`
	Destination          SlbTemplatePolicyForwardPolicySourceListDestination      `json:"destination"`
}
type SlbTemplatePolicyForwardPolicySourceListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}
type SlbTemplatePolicyForwardPolicySourceListDestination struct {
	ClassListList          []SlbTemplatePolicyForwardPolicySourceListDestinationClassListList          `json:"class-list-list"`
	WebReputationScopeList []SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList `json:"web-reputation-scope-list"`
	WebCategoryListList    []SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList    `json:"web-category-list-list"`
	Any                    SlbTemplatePolicyForwardPolicySourceListDestinationAny                      `json:"any"`
}
type SlbTemplatePolicyForwardPolicySourceListDestinationClassListList struct {
	DestClassList  string                                                                           `json:"dest-class-list"`
	Action         string                                                                           `json:"action"`
	Type           string                                                                           `json:"type"`
	Priority       int                                                                              `json:"priority"`
	Uuid           string                                                                           `json:"uuid"`
	SamplingEnable []SlbTemplatePolicyForwardPolicySourceListDestinationClassListListSamplingEnable `json:"sampling-enable"`
}
type SlbTemplatePolicyForwardPolicySourceListDestinationClassListListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}
type SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeList struct {
	WebReputationScope string                                                                                    `json:"web-reputation-scope"`
	Action             string                                                                                    `json:"action"`
	Type               string                                                                                    `json:"type"`
	Priority           int                                                                                       `json:"priority"`
	Uuid               string                                                                                    `json:"uuid"`
	SamplingEnable     []SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeListSamplingEnable `json:"sampling-enable"`
}
type SlbTemplatePolicyForwardPolicySourceListDestinationWebReputationScopeListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}
type SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListList struct {
	WebCategoryList string                                                                                 `json:"web-category-list"`
	Action          string                                                                                 `json:"action"`
	Type            string                                                                                 `json:"type"`
	Priority        int                                                                                    `json:"priority"`
	Uuid            string                                                                                 `json:"uuid"`
	SamplingEnable  []SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListListSamplingEnable `json:"sampling-enable"`
}
type SlbTemplatePolicyForwardPolicySourceListDestinationWebCategoryListListSamplingEnable struct {
	Counters1 string `json:"counters1"`
}
type SlbTemplatePolicyForwardPolicySourceListDestinationAny struct {
	Action         string                                                                 `json:"action"`
	Uuid           string                                                                 `json:"uuid"`
	SamplingEnable []SlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable `json:"sampling-enable"`
}
type SlbTemplatePolicyForwardPolicySourceListDestinationAnySamplingEnable struct {
	Counters1 string `json:"counters1"`
}
type SlbTemplatePolicySamplingEnable struct {
	Counters1 string `json:"counters1"`
}

func (p *SlbTemplatePolicy) GetId() string {
	return url.QueryEscape(p.Inst.Name)
}
func (p *SlbTemplatePolicy) getPath() string {
	return "slb/template/policy"
}
func (p *SlbTemplatePolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplatePolicy::Post")
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
func (p *SlbTemplatePolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplatePolicy::Get")
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
func (p *SlbTemplatePolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplatePolicy::Put")
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
func (p *SlbTemplatePolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("SlbTemplatePolicy::Delete")
	headers := axapi.GenRequestHeader(authToken)
	_, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
	return err
}
