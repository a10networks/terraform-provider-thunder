package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugBridgeVlanGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_bridge_vlan_group`: Debug bridge vlan traffic\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugBridgeVlanGroupCreate,
		UpdateContext: resourceDebugBridgeVlanGroupUpdate,
		ReadContext:   resourceDebugBridgeVlanGroupRead,
		DeleteContext: resourceDebugBridgeVlanGroupDelete,

		Schema: map[string]*schema.Schema{
			"dumy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dummy",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugBridgeVlanGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugBridgeVlanGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugBridgeVlanGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugBridgeVlanGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugBridgeVlanGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugBridgeVlanGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugBridgeVlanGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugBridgeVlanGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugBridgeVlanGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugBridgeVlanGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugBridgeVlanGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugBridgeVlanGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugBridgeVlanGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugBridgeVlanGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugBridgeVlanGroup(d *schema.ResourceData) edpt.DebugBridgeVlanGroup {
	var ret edpt.DebugBridgeVlanGroup
	ret.Inst.Dumy = d.Get("dumy").(int)
	//omit uuid
	return ret
}
