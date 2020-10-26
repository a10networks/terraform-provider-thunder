---
layout: "thunder"
page_title: "thunder: thunder_slb_health_stat"
sidebar_current: "docs-thunder-resource-slb-health-stat"
description: |-
    Provides details about thunder SLB health-stat resource for A10
---

# thunder\_slb\_health_stat

`thunder_slb_health_stat` Provides details about thunder SLB health-stat
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_health_stat" "health_stat" {
	sampling_enable {
		counters1 = "all"
	}
}
```

## Argument Reference

* `sampling_enable` - 
	* `counters1` - 'all’: all; 'num_burst’: Number of burst; 'max_jiffie’: Maximum number of jiffies; 'min_jiffie’: Minimum number of jiffies; 'avg_jiffie’: Average number of jiffies; 'open_socket’: Number of open sockets; 'open_socket_failed’: Number of failed open sockets; 'close_socket’: Number of closed sockets; 'connect_failed’: Number of failed connections; 'send_packet’: Number of packets sent; 'send_packet_failed’: Number of packet send failures; 'recv_packet’: Number of received packets; 'recv_packet_failed’: Number of failed packet receives; 'retry_times’: Retry times; 'timeout’: Timouet value; 'unexpected_error’: Number of unexpected errors; 'conn_imdt_succ’: Number of connection immediete success; 'sock_close_before_17’: Number of sockets closed before l7; 'sock_close_without_notify’: Number of sockets closed without notify; 'curr_health_rate’: Current health rate; 'ext_health_rate’: External health rate; 'ext_health_rate_val’: External health rate value; 'total_number’: Total number; 'status_up’: Number of status ups; 'status_down’: Number of status downs; 'status_unkn’: Number of status unknowns; 'status_other’: Number of other status; 'running_time’: Running time; 'config_health_rate’: Config health rate;


