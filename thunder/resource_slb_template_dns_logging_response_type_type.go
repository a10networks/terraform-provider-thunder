package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsLoggingResponseTypeType() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_logging_response_type_type`: response type name\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsLoggingResponseTypeTypeCreate,
		UpdateContext: resourceSlbTemplateDnsLoggingResponseTypeTypeUpdate,
		ReadContext:   resourceSlbTemplateDnsLoggingResponseTypeTypeRead,
		DeleteContext: resourceSlbTemplateDnsLoggingResponseTypeTypeDelete,

		Schema: map[string]*schema.Schema{
			"caa_type_limit_num": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Limit the field length",
			},
			"caa_type_no_limit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print the field as much as possible",
			},
			"digest": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"dnskey_type_limit_num": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Limit the field length",
			},
			"dnskey_type_no_limit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print the field as much as possible",
			},
			"ds_type_limit_num": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Limit the field length",
			},
			"ds_type_no_limit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print the field as much as possible",
			},
			"length_limit_flag": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"naptr_type_limit_num": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Limit the field length",
			},
			"naptr_type_no_limit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print the field as much as possible",
			},
			"opt_type_limit_num": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Limit the field length",
			},
			"opt_type_no_limit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print the field as much as possible",
			},
			"other_data": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"public_key": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"rdata_field": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"response_type_name": {
				Type: schema.TypeString, Required: true, Description: "'TXT': TXT; 'RRSIG': RRSIG; 'TSIG': TSIG; 'DNSKEY': DNSKEY; 'DS': DS; 'CAA': CAA; 'NAPTR': NAPTR; 'OPT': OPT;",
			},
			"rrsig_type_limit_num": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Limit the field length",
			},
			"rrsig_type_no_limit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print the field as much as possible",
			},
			"service_field": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"signature": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"tsig_type_limit_num": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Limit the field length",
			},
			"tsig_type_no_limit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print the field as much as possible",
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
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"value_field": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbTemplateDnsLoggingResponseTypeTypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingResponseTypeTypeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingResponseTypeType(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLoggingResponseTypeTypeRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsLoggingResponseTypeTypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingResponseTypeTypeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingResponseTypeType(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLoggingResponseTypeTypeRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsLoggingResponseTypeTypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingResponseTypeTypeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingResponseTypeType(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsLoggingResponseTypeTypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingResponseTypeTypeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLoggingResponseTypeType(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateDnsLoggingResponseTypeType(d *schema.ResourceData) edpt.SlbTemplateDnsLoggingResponseTypeType {
	var ret edpt.SlbTemplateDnsLoggingResponseTypeType
	ret.Inst.CaaTypeLimitNum = d.Get("caa_type_limit_num").(int)
	ret.Inst.CaaTypeNoLimit = d.Get("caa_type_no_limit").(int)
	ret.Inst.Digest = d.Get("digest").(int)
	ret.Inst.DnskeyTypeLimitNum = d.Get("dnskey_type_limit_num").(int)
	ret.Inst.DnskeyTypeNoLimit = d.Get("dnskey_type_no_limit").(int)
	ret.Inst.DsTypeLimitNum = d.Get("ds_type_limit_num").(int)
	ret.Inst.DsTypeNoLimit = d.Get("ds_type_no_limit").(int)
	ret.Inst.LengthLimitFlag = d.Get("length_limit_flag").(int)
	ret.Inst.NaptrTypeLimitNum = d.Get("naptr_type_limit_num").(int)
	ret.Inst.NaptrTypeNoLimit = d.Get("naptr_type_no_limit").(int)
	ret.Inst.OptTypeLimitNum = d.Get("opt_type_limit_num").(int)
	ret.Inst.OptTypeNoLimit = d.Get("opt_type_no_limit").(int)
	ret.Inst.OtherData = d.Get("other_data").(int)
	ret.Inst.PublicKey = d.Get("public_key").(int)
	ret.Inst.RdataField = d.Get("rdata_field").(int)
	ret.Inst.ResponseTypeName = d.Get("response_type_name").(string)
	ret.Inst.RrsigTypeLimitNum = d.Get("rrsig_type_limit_num").(int)
	ret.Inst.RrsigTypeNoLimit = d.Get("rrsig_type_no_limit").(int)
	ret.Inst.ServiceField = d.Get("service_field").(int)
	ret.Inst.Signature = d.Get("signature").(int)
	ret.Inst.TsigTypeLimitNum = d.Get("tsig_type_limit_num").(int)
	ret.Inst.TsigTypeNoLimit = d.Get("tsig_type_no_limit").(int)
	ret.Inst.TxtData = d.Get("txt_data").(int)
	ret.Inst.TxtTypeLimitNum = d.Get("txt_type_limit_num").(int)
	ret.Inst.TxtTypeNoLimit = d.Get("txt_type_no_limit").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ValueField = d.Get("value_field").(int)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
