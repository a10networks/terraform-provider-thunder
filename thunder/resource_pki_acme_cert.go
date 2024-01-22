package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePkiAcmeCert() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_pki_acme_cert`: ACME Certificate enrollment object\n\n__PLACEHOLDER__",
		CreateContext: resourcePkiAcmeCertCreate,
		UpdateContext: resourcePkiAcmeCertUpdate,
		ReadContext:   resourcePkiAcmeCertRead,
		DeleteContext: resourcePkiAcmeCertDelete,

		Schema: map[string]*schema.Schema{
			"cert_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the type of certificate",
			},
			"domain": {
				Type: schema.TypeString, Optional: true, Description: "Main domain you want to issue the cert for. CA will verify whether you control this domain",
			},
			"eab_hmac_key": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "The HMAC key for ACME External Account Binding",
			},
			"eab_key_id": {
				Type: schema.TypeString, Optional: true, Description: "The key identifier for ACME External Account Binding",
			},
			"ec_key_length": {
				Type: schema.TypeString, Optional: true, Default: "384", Description: "'256': Key size 256 bits; '384': Key size 384 bits(default);",
			},
			"ecdsa_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "ECDSA certificate",
			},
			"email": {
				Type: schema.TypeString, Optional: true, Description: "A valid email address for your ACME account. CA uses this email to send you expiration or other notices",
			},
			"enroll": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Initiates enrollment with CA. Due to CA rate limit, A10 strongly recommend you set \"run-with-staging-server\" during test",
			},
			"force": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignore the next renewal time and force to renew cert",
			},
			"log_level": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Level for logging output of ACME commands(default 1 and detailed 2, including debug messages)",
			},
			"minute": {
				Type: schema.TypeInt, Optional: true, Description: "Periodic interval in minutes",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify Certificate name to be enrolled",
			},
			"renew_before": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify interval before certificate expiry to renew the certificate",
			},
			"renew_before_type": {
				Type: schema.TypeString, Optional: true, Description: "'hour': Number of hours before cert expiry; 'day': Number of days before cert expiry; 'week': Number of weeks before cert expiry; 'month': Number of months before cert expiry(1 month=30 days);",
			},
			"renew_before_value": {
				Type: schema.TypeInt, Optional: true, Description: "Value of renewal period",
			},
			"renew_every": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify periodic interval in which to renew the certificate",
			},
			"renew_every_type": {
				Type: schema.TypeString, Optional: true, Description: "'hour': Periodic interval in hours; 'day': Periodic interval in days; 'week': Periodic interval in weeks; 'month': Periodic interval in months(1 month=30 days);",
			},
			"renew_every_value": {
				Type: schema.TypeInt, Optional: true, Description: "Value of renewal period",
			},
			"rsa_key_length": {
				Type: schema.TypeString, Optional: true, Default: "2048", Description: "'2048': Key size 2048 bits(default); '3072': Key size 3072 bits; '4096': Key size 4096 bits; '8192': Key size 8192 bits;",
			},
			"rsa_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "RSA certificate (default)",
			},
			"san_domain": {
				Type: schema.TypeString, Optional: true, Description: "Subject-alternate-name dns(s) for your cert, sperated by /",
			},
			"secret_string": {
				Type: schema.TypeString, Optional: true, Description: "The HMAC key for ACME External Account Binding",
			},
			"staging": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Run ACME operation with staging server. Due to CA rate limit, A10 strongly recommends you set this during test",
			},
			"staging_url": {
				Type: schema.TypeString, Optional: true, Description: "ACME staging directory URL. By default, use Let's encrypt as CA server",
			},
			"url": {
				Type: schema.TypeString, Optional: true, Description: "ACME directory URL. By default, use Let's encrypt as CA server",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid": {
				Type: schema.TypeInt, Optional: true, Description: "Specify ha VRRP-A vrid. It is used to sync http-01 challenge token",
			},
		},
	}
}
func resourcePkiAcmeCertCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiAcmeCertCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiAcmeCert(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiAcmeCertRead(ctx, d, meta)
	}
	return diags
}

func resourcePkiAcmeCertUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiAcmeCertUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiAcmeCert(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiAcmeCertRead(ctx, d, meta)
	}
	return diags
}
func resourcePkiAcmeCertDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiAcmeCertDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiAcmeCert(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourcePkiAcmeCertRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiAcmeCertRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiAcmeCert(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointPkiAcmeCert(d *schema.ResourceData) edpt.PkiAcmeCert {
	var ret edpt.PkiAcmeCert
	ret.Inst.CertType = d.Get("cert_type").(int)
	ret.Inst.Domain = d.Get("domain").(string)
	ret.Inst.EabHmacKey = d.Get("eab_hmac_key").(int)
	ret.Inst.EabKeyId = d.Get("eab_key_id").(string)
	ret.Inst.EcKeyLength = d.Get("ec_key_length").(string)
	ret.Inst.EcdsaType = d.Get("ecdsa_type").(int)
	ret.Inst.Email = d.Get("email").(string)
	//omit encrypted
	ret.Inst.Enroll = d.Get("enroll").(int)
	ret.Inst.Force = d.Get("force").(int)
	ret.Inst.LogLevel = d.Get("log_level").(int)
	ret.Inst.Minute = d.Get("minute").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RenewBefore = d.Get("renew_before").(int)
	ret.Inst.RenewBeforeType = d.Get("renew_before_type").(string)
	ret.Inst.RenewBeforeValue = d.Get("renew_before_value").(int)
	ret.Inst.RenewEvery = d.Get("renew_every").(int)
	ret.Inst.RenewEveryType = d.Get("renew_every_type").(string)
	ret.Inst.RenewEveryValue = d.Get("renew_every_value").(int)
	ret.Inst.RsaKeyLength = d.Get("rsa_key_length").(string)
	ret.Inst.RsaType = d.Get("rsa_type").(int)
	ret.Inst.SanDomain = d.Get("san_domain").(string)
	ret.Inst.SecretString = d.Get("secret_string").(string)
	ret.Inst.Staging = d.Get("staging").(int)
	ret.Inst.StagingUrl = d.Get("staging_url").(string)
	ret.Inst.Url = d.Get("url").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
