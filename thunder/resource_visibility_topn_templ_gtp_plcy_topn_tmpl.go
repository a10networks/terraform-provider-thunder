package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVisibilityTopnTemplGtpPlcyTopnTmpl() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_visibility_topn_templ_gtp_plcy_topn_tmpl`: Configure template for template.gtp-policy\n\n__PLACEHOLDER__",
		CreateContext: resourceVisibilityTopnTemplGtpPlcyTopnTmplCreate,
		UpdateContext: resourceVisibilityTopnTemplGtpPlcyTopnTmplUpdate,
		ReadContext:   resourceVisibilityTopnTemplGtpPlcyTopnTmplRead,
		DeleteContext: resourceVisibilityTopnTemplGtpPlcyTopnTmplDelete,

		Schema: map[string]*schema.Schema{
			"interval": {
				Type: schema.TypeString, Optional: true, Description: "'5': 5 minutes; '15': 15 minutes; '30': 30 minutes; '60': 60 minutes; 'all-time': Since template is activated;",
			},
			"metrics": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rl_message_monitor": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Track Top-N entities for GTP Message forwarded via monitor mode at rate-limit policy",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Template Name",
			},
			"topn_size": {
				Type: schema.TypeInt, Optional: true, Description: "Congure value of N for topn",
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
func resourceVisibilityTopnTemplGtpPlcyTopnTmplCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnTemplGtpPlcyTopnTmplCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnTemplGtpPlcyTopnTmpl(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnTemplGtpPlcyTopnTmplRead(ctx, d, meta)
	}
	return diags
}

func resourceVisibilityTopnTemplGtpPlcyTopnTmplUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnTemplGtpPlcyTopnTmplUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnTemplGtpPlcyTopnTmpl(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVisibilityTopnTemplGtpPlcyTopnTmplRead(ctx, d, meta)
	}
	return diags
}
func resourceVisibilityTopnTemplGtpPlcyTopnTmplDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnTemplGtpPlcyTopnTmplDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnTemplGtpPlcyTopnTmpl(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVisibilityTopnTemplGtpPlcyTopnTmplRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVisibilityTopnTemplGtpPlcyTopnTmplRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVisibilityTopnTemplGtpPlcyTopnTmpl(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVisibilityTopnTemplGtpPlcyTopnTmplMetrics3233(d []interface{}) edpt.VisibilityTopnTemplGtpPlcyTopnTmplMetrics3233 {

	count1 := len(d)
	var ret edpt.VisibilityTopnTemplGtpPlcyTopnTmplMetrics3233
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RlMessageMonitor = in["rl_message_monitor"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointVisibilityTopnTemplGtpPlcyTopnTmpl(d *schema.ResourceData) edpt.VisibilityTopnTemplGtpPlcyTopnTmpl {
	var ret edpt.VisibilityTopnTemplGtpPlcyTopnTmpl
	ret.Inst.Interval = d.Get("interval").(string)
	ret.Inst.Metrics = getObjectVisibilityTopnTemplGtpPlcyTopnTmplMetrics3233(d.Get("metrics").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.TopnSize = d.Get("topn_size").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
