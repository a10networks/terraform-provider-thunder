package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplateCsv() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_template_csv`: Specify csv template\n\n__PLACEHOLDER__",
		CreateContext: resourceTemplateCsvCreate,
		UpdateContext: resourceTemplateCsvUpdate,
		ReadContext:   resourceTemplateCsvRead,
		DeleteContext: resourceTemplateCsvDelete,

		Schema: map[string]*schema.Schema{
			"csv_name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of csv template",
			},
			"delim_char": {
				Type: schema.TypeString, Optional: true, Default: ",", Description: "enter a delimiter character, default \",\"",
			},
			"delim_num": {
				Type: schema.TypeInt, Optional: true, Default: 44, Description: "enter a delimiter number, default 44 (\",\")",
			},
			"ipv6_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Support IPv6 IP ranges",
			},
			"multiple_fields": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": {
							Type: schema.TypeInt, Optional: true, Description: "Field index number (Index of Field)",
						},
						"csv_type": {
							Type: schema.TypeString, Optional: true, Description: "'ip-from': Beginning address of IP range or subnet; 'ip-to-mask': Ending address of IP range or Mask; 'continent': Continent; 'country': Country; 'state': State or province; 'city': City;",
						},
					},
				},
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
func resourceTemplateCsvCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateCsvCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateCsv(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateCsvRead(ctx, d, meta)
	}
	return diags
}

func resourceTemplateCsvUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateCsvUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateCsv(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceTemplateCsvRead(ctx, d, meta)
	}
	return diags
}
func resourceTemplateCsvDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateCsvDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateCsv(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceTemplateCsvRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceTemplateCsvRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointTemplateCsv(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceTemplateCsvMultipleFields(d []interface{}) []edpt.TemplateCsvMultipleFields {

	count1 := len(d)
	ret := make([]edpt.TemplateCsvMultipleFields, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.TemplateCsvMultipleFields
		oi.Field = in["field"].(int)
		oi.CsvType = in["csv_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointTemplateCsv(d *schema.ResourceData) edpt.TemplateCsv {
	var ret edpt.TemplateCsv
	ret.Inst.CsvName = d.Get("csv_name").(string)
	ret.Inst.DelimChar = d.Get("delim_char").(string)
	ret.Inst.DelimNum = d.Get("delim_num").(int)
	ret.Inst.Ipv6Enable = d.Get("ipv6_enable").(int)
	ret.Inst.MultipleFields = getSliceTemplateCsvMultipleFields(d.Get("multiple_fields").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
