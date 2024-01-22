package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDnsLogging() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_dns_logging`: DNS Logging template\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDnsLoggingCreate,
		UpdateContext: resourceSlbTemplateDnsLoggingUpdate,
		ReadContext:   resourceSlbTemplateDnsLoggingRead,
		DeleteContext: resourceSlbTemplateDnsLoggingDelete,

		Schema: map[string]*schema.Schema{
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable DNS Logging template",
			},
			"dns_logging_protocol": {
				Type: schema.TypeString, Optional: true, Description: "'both': Log DNS over tcp and udp; 'tcp': Log DNS over tcp; 'udp': Log DNS over udp;",
			},
			"dns_logging_request_section": {
				Type: schema.TypeString, Optional: true, Description: "'all': Log DNS header and question section; 'header': Log DNS header information; 'question': Log DNS question section;",
			},
			"dns_logging_response_section": {
				Type: schema.TypeString, Optional: true, Description: "'all': Log DNS header information, answer, authority, additional section content; 'header': Log DNS header information; 'answer': Log DNS header information and answer section content;",
			},
			"dns_logging_type": {
				Type: schema.TypeString, Optional: true, Description: "'query': DNS Query Logging; 'response': DNS Response Logging; 'both': DNS Query and Response Logging;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DNS Logging Template Name",
			},
			"response_type": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "start config the response type detail",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceSlbTemplateDnsLoggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLogging(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLoggingRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDnsLoggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLogging(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDnsLoggingRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDnsLoggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLogging(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDnsLoggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDnsLoggingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDnsLogging(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSlbTemplateDnsLoggingResponseType1417(d []interface{}) edpt.SlbTemplateDnsLoggingResponseType1417 {

	count1 := len(d)
	var ret edpt.SlbTemplateDnsLoggingResponseType1417
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Config = in["config"].(int)
		//omit uuid
		ret.TypeList = getSliceSlbTemplateDnsLoggingResponseTypeTypeList1418(in["type_list"].([]interface{}))
	}
	return ret
}

func getSliceSlbTemplateDnsLoggingResponseTypeTypeList1418(d []interface{}) []edpt.SlbTemplateDnsLoggingResponseTypeTypeList1418 {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateDnsLoggingResponseTypeTypeList1418, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateDnsLoggingResponseTypeTypeList1418
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

func dataToEndpointSlbTemplateDnsLogging(d *schema.ResourceData) edpt.SlbTemplateDnsLogging {
	var ret edpt.SlbTemplateDnsLogging
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.DnsLoggingProtocol = d.Get("dns_logging_protocol").(string)
	ret.Inst.DnsLoggingRequestSection = d.Get("dns_logging_request_section").(string)
	ret.Inst.DnsLoggingResponseSection = d.Get("dns_logging_response_section").(string)
	ret.Inst.DnsLoggingType = d.Get("dns_logging_type").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.ResponseType = getObjectSlbTemplateDnsLoggingResponseType1417(d.Get("response_type").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
