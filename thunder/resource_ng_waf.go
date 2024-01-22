package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNgWaf() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ng_waf`: NGWAF related commands\n\n__PLACEHOLDER__",
		CreateContext: resourceNgWafCreate,
		UpdateContext: resourceNgWafUpdate,
		ReadContext:   resourceNgWafRead,
		DeleteContext: resourceNgWafDelete,

		Schema: map[string]*schema.Schema{
			"cpu": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"custom_page": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"custom_signals": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"stats_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString, Required: true, Description: "ng-waf name",
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
			"status": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceNgWafCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNgWafCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNgWaf(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNgWafRead(ctx, d, meta)
	}
	return diags
}

func resourceNgWafUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNgWafUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNgWaf(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNgWafRead(ctx, d, meta)
	}
	return diags
}
func resourceNgWafDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNgWafDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNgWaf(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNgWafRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNgWafRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNgWaf(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectNgWafCpu1076(d []interface{}) edpt.NgWafCpu1076 {

	var ret edpt.NgWafCpu1076
	return ret
}

func getObjectNgWafCustomPage1077(d []interface{}) edpt.NgWafCustomPage1077 {

	var ret edpt.NgWafCustomPage1077
	return ret
}

func getObjectNgWafCustomSignals1078(d []interface{}) edpt.NgWafCustomSignals1078 {

	var ret edpt.NgWafCustomSignals1078
	return ret
}

func getSliceNgWafStatsList(d []interface{}) []edpt.NgWafStatsList {

	count1 := len(d)
	ret := make([]edpt.NgWafStatsList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.NgWafStatsList
		oi.Name = in["name"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectNgWafStatus1079(d []interface{}) edpt.NgWafStatus1079 {

	var ret edpt.NgWafStatus1079
	return ret
}

func dataToEndpointNgWaf(d *schema.ResourceData) edpt.NgWaf {
	var ret edpt.NgWaf
	ret.Inst.Cpu = getObjectNgWafCpu1076(d.Get("cpu").([]interface{}))
	ret.Inst.CustomPage = getObjectNgWafCustomPage1077(d.Get("custom_page").([]interface{}))
	ret.Inst.CustomSignals = getObjectNgWafCustomSignals1078(d.Get("custom_signals").([]interface{}))
	ret.Inst.StatsList = getSliceNgWafStatsList(d.Get("stats_list").([]interface{}))
	ret.Inst.Status = getObjectNgWafStatus1079(d.Get("status").([]interface{}))
	//omit uuid
	return ret
}
