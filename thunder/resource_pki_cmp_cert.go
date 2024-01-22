package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePkiCmpCert() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_pki_cmp_cert`: CMP Certificate enrollment object\n\n__PLACEHOLDER__",
		CreateContext: resourcePkiCmpCertCreate,
		UpdateContext: resourcePkiCmpCertUpdate,
		ReadContext:   resourcePkiCmpCertRead,
		DeleteContext: resourcePkiCmpCertDelete,

		Schema: map[string]*schema.Schema{
			"allow_unprotected_errors": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Accept missing or invalid protection of negative responses(CA likes EJCBA tends to not protect negative responses)",
			},
			"cert_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the type of certificate",
			},
			"cmp_trusted_ca": {
				Type: schema.TypeString, Optional: true, Description: "The specific CA to trust while verifying signature of CMP response message",
			},
			"cmp_trusted_cert": {
				Type: schema.TypeString, Optional: true, Description: "The specific CMP server certificate to use and directly trust when verifying signature of CMP response message",
			},
			"ec_key_length": {
				Type: schema.TypeString, Optional: true, Default: "384", Description: "'256': Key size 256 bits; '384': Key size 384 bits(default);",
			},
			"ecdsa_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "ECDSA certificate",
			},
			"enroll": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Initiates enrollment of device with the CA",
			},
			"log_level": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Level for logging output of CMP commands(default 1 and detailed 2)",
			},
			"max_polltime": {
				Type: schema.TypeInt, Optional: true, Default: 120, Description: "Maximum time in seconds a(n) enrollment/key update may take (default 120)",
			},
			"minute": {
				Type: schema.TypeInt, Optional: true, Description: "Periodic interval in minutes",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify Certificate name to be enrolled",
			},
			"recipient_dn": {
				Type: schema.TypeString, Optional: true, Description: "Distinguished Name of the CMP message recipient, i.e., the CMP server (usually a CA or RA entity)) (DN OIDis case sensitive)",
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
				Type: schema.TypeString, Optional: true, Default: "2048", Description: "'1024': Key size 1024 bits; '2048': Key size 2048 bits(default); '4096': Key size 4096 bits; '8192': Key size 8192 bits;",
			},
			"rsa_type": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "RSA certificate (default)",
			},
			"secret": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the pre-shared secret used to enroll the device's certificate",
			},
			"secret_string": {
				Type: schema.TypeString, Optional: true, Description: "pre-shared secret",
			},
			"subject_alternate_name": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"san_type": {
							Type: schema.TypeString, Optional: true, Description: "'email': Enter e-mail address of the subject; 'dns': Enter hostname of the subject; 'ip': Enter IP address of the subject;",
						},
						"san_value": {
							Type: schema.TypeString, Optional: true, Description: "Value of subject-alternate-name",
						},
					},
				},
			},
			"subject_dn": {
				Type: schema.TypeString, Optional: true, Description: "Distinguished Name to use while enrolling the certificate(For EJBCA CA, this is the subject DN of an End Entity) (DN OID is case sensitive)",
			},
			"url": {
				Type: schema.TypeString, Optional: true, Description: "CMP server's absolute URL(http(s)://host:[port]/path), path is the location to use for the CMP server(aka CMP alias)",
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
func resourcePkiCmpCertCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiCmpCertCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiCmpCert(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiCmpCertRead(ctx, d, meta)
	}
	return diags
}

func resourcePkiCmpCertUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiCmpCertUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiCmpCert(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiCmpCertRead(ctx, d, meta)
	}
	return diags
}
func resourcePkiCmpCertDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiCmpCertDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiCmpCert(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourcePkiCmpCertRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiCmpCertRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiCmpCert(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectPkiCmpCertSubjectAlternateName(d []interface{}) edpt.PkiCmpCertSubjectAlternateName {

	count1 := len(d)
	var ret edpt.PkiCmpCertSubjectAlternateName
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SanType = in["san_type"].(string)
		ret.SanValue = in["san_value"].(string)
	}
	return ret
}

func dataToEndpointPkiCmpCert(d *schema.ResourceData) edpt.PkiCmpCert {
	var ret edpt.PkiCmpCert
	ret.Inst.AllowUnprotectedErrors = d.Get("allow_unprotected_errors").(int)
	ret.Inst.CertType = d.Get("cert_type").(int)
	ret.Inst.CmpTrustedCa = d.Get("cmp_trusted_ca").(string)
	ret.Inst.CmpTrustedCert = d.Get("cmp_trusted_cert").(string)
	ret.Inst.EcKeyLength = d.Get("ec_key_length").(string)
	ret.Inst.EcdsaType = d.Get("ecdsa_type").(int)
	//omit encrypted
	ret.Inst.Enroll = d.Get("enroll").(int)
	ret.Inst.LogLevel = d.Get("log_level").(int)
	ret.Inst.MaxPolltime = d.Get("max_polltime").(int)
	ret.Inst.Minute = d.Get("minute").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.RecipientDn = d.Get("recipient_dn").(string)
	ret.Inst.RenewBefore = d.Get("renew_before").(int)
	ret.Inst.RenewBeforeType = d.Get("renew_before_type").(string)
	ret.Inst.RenewBeforeValue = d.Get("renew_before_value").(int)
	ret.Inst.RenewEvery = d.Get("renew_every").(int)
	ret.Inst.RenewEveryType = d.Get("renew_every_type").(string)
	ret.Inst.RenewEveryValue = d.Get("renew_every_value").(int)
	ret.Inst.RsaKeyLength = d.Get("rsa_key_length").(string)
	ret.Inst.RsaType = d.Get("rsa_type").(int)
	ret.Inst.Secret = d.Get("secret").(int)
	ret.Inst.SecretString = d.Get("secret_string").(string)
	ret.Inst.SubjectAlternateName = getObjectPkiCmpCertSubjectAlternateName(d.Get("subject_alternate_name").([]interface{}))
	ret.Inst.SubjectDn = d.Get("subject_dn").(string)
	ret.Inst.Url = d.Get("url").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
