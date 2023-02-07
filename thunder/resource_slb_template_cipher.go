package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSlbTemplateCipher() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_cipher`: SSL Cipher Template\n\n",
		CreateContext: resourceSlbTemplateCipherCreate,
		UpdateContext: resourceSlbTemplateCipherUpdate,
		ReadContext:   resourceSlbTemplateCipherRead,
		DeleteContext: resourceSlbTemplateCipherDelete,
		Schema: map[string]*schema.Schema{
			"cipher_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cipher_suite": {
							Type: schema.TypeString, Optional: true, Description: "Cipher Suite",
							ValidateFunc: validation.StringInSlice([]string{"SSL3_RSA_DES_192_CBC3_SHA", "SSL3_RSA_RC4_128_MD5", "SSL3_RSA_RC4_128_SHA", "TLS1_RSA_AES_128_SHA", "TLS1_RSA_AES_256_SHA", "TLS1_RSA_AES_128_SHA256", "TLS1_RSA_AES_256_SHA256", "TLS1_DHE_RSA_AES_128_GCM_SHA256", "TLS1_DHE_RSA_AES_128_SHA", "TLS1_DHE_RSA_AES_128_SHA256", "TLS1_DHE_RSA_AES_256_GCM_SHA384", "TLS1_DHE_RSA_AES_256_SHA", "TLS1_DHE_RSA_AES_256_SHA256", "TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256", "TLS1_ECDHE_ECDSA_AES_128_SHA", "TLS1_ECDHE_ECDSA_AES_128_SHA256", "TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384", "TLS1_ECDHE_ECDSA_AES_256_SHA", "TLS1_ECDHE_RSA_AES_128_GCM_SHA256", "TLS1_ECDHE_RSA_AES_128_SHA", "TLS1_ECDHE_RSA_AES_128_SHA256", "TLS1_ECDHE_RSA_AES_256_GCM_SHA384", "TLS1_ECDHE_RSA_AES_256_SHA", "TLS1_RSA_AES_128_GCM_SHA256", "TLS1_RSA_AES_256_GCM_SHA384", "TLS1_ECDHE_RSA_AES_256_SHA384", "TLS1_ECDHE_ECDSA_AES_256_SHA384", "TLS1_ECDHE_RSA_CHACHA20_POLY1305_SHA256", "TLS1_ECDHE_ECDSA_CHACHA20_POLY1305_SHA256", "TLS1_DHE_RSA_CHACHA20_POLY1305_SHA256", "TLS1_ECDHE_SM2_WITH_SMS4_SM3", "TLS1_ECDHE_SM2_WITH_SMS4_SHA256", "TLS1_ECDHE_SM2_WITH_SMS4_GCM_SM3"}, false),
						},
						"priority": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Cipher priority (default 1)",
							ValidateFunc: validation.IntBetween(1, 100),
						},
					},
				},
			},
			"cipher13_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cipher13_suite": {
							Type: schema.TypeString, Optional: true, Description: "Cipher Suite",
							ValidateFunc: validation.StringInSlice([]string{"TLS_AES_256_GCM_SHA384", "TLS_CHACHA20_POLY1305_SHA256", "TLS_AES_128_GCM_SHA256"}, false),
						},
						"priority": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Cipher priority (default 1)",
							ValidateFunc: validation.IntBetween(1, 100),
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, ForceNew: true, Description: "Cipher Template Name",
				ValidateFunc: validation.StringLenBetween(1, 127),
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

func resourceSlbTemplateCipherCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCipherCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCipher(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateCipherRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateCipherRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCipherRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCipher(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateCipherUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCipherUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCipher(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateCipherRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateCipherDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateCipherDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateCipher(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSlbTemplateCipherCipherCfg(d []interface{}) []edpt.SlbTemplateCipherCipherCfg {
	count := len(d)
	ret := make([]edpt.SlbTemplateCipherCipherCfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateCipherCipherCfg
		oi.CipherSuite = in["cipher_suite"].(string)
		oi.Priority = in["priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateCipherCipher13Cfg(d []interface{}) []edpt.SlbTemplateCipherCipher13Cfg {
	count := len(d)
	ret := make([]edpt.SlbTemplateCipherCipher13Cfg, 0, count)
	for _, item := range d {
		if item == nil {
			continue
		}
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateCipherCipher13Cfg
		oi.Cipher13Suite = in["cipher13_suite"].(string)
		oi.Priority = in["priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSlbTemplateCipher(d *schema.ResourceData) edpt.SlbTemplateCipher {
	var ret edpt.SlbTemplateCipher
	ret.Inst.CipherCfg = getSliceSlbTemplateCipherCipherCfg(d.Get("cipher_cfg").([]interface{}))
	ret.Inst.Cipher13Cfg = getSliceSlbTemplateCipherCipher13Cfg(d.Get("cipher13_cfg").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
