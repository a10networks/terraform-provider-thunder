package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAForceSelfStandbyPersistent() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_force_self_standby_persistent`: HA VRRP-A Configured  Command to force the unit or a group to HA standby state\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpAForceSelfStandbyPersistentCreate,
		UpdateContext: resourceVrrpAForceSelfStandbyPersistentUpdate,
		ReadContext:   resourceVrrpAForceSelfStandbyPersistentRead,
		DeleteContext: resourceVrrpAForceSelfStandbyPersistentDelete,

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid": {
				Type: schema.TypeInt, Required: true, Description: "Specify one VRRP-A vrid to force into standby state",
			},
		},
	}
}
func resourceVrrpAForceSelfStandbyPersistentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAForceSelfStandbyPersistentCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAForceSelfStandbyPersistent(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAForceSelfStandbyPersistentRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAForceSelfStandbyPersistentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAForceSelfStandbyPersistentUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAForceSelfStandbyPersistent(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAForceSelfStandbyPersistentRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpAForceSelfStandbyPersistentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAForceSelfStandbyPersistentDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAForceSelfStandbyPersistent(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAForceSelfStandbyPersistentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAForceSelfStandbyPersistentRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAForceSelfStandbyPersistent(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVrrpAForceSelfStandbyPersistent(d *schema.ResourceData) edpt.VrrpAForceSelfStandbyPersistent {
	var ret edpt.VrrpAForceSelfStandbyPersistent
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
