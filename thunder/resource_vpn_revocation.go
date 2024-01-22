package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVpnRevocation() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vpn_revocation`: IPsec VPN revocation settings\n\n__PLACEHOLDER__",
		CreateContext: resourceVpnRevocationCreate,
		UpdateContext: resourceVpnRevocationUpdate,
		ReadContext:   resourceVpnRevocationRead,
		DeleteContext: resourceVpnRevocationDelete,

		Schema: map[string]*schema.Schema{
			"ca": {
				Type: schema.TypeString, Optional: true, Description: "Certificate Authority file name",
			},
			"crl": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crl_pri": {
							Type: schema.TypeString, Optional: true, Description: "Primary CRL URL (http://www.example.com/ocsp) (only .der filetypes)",
						},
						"crl_sec": {
							Type: schema.TypeString, Optional: true, Description: "Secondary CRL URL (http://www.example.com/ocsp) (only .der filetypes)",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Revocation name",
			},
			"ocsp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ocsp_pri": {
							Type: schema.TypeString, Optional: true, Description: "Primary OCSP Authentication Server",
						},
						"ocsp_sec": {
							Type: schema.TypeString, Optional: true, Description: "Secondary OCSP Authentication Server",
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
func resourceVpnRevocationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnRevocationCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnRevocation(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVpnRevocationRead(ctx, d, meta)
	}
	return diags
}

func resourceVpnRevocationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnRevocationUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnRevocation(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVpnRevocationRead(ctx, d, meta)
	}
	return diags
}
func resourceVpnRevocationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnRevocationDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnRevocation(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVpnRevocationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVpnRevocationRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVpnRevocation(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectVpnRevocationCrl(d []interface{}) edpt.VpnRevocationCrl {

	count1 := len(d)
	var ret edpt.VpnRevocationCrl
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CrlPri = in["crl_pri"].(string)
		ret.CrlSec = in["crl_sec"].(string)
	}
	return ret
}

func getObjectVpnRevocationOcsp(d []interface{}) edpt.VpnRevocationOcsp {

	count1 := len(d)
	var ret edpt.VpnRevocationOcsp
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.OcspPri = in["ocsp_pri"].(string)
		ret.OcspSec = in["ocsp_sec"].(string)
	}
	return ret
}

func dataToEndpointVpnRevocation(d *schema.ResourceData) edpt.VpnRevocation {
	var ret edpt.VpnRevocation
	ret.Inst.Ca = d.Get("ca").(string)
	ret.Inst.Crl = getObjectVpnRevocationCrl(d.Get("crl").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Ocsp = getObjectVpnRevocationOcsp(d.Get("ocsp").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
