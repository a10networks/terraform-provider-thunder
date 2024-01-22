package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LoggingNatQuotaExceeded() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_logging_nat_quota_exceeded`: Change logging level for NAT Quota Exceeded\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LoggingNatQuotaExceededCreate,
		UpdateContext: resourceCgnv6LoggingNatQuotaExceededUpdate,
		ReadContext:   resourceCgnv6LoggingNatQuotaExceededRead,
		DeleteContext: resourceCgnv6LoggingNatQuotaExceededDelete,

		Schema: map[string]*schema.Schema{
			"level": {
				Type: schema.TypeString, Optional: true, Default: "warning", Description: "'warning': Log level Warning (Default); 'critical': Log level Critical; 'notice': Log level Notice;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6LoggingNatQuotaExceededCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingNatQuotaExceededCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LoggingNatQuotaExceeded(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LoggingNatQuotaExceededRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LoggingNatQuotaExceededUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingNatQuotaExceededUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LoggingNatQuotaExceeded(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LoggingNatQuotaExceededRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LoggingNatQuotaExceededDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingNatQuotaExceededDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LoggingNatQuotaExceeded(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LoggingNatQuotaExceededRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingNatQuotaExceededRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LoggingNatQuotaExceeded(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6LoggingNatQuotaExceeded(d *schema.ResourceData) edpt.Cgnv6LoggingNatQuotaExceeded {
	var ret edpt.Cgnv6LoggingNatQuotaExceeded
	ret.Inst.Level = d.Get("level").(string)
	//omit uuid
	return ret
}
