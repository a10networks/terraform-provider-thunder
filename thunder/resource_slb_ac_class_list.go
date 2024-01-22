package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbAcClassList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_ac_class_list`: Aho-Corasic add remove config\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbAcClassListCreate,
		UpdateContext: resourceSlbAcClassListUpdate,
		ReadContext:   resourceSlbAcClassListRead,
		DeleteContext: resourceSlbAcClassListDelete,

		Schema: map[string]*schema.Schema{
			"ac_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeString, Optional: true, Description: "'add': Add the entry; 'delete': Delete the entry;",
						},
						"ac_match_type": {
							Type: schema.TypeString, Optional: true, Description: "'contains': String contains another string; 'ends-with': String ends with another string; 'equals': String equals another string; 'starts-with': String starts with another string;",
						},
						"ac_key_string": {
							Type: schema.TypeString, Optional: true, Description: "Specify key string",
						},
						"ac_key_value": {
							Type: schema.TypeString, Optional: true, Description: "Specify value string",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the class list",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
		},
	}
}
func resourceSlbAcClassListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbAcClassListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbAcClassList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbAcClassListRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbAcClassListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbAcClassListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbAcClassList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbAcClassListRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbAcClassListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbAcClassListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbAcClassList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbAcClassListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbAcClassListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbAcClassList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbAcClassListAcList(d []interface{}) []edpt.SlbAcClassListAcList {

	count1 := len(d)
	ret := make([]edpt.SlbAcClassListAcList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbAcClassListAcList
		oi.Action = in["action"].(string)
		oi.AcMatchType = in["ac_match_type"].(string)
		oi.AcKeyString = in["ac_key_string"].(string)
		oi.AcKeyValue = in["ac_key_value"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbAcClassList(d *schema.ResourceData) edpt.SlbAcClassList {
	var ret edpt.SlbAcClassList
	ret.Inst.AcList = getSliceSlbAcClassListAcList(d.Get("ac_list").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	return ret
}
