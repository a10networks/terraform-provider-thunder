package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbCommonDnsResponseRateLimiting() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_common_dns_response_rate_limiting`: DNS Response-Rate-Limiting Global Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbCommonDnsResponseRateLimitingCreate,
		UpdateContext: resourceSlbCommonDnsResponseRateLimitingUpdate,
		ReadContext:   resourceSlbCommonDnsResponseRateLimitingRead,
		DeleteContext: resourceSlbCommonDnsResponseRateLimitingDelete,

		Schema: map[string]*schema.Schema{
			"max_table_entries": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum number of entries allowed",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSlbCommonDnsResponseRateLimitingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonDnsResponseRateLimitingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonDnsResponseRateLimiting(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonDnsResponseRateLimitingRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbCommonDnsResponseRateLimitingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonDnsResponseRateLimitingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonDnsResponseRateLimiting(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbCommonDnsResponseRateLimitingRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbCommonDnsResponseRateLimitingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonDnsResponseRateLimitingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonDnsResponseRateLimiting(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbCommonDnsResponseRateLimitingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbCommonDnsResponseRateLimitingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbCommonDnsResponseRateLimiting(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbCommonDnsResponseRateLimiting(d *schema.ResourceData) edpt.SlbCommonDnsResponseRateLimiting {
	var ret edpt.SlbCommonDnsResponseRateLimiting
	ret.Inst.MaxTableEntries = d.Get("max_table_entries").(int)
	//omit uuid
	return ret
}
