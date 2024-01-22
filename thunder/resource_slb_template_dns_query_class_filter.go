package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsQueryClassFilter() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_query_class_filter`: DNS query class filter list\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsQueryClassFilterCreate,
		UpdateContext: resourceSlbTemplateDnsQueryClassFilterUpdate,
		ReadContext:   resourceSlbTemplateDnsQueryClassFilterRead,
		DeleteContext: resourceSlbTemplateDnsQueryClassFilterDelete,

		Schema: map[string]*schema.Schema{
			"query_class": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"str_query_class": {
							Type: schema.TypeString, Optional: true, Description: "'INTERNET': INTERNET query class; 'CHAOS': CHAOS query class; 'HESIOD': HESIOD query class; 'NONE': NONE query class; 'ANY': ANY query class;",
						},
						"num_query_class": {
							Type: schema.TypeInt, Optional: true, Description: "Other query class value",
						},
					},
				},
			},
			"query_class_action": {
				Type: schema.TypeString, Required: true, Description: "'allow': Allow only certain DNS query classes; 'deny': Deny only certain DNS query classes;",
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
func resourceSlbTemplateDnsQueryClassFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsQueryClassFilterCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsQueryClassFilter(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsQueryClassFilterRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsQueryClassFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsQueryClassFilterUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsQueryClassFilter(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsQueryClassFilterRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsQueryClassFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsQueryClassFilterDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsQueryClassFilter(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsQueryClassFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsQueryClassFilterRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsQueryClassFilter(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateDnsQueryClassFilterQueryClass(d []interface{}) []edpt.SlbTemplateDnsQueryClassFilterQueryClass {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsQueryClassFilterQueryClass, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsQueryClassFilterQueryClass
		oi.StrQueryClass = in["str_query_class"].(string)
		oi.NumQueryClass = in["num_query_class"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateDnsQueryClassFilter(d *schema.ResourceData) edpt.SlbTemplateDnsQueryClassFilter {
	var ret edpt.SlbTemplateDnsQueryClassFilter
	ret.Inst.QueryClass = getSliceSlbTemplateDnsQueryClassFilterQueryClass(d.Get("query_class").([]interface{}))
	ret.Inst.QueryClassAction = d.Get("query_class_action").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
