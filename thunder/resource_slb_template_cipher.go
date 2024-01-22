package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateCipher() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_cipher`: SSL Cipher Template\n\n__PLACEHOLDER__",
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
							Type: schema.TypeString, Optional: true, Description: "'SSL3_RSA_DES_192_CBC3_SHA': TLS_RSA_WITH_3DES_EDE_CBC_SHA (0x000A); 'SSL3_RSA_RC4_128_MD5': TLS_RSA_WITH_RC4_128_MD5 (0x0004); 'SSL3_RSA_RC4_128_SHA': TLS_RSA_WITH_RC4_128_SHA (0x0005); 'TLS1_RSA_AES_128_SHA': TLS_RSA_WITH_AES_128_CBC_SHA (0x002F); 'TLS1_RSA_AES_256_SHA': TLS_RSA_WITH_AES_256_CBC_SHA (0x0035); 'TLS1_RSA_AES_128_SHA256': TLS_RSA_WITH_AES_128_CBC_SHA256 (0x003C); 'TLS1_RSA_AES_256_SHA256': TLS_RSA_WITH_AES_256_CBC_SHA256 (0x003D); 'TLS1_DHE_RSA_AES_128_GCM_SHA256': TLS_DHE_RSA_WITH_AES_128_GCM_SHA256 (0x009E); 'TLS1_DHE_RSA_AES_128_SHA': TLS_DHE_RSA_WITH_AES_128_CBC_SHA (0x0033); 'TLS1_DHE_RSA_AES_128_SHA256': TLS_DHE_RSA_WITH_AES_128_CBC_SHA256 (0x0067); 'TLS1_DHE_RSA_AES_256_GCM_SHA384': TLS_DHE_RSA_WITH_AES_256_GCM_SHA384 (0x009F); 'TLS1_DHE_RSA_AES_256_SHA': TLS_DHE_RSA_WITH_AES_256_CBC_SHA (0x0039); 'TLS1_DHE_RSA_AES_256_SHA256': TLS_DHE_RSA_WITH_AES_256_CBC_SHA256 (0x006B); 'TLS1_ECDHE_ECDSA_AES_128_GCM_SHA256': TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 (0xC02B); 'TLS1_ECDHE_ECDSA_AES_128_SHA': TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA (0xC009); 'TLS1_ECDHE_ECDSA_AES_128_SHA256': TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256 (0xC023); 'TLS1_ECDHE_ECDSA_AES_256_GCM_SHA384': TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 (0xC02C); 'TLS1_ECDHE_ECDSA_AES_256_SHA': TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA (0xC00A); 'TLS1_ECDHE_RSA_AES_128_GCM_SHA256': TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256 (0xC02F); 'TLS1_ECDHE_RSA_AES_128_SHA': TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA (0xC013); 'TLS1_ECDHE_RSA_AES_128_SHA256': TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256 (0xC027); 'TLS1_ECDHE_RSA_AES_256_GCM_SHA384': TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384 (0xC030); 'TLS1_ECDHE_RSA_AES_256_SHA': TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA (0xC014); 'TLS1_RSA_AES_128_GCM_SHA256': TLS_RSA_WITH_AES_128_GCM_SHA256 (0x009C); 'TLS1_RSA_AES_256_GCM_SHA384': TLS_RSA_WITH_AES_256_GCM_SHA384 (0x009D); 'TLS1_ECDHE_RSA_AES_256_SHA384': TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384 (0xC028); 'TLS1_ECDHE_ECDSA_AES_256_SHA384': TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384 (0xC024); 'TLS1_ECDHE_RSA_CHACHA20_POLY1305_SHA256': TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256 (0xCCA8); 'TLS1_ECDHE_ECDSA_CHACHA20_POLY1305_SHA256': TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 (0xCCA9); 'TLS1_DHE_RSA_CHACHA20_POLY1305_SHA256': TLS_DHE_RSA_WITH_CHACHA20_POLY1305_SHA256 (0xCCAA); 'TLS1_ECDHE_SM2_WITH_SMS4_SM3': TLS_ECDHE_SM2_WITH_SMS4_SM3 (0xE102); 'TLS1_ECDHE_SM2_WITH_SMS4_SHA256': TLS_ECDHE_SM2_WITH_SMS4_SHA256 (0xE105); 'TLS1_ECDHE_SM2_WITH_SMS4_GCM_SM3': TLS_ECDHE_SM2_WITH_SMS4_GCM_SM3 (0xE107);",
						},
						"priority": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Cipher priority (Cipher priority (default 1))",
						},
					},
				},
			},
			"cipher13_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cipher13_suite": {
							Type: schema.TypeString, Optional: true, Description: "'TLS_AES_256_GCM_SHA384': TLS_AES_256_GCM_SHA384 (0x1302); 'TLS_CHACHA20_POLY1305_SHA256': TLS_CHACHA20_POLY1305_SHA256 (0x1303); 'TLS_AES_128_GCM_SHA256': TLS_AES_128_GCM_SHA256 (0x1301);",
						},
						"priority": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Cipher priority (Cipher priority (default 1))",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Cipher Template Name",
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

func getSliceSlbTemplateCipherCipherCfg(d []interface{}) []edpt.SlbTemplateCipherCipherCfg {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateCipherCipherCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SlbTemplateCipherCipherCfg
		oi.CipherSuite = in["cipher_suite"].(string)
		oi.Priority = in["priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSlbTemplateCipherCipher13Cfg(d []interface{}) []edpt.SlbTemplateCipherCipher13Cfg {

	count1 := len(d)
	ret := make([]edpt.SlbTemplateCipherCipher13Cfg, 0, count1)
	for _, item := range d {
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
