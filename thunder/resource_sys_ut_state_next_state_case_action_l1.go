package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtStateNextStateCaseActionL1() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_state_next_state_case_action_l1`: L1 packet paramters\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtStateNextStateCaseActionL1Create,
		UpdateContext: resourceSysUtStateNextStateCaseActionL1Update,
		ReadContext:   resourceSysUtStateNextStateCaseActionL1Read,
		DeleteContext: resourceSysUtStateNextStateCaseActionL1Delete,

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
			"case_number": {
				Type: schema.TypeString, Required: true, Description: "CaseNumber",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
			"direction": {
				Type: schema.TypeString, Required: true, Description: "Direction",
			},
		},
	}
}
func resourceSysUtStateNextStateCaseActionL1Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionL1Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionL1(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtStateNextStateCaseActionL1Read(ctx, d, meta)
	}
	return diags
}

func resourceSysUtStateNextStateCaseActionL1Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionL1Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionL1(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtStateNextStateCaseActionL1Read(ctx, d, meta)
	}
	return diags
}
func resourceSysUtStateNextStateCaseActionL1Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionL1Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionL1(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtStateNextStateCaseActionL1Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionL1Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionL1(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSysUtStateNextStateCaseActionL1EthList(d []interface{}) []edpt.SysUtStateNextStateCaseActionL1EthList {

	count1 := len(d)
	ret := make([]edpt.SysUtStateNextStateCaseActionL1EthList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateNextStateCaseActionL1EthList
		oi.EthernetStart = in["ethernet_start"].(int)
		oi.EthernetEnd = in["ethernet_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSysUtStateNextStateCaseActionL1Trunk_list(d []interface{}) []edpt.SysUtStateNextStateCaseActionL1Trunk_list {

	count1 := len(d)
	ret := make([]edpt.SysUtStateNextStateCaseActionL1Trunk_list, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtStateNextStateCaseActionL1Trunk_list
		oi.TrunkStart = in["trunk_start"].(int)
		oi.TrunkEnd = in["trunk_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSysUtStateNextStateCaseActionL1(d *schema.ResourceData) edpt.SysUtStateNextStateCaseActionL1 {
	var ret edpt.SysUtStateNextStateCaseActionL1
	ret.Inst.Auto = d.Get("auto").(int)
	ret.Inst.EthList = getSliceSysUtStateNextStateCaseActionL1EthList(d.Get("eth_list").([]interface{}))
	ret.Inst.Length = d.Get("length").(int)
	ret.Inst.Trunk_list = getSliceSysUtStateNextStateCaseActionL1Trunk_list(d.Get("trunk_list").([]interface{}))
	//omit uuid
	ret.Inst.Value = d.Get("value").(int)
	ret.Inst.CaseNumber = d.Get("case_number").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Direction = d.Get("direction").(string)
	return ret
}
