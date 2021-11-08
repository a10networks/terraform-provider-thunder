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
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_fw_template_logging" "resourceFwTemplateLoggingTest" {
	name = "string"
resolution = "string"
include_dest_fqdn = 0
merged_style = 0
log {  
 	http_requests =  "string" 
	}
include_radius_attribute {  
 attr-cfg {   
	attr =  "string" 
	attr_event =  "string" 
	}
	no_quote =  0 
	framed_ipv6_prefix =  0 
	prefix_length =  "string" 
	insert_if_not_existing =  0 
	zero_in_custom_attr =  0 
	}
include_http {  
 header-cfg {   
	http_header =  "string" 
	max_length =  0 
	custom_header_name =  "string" 
	custom_max_length =  0 
	}
	l4_session_info =  0 
	method =  0 
	request_number =  0 
	file_extension =  0 
	}
rule {  
 rule_http_requests {  
 dest-port {   
	dest_port_number =  0 
	include_byte_count =  0 
	}
	log_every_http_request =  0 
	max_url_len =  0 
	include_all_headers =  0 
	disable_sequence_check =  0 
	}
	}
facility = "string"
severity = "string"
format = "string"
service_group = "string"
uuid = "string"
user_tag = "string"
source_address {  
 	ip =  "string" 
	ipv6 =  "string" 
	uuid =  "string" 
	}
 
}

```

## Argument Reference

* `logging` - Logging Template
* `name` - Logging Template Name
* `resolution` - 'seconds': Logging timestamp resolution in seconds (default); '10-milliseconds': Logging timestamp resolution in 10s of milli-seconds;
* `include-dest-fqdn` - Include destination FQDN string
* `merged-style` - Merge creation and deletion of session logs to one
* `http-requests` - 'host': Log the HTTP Host Header; 'url': Log the HTTP Request URL;
* `attr` - 'imei': Include IMEI; 'imsi': Include IMSI; 'msisdn': Include MSISDN; 'custom1': Customized attribute 1; 'custom2': Customized attribute 2; 'custom3': Customized attribute 3; 'custom4': Customized attribute 4; 'custom5': Customized attribute 5; 'custom6': Customized attribute 6;
* `attr-event` - 'http-requests': Include in HTTP request logs; 'sessions': Include in session logs; 'limit-policy': Include in limit policy logs;
* `no-quote` - No quotation marks for RADIUS attributes in logs
* `framed-ipv6-prefix` - Include radius attributes for the prefix
* `prefix-length` - '32': Prefix length 32; '48': Prefix length 48; '64': Prefix length 64; '80': Prefix length 80; '96': Prefix length 96; '112': Prefix length 112;
* `insert-if-not-existing` - Configure what string is to be inserted for custom RADIUS attributes
* `zero-in-custom-attr` - Insert 0000 for standard and custom attributes in log string
* `http-header` - 'cookie': Log HTTP Cookie Header; 'referer': Log HTTP Referer Header; 'user-agent': Log HTTP User-Agent Header; 'header1': Log HTTP Header 1; 'header2': Log HTTP Header 2; 'header3': Log HTTP Header 3;
* `max-length` - Max length for a HTTP request log (Max header length (Default: 100 char))
* `custom-header-name` - Header name
* `custom-max-length` - Max length for a HTTP request log (Max header length (Default: 100 char))
* `l4-session-info` - Log the L4 session information of the HTTP request
* `method` - Log the HTTP Request Method
* `request-number` - HTTP Request Number
* `file-extension` - HTTP file extension
* `include-byte-count` - Include the byte count of HTTP Request/Response in FW session deletion logs
* `log-every-http-request` - Log every HTTP request in an HTTP 1.1 session (Default: Log the first HTTP request in a session)
* `max-url-len` - Max length of URL log (Max URL length (Default: 128 char))
* `include-all-headers` - Include all configured headers despite of absence in HTTP request
* `disable-sequence-check` - Disable http packet sequence check and don't drop out of order packets
* `facility` - 'kernel': 0: Kernel; 'user': 1: User-level; 'mail': 2: Mail; 'daemon': 3: System daemons; 'security-authorization': 4: Security/authorization; 'syslog': 5: Syslog internal; 'line-printer': 6: Line printer; 'news': 7: Network news; 'uucp': 8: UUCP subsystem; 'cron': 9: Time-related; 'security-authorization-private': 10: Private security/authorization; 'ftp': 11: FTP; 'ntp': 12: NTP; 'audit': 13: Audit; 'alert': 14: Alert; 'clock': 15: Clock-related; 'local0': 16: Local use 0; 'local1': 17: Local use 1; 'local2': 18: Local use 2; 'local3': 19: Local use 3; 'local4': 20: Local use 4; 'local5': 21: Local use 5; 'local6': 22: Local use 6; 'local7': 23: Local use 7;
* `severity` - 'emergency': 0: Emergency; 'alert': 1: Alert; 'critical': 2: Critical; 'error': 3: Error; 'warning': 4: Warning; 'notice': 5: Notice; 'informational': 6: Informational; 'debug': 7: Debug;
* `format` - 'ascii': A10 Text logging format (ASCII); 'cef': Common Event Format for logging (default);
* `service-group` - Bind a Service Group to the logging template (Service Group Name)
* `uuid` - uuid of the object
* `user-tag` - Customized tag
* `ip` - Specify source IP address
* `ipv6` - Specify source IPv6 address

