package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtEventActionL1() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_event_action_l1`: L1 packet paramters\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtEventActionL1Create,
		UpdateContext: resourceSysUtEventActionL1Update,
		ReadContext:   resourceSysUtEventActionL1Read,
		DeleteContext: resourceSysUtEventActionL1Delete,

		Schema: map[string]*schema.Schema{
			"auto": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Auto calculate pkt len",
			},
			"eth_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ethernet_start": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet port (Interface number)",
						},
						"ethernet_end": {
							Type: schema.TypeInt, Optional: true, Description: "Ethernet port",
						},
					},
				},
			},
			"length": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "packet length",
			},
			"trunk_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trunk_start": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk groups",
						},
						"trunk_end": {
							Type: schema.TypeInt, Optional: true, Description: "Trunk Group",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value": {
				Type: schema.TypeInt, Optional: true, Description: "Total packet length starting at L2 header",
			},
			"event_number": {
				Type: schema.TypeString, Required: true, Description: "EventNumber",
			},
			"direction": {
				Type: schema.TypeString, Required: true, Description: "Direction",
			},
		},
	}
}
func resourceSysUtEventActionL1Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionL1Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionL1(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtEventActionL1Read(ctx, d, meta)
	}
	return diags
}

func resourceSysUtEventActionL1Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionL1Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionL1(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtEventActionL1Read(ctx, d, meta)
	}
	return diags
}
func resourceSysUtEventActionL1Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionL1Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionL1(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtEventActionL1Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtEventActionL1Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtEventActionL1(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSysUtEventActionL1EthList(d []interface{}) []edpt.SysUtEventActionL1EthList {

	count1 := len(d)
	ret := make([]edpt.SysUtEventActionL1EthList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtEventActionL1EthList
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSysUtEventActionL1Trunk_list(d []interface{}) []edpt.SysUtEventActionL1Trunk_list {

	count1 := len(d)
	ret := make([]edpt.SysUtEventActionL1Trunk_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtEventActionL1Trunk_list
		oi.TrunkStart = in["trunk_start"].(int)
		oi.TrunkEnd = in["trunk_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSysUtEventActionL1(d *schema.ResourceData) edpt.SysUtEventActionL1 {
	var ret edpt.SysUtEventActionL1
	ret.Inst.Auto = d.Get("auto").(int)
	ret.Inst.EthList = getSliceSysUtEventActionL1EthList(d.Get("eth_list").([]interface{}))
	ret.Inst.Length = d.Get("length").(int)
	ret.Inst.Trunk_list = getSliceSysUtEventActionL1Trunk_list(d.Get("trunk_list").([]interface{}))
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	ret.Inst.EventNumber = d.Get("event_number").(string)
	ret.Inst.Direction = d.Get("direction").(string)
	return ret
}
