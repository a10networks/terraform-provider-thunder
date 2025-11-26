package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsLabelCountFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_label_count_filter`: fqdn label count filter\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsLabelCountFilterCreate,
		UpdateContext: resourceSlbTemplateDnsLabelCountFilterUpdate,
		ReadContext:   resourceSlbTemplateDnsLabelCountFilterRead,
		DeleteContext: resourceSlbTemplateDnsLabelCountFilterDelete,

		Schema: map[string]*schema.Schema{
			"drop_log_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable the log when hit the rule",
			},
			"label_count_filter_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': drop; 'ignore': ignore;",
			},
			"max_fqdn_label_count": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum number of FQDN labels per FQDN",
			},
			"min_fqdn_label_count": {
				Type: schema.TypeInt, Optional: true, Description: "Minimum number of FQDN labels per FQDN",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"dns_name": {
				Type: schema.TypeString, Required: true, Description: "Dns_name",
			},
		},
	}
}
func resourceSlbTemplateDnsLabelCountFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLabelCountFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLabelCountFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLabelCountFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsLabelCountFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLabelCountFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLabelCountFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLabelCountFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsLabelCountFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLabelCountFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLabelCountFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsLabelCountFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLabelCountFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLabelCountFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateDnsLabelCountFilter(d *schema.ResourceData) edpt.SlbTemplateDnsLabelCountFilter {
	var ret edpt.SlbTemplateDnsLabelCountFilter
	ret.Inst.DropLogEnable = d.Get("drop_log_enable").(int)
	ret.Inst.LabelCountFilterAction = d.Get("label_count_filter_action").(string)
	ret.Inst.MaxFqdnLabelCount = d.Get("max_fqdn_label_count").(int)
	ret.Inst.MinFqdnLabelCount = d.Get("min_fqdn_label_count").(int)
	//omit uuid
	ret.Inst.Dns_name = d.Get("dns_name").(string)
	return ret
}
