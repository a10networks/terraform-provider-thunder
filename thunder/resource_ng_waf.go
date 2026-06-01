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

func getObjectNgWafCustomPage1157(d []interface{}) edpt.NgWafCustomPage1157 {

	var ret edpt.NgWafCustomPage1157
	return ret
}

func getObjectNgWafCustomSignals1158(d []interface{}) edpt.NgWafCustomSignals1158 {

	var ret edpt.NgWafCustomSignals1158
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

func getObjectNgWafStatus1159(d []interface{}) edpt.NgWafStatus1159 {

	var ret edpt.NgWafStatus1159
	return ret
}

func dataToEndpointNgWaf(d *schema.ResourceData) edpt.NgWaf {
    var ret edpt.NgWaf
    added := false

    if _, ok := d.GetOk("custom_page"); ok {
        ret.Inst.CustomPage = edpt.NgWafCustomPage1157{}
        added = true
    }

    if _, ok := d.GetOk("custom_signals"); ok {
        ret.Inst.CustomSignals = edpt.NgWafCustomSignals1158{}
        added = true
    }

    if v, ok := d.GetOk("stats_list"); ok {
        ret.Inst.StatsList = getSliceNgWafStatsList(v.([]interface{}))
        if len(ret.Inst.StatsList) > 0 {
            added = true
        }
    }

    if _, ok := d.GetOk("status"); ok {
        ret.Inst.Status = edpt.NgWafStatus1159{}
        added = true
    }

    // Ensure AXAPI gets at least one object
    if !added {
        ret.Inst.Status = edpt.NgWafStatus1159{}
    }

    return ret
}

