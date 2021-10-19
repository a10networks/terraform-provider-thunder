---
layout: "thunder"
page_title: "thunder: thunder_slb_service_group"
sidebar_current: "docs-thunder-resource-slb-service-group"
description: |-
    Provides details about thunder slb service group resource for A10
---

# thunder\_slb\_service\_group

`thunder_slb_service_group` Provides details about thunder slb service group
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_slb_service_group" "resourceSlbServiceGroupTest" {
	name = "string"
protocol = "string"
template_port = "string"
template_server = "string"
template_policy = "string"
shared_partition_policy_template = 0
template_policy_shared = "string"
lb_method = "string"
lc_method = "string"
stateless_lb_method = "string"
llb_method = "string"
link_probe_template = "string"
lclb_method = "string"
pseudo_round_robin = 0
stateless_auto_switch = 0
stateless_lb_method2 = "string"
conn_rate = 0
conn_rate_duration = 0
conn_revert_rate = 0
conn_rate_revert_duration = 0
conn_rate_grace_period = 0
conn_rate_log = 0
l4_session_usage = 0
l4_session_usage_duration = 0
l4_session_usage_revert_rate = 0
l4_session_revert_duration = 0
l4_session_usage_grace_period = 0
l4_session_usage_log = 0
min_active_member = 0
min_active_member_action = "string"
reset_on_server_selection_fail = 0
priority_affinity = 0
reset_priority_affinity = 0
backup_server_event_log = 0
strict_select = 0
stats_data_action = "string"
extended_stats = 0
traffic_replication_mirror = 0
traffic_replication_mirror_da_repl = 0
traffic_replication_mirror_ip_repl = 0
traffic_replication_mirror_sa_da_repl = 0
traffic_replication_mirror_sa_repl = 0
health_check = "string"
shared_partition_svcgrp_health_check = 0
svcgrp_health_check_shared = "string"
health_check_disable = 0
priorities {   
	priority =  0 
	priority_action =  "string" 
	}
sample_rsp_time = 0
rpt_ext_server = 0
report_delay = 0
top_slowest = 0
top_fastest = 0
persist_scoring = "string"
uuid = "string"
user_tag = "string"
sampling-enable {   
	counters1 =  "string" 
	}
reset {  
 	auto_switch =  0 
	}
member-list {   
	name =  "string" 
	port =  0 
	fqdn_name =  "string" 
	resolve_as =  "string" 
	host =  "string" 
	server_ipv6_addr =  "string" 
	member_state =  "string" 
	member_stats_data_disable =  0 
	member_template =  "string" 
	member_priority =  0 
	uuid =  "string" 
	user_tag =  "string" 
sampling-enable {   
	counters1 =  "string" 
	}
	}
 
}

