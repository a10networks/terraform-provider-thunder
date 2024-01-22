package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6SctpGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_sctp_global`: Configure CGNv6 SCTP NAT global parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6SctpGlobalCreate,
		UpdateContext: resourceCgnv6SctpGlobalUpdate,
		ReadContext:   resourceCgnv6SctpGlobalRead,
		DeleteContext: resourceCgnv6SctpGlobalDelete,

		Schema: map[string]*schema.Schema{
			"half_open_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 4, Description: "Set SCTP half-open timeout (SCTP half-open timeout in seconds (default 4))",
			},
			"idle_timeout": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "SCTP idle timeout in minutes (default 5)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6SctpGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6SctpGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6SctpGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6SctpGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6SctpGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6SctpGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6SctpGlobal(d *schema.ResourceData) edpt.Cgnv6SctpGlobal {
	var ret edpt.Cgnv6SctpGlobal
	ret.Inst.HalfOpenTimeout = d.Get("half_open_timeout").(int)
	ret.Inst.IdleTimeout = d.Get("idle_timeout").(int)
	//omit uuid
	return ret
}
