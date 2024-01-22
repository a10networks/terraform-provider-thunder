package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtTemplateL1() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_template_l1`: L1 packet paramters\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtTemplateL1Create,
		UpdateContext: resourceSysUtTemplateL1Update,
		ReadContext:   resourceSysUtTemplateL1Read,
		DeleteContext: resourceSysUtTemplateL1Delete,

		Schema: map[string]*schema.Schema{
			"auto": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Auto calculate pkt len",
			},
			"drop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet drop. Only allowed for output spec",
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSysUtTemplateL1Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL1Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL1(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateL1Read(ctx, d, meta)
	}
	return diags
}

func resourceSysUtTemplateL1Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL1Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL1(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateL1Read(ctx, d, meta)
	}
	return diags
}
func resourceSysUtTemplateL1Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL1Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL1(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtTemplateL1Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateL1Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateL1(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSysUtTemplateL1EthList(d []interface{}) []edpt.SysUtTemplateL1EthList {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateL1EthList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateL1EthList
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSysUtTemplateL1Trunk_list(d []interface{}) []edpt.SysUtTemplateL1Trunk_list {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateL1Trunk_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateL1Trunk_list
		oi.TrunkStart = in["trunk_start"].(int)
		oi.TrunkEnd = in["trunk_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSysUtTemplateL1(d *schema.ResourceData) edpt.SysUtTemplateL1 {
	var ret edpt.SysUtTemplateL1
	ret.Inst.Auto = d.Get("auto").(int)
	ret.Inst.Drop = d.Get("drop").(int)
	ret.Inst.EthList = getSliceSysUtTemplateL1EthList(d.Get("eth_list").([]interface{}))
	ret.Inst.Length = d.Get("length").(int)
	ret.Inst.Trunk_list = getSliceSysUtTemplateL1Trunk_list(d.Get("trunk_list").([]interface{}))
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
