package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsRpz() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_rpz`: Response Policy Zone File\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsRpzCreate,
		UpdateContext: resourceSlbTemplateDnsRpzUpdate,
		ReadContext:   resourceSlbTemplateDnsRpzRead,
		DeleteContext: resourceSlbTemplateDnsRpzDelete,

		Schema: map[string]*schema.Schema{
			"logging": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Log RPZ triggered action",
						},
						"rpz_action": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"str_rpz_action": {
										Type: schema.TypeString, Optional: true, Description: "'drop': Log RPZ due to drop action; 'pass-thru': Log RPZ due to pass-thru action; 'nxdomain': Log RPZ due to nxdomain action; 'nodata': Log RPZ due to nodata action; 'tcp-only': Log RPZ due to tcp-only action; 'local-data': Log RPZ due to local-data action;",
									},
								},
							},
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Optional: true, Description: "Specify a Response Policy Zone name",
			},
			"seq_id": {
				Type: schema.TypeInt, Required: true, Description: "sequential id of RPZ",
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
func resourceSlbTemplateDnsRpzCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRpzCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRpz(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsRpzRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsRpzUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRpzUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRpz(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsRpzRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsRpzDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRpzDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRpz(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsRpzRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRpzRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRpz(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbTemplateDnsRpzLogging1422(d []interface{}) edpt.SlbTemplateDnsRpzLogging1422 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsRpzLogging1422
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Enable = in["enable"].(int)
		ret.RpzAction = getSliceSlbTemplateDnsRpzLoggingRpzAction1423(in["rpz_action"].([]interface{}))
		//omit uuid
	}
	return ret
}

func getSliceSlbTemplateDnsRpzLoggingRpzAction1423(d []interface{}) []edpt.SlbTemplateDnsRpzLoggingRpzAction1423 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsRpzLoggingRpzAction1423, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsRpzLoggingRpzAction1423
		oi.StrRpzAction = in["str_rpz_action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateDnsRpz(d *schema.ResourceData) edpt.SlbTemplateDnsRpz {
	var ret edpt.SlbTemplateDnsRpz
	ret.Inst.Logging = getObjectSlbTemplateDnsRpzLogging1422(d.Get("logging").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.SeqId = d.Get("seq_id").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
