package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwPrioritizeControlTraffic() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_prioritize_control_traffic`: Configure Permit internal policy lookup ahead of external policy lookup for internal traffic\n\n__PLACEHOLDER__",
		CreateContext: resourceFwPrioritizeControlTrafficCreate,
		UpdateContext: resourceFwPrioritizeControlTrafficUpdate,
		ReadContext:   resourceFwPrioritizeControlTrafficRead,
		DeleteContext: resourceFwPrioritizeControlTrafficDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'disable': disable; 'enable': enable;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwPrioritizeControlTrafficCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwPrioritizeControlTrafficCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwPrioritizeControlTraffic(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwPrioritizeControlTrafficRead(ctx, d, meta)
	}
	return diags
}

func resourceFwPrioritizeControlTrafficUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwPrioritizeControlTrafficUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwPrioritizeControlTraffic(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwPrioritizeControlTrafficRead(ctx, d, meta)
	}
	return diags
}
func resourceFwPrioritizeControlTrafficDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwPrioritizeControlTrafficDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwPrioritizeControlTraffic(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwPrioritizeControlTrafficRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwPrioritizeControlTrafficRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwPrioritizeControlTraffic(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwPrioritizeControlTraffic(d *schema.ResourceData) edpt.FwPrioritizeControlTraffic {
	var ret edpt.FwPrioritizeControlTraffic
	ret.Inst.Action = d.Get("action").(string)
	//omit uuid
	return ret
}
