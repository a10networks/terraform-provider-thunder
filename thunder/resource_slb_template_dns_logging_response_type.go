package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsLoggingResponseType() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_logging_response_type`: Response rdata field length limit type by type\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsLoggingResponseTypeCreate,
		UpdateContext: resourceSlbTemplateDnsLoggingResponseTypeUpdate,
		ReadContext:   resourceSlbTemplateDnsLoggingResponseTypeRead,
		DeleteContext: resourceSlbTemplateDnsLoggingResponseTypeDelete,

		Schema: map[string]*schema.Schema{
			"config": {
				Type: schema.TypeInt, Required: true, Description: "start config the response type detail",
			},
			"type_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"response_type_name": {
							Type: schema.TypeString, Required: true, Description: "'TXT': TXT; 'RRSIG': RRSIG; 'TSIG': TSIG; 'DNSKEY': DNSKEY; 'DS': DS; 'CAA': CAA; 'NAPTR': NAPTR; 'OPT': OPT;",
						},
						"length_limit_flag": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"txt_data": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"txt_type_limit_num": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Limit the field length",
						},
						"txt_type_no_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print the field as much as possible",
						},
						"signature": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"rrsig_type_limit_num": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Limit the field length",
						},
						"rrsig_type_no_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print the field as much as possible",
						},
						"other_data": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"tsig_type_limit_num": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Limit the field length",
						},
						"tsig_type_no_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print the field as much as possible",
						},
						"public_key": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"dnskey_type_limit_num": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Limit the field length",
						},
						"dnskey_type_no_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print the field as much as possible",
						},
						"digest": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"ds_type_limit_num": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Limit the field length",
						},
						"ds_type_no_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print the field as much as possible",
						},
						"value_field": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"caa_type_limit_num": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Limit the field length",
						},
						"caa_type_no_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print the field as much as possible",
						},
						"service_field": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"naptr_type_limit_num": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Limit the field length",
						},
						"naptr_type_no_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print the field as much as possible",
						},
						"rdata_field": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"opt_type_limit_num": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Limit the field length",
						},
						"opt_type_no_limit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print the field as much as possible",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbTemplateDnsLoggingResponseTypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingResponseTypeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingResponseType(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLoggingResponseTypeRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsLoggingResponseTypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingResponseTypeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingResponseType(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLoggingResponseTypeRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsLoggingResponseTypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingResponseTypeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingResponseType(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsLoggingResponseTypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingResponseTypeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingResponseType(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateDnsLoggingResponseTypeTypeList(d []interface{}) []edpt.SlbTemplateDnsLoggingResponseTypeTypeList {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsLoggingResponseTypeTypeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsLoggingResponseTypeTypeList
		oi.ResponseTypeName = in["response_type_name"].(string)
		oi.LengthLimitFlag = in["length_limit_flag"].(int)
		oi.TxtData = in["txt_data"].(int)
		oi.TxtTypeLimitNum = in["txt_type_limit_num"].(int)
		oi.TxtTypeNoLimit = in["txt_type_no_limit"].(int)
		oi.Signature = in["signature"].(int)
		oi.RrsigTypeLimitNum = in["rrsig_type_limit_num"].(int)
		oi.RrsigTypeNoLimit = in["rrsig_type_no_limit"].(int)
		oi.OtherData = in["other_data"].(int)
		oi.TsigTypeLimitNum = in["tsig_type_limit_num"].(int)
		oi.TsigTypeNoLimit = in["tsig_type_no_limit"].(int)
		oi.PublicKey = in["public_key"].(int)
		oi.DnskeyTypeLimitNum = in["dnskey_type_limit_num"].(int)
		oi.DnskeyTypeNoLimit = in["dnskey_type_no_limit"].(int)
		oi.Digest = in["digest"].(int)
		oi.DsTypeLimitNum = in["ds_type_limit_num"].(int)
		oi.DsTypeNoLimit = in["ds_type_no_limit"].(int)
		oi.ValueField = in["value_field"].(int)
		oi.CaaTypeLimitNum = in["caa_type_limit_num"].(int)
		oi.CaaTypeNoLimit = in["caa_type_no_limit"].(int)
		oi.ServiceField = in["service_field"].(int)
		oi.NaptrTypeLimitNum = in["naptr_type_limit_num"].(int)
		oi.NaptrTypeNoLimit = in["naptr_type_no_limit"].(int)
		oi.RdataField = in["rdata_field"].(int)
		oi.OptTypeLimitNum = in["opt_type_limit_num"].(int)
		oi.OptTypeNoLimit = in["opt_type_no_limit"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateDnsLoggingResponseType(d *schema.ResourceData) edpt.SlbTemplateDnsLoggingResponseType {
	var ret edpt.SlbTemplateDnsLoggingResponseType
	ret.Inst.Config = d.Get("config").(int)
	ret.Inst.TypeList = getSliceSlbTemplateDnsLoggingResponseTypeTypeList(d.Get("type_list").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
