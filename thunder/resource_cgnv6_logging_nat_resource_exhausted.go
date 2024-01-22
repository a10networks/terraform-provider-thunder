package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6LoggingNatResourceExhausted() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_logging_nat_resource_exhausted`: Change logging level for NAT Resource Exhausted\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6LoggingNatResourceExhaustedCreate,
		UpdateContext: resourceCgnv6LoggingNatResourceExhaustedUpdate,
		ReadContext:   resourceCgnv6LoggingNatResourceExhaustedRead,
		DeleteContext: resourceCgnv6LoggingNatResourceExhaustedDelete,

		Schema: map[string]*schema.Schema{
			"level": {
				Type: schema.TypeString, Optional: true, Default: "critical", Description: "'warning': Log level Warning; 'critical': Log level Critical (Default); 'notice': Log level Notice;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6LoggingNatResourceExhaustedCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingNatResourceExhaustedCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LoggingNatResourceExhausted(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LoggingNatResourceExhaustedRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6LoggingNatResourceExhaustedUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingNatResourceExhaustedUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LoggingNatResourceExhausted(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6LoggingNatResourceExhaustedRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6LoggingNatResourceExhaustedDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingNatResourceExhaustedDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LoggingNatResourceExhausted(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6LoggingNatResourceExhaustedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6LoggingNatResourceExhaustedRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6LoggingNatResourceExhausted(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6LoggingNatResourceExhausted(d *schema.ResourceData) edpt.Cgnv6LoggingNatResourceExhausted {
	var ret edpt.Cgnv6LoggingNatResourceExhausted
	ret.Inst.Level = d.Get("level").(string)
	//omit uuid
	return ret
}
