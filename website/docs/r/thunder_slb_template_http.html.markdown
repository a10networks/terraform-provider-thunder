---
layout: "thunder"
page_title: "thunder: thunder_slb_template_http"
sidebar_current: "docs-thunder-resource-slb-template-http"
description: |-
    Provides details about thunder slb template http resource for A10
---

# thunder\_slb\_template\_http

`thunder_slb_template_http` Provides details about thunder slb template http
## Example Usage


```hcl
provider "thunder" {
  address  = var.address
  username = var.username
  password = var.password
}


resource "thunder_slb_template_http" "resourceSlbTemplateHttpTest" {
	name = "string"
compression_auto_disable_on_high_cpu = 0
compression-content-type {   
	content_type =  "string" 
	}
compression_enable = 0
compression-exclude-content-type {   
	exclude_content_type =  "string" 
	}
compression-exclude-uri {   
	exclude_uri =  "string" 
	}
compression_keep_accept_encoding = 0
compression_keep_accept_encoding_enable = 0
compression_level = 0
compression_minimum_content_length = 0
default_charset = "string"
max_concurrent_streams = 0
frame_limit = 0
failover_url = "string"
host-switching {   
	host_switching_type =  "string" 
	host_match_string =  "string" 
	host_service_group =  "string" 
	}
insert_client_ip = 0
insert_client_ip_header_name = "string"
client_ip_hdr_replace = 0
insert_client_port = 0
insert_client_port_header_name = "string"
client_port_hdr_replace = 0
log_retry = 0
non_http_bypass = 0
bypass_sg = "string"
redirect = 0
rd_simple_loc = "string"
rd_secure = 0
rd_port = 0
rd_resp_code = "string"
redirect_rewrite {  
 match-list {   
	redirect_match =  "string" 
	rewrite_to =  "string" 
	}
	redirect_secure =  0 
	redirect_secure_port =  0 
	}
request-header-erase-list {   
	request_header_erase =  "string" 
	}
request-header-insert-list {   
	request_header_insert =  "string" 
	request_header_insert_type =  "string" 
	}
response-content-replace-list {   
	response_content_replace =  "string" 
	response_new_string =  "string" 
	}
response-header-erase-list {   
	response_header_erase =  "string" 
	}
response-header-insert-list {   
	response_header_insert =  "string" 
	response_header_insert_type =  "string" 
	}
request_timeout = 0
retry_on_5xx = 0
retry_on_5xx_val = 0
retry_on_5xx_per_req = 0
retry_on_5xx_per_req_val = 0
strict_transaction_switch = 0
template {  
 	logging =  "string" 
	}
term_11client_hdr_conn_close = 0
persist_on_401 = 0
cont_100_wait_for_req_complete = 0
url_hash_persist = 0
url_hash_offset = 0
url_hash_first = 0
url_hash_last = 0
use_server_status = 0
url-switching {   
	url_switching_type =  "string" 
	url_match_string =  "string" 
	url_service_group =  "string" 
	}
req_hdr_wait_time = 0
req_hdr_wait_time_val = 0
request_line_case_insensitive = 0
keep_client_alive = 0
cookie_format = "string"
prefix = "string"
cookie_samesite = "string"
uuid = "string"
user_tag = "string"
 
}

```

## Argument Reference

