package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNetworkSpanningTreeModeMstp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_network_spanning_tree_mode_mstp`: Configure spanning tree protocol MST\n\n__PLACEHOLDER__",
		CreateContext: resourceNetworkSpanningTreeModeMstpCreate,
		UpdateContext: resourceNetworkSpanningTreeModeMstpUpdate,
		ReadContext:   resourceNetworkSpanningTreeModeMstpRead,
		DeleteContext: resourceNetworkSpanningTreeModeMstpDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable spanning tree MSTP mode",
			},
			"instance_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_start": {
							Type: schema.TypeInt, Required: true, Description: "Instance ID",
						},
						"vlan_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vlan_start": {
										Type: schema.TypeInt, Optional: true, Description: "VLAN ID",
									},
									"vlan_end": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"priority": {
							Type: schema.TypeInt, Optional: true, Default: 32768, Description: "Set instance priority",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"priority": {
				Type: schema.TypeInt, Optional: true, Default: 32768, Description: "Set bridge priority",
			},
			"revision_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"revision": {
							Type: schema.TypeInt, Optional: true, Description: "Set MSTP revision level and name",
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "Set MSTP revision name",
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
func resourceNetworkSpanningTreeModeMstpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeMstpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeMstp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkSpanningTreeModeMstpRead(ctx, d, meta)
	}
	return diags
}

func resourceNetworkSpanningTreeModeMstpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeMstpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeMstp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNetworkSpanningTreeModeMstpRead(ctx, d, meta)
	}
	return diags
}
func resourceNetworkSpanningTreeModeMstpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeMstpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeMstp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNetworkSpanningTreeModeMstpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNetworkSpanningTreeModeMstpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNetworkSpanningTreeModeMstp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceNetworkSpanningTreeModeMstpInstanceList(d []interface{}) []edpt.NetworkSpanningTreeModeMstpInstanceList {

	count1 := len(d)
	ret := make([]edpt.NetworkSpanningTreeModeMstpInstanceList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkSpanningTreeModeMstpInstanceList
		oi.InstanceStart = in["instance_start"].(int)
		oi.VlanList = getSliceNetworkSpanningTreeModeMstpInstanceListVlanList(in["vlan_list"].([]interface{}))
		oi.Priority = in["priority"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceNetworkSpanningTreeModeMstpInstanceListVlanList(d []interface{}) []edpt.NetworkSpanningTreeModeMstpInstanceListVlanList {

	count1 := len(d)
	ret := make([]edpt.NetworkSpanningTreeModeMstpInstanceListVlanList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NetworkSpanningTreeModeMstpInstanceListVlanList
		oi.VlanStart = in["vlan_start"].(int)
		oi.VlanEnd = in["vlan_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectNetworkSpanningTreeModeMstpRevisionCfg(d []interface{}) edpt.NetworkSpanningTreeModeMstpRevisionCfg {

	count1 := len(d)
	var ret edpt.NetworkSpanningTreeModeMstpRevisionCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Revision = in["revision"].(int)
		ret.Name = in["name"].(string)
	}
	return ret
}

func dataToEndpointNetworkSpanningTreeModeMstp(d *schema.ResourceData) edpt.NetworkSpanningTreeModeMstp {
	var ret edpt.NetworkSpanningTreeModeMstp
	ret.Inst.Action = d.Get("action").(int)
	ret.Inst.InstanceList = getSliceNetworkSpanningTreeModeMstpInstanceList(d.Get("instance_list").([]interface{}))
	ret.Inst.Priority = d.Get("priority").(int)
	ret.Inst.RevisionCfg = getObjectNetworkSpanningTreeModeMstpRevisionCfg(d.Get("revision_cfg").([]interface{}))
	//omit uuid
	return ret
}
