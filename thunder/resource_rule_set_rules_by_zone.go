package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRuleSetRulesByZone() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_rule_set_rules_by_zone`: \n\n__PLACEHOLDER__",
		CreateContext: resourceRuleSetRulesByZoneCreate,
		UpdateContext: resourceRuleSetRulesByZoneUpdate,
		ReadContext:   resourceRuleSetRulesByZoneRead,
		DeleteContext: resourceRuleSetRulesByZoneDelete,

		Schema: map[string]*schema.Schema{
			"sampling_enable": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"counters1": {
							Type: schema.TypeString, Optional: true, Description: "'all': all; 'dummy': Entry for a10countergen;",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceRuleSetRulesByZoneCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRulesByZoneCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRulesByZone(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRuleSetRulesByZoneRead(ctx, d, meta)
	}
	return diags
}

func resourceRuleSetRulesByZoneUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRulesByZoneUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRulesByZone(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRuleSetRulesByZoneRead(ctx, d, meta)
	}
	return diags
}
func resourceRuleSetRulesByZoneDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRulesByZoneDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRulesByZone(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRuleSetRulesByZoneRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRulesByZoneRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRulesByZone(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRuleSetRulesByZoneSamplingEnable(d []interface{}) []edpt.RuleSetRulesByZoneSamplingEnable {

	count1 := len(d)
	ret := make([]edpt.RuleSetRulesByZoneSamplingEnable, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RuleSetRulesByZoneSamplingEnable
		oi.Counters1 = in["counters1"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRuleSetRulesByZone(d *schema.ResourceData) edpt.RuleSetRulesByZone {
	var ret edpt.RuleSetRulesByZone
	ret.Inst.SamplingEnable = getSliceRuleSetRulesByZoneSamplingEnable(d.Get("sampling_enable").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
