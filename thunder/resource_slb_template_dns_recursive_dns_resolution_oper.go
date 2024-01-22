package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsRecursiveDnsResolutionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_slb_template_dns_recursive_dns_resolution_oper`: Operational Status for the object recursive-dns-resolution\n\n__PLACEHOLDER__",
		ReadContext: resourceSlbTemplateDnsRecursiveDnsResolutionOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gwhc_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"gwhc_up_retries": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"gwhc_down_retries": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}

func resourceSlbTemplateDnsRecursiveDnsResolutionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsRecursiveDnsResolutionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsRecursiveDnsResolutionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SlbTemplateDnsRecursiveDnsResolutionOperOper := setObjectSlbTemplateDnsRecursiveDnsResolutionOperOper(res)
		d.Set("oper", SlbTemplateDnsRecursiveDnsResolutionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSlbTemplateDnsRecursiveDnsResolutionOperOper(ret edpt.DataSlbTemplateDnsRecursiveDnsResolutionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"gwhc_status":       ret.DtSlbTemplateDnsRecursiveDnsResolutionOper.Oper.GwhcStatus,
			"gwhc_up_retries":   ret.DtSlbTemplateDnsRecursiveDnsResolutionOper.Oper.GwhcUpRetries,
			"gwhc_down_retries": ret.DtSlbTemplateDnsRecursiveDnsResolutionOper.Oper.GwhcDownRetries,
		},
	}
}

func getObjectSlbTemplateDnsRecursiveDnsResolutionOperOper(d []interface{}) edpt.SlbTemplateDnsRecursiveDnsResolutionOperOper {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsRecursiveDnsResolutionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GwhcStatus = in["gwhc_status"].(string)
		ret.GwhcUpRetries = in["gwhc_up_retries"].(int)
		ret.GwhcDownRetries = in["gwhc_down_retries"].(int)
	}
	return ret
}

func dataToEndpointSlbTemplateDnsRecursiveDnsResolutionOper(d *schema.ResourceData) edpt.SlbTemplateDnsRecursiveDnsResolutionOper {
	var ret edpt.SlbTemplateDnsRecursiveDnsResolutionOper

	ret.Oper = getObjectSlbTemplateDnsRecursiveDnsResolutionOperOper(d.Get("oper").([]interface{}))

	ret.Name = d.Get("name").(string)
	return ret
}
