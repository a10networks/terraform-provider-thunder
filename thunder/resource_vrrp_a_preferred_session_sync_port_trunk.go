package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAPreferredSessionSyncPortTrunk() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_preferred_session_sync_port_trunk`: preferred-session-sync-port trunk\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpAPreferredSessionSyncPortTrunkCreate,
		UpdateContext: resourceVrrpAPreferredSessionSyncPortTrunkUpdate,
		ReadContext:   resourceVrrpAPreferredSessionSyncPortTrunkRead,
		DeleteContext: resourceVrrpAPreferredSessionSyncPortTrunkDelete,

		Schema: map[string]*schema.Schema{
			"pre_trunk": {
				Type: schema.TypeInt, Required: true, Description: "Trunk Interface number",
			},
			"pre_vlan": {
				Type: schema.TypeInt, Optional: true, Description: "Interface VLAN (VLAN ID)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceVrrpAPreferredSessionSyncPortTrunkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAPreferredSessionSyncPortTrunkCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAPreferredSessionSyncPortTrunk(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAPreferredSessionSyncPortTrunkRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAPreferredSessionSyncPortTrunkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAPreferredSessionSyncPortTrunkUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAPreferredSessionSyncPortTrunk(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAPreferredSessionSyncPortTrunkRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpAPreferredSessionSyncPortTrunkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAPreferredSessionSyncPortTrunkDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAPreferredSessionSyncPortTrunk(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAPreferredSessionSyncPortTrunkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAPreferredSessionSyncPortTrunkRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAPreferredSessionSyncPortTrunk(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVrrpAPreferredSessionSyncPortTrunk(d *schema.ResourceData) edpt.VrrpAPreferredSessionSyncPortTrunk {
	var ret edpt.VrrpAPreferredSessionSyncPortTrunk
	ret.Inst.PreTrunk = d.Get("pre_trunk").(int)
	ret.Inst.PreVlan = d.Get("pre_vlan").(int)
	//omit uuid
	return ret
}
