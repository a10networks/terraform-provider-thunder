package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRuleSetRuleOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_rule_set_rule_oper`: Operational Status for the object rule\n\n__PLACEHOLDER__",
		ReadContext: resourceRuleSetRuleOperRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString, Required: true, Description: "Rule name",
			},
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hitcount": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"last_hitcount_time": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"action": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"permitbytes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"denybytes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"resetbytes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"totalbytes": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"permitpackets": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"denypackets": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"resetpackets": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"totalpackets": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"activesessiontcp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"activesessionudp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"activesessionicmp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"activesessionsctp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"activesessionother": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"activesessiontotal": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sessiontcp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sessionudp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sessionicmp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sessionsctp": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sessionother": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sessiontotal": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"ratelimitdrops": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceRuleSetRuleOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRuleSetRuleOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRuleSetRuleOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		RuleSetRuleOperOper := setObjectRuleSetRuleOperOper(res)
		d.Set("oper", RuleSetRuleOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectRuleSetRuleOperOper(ret edpt.DataRuleSetRuleOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hitcount":           ret.DtRuleSetRuleOper.Oper.Hitcount,
			"last_hitcount_time": ret.DtRuleSetRuleOper.Oper.LastHitcountTime,
			"action":             ret.DtRuleSetRuleOper.Oper.Action,
			"status":             ret.DtRuleSetRuleOper.Oper.Status,
			"permitbytes":        ret.DtRuleSetRuleOper.Oper.Permitbytes,
			"denybytes":          ret.DtRuleSetRuleOper.Oper.Denybytes,
			"resetbytes":         ret.DtRuleSetRuleOper.Oper.Resetbytes,
			"totalbytes":         ret.DtRuleSetRuleOper.Oper.Totalbytes,
			"permitpackets":      ret.DtRuleSetRuleOper.Oper.Permitpackets,
			"denypackets":        ret.DtRuleSetRuleOper.Oper.Denypackets,
			"resetpackets":       ret.DtRuleSetRuleOper.Oper.Resetpackets,
			"totalpackets":       ret.DtRuleSetRuleOper.Oper.Totalpackets,
			"activesessiontcp":   ret.DtRuleSetRuleOper.Oper.Activesessiontcp,
			"activesessionudp":   ret.DtRuleSetRuleOper.Oper.Activesessionudp,
			"activesessionicmp":  ret.DtRuleSetRuleOper.Oper.Activesessionicmp,
			"activesessionsctp":  ret.DtRuleSetRuleOper.Oper.Activesessionsctp,
			"activesessionother": ret.DtRuleSetRuleOper.Oper.Activesessionother,
			"activesessiontotal": ret.DtRuleSetRuleOper.Oper.Activesessiontotal,
			"sessiontcp":         ret.DtRuleSetRuleOper.Oper.Sessiontcp,
			"sessionudp":         ret.DtRuleSetRuleOper.Oper.Sessionudp,
			"sessionicmp":        ret.DtRuleSetRuleOper.Oper.Sessionicmp,
			"sessionsctp":        ret.DtRuleSetRuleOper.Oper.Sessionsctp,
			"sessionother":       ret.DtRuleSetRuleOper.Oper.Sessionother,
			"sessiontotal":       ret.DtRuleSetRuleOper.Oper.Sessiontotal,
			"ratelimitdrops":     ret.DtRuleSetRuleOper.Oper.Ratelimitdrops,
		},
	}
}

func getObjectRuleSetRuleOperOper(d []interface{}) edpt.RuleSetRuleOperOper {

	count1 := len(d)
	var ret edpt.RuleSetRuleOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Hitcount = in["hitcount"].(int)
		ret.LastHitcountTime = in["last_hitcount_time"].(string)
		ret.Action = in["action"].(string)
		ret.Status = in["status"].(string)
		ret.Permitbytes = in["permitbytes"].(int)
		ret.Denybytes = in["denybytes"].(int)
		ret.Resetbytes = in["resetbytes"].(int)
		ret.Totalbytes = in["totalbytes"].(int)
		ret.Permitpackets = in["permitpackets"].(int)
		ret.Denypackets = in["denypackets"].(int)
		ret.Resetpackets = in["resetpackets"].(int)
		ret.Totalpackets = in["totalpackets"].(int)
		ret.Activesessiontcp = in["activesessiontcp"].(int)
		ret.Activesessionudp = in["activesessionudp"].(int)
		ret.Activesessionicmp = in["activesessionicmp"].(int)
		ret.Activesessionsctp = in["activesessionsctp"].(int)
		ret.Activesessionother = in["activesessionother"].(int)
		ret.Activesessiontotal = in["activesessiontotal"].(int)
		ret.Sessiontcp = in["sessiontcp"].(int)
		ret.Sessionudp = in["sessionudp"].(int)
		ret.Sessionicmp = in["sessionicmp"].(int)
		ret.Sessionsctp = in["sessionsctp"].(int)
		ret.Sessionother = in["sessionother"].(int)
		ret.Sessiontotal = in["sessiontotal"].(int)
		ret.Ratelimitdrops = in["ratelimitdrops"].(int)
	}
	return ret
}

func dataToEndpointRuleSetRuleOper(d *schema.ResourceData) edpt.RuleSetRuleOper {
	var ret edpt.RuleSetRuleOper

	ret.Name = d.Get("name").(string)

	ret.Oper = getObjectRuleSetRuleOperOper(d.Get("oper").([]interface{}))
	return ret
}
