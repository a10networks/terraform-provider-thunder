package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsRecursiveDnsResolution() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_recursive_dns_resolution`: Recursive DNS resolver configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsRecursiveDnsResolutionCreate,
		UpdateContext: resourceSlbTemplateDnsRecursiveDnsResolutionUpdate,
		ReadContext:   resourceSlbTemplateDnsRecursiveDnsResolutionRead,
		DeleteContext: resourceSlbTemplateDnsRecursiveDnsResolutionDelete,

		Schema: map[string]*schema.Schema{
			"csubnet_retry": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "retry when server REFUSED AX inserted EDNS(0) subnet, works only when insert-client-subnet is configured",
			},
			"default_recursive": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Default recursive mode, forward query to bound service-group if hostnames matched",
			},
			"dnssec_validation": {
				Type: schema.TypeString, Optional: true, Default: "disabled", Description: "'enabled': Enable DNSSEC validation; 'disabled': Disable DNSSEC validation;",
			},
			"fast_ns_selection": {
				Type: schema.TypeString, Optional: true, Default: "enabled", Description: "'enabled': Enable fast NS selection; 'disabled': Disable fast NS selection;",
			},
			"force_cname_resolution": {
				Type: schema.TypeString, Optional: true, Default: "enabled", Description: "'enabled': Force CNAME resolution always; 'disabled': Use answer record in CNAME response if it exists, else resolve;",
			},
			"full_response": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Serve all records (authority and additional) when applicable",
			},
			"gateway_health_check": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"query_name": {
							Type: schema.TypeString, Optional: true, Default: "a10networks.com", Description: "Specify the query name used in probe queries, default \"a10networks.com\"",
						},
						"retry": {
							Type: schema.TypeInt, Optional: true, Default: 6, Description: "Maximum number of DNS query retries at each server level before health check fails, default 6 (Retry count (default 6))",
						},
						"timeout": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify the health check timeout before retrying or finish, default is 5 sec (Timeout value, in seconds (default 5))",
						},
						"interval": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Specify the health check interval, default is 10 sec (Interval value, in seconds (default 10))",
						},
						"up_retry": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify number of times that health check consecutively passes before declaring gateway UP, default 1 (up-retry count (default 1))",
						},
						"retry_multi": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify number of times that health check consecutively fails before declaring gateway DOWN, default 1 (retry-multi count (default 1))",
						},
						"gwhc_ns_cache_lookup": {
							Type: schema.TypeString, Optional: true, Default: "disabled", Description: "'disabled': Disable NS Cache Lookup; 'enabled': Enable NS Cache Lookup;",
						},
						"str_query_type": {
							Type: schema.TypeString, Optional: true, Default: "A", Description: "'A': Address record; 'AAAA': IPv6 Address record; 'CNAME': Canonical name record; 'MX': Mail exchange record; 'NS': Name server record; 'SRV': Service locator; 'PTR': PTR resource record; 'SOA': Start of authority record; 'TXT': Text record;",
						},
						"num_query_type": {
							Type: schema.TypeInt, Optional: true, Description: "Other record type value",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"host_list_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostnames": {
							Type: schema.TypeString, Optional: true, Description: "Hostnames class-list name (dns type), perform resolution while query name matched",
						},
					},
				},
			},
			"ipv4_nat_pool": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 Source NAT pool or pool group",
			},
			"ipv6_nat_pool": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 Source NAT pool or pool group",
			},
			"lookup_order": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"query_type": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"str_query_type": {
										Type: schema.TypeString, Optional: true, Description: "'A': Address record; 'AAAA': IPv6 Address record; 'CNAME': Canonical name record; 'MX': Mail exchange record; 'NS': Name server record; 'SRV': Service locator; 'PTR': PTR resource record; 'SOA': Start of authority record; 'TXT': Text record; 'ANY': All cached record;",
									},
									"num_query_type": {
										Type: schema.TypeInt, Optional: true, Description: "Other query type value",
									},
									"order": {
										Type: schema.TypeString, Optional: true, Description: "'ipv4-precede-ipv6': Recursive lookup via IPv4 then IPv6; 'ipv6-precede-ipv4': Recursive lookup via IPv6 then IPv4;",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"max_trials": {
				Type: schema.TypeInt, Optional: true, Default: 255, Description: "Total number of times to try DNS query to server before closing client connection, default 255",
			},
			"ns_cache_lookup": {
				Type: schema.TypeString, Optional: true, Default: "enabled", Description: "'disabled': Disable NS Cache Lookup; 'enabled': Enable NS Cache Lookup;",
			},
			"request_for_pending_resolution": {
				Type: schema.TypeString, Optional: true, Default: "respond-with-servfail", Description: "'drop': Drop of the request during ongoing; 'respond-with-servfail': Respond with SERVFAIL of the request during ongoing; 'start-new-resolution': Start new resolution of the request during ongoing;",
			},
			"retries_per_level": {
				Type: schema.TypeInt, Optional: true, Default: 6, Description: "Number of DNS query retries at each server level before closing client connection, default 6",
			},
			"udp_initial_interval": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "UDP DNS Retry Interval value 1-6, default is 5 sec (1-6, default is 5sec)",
			},
			"udp_retry_interval": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "UDP DNS Retry Interval value 1-6, default is 1 sec (1-6 , default is 1 sec)",
			},
			"use_client_qid": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use client side query id for recursive query",
			},
			"use_service_group_response": {
				Type: schema.TypeString, Optional: true, Default: "enabled", Description: "'disabled': Start Recursive Resolver if Server response doesnt have final answer; 'enabled': Forward Backend Server response to client and dont start recursive resolver;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbTemplateDnsRecursiveDnsResolutionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRecursiveDnsResolutionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRecursiveDnsResolution(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsRecursiveDnsResolutionRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsRecursiveDnsResolutionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRecursiveDnsResolutionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRecursiveDnsResolution(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsRecursiveDnsResolutionRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsRecursiveDnsResolutionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRecursiveDnsResolutionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRecursiveDnsResolution(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsRecursiveDnsResolutionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRecursiveDnsResolutionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRecursiveDnsResolution(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1419(d []interface{}) edpt.SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1419 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1419
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.QueryName = in["query_name"].(string)
		ret.Retry = in["retry"].(int)
		ret.Timeout = in["timeout"].(int)
		ret.Interval = in["interval"].(int)
		ret.UpRetry = in["up_retry"].(int)
		ret.RetryMulti = in["retry_multi"].(int)
		ret.GwhcNsCacheLookup = in["gwhc_ns_cache_lookup"].(string)
		ret.StrQueryType = in["str_query_type"].(string)
		ret.NumQueryType = in["num_query_type"].(int)
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsRecursiveDnsResolutionHostListCfg(d []interface{}) []edpt.SlbTemplateDnsRecursiveDnsResolutionHostListCfg {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsRecursiveDnsResolutionHostListCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsRecursiveDnsResolutionHostListCfg
		oi.Hostnames = in["hostnames"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectSlbTemplateDnsRecursiveDnsResolutionLookupOrder1420(d []interface{}) edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrder1420 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrder1420
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.QueryType = getSliceSlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1421(in["query_type"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1421(d []interface{}) []edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1421 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1421, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsRecursiveDnsResolutionLookupOrderQueryType1421
		oi.StrQueryType = in["str_query_type"].(string)
		oi.NumQueryType = in["num_query_type"].(int)
		oi.Order = in["order"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateDnsRecursiveDnsResolution(d *schema.ResourceData) edpt.SlbTemplateDnsRecursiveDnsResolution {
	var ret edpt.SlbTemplateDnsRecursiveDnsResolution
	ret.Inst.CsubnetRetry = d.Get("csubnet_retry").(int)
	ret.Inst.DefaultRecursive = d.Get("default_recursive").(int)
	ret.Inst.DnssecValidation = d.Get("dnssec_validation").(string)
	ret.Inst.FastNsSelection = d.Get("fast_ns_selection").(string)
	ret.Inst.ForceCnameResolution = d.Get("force_cname_resolution").(string)
	ret.Inst.FullResponse = d.Get("full_response").(int)
	ret.Inst.GatewayHealthCheck = getObjectSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck1419(d.Get("gateway_health_check").([]interface{}))
	ret.Inst.HostListCfg = getSliceSlbTemplateDnsRecursiveDnsResolutionHostListCfg(d.Get("host_list_cfg").([]interface{}))
	ret.Inst.Ipv4NatPool = d.Get("ipv4_nat_pool").(string)
	ret.Inst.Ipv6NatPool = d.Get("ipv6_nat_pool").(string)
	ret.Inst.LookupOrder = getObjectSlbTemplateDnsRecursiveDnsResolutionLookupOrder1420(d.Get("lookup_order").([]interface{}))
	ret.Inst.MaxTrials = d.Get("max_trials").(int)
	ret.Inst.NsCacheLookup = d.Get("ns_cache_lookup").(string)
	ret.Inst.RequestForPendingResolution = d.Get("request_for_pending_resolution").(string)
	ret.Inst.RetriesPerLevel = d.Get("retries_per_level").(int)
	ret.Inst.UdpInitialInterval = d.Get("udp_initial_interval").(int)
	ret.Inst.UdpRetryInterval = d.Get("udp_retry_interval").(int)
	ret.Inst.UseClientQid = d.Get("use_client_qid").(int)
	ret.Inst.UseServiceGroupResponse = d.Get("use_service_group_response").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
