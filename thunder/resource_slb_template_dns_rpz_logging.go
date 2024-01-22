package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsRpzLogging() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_rpz_logging`: Log RPZ triggered action\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsRpzLoggingCreate,
		UpdateContext: resourceSlbTemplateDnsRpzLoggingUpdate,
		ReadContext:   resourceSlbTemplateDnsRpzLoggingRead,
		DeleteContext: resourceSlbTemplateDnsRpzLoggingDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Required: true, Description: "Log RPZ triggered action",
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
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
			"seq_id": {
				Type: schema.TypeString, Required: true, Description: "SeqId",
			},
		},
	}
}
func resourceSlbTemplateDnsRpzLoggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRpzLoggingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRpzLogging(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsRpzLoggingRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsRpzLoggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRpzLoggingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRpzLogging(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsRpzLoggingRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsRpzLoggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRpzLoggingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRpzLogging(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsRpzLoggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRpzLoggingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRpzLogging(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateDnsRpzLoggingRpzAction(d []interface{}) []edpt.SlbTemplateDnsRpzLoggingRpzAction {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsRpzLoggingRpzAction, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsRpzLoggingRpzAction
		oi.StrRpzAction = in["str_rpz_action"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateDnsRpzLogging(d *schema.ResourceData) edpt.SlbTemplateDnsRpzLogging {
	var ret edpt.SlbTemplateDnsRpzLogging
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.RpzAction = getSliceSlbTemplateDnsRpzLoggingRpzAction(d.Get("rpz_action").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.SeqId = d.Get("seq_id").(string)
	return ret
}
