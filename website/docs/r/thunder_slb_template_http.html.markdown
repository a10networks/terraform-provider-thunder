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
  address  = "129.213.82.65"
  username = "admin"
  password = "admin"
}

resource "thunder_slb_template_http" "Slb_Template_Http_Test" {

keep_client_alive = 0
compression_auto_disable_on_high_cpu = 0
req_hdr_wait_time = 0
compression-exclude-uri {   
        exclude_uri =  "string" 
        }
compression_enable = 0
compression_keep_accept_encoding = 0
failover_url = "string"
redirect_rewrite {  
        redirect_secure_port =  0 
        redirect_secure =  0 
match-list {   
        rewrite_to =  "string" 
        redirect_match =  "string" 
        }
        }
request-header-erase-list {   
        request_header_erase =  "string" 
        }
rd_port = 0
default_charset = "string"
url_hash_last = 0
client_ip_hdr_replace = 0
use_server_status = 0
req_hdr_wait_time_val = 0
response-header-insert-list {   
        response_header_insert_type =  "string" 
        response_header_insert =  "string" 
        }
persist_on_401 = 0
redirect = 0
insert_client_port = 0
retry_on_5xx_per_req_val = 0
url_hash_offset = 0
rd_simple_loc = "string"
log_retry = 0
non_http_bypass = 0
retry_on_5xx_per_req = 0
insert_client_ip = 0
template {  
        logging =  "string" 
        }
frame_limit = 0
url-switching {   
        url_service_group =  "string" 
        url_match_string =  "string" 
        url_switching_type =  "string" 
        }
insert_client_port_header_name = "string"
strict_transaction_switch = 0
response-content-replace-list {   
        response_new_string =  "string" 
        response_content_replace =  "string" 
        }
100_cont_wait_for_req_complete = 0
max_concurrent_streams = 0
request-header-insert-list {   
        request_header_insert =  "string" 
        request_header_insert_type =  "string" 
        }
request_timeout = 0
compression_minimum_content_length = 0
compression_level = 0
request_line_case_insensitive = 0
url_hash_persist = 0
response-header-erase-list {   
        response_header_erase =  "string" 
        }
uuid = "string"
bypass_sg = "string"
name = "string"
retry_on_5xx_val = 0
host-switching {   
        host_switching_type =  "string" 
        host_service_group =  "string" 
        host_match_string =  "string" 
        }
url_hash_first = 0
compression_keep_accept_encoding_enable = 0
user_tag = "string"
compression-content-type {   
        content_type =  "string" 
        }
client_port_hdr_replace = 0
insert_client_ip_header_name = "string"
rd_secure = 0
retry_on_5xx = 0
cookie_format = "string"
term_11client_hdr_conn_close = 0
compression-exclude-content-type {   
        exclude_content_type =  "string" 
        }
rd_resp_code = "string"
 
}
```

## Argument Reference

* `100_cont_wait_for_req_complete` - When REQ has Expect 100 and response is not 100, then wait for whole request to be sent
* `bypass_sg` - Select service group for non-http traffic (Service Group Name)
* `client_ip_hdr_replace` - Replace the existing header
* `client_port_hdr_replace` - Replace the existing header
* `compression_auto_disable_on_high_cpu` - Auto-disable software compression on high cpu usage (Disable compression if cpu usage is above threshold. Default is off.)
* `compression_enable` - Enable Compression
* `compression_keep_accept_encoding` - Keep accept encoding
* `compression_keep_accept_encoding_enable` - Enable Server Accept Encoding
* `compression_level` - compression level, default 1 (compression level value, default is 1)
* `compression_minimum_content_length` - Minimum Content Length (Minimum content length for compression in bytes. Default is 120.)
* `cookie_format` - ‘rfc6265’: Follow rfc6265;
* `failover_url` - Failover to this URL (Failover URL Name)
* `insert_client_ip` - Insert Client IP address into HTTP header
* `insert_client_ip_header_name` - HTTP Header Name for inserting Client IP
* `insert_client_port` - Insert Client Port address into HTTP header
* `insert_client_port_header_name` - HTTP Header Name for inserting Client Port
* `keep_client_alive` - Keep client alive
* `log_retry` - log when HTTP request retry
* `max_concurrent_streams` - (http2 only) Max concurrent streams, default 100
* `name` - HTTP Template Name
* `non_http_bypass` - Bypass non-http traffic instead of dropping
* `persist_on_401` - Persist to the same server if the response code is 401
* `rd_port` - Port (Port Number)
* `rd_resp_code` - ‘301’: Moved Permanently; ‘302’: Found; ‘303’: See Other; ‘307’: Temporary Redirect;
* `rd_secure` - Use HTTPS
* `rd_simple_loc` - Redirect location tag absolute URI string
* `redirect` - Automatically send a redirect response
* `req_hdr_wait_time` - HTTP request header wait time before abort connection
* `req_hdr_wait_time_val` - Number of seconds wait for client request header (default is 7)
* `request_line_case_insensitive` - Parse http request line as case insensitive
* `request_timeout` - Request timeout if response not received (timeout in seconds)
* `retry_on_5xx` - Retry http request on HTTP 5xx code
* `retry_on_5xx_per_req` - Retry http request on HTTP 5xx code for each request
* `retry_on_5xx_per_req_val` - Number of times to retry (default is 3)
* `retry_on_5xx_val` - Number of times to retry (default is 3)
* `strict_transaction_switch` - Force server selection on every HTTP request
* `term_11client_hdr_conn_close` - Terminate HTTP 1.1 client when req has Connection: close
* `url_hash_first` - Use the begining part of URL to calculate hash value (URL string length to calculate hash value)
* `url_hash_last` - Use the end part of URL to calculate hash value (URL string length to calculate hash value)
* `url_hash_offset` - Skip part of URL to calculate hash value (Offset of the URL string)
* `url_hash_persist` - Use URL’s hash value to select server
* `use_server_status` - Use Server-Status header to do URL hashing
* `user_tag` - Customized tag
* `uuid` - uuid of the object
* `request_header_erase` - Erase a header from HTTP request (Header Name)
* `redirect_secure` - Use HTTPS
* `redirect_secure_port` - Port (Port Number)
* `redirect_match` - URL Matching (Pattern URL String)
* `rewrite_to` - Rewrite to Destination URL String
* `response_header_insert` - Insert a header into HTTP response (Header Content (Format: “[name]: [value]”))
* `response_header_insert_type` - ‘insert-if-not-exist’: Only insert the header when it does not exist; ‘insert-always’: Always insert the header even when there is a header with the same name;
* `response_header_erase` - Erase a header from HTTP response (Header Name)
* `logging` - Logging template (Logging Config name)
* `url_match_string` - URL String
* `url_service_group` - Create a Service Group comprising Servers (Service Group Name)
* `url_switching_type` - ‘contains’: Select service group if URL string contains another string; ‘ends-with’: Select service group if URL string ends with another string; ‘equals’: Select service group if URL string equals another string; ‘starts-with’: Select service group if URL string starts with another string; ‘url-case-insensitive’: Case insensitive URL switching; ‘url-hits-enable’: Enables URL Hits;
* `host_match_string` - Hostname String
* `host_service_group` - Create a Service Group comprising Servers (Service Group Name)
* `host_switching_type` - ‘contains’: Select service group if hostname contains another string; ‘ends-with’: Select service group if hostname ends with another string; ‘equals’: Select service group if hostname equals another string; ‘starts-with’: Select service group if hostname starts with another string; ‘host-hits-enable’: Enables Host Hits counters;
* `response_content_replace` - replace the data from HTTP response content (String in the http content need to be replaced)
* `response_new_string` - String will be in the http content
* `request_header_insert` - Insert a header into HTTP request (Header Content (Format: “[name]: [value]”))
* `request_header_insert_type` - ‘insert-if-not-exist’: Only insert the header when it does not exist; ‘insert-always’: Always insert the header even when there is a header with the same name;
* `content_type` - Compression content-type
* `exclude_uri` - Compression exclude uri
* `exclude_content_type` - Compression exclude content-type (Compression exclude content type)
