package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGslbTemplateCsv() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_template_csv`: Specify csv template\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbTemplateCsvCreate,
		UpdateContext: resourceGslbTemplateCsvUpdate,
		ReadContext:   resourceGslbTemplateCsvRead,
		DeleteContext: resourceGslbTemplateCsvDelete,
		Schema: map[string]*schema.Schema{
			"csv_name": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Specify name of csv template",
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
			"delim_char": {
				Type: schema.TypeString, Optional: true, Default: ",", Description: "enter a delimiter character, default \",\"",
				ValidateFunc:  validation.StringLenBetween(1, 1),
				ConflictsWith: []string{"delim_num"},
			},
			"delim_num": {
				Type: schema.TypeInt, Optional: true, Default: 44, Description: "enter a delimiter number, default 44 (\",\")",
				ValidateFunc:  validation.IntBetween(0, 255),
				ConflictsWith: []string{"delim_char"},
			},
			"ipv6_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Support IPv6 IP ranges",
				ValidateFunc: validation.IntBetween(0, 1),
			},
			"multiple_fields": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": {
							Type: schema.TypeInt, Optional: true, Description: "Field index number (Index of Field)",
							ValidateFunc: validation.IntBetween(1, 64),
						},
						"csv_type": {
							Type: schema.TypeString, Optional: true, Description: "'ip-from': Beginning address of IP range or subnet; 'ip-to-mask': Ending address of IP range or Mask; 'continent': Continent; 'country': Country; 'state': State or province; 'city': City;",
							ValidateFunc: validation.StringInSlice([]string{"ip-from", "ip-to-mask", "continent", "country", "state", "city"}, false),
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
				ValidateFunc: validation.StringLenBetween(1, 127),
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}

func resourceGslbTemplateCsvCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbTemplateCsvCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbTemplateCsv(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbTemplateCsvRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbTemplateCsvRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbTemplateCsvRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbTemplateCsv(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbTemplateCsvUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbTemplateCsvUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbTemplateCsv(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbTemplateCsvRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbTemplateCsvDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbTemplateCsvDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbTemplateCsv(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbTemplateCsvMultipleFields(d []interface{}) []edpt.GslbTemplateCsvMultipleFields {
	count := len(d)
	ret := make([]edpt.GslbTemplateCsvMultipleFields, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbTemplateCsvMultipleFields
		oi.Field = in["field"].(int)
		oi.CsvType = in["csv_type"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbTemplateCsv(d *schema.ResourceData) edpt.GslbTemplateCsv {
	var ret edpt.GslbTemplateCsv
	ret.Inst.CsvName = d.Get("csv_name").(string)
	ret.Inst.DelimChar = d.Get("delim_char").(string)
	ret.Inst.DelimNum = d.Get("delim_num").(int)
	ret.Inst.Ipv6Enable = d.Get("ipv6_enable").(int)
	ret.Inst.MultipleFields = getSliceGslbTemplateCsvMultipleFields(d.Get("multiple_fields").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
