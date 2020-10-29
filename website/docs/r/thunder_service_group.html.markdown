---
layout: "thunder"
page_title: "thunder: thunder_service_group"
sidebar_current: "docs-thunder-resource-service-group"
description: |-
	Provides details about thunder service group resource for A10
---

# thunder\_service\_group

`thunder_service_group` Provides details about thunder slb service group
## Example Usage


```hcl
provider "thunder" {
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_service_group" "Service_Group_Test" {

conn_rate = 0
reset_on_server_selection_fail = 0
health_check_disable = 0
protocol = "string"
llb_method = "string"
traffic_replication_mirror_ip_repl = 0
reset_priority_affinity = 0
priorities {   
        priority =  0 
        priority_action =  "string" 
        }
min_active_member = 0
member-list {   
        member_priority =  0 
        uuid =  "string" 
        fqdn_name =  "string" 
        resolve_as =  "string" 
sampling-enable {   
        counters1 =  "string" 
        }
        member_template =  "string" 
        name =  "string" 
        host =  "string" 
        user_tag =  "string" 
        member_state =  "string" 
        server_ipv6_addr =  "string" 
        port =  0 
        member_stats_data_disable =  0 
        }
stats_data_action = "string"
traffic_replication_mirror_da_repl = 0
template_policy_shared = "string"
rpt_ext_server = 0
template_port = "string"
conn_rate_grace_period = 0
l4_session_usage_duration = 0
uuid = "string"
backup_server_event_log = 0
lc_method = "string"
pseudo_round_robin = 0
shared_partition_policy_template = 0
l4_session_usage_revert_rate = 0
shared_partition_svcgrp_health_check = 0
svcgrp_health_check_shared = "string"
traffic_replication_mirror = 0
l4_session_revert_duration = 0
link_probe_template = "string"
traffic_replication_mirror_sa_da_repl = 0
lb_method = "string"
stateless_auto_switch = 0
min_active_member_action = "string"
l4_session_usage = 0
extended_stats = 0
conn_rate_revert_duration = 0
strict_select = 0
name = "string"
reset {  
        auto_switch =  0 
        }
traffic_replication_mirror_sa_repl = 0
report_delay = 0
conn_rate_log = 0
l4_session_usage_log = 0
conn_rate_duration = 0
stateless_lb_method = "string"
template_policy = "string"
stateless_lb_method2 = "string"
user_tag = "string"
sample_rsp_time = 0
sampling-enable {   
        counters1 =  "string" 
        }
top_fastest = 0
conn_revert_rate = 0
l4_session_usage_grace_period = 0
priority_affinity = 0
top_slowest = 0
health_check = "string"
 
}

```

## Argument Reference

