package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTrunkIpv6Rip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_trunk_ipv6_rip`: RIP\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTrunkIpv6RipCreate,
		UpdateContext: resourceInterfaceTrunkIpv6RipUpdate,
		ReadContext:   resourceInterfaceTrunkIpv6RipRead,
		DeleteContext: resourceInterfaceTrunkIpv6RipDelete,

		Schema: map[string]*schema.Schema{
			"split_horizon_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state": {
							Type: schema.TypeString, Optional: true, Default: "poisoned", Description: "'poisoned': Perform split horizon with poisoned reverse; 'disable': Disable split horizon; 'enable': Perform split horizon without poisoned reverse;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceTrunkIpv6RipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6RipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6Rip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpv6RipRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTrunkIpv6RipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6RipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6Rip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpv6RipRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTrunkIpv6RipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6RipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6Rip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTrunkIpv6RipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6RipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6Rip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectInterfaceTrunkIpv6RipSplitHorizonCfg(d []interface{}) edpt.InterfaceTrunkIpv6RipSplitHorizonCfg {

	count1 := len(d)
	var ret edpt.InterfaceTrunkIpv6RipSplitHorizonCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.State = in["state"].(string)
	}
	return ret
}

func dataToEndpointInterfaceTrunkIpv6Rip(d *schema.ResourceData) edpt.InterfaceTrunkIpv6Rip {
	var ret edpt.InterfaceTrunkIpv6Rip
	ret.Inst.SplitHorizonCfg = getObjectInterfaceTrunkIpv6RipSplitHorizonCfg(d.Get("split_horizon_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
