---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "thunder_slb_template_policy_forward_policy Resource - terraform-provider-thunder"
subcategory: ""
description: |-
  thunder_slb_template_policy_forward_policy: Forward Policy commands
  PLACEHOLDER
---

# thunder_slb_template_policy_forward_policy (Resource)

`thunder_slb_template_policy_forward_policy`: Forward Policy commands

__PLACEHOLDER__

## Example Usage

```terraform
provider "thunder" {
  address  = var.dut9049
  username = var.username
  password = var.password
}
resource "thunder_slb_template_policy_forward_policy" "thunder_slb_template_policy_forward_policy" {

  name           = "test-policy"
  acos_event_log = 0
  san_filtering {
    ssli_url_filtering_san = "enable-san"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name

### Optional

- `acos_event_log` (Number) Enable acos event logging
- `action_list` (Block List) (see [below for nested schema](#nestedblock--action_list))
- `dual_stack_action_list` (Block List) (see [below for nested schema](#nestedblock--dual_stack_action_list))
- `enable_adv_match` (Number) Enable adv-match rules and deactive all the other kinds of destination rules
- `filtering` (Block List) (see [below for nested schema](#nestedblock--filtering))
- `forward_http_connect_to_icap` (Number) Forward HTTP CONNECT request to ICAP server
- `local_logging` (Number) Enable local logging
- `no_client_conn_reuse` (Number) Inspects only first request of a connection
- `reqmod_icap` (String) ICAP reqmod template (Reqmod ICAP Template Name)
- `require_web_category` (Number) Wait for web category to be resolved before taking proxy decision
- `san_filtering` (Block List) (see [below for nested schema](#nestedblock--san_filtering))
- `source_list` (Block List) (see [below for nested schema](#nestedblock--source_list))
- `uuid` (String) uuid of the object

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--action_list"></a>
### Nested Schema for `action_list`

Required:

- `name` (String) Action policy name

Optional:

- `action1` (String) 'forward-to-internet': Forward request to Internet; 'forward-to-service-group': Forward request to service group; 'forward-to-proxy': Forward request to HTTP proxy server; 'drop': Drop request;
- `drop_message` (String) drop-message sent to the client as webpage(html tags are included and quotation marks are required for white spaces)
- `drop_redirect_url` (String) Specify URL to which client request is redirected upon being dropped
- `drop_response_code` (Number) Specify response code for drop action
- `fake_sg` (String) service group to forward the packets to Internet
- `fall_back` (String) Fallback service group for Internet
- `fall_back_snat` (String) Source NAT pool or pool group for fallback server
- `fall_back_snat_pt_only` (Number) Source port translation only for fallback server
- `forward_snat` (String) Source NAT pool or pool group
- `forward_snat_pt_only` (Number) Source port translation only
- `http_status_code` (String) '301': Moved permanently; '302': Found;
- `log` (Number) enable logging
- `proxy_chaining` (Number) Enable proxy chaining feature
- `proxy_chaining_bypass` (Number) Forward all https packets to upstream proxy
- `real_sg` (String) service group to forward the packets
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--action_list--sampling_enable))
- `support_cert_fetch` (Number) Fetch server certificate by upstream proxy
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

<a id="nestedblock--action_list--sampling_enable"></a>
### Nested Schema for `action_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'hits': Number of requests matching this destination rule;



<a id="nestedblock--dual_stack_action_list"></a>
### Nested Schema for `dual_stack_action_list`

Required:

- `name` (String) Action name

Optional:

- `fall_back` (String) Fallback service group
- `fall_back_snat` (String) Source NAT pool or pool group for fallback
- `ipv4` (String) IPv4 service group to forward
- `ipv4_snat` (String) IPv4 source NAT pool or pool group
- `ipv6` (String) IPv6 service group to forward
- `ipv6_snat` (String) IPv6 source NAT pool or pool group
- `log` (Number) enable logging
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--dual_stack_action_list--sampling_enable))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

<a id="nestedblock--dual_stack_action_list--sampling_enable"></a>
### Nested Schema for `dual_stack_action_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'hits': Number of requests forward by this action;



<a id="nestedblock--filtering"></a>
### Nested Schema for `filtering`

Optional:

- `ssli_url_filtering` (String) 'bypassed-sni-disable': Disable SNI filtering for bypassed URL's(enabled by default); 'intercepted-sni-enable': Enable SNI filtering for intercepted URL's(disabled by default); 'intercepted-http-disable': Disable HTTP(host/URL) filtering for intercepted URL's(enabled by default); 'no-sni-allow': Allow connection if SNI filtering is enabled and SNI header is not present(Drop by default);


<a id="nestedblock--san_filtering"></a>
### Nested Schema for `san_filtering`

Optional:

- `ssli_url_filtering_san` (String) 'enable-san': Enable SAN filtering(disabled by default); 'bypassed-san-disable': Disable SAN filtering for bypassed URL's(enabled by default); 'intercepted-san-enable': Enable SAN filtering for intercepted URL's(disabled by default); 'no-san-allow': Allow connection if SAN filtering is enabled and SAN field is not present(Drop by default);


<a id="nestedblock--source_list"></a>
### Nested Schema for `source_list`

Required:

- `name` (String) source destination match rule name

Optional:

- `destination` (Block List, Max: 1) (see [below for nested schema](#nestedblock--source_list--destination))
- `match_any` (Number) Match any source
- `match_authorize_policy` (String) Authorize-policy for user and group based policy
- `match_class_list` (String) Class List Name
- `priority` (Number) Priority of the source(higher the number higher the priority, default 0)
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--source_list--sampling_enable))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

<a id="nestedblock--source_list--destination"></a>
### Nested Schema for `source_list.destination`

