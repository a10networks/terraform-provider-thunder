package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePkiScepCert() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_pki_scep_cert`: SCEP Certificate enrollment object\n\n__PLACEHOLDER__",
		CreateContext: resourcePkiScepCertCreate,
		UpdateContext: resourcePkiScepCertUpdate,
		ReadContext:   resourcePkiScepCertRead,
		DeleteContext: resourcePkiScepCertDelete,

		Schema: map[string]*schema.Schema{
			"days": {
				Type: schema.TypeInt, Optional: true, Default: 1825, Description: "Validity of self-signed certificate (default 1825)",
			},
			"dn": {
				Type: schema.TypeString, Optional: true, Description: "Specify the Distinguished-Name to use while enrolling the certificate (Format: \"cn=user, dc=example, dc=com\")",
			},
			"end_date": {
				Type: schema.TypeString, Optional: true, Description: "End date of self-signed certificate in YYMMDDHHMMSS format specified in UTC time",
			},
			"enroll": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Initiates enrollment of device with the CA",
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Interval time in seconds to poll when SCEP response is PENDING (default 5)",
			},
			"key_length": {
				Type: schema.TypeString, Optional: true, Default: "2048", Description: "'1024': Key size 1024 bits; '2048': Key size 2048 bits(default); '4096': Key size 4096 bits; '8192': Key size 8192 bits;",
			},
			"log_level": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "level for logging output of scepclient commands(default 1 and detailed 4)",
			},
			"max_polltime": {
				Type: schema.TypeInt, Optional: true, Default: 180, Description: "Maximum time in seconds to poll when SCEP response is PENDING (default 180)",
			},
			"method": {
				Type: schema.TypeString, Optional: true, Default: "GET", Description: "'GET': GET request; 'POST': POST request;",
			},
			"minute": {
				Type: schema.TypeInt, Optional: true, Description: "Periodic interval in minutes",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify Certificate name to be enrolled",
			},
			"password": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the password used to enroll the device's certificate",
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
			"secret_string": {
				Type: schema.TypeString, Optional: true, Description: "secret password",
			},
			"start_date": {
				Type: schema.TypeString, Optional: true, Description: "Start date of self-signed certificate in YYMMDDHHMMSS format specified in UTC time",
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
			"url": {
				Type: schema.TypeString, Optional: true, Description: "Specify the Enrollment Agent's absolute URL (Format: http://host/path)",
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
func resourcePkiScepCertCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiScepCertCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiScepCert(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiScepCertRead(ctx, d, meta)
	}
	return diags
}

func resourcePkiScepCertUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiScepCertUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiScepCert(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourcePkiScepCertRead(ctx, d, meta)
	}
	return diags
}
func resourcePkiScepCertDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiScepCertDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiScepCert(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourcePkiScepCertRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourcePkiScepCertRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointPkiScepCert(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectPkiScepCertSubjectAlternateName(d []interface{}) edpt.PkiScepCertSubjectAlternateName {

	count1 := len(d)
	var ret edpt.PkiScepCertSubjectAlternateName
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SanType = in["san_type"].(string)
		ret.SanValue = in["san_value"].(string)
	}
	return ret
}

func dataToEndpointPkiScepCert(d *schema.ResourceData) edpt.PkiScepCert {
	var ret edpt.PkiScepCert
	ret.Inst.Days = d.Get("days").(int)
	ret.Inst.Dn = d.Get("dn").(string)
	//omit encrypted
	ret.Inst.EndDate = d.Get("end_date").(string)
	ret.Inst.Enroll = d.Get("enroll").(int)
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.KeyLength = d.Get("key_length").(string)
	ret.Inst.LogLevel = d.Get("log_level").(int)
	ret.Inst.MaxPolltime = d.Get("max_polltime").(int)
	ret.Inst.Method = d.Get("method").(string)
	ret.Inst.Minute = d.Get("minute").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Password = d.Get("password").(int)
	ret.Inst.RenewBefore = d.Get("renew_before").(int)
	ret.Inst.RenewBeforeType = d.Get("renew_before_type").(string)
	ret.Inst.RenewBeforeValue = d.Get("renew_before_value").(int)
	ret.Inst.RenewEvery = d.Get("renew_every").(int)
	ret.Inst.RenewEveryType = d.Get("renew_every_type").(string)
	ret.Inst.RenewEveryValue = d.Get("renew_every_value").(int)
	ret.Inst.SecretString = d.Get("secret_string").(string)
	ret.Inst.StartDate = d.Get("start_date").(string)
	ret.Inst.SubjectAlternateName = getObjectPkiScepCertSubjectAlternateName(d.Get("subject_alternate_name").([]interface{}))
	ret.Inst.Url = d.Get("url").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