```

## Argument Reference

* `service-group` - Service Group
* `name` - Member name
* `protocol` - 'tcp': TCP LB service; 'udp': UDP LB service;
* `template-port` - Port template (Port template name)
* `template-server` - Server template (Server template name)
* `template-policy` - Policy template (Policy template name)
* `shared-partition-policy-template` - Reference a policy template from shared partition
* `template-policy-shared` - Policy template
* `lb-method` - 'dst-ip-hash': Load-balancing based on only Dst IP and Port hash; 'dst-ip-only-hash': Load-balancing based on only Dst IP hash; 'fastest-response': Fastest response time on service port level; 'least-request': Least request on service port level; 'src-ip-hash': Load-balancing based on only Src IP and Port hash; 'src-ip-only-hash': Load-balancing based on only Src IP hash; 'weighted-rr': Weighted round robin on server level; 'service-weighted-rr': Weighted round robin on service port level; 'round-robin': Round robin on server level; 'round-robin-strict': Strict mode round robin on server level; 'odd-even-hash': odd/even hash based of client src-ip;
* `lc-method` - 'least-connection': Least connection on server level; 'service-least-connection': Least connection on service port level; 'weighted-least-connection': Weighted least connection on server level; 'service-weighted-least-connection': Weighted least connection on service port level;
* `stateless-lb-method` - 'stateless-dst-ip-hash': Stateless load-balancing based on Dst IP and Dst port hash; 'stateless-per-pkt-round-robin': Stateless load-balancing using per-packet round-robin; 'stateless-src-dst-ip-hash': Stateless load-balancing based on IP and port hash for both Src and Dst; 'stateless-src-dst-ip-only-hash': Stateless load-balancing based on only IP hash for both Src and Dst; 'stateless-src-ip-hash': Stateless load-balancing based on Src IP and Src port hash; 'stateless-src-ip-only-hash': Stateless load-balancing based on only Src IP hash; 'stateless-per-pkt-weighted-rr': Stateless load-balancing using per-packet weighted round robin on server level; 'stateless-per-pkt-service-weighted-rr': Stateless load-balancing using per-packet weighted round robin on service port level;
* `llb-method` - 'next-hop-link': Server selection w/ link probe template on service port level;
* `link-probe-template` - Link Probe template (Link Probe template name)
* `lclb-method` - 'link-cost-load-balance': Link cost load balance;
* `pseudo-round-robin` - PRR, select the oldest node for sub-select
* `stateless-auto-switch` - Enable auto stateless method
* `stateless-lb-method2` - 'stateless-dst-ip-hash': Stateless load-balancing based on Dst IP and Dst port hash; 'stateless-per-pkt-round-robin': Stateless load-balancing using per-packet round-robin; 'stateless-src-dst-ip-hash': Stateless load-balancing based on IP and port hash for both Src and Dst; 'stateless-src-dst-ip-only-hash': Stateless load-balancing based on only IP hash for both Src and Dst; 'stateless-src-ip-hash': Stateless load-balancing based on Src IP and Src port hash; 'stateless-src-ip-only-hash': Stateless load-balancing based on only Src IP hash; 'stateless-per-pkt-weighted-rr': Stateless load-balancing using per-packet weighted round robin on server level; 'stateless-per-pkt-service-weighted-rr': Stateless load-balancing using per-packet weighted round robin on service port level;
* `conn-rate` - Dynamically enable stateless method by conn-rate (Rate to trigger stateless method(conn/sec))
* `conn-rate-duration` - Period that trigger condition consistently happens(seconds)
* `conn-revert-rate` - Rate to revert to statelful method (conn/sec)
* `conn-rate-revert-duration` - Period that revert condition consistently happens(seconds)
* `conn-rate-grace-period` - Define the grace period during transition (Define the grace period during transition(seconds))
* `conn-rate-log` - Send log if transition happens
* `l4-session-usage` - Dynamically enable stateless method by session usage (Usage to trigger stateless method)
* `l4-session-usage-duration` - Period that trigger condition consistently happens(seconds)
* `l4-session-usage-revert-rate` - Usage to revert to statelful method
* `l4-session-revert-duration` - Period that revert condition consistently happens(seconds)
* `l4-session-usage-grace-period` - Define the grace period during transition (Define the grace period during transition(seconds))
* `l4-session-usage-log` - Send log if transition happens
* `min-active-member` - Minimum Active Member Per Priority (Minimum Active Member before Action)
* `min-active-member-action` - 'dynamic-priority': dynamic change member priority to met the min-active-member requirement; 'skip-pri-set': Skip Current Priority Set If Min not met;
* `reset-on-server-selection-fail` - Send reset to client if server selection fails
* `priority-affinity` - Priority affinity. Persist to the same priority if possible.
* `reset-priority-affinity` - Reset
* `backup-server-event-log` - Send log info on back up server events
* `strict-select` - strict selection
* `stats-data-action` - 'stats-data-enable': Enable statistical data collection for service group; 'stats-data-disable': Disable statistical data collection for service group;
* `extended-stats` - Enable extended statistics on service group
* `traffic-replication-mirror` - Mirror Bi-directional Packet
* `traffic-replication-mirror-da-repl` - Replace Destination MAC
* `traffic-replication-mirror-ip-repl` - Replaces IP with server-IP
* `traffic-replication-mirror-sa-da-repl` - Replace Source MAC and Destination MAC
* `traffic-replication-mirror-sa-repl` - Replace Source MAC
* `health-check` - Health Check (Monitor Name)
* `shared-partition-svcgrp-health-check` - Reference a health-check from shared partition
* `svcgrp-health-check-shared` - Health Check (Monitor Name)
* `health-check-disable` - Disable health check
* `priority` - Priority option. Define different action for each priority node. (Priority in the Group)
* `priority-action` - 'drop': Drop request when all priority nodes fail; 'drop-if-exceed-limit': Drop request when connection over limit; 'proceed': Proceed to next priority when all priority nodes fail(default); 'reset': Send client reset when all priority nodes fail; 'reset-if-exceed-limit': Send client reset when connection over limit;
* `sample-rsp-time` - sample server response time
* `rpt-ext-server` - Report top 10 fastest/slowest servers
* `report-delay` - Reporting frequency (in minutes)
* `top-slowest` - Report top 10 slowest servers
* `top-fastest` - Report top 10 fastest servers
* `persist-scoring` - 'global': Use Global Configuration; 'enable': Enable persist-scoring; 'disable': Disable persist-scoring;
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `counters1` - 'all': all; 'total_fwd_bytes': Bytes processed in forward direction; 'total_fwd_pkts': Packets processed in forward direction; 'total_rev_bytes': Bytes processed in reverse direction; 'total_rev_pkts': Packets processed in reverse direction; 'total_conn': Total established connections; 'total_rev_pkts_inspected': Total reverse packets inspected; 'total_rev_pkts_inspected_status_code_2xx': Total reverse packets inspected status code 2xx; 'total_rev_pkts_inspected_status_code_non_5xx': Total reverse packets inspected status code non 5xx; 'curr_req': Current requests; 'total_req': Total requests; 'total_req_succ': Total requests successful; 'peak_conn': Peak connections; 'response_time': Response time; 'fastest_rsp_time': Fastest response time; 'slowest_rsp_time': Slowest response time; 'curr_ssl_conn': Current SSL connections; 'total_ssl_conn': Total SSL connections; 'curr_conn_overflow': Current connection counter overflow count; 'state_flaps': State flaps count;
* `auto-switch` - Reset auto stateless state
* `port` - Port number
* `fqdn-name` - Server hostname - Not applicable if real server is already defined
* `resolve-as` - 'resolve-to-ipv4': Use A Query only to resolve FQDN; 'resolve-to-ipv6': Use AAAA Query only to resolve FQDN; 'resolve-to-ipv4-and-ipv6': Use A as well as AAAA Query to resolve FQDN;
* `host` - IP Address - Not applicable if real server is already defined
* `server-ipv6-addr` - IPV6 Address - Not applicable if real server is already defined
* `member-state` - 'enable': Enable member service port; 'disable': Disable member service port; 'disable-with-health-check': disable member service port, but health check work;
* `member-stats-data-disable` - Disable statistical data collection
* `member-template` - Real server port template (Real server port template name)
* `member-priority` - Priority of Port in the Group (Priority of Port in the Group, default is 1)

