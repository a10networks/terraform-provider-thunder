---
layout: "thunder"
page_title: "thunder: thunder_fw_template_logging"
sidebar_current: "docs-thunder-resource-fw-template-logging"
description: |-
	Provides details about thunder fw template logging resource for A10
---

# thunder\_fw\_template\_logging

`thunder_fw_template_logging` Provides details about thunder fw template logging
## Example Usage


```hcl
provider "thunder" {
  address  = "192.0.2.65"
  username = "admin"
  password = "admin"
}

resource "thunder_fw_template_logging" "FwTemplateTest" {
	name = "a"
	user_tag = "a" 
}
```

## Argument Reference

* `facility` - ‘kernel’: 0: Kernel; ‘user’: 1: User-level; ‘mail’: 2: Mail; ‘daemon’: 3: System daemons; ‘security-authorization’: 4: Security/authorization; ‘syslog’: 5: Syslog internal; ‘line-printer’: 6: Line printer; ‘news’: 7: Network news; ‘uucp’: 8: UUCP subsystem; ‘cron’: 9: Time-related; ‘security-authorization-private’: 10: Private security/authorization; ‘ftp’: 11: FTP; ‘ntp’: 12: NTP; ‘audit’: 13: Audit; ‘alert’: 14: Alert; ‘clock’: 15: Clock-related; ‘local0’: 16: Local use 0; ‘local1’: 17: Local use 1; ‘local2’: 18: Local use 2; ‘local3’: 19: Local use 3; ‘local4’: 20: Local use 4; ‘local5’: 21: Local use 5; ‘local6’: 22: Local use 6; ‘local7’: 23: Local use 7;
* `format` - ‘ascii’: A10 Text logging format (ASCII); ‘cef’: Common Event Format for logging (default);
* `include_dest_fqdn` - Include destination FQDN string
* `merged_style` - Merge creation and deletion of session logs to one
* `name` - Logging Template Name
* `resolution` - ‘seconds’: Logging timestamp resolution in seconds (default); ‘10-milliseconds’: Logging timestamp resolution in 10s of milli-seconds;
* `service_group` - Bind a Service Group to the logging template (Service Group Name)
* `severity` - ‘emergency’: 0: Emergency; ‘alert’: 1: Alert; ‘critical’: 2: Critical; ‘error’: 3: Error; ‘warning’: 4: Warning; ‘notice’: 5: Notice; ‘informational’: 6: Informational; ‘debug’: 7: Debug;
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `ip` - Specify source IP address
* `ipv6` - Specify source IPv6 address
* `framed_ipv6_prefix` - Include radius attributes for the prefix
* `insert_if_not_existing` - Configure what string is to be inserted for custom RADIUS attributes
* `no_quote` - No quotation marks for RADIUS attributes in logs
* `prefix_length` - ‘32’: Prefix length 32; ‘48’: Prefix length 48; ‘64’: Prefix length 64; ‘80’: Prefix length 80; ‘96’: Prefix length 96; ‘112’: Prefix length 112;
* `zero_in_custom_attr` - Insert 0000 for standard and custom attributes in log string
* `attr` - ‘imei’: Include IMEI; ‘imsi’: Include IMSI; ‘msisdn’: Include MSISDN; ‘custom1’: Customized attribute 1; ‘custom2’: Customized attribute 2; ‘custom3’: Customized attribute 3;
* `attr_event` - ‘http-requests’: Include in HTTP request logs; ‘sessions’: Include in session logs;
* `disable_sequence_check` - Disable http packet sequence check and don’t drop out of order packets
* `include_all_headers` - Include all configured headers despite of absence in HTTP request
* `log_every_http_request` - Log every HTTP request in an HTTP 1.1 session (Default: Log the first HTTP request in a session)
* `max_url_len` - Max length of URL log (Max URL length (Default: 100 char))
* `include_byte_count` - Include the byte count of HTTP Request/Response in FW session deletion logs
* `file_extension` - HTTP file extension
* `l4_session_info` - Log the L4 session information of the HTTP request
* `method` - Log the HTTP Request Method
* `request_number` - HTTP Request Number
* `custom_header_name` - Header name
* `custom_max_length` - Max length for a HTTP request log (Max header length (Default: 100 char))
* `http_header` - ‘cookie’: Log HTTP Cookie Header; ‘referer’: Log HTTP Referer Header; ‘user-agent’: Log HTTP User-Agent Header; ‘header1’: Log HTTP Header 1; ‘header2’: Log HTTP Header 2; ‘header3’: Log HTTP Header 3;
* `max_length` - Max length for a HTTP request log (Max header length (Default: 100 char))
* `http_requests` - ‘host’: Log the HTTP Host Header; ‘url’: Log the HTTP Request URL;

