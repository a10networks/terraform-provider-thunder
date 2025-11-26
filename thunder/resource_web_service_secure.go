package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWebServiceSecure() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_web_service_secure`: Web-service secure operation\n\n__PLACEHOLDER__",
		CreateContext: resourceWebServiceSecureCreate,
		UpdateContext: resourceWebServiceSecureUpdate,
		ReadContext:   resourceWebServiceSecureRead,
		DeleteContext: resourceWebServiceSecureDelete,

		Schema: map[string]*schema.Schema{
			"certificate": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"load": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load WEB certificate",
						},
						"use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
						},
						"file_url": {
							Type: schema.TypeString, Optional: true, Description: "File URL",
						},
					},
				},
			},
			"generate": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_name": {
							Type: schema.TypeString, Optional: true, Description: "The domain name",
						},
						"country": {
							Type: schema.TypeString, Optional: true, Description: "The country name",
						},
						"state": {
							Type: schema.TypeString, Optional: true, Description: "The location",
						},
					},
				},
			},
			"private_key": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"load": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Load WEB private-key",
						},
						"use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
						},
						"file_url": {
							Type: schema.TypeString, Optional: true, Description: "File URL",
						},
					},
				},
			},
			"regenerate": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_name": {
							Type: schema.TypeString, Optional: true, Description: "The domain name",
						},
						"country": {
							Type: schema.TypeString, Optional: true, Description: "The country name",
						},
						"state": {
							Type: schema.TypeString, Optional: true, Description: "The location",
						},
					},
				},
			},
			"restart": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Restart WEB service",
			},
			"wipe": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Wipe WEB private-key and certificate",
			},
		},
	}
}
func resourceWebServiceSecureCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecureCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecure(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebServiceSecureRead(ctx, d, meta)
	}
	return diags
}

func resourceWebServiceSecureUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecureUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecure(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWebServiceSecureRead(ctx, d, meta)
	}
	return diags
}
func resourceWebServiceSecureDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecureDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecure(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceWebServiceSecureRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWebServiceSecureRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWebServiceSecure(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectWebServiceSecureCertificate3771(d []interface{}) edpt.WebServiceSecureCertificate3771 {

	count1 := len(d)
	var ret edpt.WebServiceSecureCertificate3771
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Load = in["load"].(int)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.FileUrl = in["file_url"].(string)
	}
	return ret
}

func getObjectWebServiceSecureGenerate3772(d []interface{}) edpt.WebServiceSecureGenerate3772 {

	count1 := len(d)
	var ret edpt.WebServiceSecureGenerate3772
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DomainName = in["domain_name"].(string)
		ret.Country = in["country"].(string)
		ret.State = in["state"].(string)
	}
	return ret
}

func getObjectWebServiceSecurePrivateKey3773(d []interface{}) edpt.WebServiceSecurePrivateKey3773 {

	count1 := len(d)
	var ret edpt.WebServiceSecurePrivateKey3773
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Load = in["load"].(int)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.FileUrl = in["file_url"].(string)
	}
	return ret
}

func getObjectWebServiceSecureRegenerate3774(d []interface{}) edpt.WebServiceSecureRegenerate3774 {

	count1 := len(d)
	var ret edpt.WebServiceSecureRegenerate3774
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.DomainName = in["domain_name"].(string)
		ret.Country = in["country"].(string)
		ret.State = in["state"].(string)
	}
	return ret
}

func dataToEndpointWebServiceSecure(d *schema.ResourceData) edpt.WebServiceSecure {
	var ret edpt.WebServiceSecure
	ret.Inst.Certificate = getObjectWebServiceSecureCertificate3771(d.Get("certificate").([]interface{}))
	ret.Inst.Generate = getObjectWebServiceSecureGenerate3772(d.Get("generate").([]interface{}))
	ret.Inst.PrivateKey = getObjectWebServiceSecurePrivateKey3773(d.Get("private_key").([]interface{}))
	ret.Inst.Regenerate = getObjectWebServiceSecureRegenerate3774(d.Get("regenerate").([]interface{}))
	ret.Inst.Restart = d.Get("restart").(int)
	ret.Inst.Wipe = d.Get("wipe").(int)
	return ret
}
