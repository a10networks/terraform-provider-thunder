package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDnssecTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_dnssec_template`: template Settings\n\n__PLACEHOLDER__",
		CreateContext: resourceDnssecTemplateCreate,
		UpdateContext: resourceDnssecTemplateUpdate,
		ReadContext:   resourceDnssecTemplateRead,
		DeleteContext: resourceDnssecTemplateDelete,

		Schema: map[string]*schema.Schema{
			"algorithm": {
				Type: schema.TypeString, Optional: true, Description: "'RSASHA1': RSASHA1 algorithm; 'RSASHA256': RSASHA256 algorithm; 'RSASHA512': RSASHA512 algorithm;",
			},
			"combinations_limit": {
				Type: schema.TypeInt, Optional: true, Description: "the max number of combinations per RRset (Default value is 31)",
			},
			"dnskey_ttl_k": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "The TTL value of DNSKEY RR",
			},
			"dnskey_ttl_v": {
				Type: schema.TypeInt, Optional: true, Default: 14400, Description: "in seconds, 14400 seconds by default",
			},
			"dnssec_temp_name": {
				Type: schema.TypeString, Required: true, Description: "DNSSEC Template Name",
			},
			"dnssec_template_ksk": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ksk_keysize_k": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the number of bits in the DNSSEC KSK keys",
						},
						"ksk_keysize_v": {
							Type: schema.TypeInt, Optional: true, Description: "Default size is 2048 and must be an exact multiple of 64",
						},
						"ksk_lifetime_k": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set the lifetime for DNSSEC KSK keys in days",
						},
						"ksk_lifetime_v": {
							Type: schema.TypeInt, Optional: true, Description: "Default value is 365 days",
						},
						"ksk_rollover_time_k": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set the rollover time in days",
						},
						"zsk_rollover_time_v": {
							Type: schema.TypeInt, Optional: true, Default: 358, Description: "7 days less than the lifetime by default",
						},
					},
				},
			},
			"dnssec_template_zsk": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zsk_keysize_k": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the number of bits in the DNSSEC ZSK keys",
						},
						"zsk_keysize_v": {
							Type: schema.TypeInt, Optional: true, Description: "Default size is 2048 and must be an exact multiple of 64",
						},
						"zsk_lifetime_k": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set the lifetime for DNSSEC ZSK keys in days",
						},
						"zsk_lifetime_v": {
							Type: schema.TypeInt, Optional: true, Default: 90, Description: "Default value is 90 days",
						},
						"zsk_rollover_time_k": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set the rollover time in days",
						},
						"zsk_rollover_time_v": {
							Type: schema.TypeInt, Optional: true, Default: 83, Description: "7 days less than the lifetime by default",
						},
					},
				},
			},
			"enable_nsec3": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable NSEC3 support. disabled by default",
			},
			"hsm": {
				Type: schema.TypeString, Optional: true, Description: "specify the HSM template",
			},
			"return_nsec_on_failure": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "return NSEC/NSEC3 or not on failure case. return by default",
			},
			"signature_validity_period_k": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "The period that a signature is valid",
			},
			"signature_validity_period_v": {
				Type: schema.TypeInt, Optional: true, Default: 10, Description: "in days, 10 days by default",
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
func resourceDnssecTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDnssecTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceDnssecTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDnssecTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceDnssecTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDnssecTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssecTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDnssecTemplateDnssecTemplateKsk(d []interface{}) edpt.DnssecTemplateDnssecTemplateKsk {

	count1 := len(d)
	var ret edpt.DnssecTemplateDnssecTemplateKsk
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KskKeysizeK = in["ksk_keysize_k"].(int)
		ret.KskKeysizeV = in["ksk_keysize_v"].(int)
		ret.KskLifetimeK = in["ksk_lifetime_k"].(int)
		ret.KskLifetimeV = in["ksk_lifetime_v"].(int)
		ret.KskRolloverTimeK = in["ksk_rollover_time_k"].(int)
		ret.ZskRolloverTimeV = in["zsk_rollover_time_v"].(int)
	}
	return ret
}

func getObjectDnssecTemplateDnssecTemplateZsk(d []interface{}) edpt.DnssecTemplateDnssecTemplateZsk {

	count1 := len(d)
	var ret edpt.DnssecTemplateDnssecTemplateZsk
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZskKeysizeK = in["zsk_keysize_k"].(int)
		ret.ZskKeysizeV = in["zsk_keysize_v"].(int)
		ret.ZskLifetimeK = in["zsk_lifetime_k"].(int)
		ret.ZskLifetimeV = in["zsk_lifetime_v"].(int)
		ret.ZskRolloverTimeK = in["zsk_rollover_time_k"].(int)
		ret.ZskRolloverTimeV = in["zsk_rollover_time_v"].(int)
	}
	return ret
}

func dataToEndpointDnssecTemplate(d *schema.ResourceData) edpt.DnssecTemplate {
	var ret edpt.DnssecTemplate
	ret.Inst.Algorithm = d.Get("algorithm").(string)
	ret.Inst.CombinationsLimit = d.Get("combinations_limit").(int)
	ret.Inst.DnskeyTtlK = d.Get("dnskey_ttl_k").(int)
	ret.Inst.DnskeyTtlV = d.Get("dnskey_ttl_v").(int)
	ret.Inst.DnssecTempName = d.Get("dnssec_temp_name").(string)
	ret.Inst.DnssecTemplateKsk = getObjectDnssecTemplateDnssecTemplateKsk(d.Get("dnssec_template_ksk").([]interface{}))
	ret.Inst.DnssecTemplateZsk = getObjectDnssecTemplateDnssecTemplateZsk(d.Get("dnssec_template_zsk").([]interface{}))
	ret.Inst.EnableNsec3 = d.Get("enable_nsec3").(int)
	ret.Inst.Hsm = d.Get("hsm").(string)
	ret.Inst.ReturnNsecOnFailure = d.Get("return_nsec_on_failure").(int)
	ret.Inst.SignatureValidityPeriodK = d.Get("signature_validity_period_k").(int)
	ret.Inst.SignatureValidityPeriodV = d.Get("signature_validity_period_v").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
