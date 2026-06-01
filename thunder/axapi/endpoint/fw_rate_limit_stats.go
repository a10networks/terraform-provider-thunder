package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_8-219
type FwRateLimitStats struct {
	Stats FwRateLimitStatsStats `json:"stats"`
}
type DataFwRateLimitStats struct {
	DtFwRateLimitStats FwRateLimitStats `json:"rate-limit"`
}

type FwRateLimitStatsStats struct {
	Ratelimit_used_total_mem                      int `json:"ratelimit_used_total_mem"`
	Ratelimit_entry_count_t2_key                  int `json:"ratelimit_entry_count_t2_key"`
	Ratelimit_entry_count_fw_rule_uid             int `json:"ratelimit_entry_count_fw_rule_uid"`
	Ratelimit_entry_count_ip_addr                 int `json:"ratelimit_entry_count_ip_addr"`
	Ratelimit_entry_count_ip6_addr                int `json:"ratelimit_entry_count_ip6_addr"`
	Ratelimit_entry_count_session_id              int `json:"ratelimit_entry_count_session_id"`
	Ratelimit_entry_count_rule_ipv4_prefix        int `json:"ratelimit_entry_count_rule_ipv4_prefix"`
	Ratelimit_entry_count_rule_ipv6_prefix        int `json:"ratelimit_entry_count_rule_ipv6_prefix"`
	Ratelimit_entry_count_parent_uid              int `json:"ratelimit_entry_count_parent_uid"`
	Ratelimit_entry_count_parent_ipv4_prefix      int `json:"ratelimit_entry_count_parent_ipv4_prefix"`
	Ratelimit_entry_count_parent_ipv6_prefix      int `json:"ratelimit_entry_count_parent_ipv6_prefix"`
	Ratelimit_entry_count_allocated               int `json:"ratelimit_entry_count_allocated"`
	Ratelimit_entry_count_freed                   int `json:"ratelimit_entry_count_freed"`
	Ratelimit_entry_count_rule_ip                 int `json:"ratelimit_entry_count_rule_ip"`
	Ratelimit_entry_count_parent_ip               int `json:"ratelimit_entry_count_parent_ip"`
	Ratelimit_entry_count_radius_usergroup        int `json:"ratelimit_entry_count_radius_usergroup"`
	Ratelimit_entry_count_parent_radius_usergroup int `json:"ratelimit_entry_count_parent_radius_usergroup"`
	Ratelimit_entry_count_radius_userid           int `json:"ratelimit_entry_count_radius_userid"`
	Ratelimit_entry_count_parent_radius_userid    int `json:"ratelimit_entry_count_parent_radius_userid"`
}

func (p *FwRateLimitStats) GetId() string {
	return "1"
}

func (p *FwRateLimitStats) getPath() string {
	return "fw/rate-limit/stats"
}

func (p *FwRateLimitStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataFwRateLimitStats, error) {
	logger.Println("FwRateLimitStats::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
	var payload DataFwRateLimitStats
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
