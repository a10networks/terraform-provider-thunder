package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkVlanGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_vlan_global`: Configure global options for vlan\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkVlanGlobalCreate,
		UpdateContext: resourceNetworkVlanGlobalUpdate,
		ReadContext:   resourceNetworkVlanGlobalRead,
		DeleteContext: resourceNetworkVlanGlobalDelete,

		Schema: map[string]*schema.Schema{
			"enable_def_vlan_l2_forwarding": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable layer 2 forwarding on default vlan",
			},
			"l3_vlan_fwd_disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable L3 forwarding between VLANs",
			},
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'xparent_vlan_list_err': Transparent Mode VLAN List Errors;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetworkVlanGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVlanGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVlanGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVlanGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkVlanGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVlanGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVlanGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkVlanGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkVlanGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVlanGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVlanGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkVlanGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkVlanGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkVlanGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceNetworkVlanGlobalSamplingEnable(d []interface{}) []edpt.NetworkVlanGlobalSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.NetworkVlanGlobalSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkVlanGlobalSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkVlanGlobal(d *schema.ResourceData) edpt.NetworkVlanGlobal {
	var ret edpt.NetworkVlanGlobal
	ret.Inst.EnableDefVlanL2Forwarding = d.Get("enable_def_vlan_l2_forwarding").(int)
	ret.Inst.L3VlanFwdDisable = d.Get("l3_vlan_fwd_disable").(int)
	ret.Inst.SamplingEnable = getSliceNetworkVlanGlobalSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	return ret
}
