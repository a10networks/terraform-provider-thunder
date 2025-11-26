package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFwFastForwarding() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_fw_fast_forwarding`: Fast packets forwarding for established CGN/FW sessions\n\n__PLACEHOLDER__",
		CreateContext: resourceFwFastForwardingCreate,
		UpdateContext: resourceFwFastForwardingUpdate,
		ReadContext:   resourceFwFastForwardingRead,
		DeleteContext: resourceFwFastForwardingDelete,

		Schema: map[string]*schema.Schema{
			"toggle": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': Enable fast forwarding path; 'disable': Disable fast forwarding path;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceFwFastForwardingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwFastForwardingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwFastForwarding(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwFastForwardingRead(ctx, d, meta)
	}
	return diags
}

func resourceFwFastForwardingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwFastForwardingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwFastForwarding(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceFwFastForwardingRead(ctx, d, meta)
	}
	return diags
}
func resourceFwFastForwardingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwFastForwardingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwFastForwarding(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceFwFastForwardingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceFwFastForwardingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointFwFastForwarding(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointFwFastForwarding(d *schema.ResourceData) edpt.FwFastForwarding {
	var ret edpt.FwFastForwarding
	ret.Inst.Toggle = d.Get("toggle").(string)
	//omit uuid
	return ret
}
