package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_recursive_dns_resolution_gateway_health_check`: Enable health check to recursive gateway based on current configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheckCreate,
		UpdateContext: resourceSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheckUpdate,
		ReadContext:   resourceSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheckRead,
		DeleteContext: resourceSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheckDelete,

		Schema: map[string]*schema.Schema{
			"gwhc_ns_cache_lookup": {
				Type: schema.TypeString, Optional: true, Default: "disabled", Description: "'disabled': Disable NS Cache Lookup; 'enabled': Enable NS Cache Lookup;",
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "Specify the health check interval, default is 10 sec (Interval value, in seconds (default 10))",
			},
			"num_query_type": {
				Type: schema.TypeInt, Optional: true, Description: "Other record type value",
			},
			"query_name": {
				Type: schema.TypeString, Optional: true, Default: "a10networks.com", Description: "Specify the query name used in probe queries, default \"a10networks.com\"",
			},
			"retry": {
				Type: schema.TypeInt, Optional: true, Default: 6, Description: "Maximum number of DNS query retries at each server level before health check fails, default 6 (Retry count (default 6))",
			},
			"retry_multi": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify number of times that health check consecutively fails before declaring gateway DOWN, default 1 (retry-multi count (default 1))",
			},
			"str_query_type": {
				Type: schema.TypeString, Optional: true, Default: "A", Description: "'A': Address record; 'AAAA': IPv6 Address record; 'CNAME': Canonical name record; 'MX': Mail exchange record; 'NS': Name server record; 'SRV': Service locator; 'PTR': PTR resource record; 'SOA': Start of authority record; 'TXT': Text record;",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Specify the health check timeout before retrying or finish, default is 5 sec (Timeout value, in seconds (default 5))",
			},
			"up_retry": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Specify number of times that health check consecutively passes before declaring gateway UP, default 1 (up-retry count (default 1))",
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
func resourceSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheckCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheckRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheckUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheckRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheckDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheckRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck(d *schema.ResourceData) edpt.SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck {
	var ret edpt.SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck
	ret.Inst.GwhcNsCacheLookup = d.Get("gwhc_ns_cache_lookup").(string)
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.NumQueryType = d.Get("num_query_type").(int)
	ret.Inst.QueryName = d.Get("query_name").(string)
	ret.Inst.Retry = d.Get("retry").(int)
	ret.Inst.RetryMulti = d.Get("retry_multi").(int)
	ret.Inst.StrQueryType = d.Get("str_query_type").(string)
	ret.Inst.Timeout = d.Get("timeout").(int)
	ret.Inst.UpRetry = d.Get("up_retry").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