* `http` - HTTP
* `name` - HTTP Template Name
* `compression-auto-disable-on-high-cpu` - Auto-disable software compression on high cpu usage (Disable compression if cpu usage is above threshold. Default is off.)
* `content-type` - Compression content-type
* `compression-enable` - Enable Compression
* `exclude-content-type` - Compression exclude content-type (Compression exclude content type)
* `exclude-uri` - Compression exclude uri
* `compression-keep-accept-encoding` - Keep accept encoding
* `compression-keep-accept-encoding-enable` - Enable Server Accept Encoding
* `compression-level` - compression level, default 1 (compression level value, default is 1)
* `compression-minimum-content-length` - Minimum Content Length (Minimum content length for compression in bytes. Default is 120.)
* `default-charset` - 'iso-8859-1': Use ISO-8859-1 as the default charset; 'utf-8': Use UTF-8 as the default charset; 'us-ascii': Use US-ASCII as the default charset;
* `max-concurrent-streams` - (http2 only) Max concurrent streams, default 100
* `frame-limit` - Limit the number of CONTINUATION, PING, PRIORITY, RESET, SETTINGS and empty frames in one HTTP2 connection, default 10000
* `failover-url` - Failover to this URL (Failover URL Name)
* `host-switching-type` - 'contains': Select service group if hostname contains another string; 'ends-with': Select service group if hostname ends with another string; 'equals': Select service group if hostname equals another string; 'starts-with': Select service group if hostname starts with another string; 'regex-match': Select service group if URL string matches with regular expression; 'host-hits-enable': Enables Host Hits counters;
* `host-match-string` - Hostname String
* `host-service-group` - Create a Service Group comprising Servers (Service Group Name)
* `insert-client-ip` - Insert Client IP address into HTTP header
* `insert-client-ip-header-name` - HTTP Header Name for inserting Client IP
* `client-ip-hdr-replace` - Replace the existing header
* `insert-client-port` - Insert Client Port address into HTTP header
* `insert-client-port-header-name` - HTTP Header Name for inserting Client Port
* `client-port-hdr-replace` - Replace the existing header
* `log-retry` - log when HTTP request retry
* `non-http-bypass` - Bypass non-http traffic instead of dropping
* `bypass-sg` - Select service group for non-http traffic (Service Group Name)
* `redirect` - Automatically send a redirect response
* `rd-simple-loc` - Redirect location tag absolute URI string
* `rd-secure` - Use HTTPS
* `rd-port` - Port (Port Number)
* `rd-resp-code` - '301': Moved Permanently; '302': Found; '303': See Other; '307': Temporary Redirect;
* `redirect-match` - URL Matching (Pattern URL String)
* `rewrite-to` - Rewrite to Destination URL String
* `redirect-secure` - Use HTTPS
* `redirect-secure-port` - Port (Port Number)
* `request-header-erase` - Erase a header from HTTP request (Header Name)
* `request-header-insert` - Insert a header into HTTP request (Header Content (Format: "[name]:[value]"))
* `request-header-insert-type` - 'insert-if-not-exist': Only insert the header when it does not exist; 'insert-always': Always insert the header even when there is a header with the same name;
* `response-content-replace` - replace the data from HTTP response content (String in the http content need to be replaced)
* `response-new-string` - String will be in the http content
* `response-header-erase` - Erase a header from HTTP response (Header Name)
* `response-header-insert` - Insert a header into HTTP response (Header Content (Format: "[name]:[value]"))
* `response-header-insert-type` - 'insert-if-not-exist': Only insert the header when it does not exist; 'insert-always': Always insert the header even when there is a header with the same name;
* `request-timeout` - Request timeout if response not received (timeout in seconds)
* `retry-on-5xx` - Retry http request on HTTP 5xx code and request timeout
* `retry-on-5xx-val` - Number of times to retry (default is 3)
* `retry-on-5xx-per-req` - Retry http request on HTTP 5xx code for each request
* `retry-on-5xx-per-req-val` - Number of times to retry (default is 3)
* `strict-transaction-switch` - Force server selection on every HTTP request
* `logging` - Logging template (Logging Config name)
* `term-11client-hdr-conn-close` - Terminate HTTP 1.1 client when req has Connection: close
* `persist-on-401` - Persist to the same server if the response code is 401
* `100-cont-wait-for-req-complete` - When REQ has Expect 100 and response is not 100, then wait for whole request to be sent
* `url-hash-persist` - Use URL's hash value to select server
* `url-hash-offset` - Skip part of URL to calculate hash value (Offset of the URL string)
* `url-hash-first` - Use the begining part of URL to calculate hash value (URL string length to calculate hash value)
* `url-hash-last` - Use the end part of URL to calculate hash value (URL string length to calculate hash value)
* `use-server-status` - Use Server-Status header to do URL hashing
* `url-switching-type` - 'contains': Select service group if URL string contains another string; 'ends-with': Select service group if URL string ends with another string; 'equals': Select service group if URL string equals another string; 'starts-with': Select service group if URL string starts with another string; 'regex-match': Select service group if URL string matches with regular expression; 'url-case-insensitive': Case insensitive URL switching; 'url-hits-enable': Enables URL Hits;
* `url-match-string` - URL String
* `url-service-group` - Create a Service Group comprising Servers (Service Group Name)
* `req-hdr-wait-time` - HTTP request header wait time before abort connection
* `req-hdr-wait-time-val` - Number of seconds wait for client request header (default is 7)
* `request-line-case-insensitive` - Parse http request line as case insensitive
* `keep-client-alive` - Keep client alive
* `cookie-format` - 'rfc6265': Follow rfc6265;
* `prefix` - 'host': the cookie will have been set with a Secure attribute, a Path attribute with a value of /, and no Domain attribute; 'secure': the cookie will have been set with a Secure attribute; 'check': check server prefix and enforce prefix format;
* `cookie-samesite` - 'none': none; 'lax': lax; 'strict': strict;
* `uuid` - uuid of the object
* `user-tag` - Customized tag