* `backup_server_event_log` - Send log info on back up server events
* `conn_rate` - Dynamically enable stateless method by conn-rate (Rate to trigger stateless method(conn/sec))
* `conn_rate_duration` - Period that trigger condition consistently happens(seconds)
* `conn_rate_grace_period` - Define the grace period during transition (Define the grace period during transition(seconds))
* `conn_rate_log` - Send log if transition happens
* `conn_rate_revert_duration` - Period that revert condition consistently happens(seconds)
* `conn_revert_rate` - Rate to revert to statelful method (conn/sec)
* `extended_stats` - Enable extended statistics on service group
* `health_check` - Health Check (Monitor Name)
* `health_check_disable` - Disable health check
* `l4_session_revert_duration` - Period that revert condition consistently happens(seconds)
* `l4_session_usage` - Dynamically enable stateless method by session usage (Usage to trigger stateless method)
* `l4_session_usage_duration` - Period that trigger condition consistently happens(seconds)
* `l4_session_usage_grace_period` - Define the grace period during transition (Define the grace period during transition(seconds))
* `l4_session_usage_log` - Send log if transition happens
* `l4_session_usage_revert_rate` - Usage to revert to statelful method
* `lb_method` - ‘dst-ip-hash’: Load-balancing based on only Dst IP and Port hash; ‘dst-ip-only-hash’: Load-balancing based on only Dst IP hash; ‘fastest-response’: Fastest response time on service port level; ‘least-request’: Least request on service port level; ‘src-ip-hash’: Load-balancing based on only Src IP and Port hash; ‘src-ip-only-hash’: Load-balancing based on only Src IP hash; ‘weighted-rr’: Weighted round robin on server level; ‘round-robin’: Round robin on server level; ‘round-robin-strict’: Strict mode round robin on server level; ‘odd-even-hash’: odd/even hash based of client src-ip;
* `lc_method` - ‘least-connection’: Least connection on server level; ‘service-least-connection’: Least connection on service port level; ‘weighted-least-connection’: Weighted least connection on server level; ‘service-weighted-least-connection’: Weighted least connection on service port level;
* `min_active_member` - Minimum Active Member Per Priority (Minimum Active Member before Action)
* `min_active_member_action` - ‘dynamic-priority’: dynamic change member priority to met the min-active-member requirement; ‘skip-pri-set’: Skip Current Priority Set If Min not met;
* `name` - Member name
* `priority_affinity` - Priority affinity. Persist to the same priority if possible.
* `protocol` - ‘tcp’: TCP LB service; ‘udp’: UDP LB service;
* `pseudo_round_robin` - PRR, select the oldest node for sub-select
* `report_delay` - Reporting frequency (in minutes)
* `reset_on_server_selection_fail` - Send reset to client if server selection fails
* `reset_priority_affinity` - Reset
* `rpt_ext_server` - Report top 10 fastest/slowest servers
* `sample_rsp_time` - sample server response time
* `shared_partition_policy_template` - Reference a policy template from shared partition
* `shared_partition_svcgrp_health_check` - Reference a health-check from shared partition
* `stateless_auto_switch` - Enable auto stateless method
* `stateless_lb_method` - ‘stateless-dst-ip-hash’: Stateless load-balancing based on Dst IP and Dst port hash; ‘stateless-per-pkt-round-robin’: Stateless load-balancing using per-packet round-robin; ‘stateless-src-dst-ip-hash’: Stateless load-balancing based on IP and port hash for both Src and Dst; ‘stateless-src-dst-ip-only-hash’: Stateless load-balancing based on only IP hash for both Src and Dst; ‘stateless-src-ip-hash’: Stateless load-balancing based on Src IP and Src port hash; ‘stateless-src-ip-only-hash’: Stateless load-balancing based on only Src IP hash;
* `stateless_lb_method2` - ‘stateless-dst-ip-hash’: Stateless load-balancing based on Dst IP and Dst port hash; ‘stateless-per-pkt-round-robin’: Stateless load-balancing using per-packet round-robin; ‘stateless-src-dst-ip-hash’: Stateless load-balancing based on IP and port hash for both Src and Dst; ‘stateless-src-dst-ip-only-hash’: Stateless load-balancing based on only IP hash for both Src and Dst; ‘stateless-src-ip-hash’: Stateless load-balancing based on Src IP and Src port hash; ‘stateless-src-ip-only-hash’: Stateless load-balancing based on only Src IP hash;
* `stats_data_action` - ‘stats-data-enable’: Enable statistical data collection for service group; ‘stats-data-disable’: Disable statistical data collection for service group;
* `strict_select` - strict selection
* `svcgrp_health_check_shared` - Health Check (Monitor Name)
* `template_policy` - Policy template (Policy template name)
* `template_policy_shared` - Policy template
* `template_port` - Port template (Port template name)
* `template_server` - Server template (Server template name)
* `top_fastest` - Report top 10 fastest servers
* `top_slowest` - Report top 10 slowest servers
* `traffic_replication_mirror` - Mirror Bi-directional Packet
* `traffic_replication_mirror_da_repl` - Replace Destination MAC
* `traffic_replication_mirror_ip_repl` - Replaces IP with server-IP
* `traffic_replication_mirror_sa_da_repl` - Replace Source MAC and Destination MAC
* `traffic_replication_mirror_sa_repl` - Replace Source MAC
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `fqdn_name` - Server hostname - Not applicable if real server is already defined
* `host` - IP Address - Not applicable if real server is already defined
* `member_priority` - Priority of Port in the Group (Priority of Port in the Group, default is 1)
* `member_state` - ‘enable’: Enable member service port; ‘disable’: Disable member service port; ‘disable-with-health-check’: disable member service port, but health check work;
* `member_stats_data_disable` - Disable statistical data collection
* `member_template` - Real server port template (Real server port template name)
* `port` - Port number
* `resolve_as` - ‘resolve-to-ipv4’: Use A Query only to resolve FQDN; ‘resolve-to-ipv6’: Use AAAA Query only to resolve FQDN; ‘resolve-to-ipv4-and-ipv6’: Use A as well as AAAA Query to resolve FQDN;
* `server_ipv6_addr` - IPV6 Address - Not applicable if real server is already defined
* `counters1` - ‘all’: all; ‘server_selection_fail_drop’: Service selection fail drop; ‘server_selection_fail_reset’: Service selection fail reset; ‘service_peak_conn’: Service peak connection;
* `auto_switch` - Reset auto stateless state
* `priority` - Priority option. Define different action for each priority node. (Priority in the Group)
* `priority_action` - ‘drop’: Drop request when all priority nodes fail; ‘drop-if-exceed-limit’: Drop request when connection over limit; ‘proceed’: Proceed to next priority when all priority nodes fail(default); ‘reset’: Send client reset when all priority nodes fail; ‘reset-if-exceed-limit’: Send client reset when connection over limit;
