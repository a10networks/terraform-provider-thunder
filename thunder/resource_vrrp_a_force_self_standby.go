package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVrrpAForceSelfStandby() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vrrp_a_force_self_standby`: HA VRRP-A Operational Command to force the unit or a group to HA standby state\n\n__PLACEHOLDER__",
		CreateContext: resourceVrrpAForceSelfStandbyCreate,
		UpdateContext: resourceVrrpAForceSelfStandbyUpdate,
		ReadContext:   resourceVrrpAForceSelfStandbyRead,
		DeleteContext: resourceVrrpAForceSelfStandbyDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Description: "'enable': enable vrrp-a force-self-standby; 'disable': disable vrrp-a force-self-standby;",
			},
			"all_partitions": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "force all partitions in standby state",
			},
			"skip_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "skip single device check",
			},
			"vrid": {
				Type: schema.TypeInt, Optional: true, Description: "Specify one VRRP-A vrid to force into standby state",
			},
		},
	}
}
func resourceVrrpAForceSelfStandbyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAForceSelfStandbyCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAForceSelfStandby(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAForceSelfStandbyRead(ctx, d, meta)
	}
	return diags
}

func resourceVrrpAForceSelfStandbyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAForceSelfStandbyUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAForceSelfStandby(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVrrpAForceSelfStandbyRead(ctx, d, meta)
	}
	return diags
}
func resourceVrrpAForceSelfStandbyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAForceSelfStandbyDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAForceSelfStandby(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVrrpAForceSelfStandbyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVrrpAForceSelfStandbyRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVrrpAForceSelfStandby(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointVrrpAForceSelfStandby(d *schema.ResourceData) edpt.VrrpAForceSelfStandby {
	var ret edpt.VrrpAForceSelfStandby
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.AllPartitions = d.Get("all_partitions").(int)
	ret.Inst.SkipCheck = d.Get("skip_check").(int)
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