Optional:

- `adv_match_list` (Block List) (see [below for nested schema](#nestedblock--source_list--destination--adv_match_list))
- `any` (Block List, Max: 1) (see [below for nested schema](#nestedblock--source_list--destination--any))
- `class_list_list` (Block List) (see [below for nested schema](#nestedblock--source_list--destination--class_list_list))
- `web_category_list_list` (Block List) (see [below for nested schema](#nestedblock--source_list--destination--web_category_list_list))
- `web_reputation_scope_list` (Block List) (see [below for nested schema](#nestedblock--source_list--destination--web_reputation_scope_list))

<a id="nestedblock--source_list--destination--adv_match_list"></a>
### Nested Schema for `source_list.destination.adv_match_list`

Required:

- `priority` (Number) Rule priority (1000 is highest)

Optional:

- `action` (String) Forwading action of this rule
- `disable_reqmod_icap` (Number) Disable REQMOD ICAP template
- `disable_respmod_icap` (Number) Disable RESPMOD ICAP template
- `dual_stack_action` (String) Forwarding action of this rule
- `match_host` (String) Match request host (HTTP stage) or SNI/SAN (SSL stage)
- `match_http_content_encoding` (String) Match the value of HTTP header "Content-Encoding"
- `match_http_content_length_range_begin` (Number) Match the value of HTTP header "Content-Length" with an inclusive range
- `match_http_content_length_range_end` (Number) End of the "Content-Length" range
- `match_http_content_type` (String) Match the value of HTTP header "Content-Type"
- `match_http_header` (String) Matching the name of all request headers
- `match_http_method_connect` (Number) Match HTTP request method CONNECT
- `match_http_method_delete` (Number) Match HTTP request method DELETE
- `match_http_method_get` (Number) Match HTTP request method GET
- `match_http_method_head` (Number) Match HTTP request method HEAD
- `match_http_method_options` (Number) Match HTTP request method OPTIONS
- `match_http_method_patch` (Number) Match HTTP request method PATCH
- `match_http_method_post` (Number) Match HTTP request method POST
- `match_http_method_put` (Number) Match HTTP request method PUT
- `match_http_method_trace` (Number) Match HTTP request method TRACE
- `match_http_request_file_extension` (String) Match file extension of URL in HTTP request line
- `match_http_url` (String) Match URL in HTTP request line
- `match_http_url_regex` (String) Match URI in HTTP request line by given regular expression
- `match_http_user_agent` (String) Matching the value of HTTP header "User-Agent"
- `match_server_address` (String) Match target server IP address
- `match_server_port` (Number) Match target server port number
- `match_server_port_range_begin` (Number) Math targer server port range inclusively
- `match_server_port_range_end` (Number) End of port range
- `match_time_range` (String) Enable rule in this time-range
- `match_web_category_list` (String) Match web-category list
- `match_web_reputation_scope` (String) Match web-reputation scope
- `notify_page` (String) Send notify-page to client
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--source_list--destination--adv_match_list--sampling_enable))
- `user_tag` (String) Customized tag
- `uuid` (String) uuid of the object

<a id="nestedblock--source_list--destination--adv_match_list--sampling_enable"></a>
### Nested Schema for `source_list.destination.adv_match_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'hits': Number of requests hit this rule;



<a id="nestedblock--source_list--destination--any"></a>
### Nested Schema for `source_list.destination.any`

Optional:

- `action` (String) Action to be performed
- `dual_stack_action` (String) Dual-stack action to be performed
- `sampling_enable` (Block List) (see [below for nested schema](#nestedblock--source_list--destination--any--sampling_enable))
- `uuid` (String) uuid of the object

<a id="nestedblock--source_list--destination--any--sampling_enable"></a>
### Nested Schema for `source_list.destination.any.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'hits': Number of requests matching this destination rule;



<a id="nestedblock--source_list--destination--class_list_list"></a>
### Nested Schema for `source_list.destination.class_list_list`

Required:

- `dest_class_list` (String) Destination Class List Name

Optional:

- `action` (String) Action to be performed
- `dual_stack_action` (String) Dual-stack action to be performed
- `priority` (Number) Priority value of the action(higher the number higher the priority)
- `type` (String) 'host': Match hostname; 'url': Match URL; 'ip': Match destination IP address;
- `uuid` (String) uuid of the object


<a id="nestedblock--source_list--destination--web_category_list_list"></a>
### Nested Schema for `source_list.destination.web_category_list_list`

Required:

- `web_category_list` (String) Destination Web Category List Name

Optional:

- `action` (String) Action to be performed
- `dual_stack_action` (String) Dual-stack action to be performed
- `priority` (Number) Priority value of the action(higher the number higher the priority)
- `type` (String) 'host': Match hostname; 'url': match URL;
- `uuid` (String) uuid of the object


<a id="nestedblock--source_list--destination--web_reputation_scope_list"></a>
### Nested Schema for `source_list.destination.web_reputation_scope_list`

Required:

- `web_reputation_scope` (String) Destination Web Reputation Scope Name

Optional:

- `action` (String) Action to be performed
- `dual_stack_action` (String) Dual-stack action to be performed
- `priority` (Number) Priority value of the action(higher the number higher the priority)
- `type` (String) 'host': Match hostname; 'url': match URL;
- `uuid` (String) uuid of the object



<a id="nestedblock--source_list--sampling_enable"></a>
### Nested Schema for `source_list.sampling_enable`

Optional:

- `counters1` (String) 'all': all; 'hits': Number of requests matching this source rule; 'destination-match-not-found': Number of requests without matching destination rule; 'no-host-info': Failed to parse ip or host information from request;

