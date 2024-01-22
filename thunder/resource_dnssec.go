package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDnssec() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_dnssec`: Domain Name System Security Extensions commands\n\n__PLACEHOLDER__",
		CreateContext: resourceDnssecCreate,
		UpdateContext: resourceDnssecUpdate,
		ReadContext:   resourceDnssecRead,
		DeleteContext: resourceDnssecDelete,

		Schema: map[string]*schema.Schema{
			"dnskey": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key_delete": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete the DNSKEY file",
						},
						"zone_name": {
							Type: schema.TypeString, Optional: true, Description: "DNS zone name of the child zone",
						},
					},
				},
			},
			"ds": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ds_delete": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete the DS file",
						},
						"zone_name": {
							Type: schema.TypeString, Optional: true, Description: "DNS zone name of the child zone",
						},
					},
				},
			},
			"key_rollover": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zone_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify the name for the DNS zone",
						},
						"dnssec_key_type": {
							Type: schema.TypeString, Optional: true, Description: "'ZSK': Zone Signing Key; 'KSK': Key Signing Key;",
						},
						"zsk_start": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "start ZSK rollover in emergency mode",
						},
						"ksk_start": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "start KSK rollover in emergency mode",
						},
						"ds_ready_in_parent_zone": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "DS RR is already ready in the parent zone",
						},
					},
				},
			},
			"sign_zone_now": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zone_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify the name for the DNS zone, empty means sign all zones",
						},
					},
				},
			},
			"standalone": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Run DNSSEC in standalone mode, in GSLB group mode by default",
			},
			"template_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dnssec_temp_name": {
							Type: schema.TypeString, Required: true, Description: "DNSSEC Template Name",
						},
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
						"enable_nsec3": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "enable NSEC3 support. disabled by default",
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
						"hsm": {
							Type: schema.TypeString, Optional: true, Description: "specify the HSM template",
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
		},
	}
}
func resourceDnssecCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssec(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDnssecRead(ctx, d, meta)
	}
	return diags
}

func resourceDnssecUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssec(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDnssecRead(ctx, d, meta)
	}
	return diags
}
func resourceDnssecDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssec(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDnssecRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDnssecRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDnssec(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDnssecDnskey343(d []interface{}) edpt.DnssecDnskey343 {

	count1 := len(d)
	var ret edpt.DnssecDnskey343
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.KeyDelete = in["key_delete"].(int)
		ret.ZoneName = in["zone_name"].(string)
	}
	return ret
}

func getObjectDnssecDs344(d []interface{}) edpt.DnssecDs344 {

	count1 := len(d)
	var ret edpt.DnssecDs344
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DsDelete = in["ds_delete"].(int)
		ret.ZoneName = in["zone_name"].(string)
	}
	return ret
}

func getObjectDnssecKeyRollover345(d []interface{}) edpt.DnssecKeyRollover345 {

	count1 := len(d)
	var ret edpt.DnssecKeyRollover345
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneName = in["zone_name"].(string)
		ret.DnssecKeyType = in["dnssec_key_type"].(string)
		ret.ZskStart = in["zsk_start"].(int)
		ret.KskStart = in["ksk_start"].(int)
		ret.DsReadyInParentZone = in["ds_ready_in_parent_zone"].(int)
	}
	return ret
}

func getObjectDnssecSignZoneNow346(d []interface{}) edpt.DnssecSignZoneNow346 {

	count1 := len(d)
	var ret edpt.DnssecSignZoneNow346
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ZoneName = in["zone_name"].(string)
	}
	return ret
}

func getSliceDnssecTemplateList(d []interface{}) []edpt.DnssecTemplateList {

	count1 := len(d)
	ret := make([]edpt.DnssecTemplateList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DnssecTemplateList
		oi.DnssecTempName = in["dnssec_temp_name"].(string)
		oi.Algorithm = in["algorithm"].(string)
		oi.CombinationsLimit = in["combinations_limit"].(int)
		oi.DnskeyTtlK = in["dnskey_ttl_k"].(int)
		oi.DnskeyTtlV = in["dnskey_ttl_v"].(int)
		oi.EnableNsec3 = in["enable_nsec3"].(int)
		oi.ReturnNsecOnFailure = in["return_nsec_on_failure"].(int)
		oi.SignatureValidityPeriodK = in["signature_validity_period_k"].(int)
		oi.SignatureValidityPeriodV = in["signature_validity_period_v"].(int)
		oi.Hsm = in["hsm"].(string)
		oi.DnssecTemplateZsk = getObjectDnssecTemplateListDnssecTemplateZsk(in["dnssec_template_zsk"].([]interface{}))
		oi.DnssecTemplateKsk = getObjectDnssecTemplateListDnssecTemplateKsk(in["dnssec_template_ksk"].([]interface{}))
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDnssecTemplateListDnssecTemplateZsk(d []interface{}) edpt.DnssecTemplateListDnssecTemplateZsk {

	count1 := len(d)
	var ret edpt.DnssecTemplateListDnssecTemplateZsk
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

func getObjectDnssecTemplateListDnssecTemplateKsk(d []interface{}) edpt.DnssecTemplateListDnssecTemplateKsk {

	count1 := len(d)
	var ret edpt.DnssecTemplateListDnssecTemplateKsk
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

func dataToEndpointDnssec(d *schema.ResourceData) edpt.Dnssec {
	var ret edpt.Dnssec
	ret.Inst.Dnskey = getObjectDnssecDnskey343(d.Get("dnskey").([]interface{}))
	ret.Inst.Ds = getObjectDnssecDs344(d.Get("ds").([]interface{}))
	ret.Inst.KeyRollover = getObjectDnssecKeyRollover345(d.Get("key_rollover").([]interface{}))
	ret.Inst.SignZoneNow = getObjectDnssecSignZoneNow346(d.Get("sign_zone_now").([]interface{}))
	ret.Inst.Standalone = d.Get("standalone").(int)
	ret.Inst.TemplateList = getSliceDnssecTemplateList(d.Get("template_list").([]interface{}))
	//omit uuid
	return ret
}
