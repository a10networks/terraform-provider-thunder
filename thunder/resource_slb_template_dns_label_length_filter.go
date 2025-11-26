package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsLabelLengthFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_label_length_filter`: fqdn label length filter\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsLabelLengthFilterCreate,
		UpdateContext: resourceSlbTemplateDnsLabelLengthFilterUpdate,
		ReadContext:   resourceSlbTemplateDnsLabelLengthFilterRead,
		DeleteContext: resourceSlbTemplateDnsLabelLengthFilterDelete,

		Schema: map[string]*schema.Schema{
			"drop_log_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable the log when hit the rule",
			},
			"fqdn_label_length": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"length": {
							Type: schema.TypeInt, Optional: true, Description: "fqdn label length",
						},
						"suffix": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"label_length_filter_action": {
				Type: schema.TypeString, Optional: true, Default: "drop", Description: "'drop': drop; 'ignore': ignore;",
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
func resourceSlbTemplateDnsLabelLengthFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLabelLengthFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLabelLengthFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLabelLengthFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsLabelLengthFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLabelLengthFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLabelLengthFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLabelLengthFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsLabelLengthFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLabelLengthFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLabelLengthFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsLabelLengthFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLabelLengthFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLabelLengthFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateDnsLabelLengthFilterFqdnLabelLength(d []interface{}) []edpt.SlbTemplateDnsLabelLengthFilterFqdnLabelLength {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsLabelLengthFilterFqdnLabelLength, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsLabelLengthFilterFqdnLabelLength
		oi.Length = in["length"].(int)
		oi.Suffix = in["suffix"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateDnsLabelLengthFilter(d *schema.ResourceData) edpt.SlbTemplateDnsLabelLengthFilter {
	var ret edpt.SlbTemplateDnsLabelLengthFilter
	ret.Inst.DropLogEnable = d.Get("drop_log_enable").(int)
	ret.Inst.FqdnLabelLength = getSliceSlbTemplateDnsLabelLengthFilterFqdnLabelLength(d.Get("fqdn_label_length").([]interface{}))
	ret.Inst.LabelLengthFilterAction = d.Get("label_length_filter_action").(string)
	//omit uuid
	ret.Inst.Dns_name = d.Get("dns_name").(string)
	return ret
}
