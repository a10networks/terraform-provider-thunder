package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpASessionSync() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_session_sync`: Enable/Disable vrrp-a session sync\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpASessionSyncCreate,
		UpdateContext: resourceVrrpASessionSyncUpdate,
		ReadContext:   resourceVrrpASessionSyncRead,
		DeleteContext: resourceVrrpASessionSyncDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "enable", Description: "'enable': enable vrrp-a session sync (default); 'disable': disable vrrp-a session sync;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVrrpASessionSyncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpASessionSyncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpASessionSync(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpASessionSyncRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpASessionSyncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpASessionSyncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpASessionSync(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpASessionSyncRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpASessionSyncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpASessionSyncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpASessionSync(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpASessionSyncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpASessionSyncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpASessionSync(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVrrpASessionSync(d *schema.ResourceData) edpt.VrrpASessionSync {
	var ret edpt.VrrpASessionSync
	ret.Inst.Action = d.Get("action").(string)
	//omit uuid
	return ret
}
