package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkBpduFwdGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_bpdu_fwd_group`: STP BPDU forward Group Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkBpduFwdGroupCreate,
		UpdateContext: resourceNetworkBpduFwdGroupUpdate,
		ReadContext:   resourceNetworkBpduFwdGroupRead,
		DeleteContext: resourceNetworkBpduFwdGroupDelete,

		Schema: map[string]*schema.Schema{
			"bpdu_fwd_group_number": {
				Type: schema.TypeInt, Required: true, Description: "",
			},
			"ethernet_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ethernet_start": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet Port (Interface number)",
						},
						"ethernet_end": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet Port",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNetworkBpduFwdGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkBpduFwdGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkBpduFwdGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkBpduFwdGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkBpduFwdGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkBpduFwdGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkBpduFwdGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkBpduFwdGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkBpduFwdGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkBpduFwdGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkBpduFwdGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkBpduFwdGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkBpduFwdGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkBpduFwdGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceNetworkBpduFwdGroupEthernetList(d []interface{}) []edpt.NetworkBpduFwdGroupEthernetList {

	count1 := len(d)
	ret := make([]edpt.NetworkBpduFwdGroupEthernetList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkBpduFwdGroupEthernetList
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointNetworkBpduFwdGroup(d *schema.ResourceData) edpt.NetworkBpduFwdGroup {
	var ret edpt.NetworkBpduFwdGroup
	ret.Inst.BpduFwdGroupNumber = d.Get("bpdu_fwd_group_number").(int)
	ret.Inst.EthernetList = getSliceNetworkBpduFwdGroupEthernetList(d.Get("ethernet_list").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
