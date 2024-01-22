package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodHttps() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_https`: HTTPS type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodHttpsCreate,
		UpdateContext: resourceHealthMonitorMethodHttpsUpdate,
		ReadContext:   resourceHealthMonitorMethodHttpsRead,
		DeleteContext: resourceHealthMonitorMethodHttpsDelete,

		Schema: map[string]*schema.Schema{
			"cert": {
				Type: schema.TypeString, Optional: true, Description: "Specify client certificate (Certificate name)",
			},
			"cert_key_shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Select shared partition",
			},
			"disable_sslv2hello": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable SSLv2Hello for HTTPs",
			},
			"http_version": {
				Type: schema.TypeString, Optional: true, Description: "'http-version2': HTTP version 2 for HTTPs; 'http-version3': HTTP version 3 for HTTPs;",
			},
			"https": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "HTTPS type",
			},
			"https_expect": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify what you expect from the response message",
			},
			"https_host": {
				Type: schema.TypeString, Optional: true, Description: "Specify \"Host:\" header used in request (enclose IPv6 address in [])",
			},
			"https_kerberos_auth": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Https Kerberos Auth",
			},
			"https_kerberos_kdc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"https_kerberos_hostip": {
							Type: schema.TypeString, Optional: true, Description: "Kdc's hostname(length:1-31) or IP address",
						},
						"https_kerberos_hostipv6": {
							Type: schema.TypeString, Optional: true, Description: "Server's IPV6 address",
						},
						"https_kerberos_port": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the kdc port",
						},
						"https_kerberos_portv6": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the kdc port",
						},
					},
				},
			},
			"https_kerberos_realm": {
				Type: schema.TypeString, Optional: true, Description: "Specify realm of Kerberos server",
			},
			"https_maintenance_code": {
				Type: schema.TypeString, Optional: true, Description: "Specify response code for maintenance (Format is xx,xx-xx (xx between [100, 899])",
			},
			"https_password": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the user password",
			},
			"https_password_string": {
				Type: schema.TypeString, Optional: true, Description: "Configure password, '' means empty password",
			},
			"https_postdata": {
				Type: schema.TypeString, Optional: true, Description: "Specify the HTTP post data (Input post data here)",
			},
			"https_postfile": {
				Type: schema.TypeString, Optional: true, Description: "Specify the HTTP post data (Input post data file name here)",
			},
			"https_response_code": {
				Type: schema.TypeString, Optional: true, Description: "Specify response code range (e.g. 200,400-430) (Format is xx,xx-xx (xx between [100, 899])",
			},
			"https_server_cert_name": {
				Type: schema.TypeString, Optional: true, Description: "Expect Server Cert commonName",
			},
			"https_text": {
				Type: schema.TypeString, Optional: true, Description: "Specify text expected",
			},
			"https_url": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify URL string, default is GET /",
			},
			"https_username": {
				Type: schema.TypeString, Optional: true, Description: "Specify the username",
			},
			"key": {
				Type: schema.TypeString, Optional: true, Description: "Specify client private key (Key name)",
			},
			"key_pass_phrase": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Client private key password phrase",
			},
			"key_phrase": {
				Type: schema.TypeString, Optional: true, Description: "Password Phrase",
			},
			"maintenance": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify response text for maintenance",
			},
			"maintenance_text": {
				Type: schema.TypeString, Optional: true, Description: "Specify text for maintenance",
			},
			"maintenance_text_regex": {
				Type: schema.TypeString, Optional: true, Description: "Specify Regex text for maintenance",
			},
			"post_path": {
				Type: schema.TypeString, Optional: true, Description: "Specify URL path, default is \"/\"",
			},
			"post_type": {
				Type: schema.TypeString, Optional: true, Description: "'postdata': Specify the HTTP post data; 'postfile': Specify the HTTP post data;",
			},
			"response_code_regex": {
				Type: schema.TypeString, Optional: true, Description: "Specify response code range with Regex (code with Regex, such as [2-5][0-9][0-9])",
			},
			"sni": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Server Name Indication for HTTPs",
			},
			"text_regex": {
				Type: schema.TypeString, Optional: true, Description: "Specify text expected  with Regex",
			},
			"url_path": {
				Type: schema.TypeString, Optional: true, Description: "Specify URL path, default is \"/\"",
			},
			"url_type": {
				Type: schema.TypeString, Optional: true, Description: "'GET': HTTP GET method; 'POST': HTTP POST method; 'HEAD': HTTP HEAD method;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"web_port": {
				Type: schema.TypeInt, Optional: true, Default: 443, Description: "Specify HTTPS port (Port Number (default 443))",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceHealthMonitorMethodHttpsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodHttpsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodHttps(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodHttpsRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodHttpsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodHttpsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodHttps(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodHttpsRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodHttpsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodHttpsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodHttps(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodHttpsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodHttpsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodHttps(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectHealthMonitorMethodHttpsHttpsKerberosKdc(d []interface{}) edpt.HealthMonitorMethodHttpsHttpsKerberosKdc {

	count1 := len(d)
	var ret edpt.HealthMonitorMethodHttpsHttpsKerberosKdc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HttpsKerberosHostip = in["https_kerberos_hostip"].(string)
		ret.HttpsKerberosHostipv6 = in["https_kerberos_hostipv6"].(string)
		ret.HttpsKerberosPort = in["https_kerberos_port"].(int)
		ret.HttpsKerberosPortv6 = in["https_kerberos_portv6"].(int)
	}
	return ret
}

func dataToEndpointHealthMonitorMethodHttps(d *schema.ResourceData) edpt.HealthMonitorMethodHttps {
	var ret edpt.HealthMonitorMethodHttps
	ret.Inst.Cert = d.Get("cert").(string)
	ret.Inst.CertKeyShared = d.Get("cert_key_shared").(int)
	ret.Inst.DisableSslv2hello = d.Get("disable_sslv2hello").(int)
	ret.Inst.HttpVersion = d.Get("http_version").(string)
	ret.Inst.Https = d.Get("https").(int)
	//omit https_encrypted
	ret.Inst.HttpsExpect = d.Get("https_expect").(int)
	ret.Inst.HttpsHost = d.Get("https_host").(string)
	ret.Inst.HttpsKerberosAuth = d.Get("https_kerberos_auth").(int)
	ret.Inst.HttpsKerberosKdc = getObjectHealthMonitorMethodHttpsHttpsKerberosKdc(d.Get("https_kerberos_kdc").([]interface{}))
	ret.Inst.HttpsKerberosRealm = d.Get("https_kerberos_realm").(string)
	//omit https_key_encrypted
	ret.Inst.HttpsMaintenanceCode = d.Get("https_maintenance_code").(string)
	ret.Inst.HttpsPassword = d.Get("https_password").(int)
	ret.Inst.HttpsPasswordString = d.Get("https_password_string").(string)
	ret.Inst.HttpsPostdata = d.Get("https_postdata").(string)
	ret.Inst.HttpsPostfile = d.Get("https_postfile").(string)
	ret.Inst.HttpsResponseCode = d.Get("https_response_code").(string)
	ret.Inst.HttpsServerCertName = d.Get("https_server_cert_name").(string)
	ret.Inst.HttpsText = d.Get("https_text").(string)
	ret.Inst.HttpsUrl = d.Get("https_url").(int)
	ret.Inst.HttpsUsername = d.Get("https_username").(string)
	ret.Inst.Key = d.Get("key").(string)
	ret.Inst.KeyPassPhrase = d.Get("key_pass_phrase").(int)
	ret.Inst.KeyPhrase = d.Get("key_phrase").(string)
	ret.Inst.Maintenance = d.Get("maintenance").(int)
	ret.Inst.MaintenanceText = d.Get("maintenance_text").(string)
	ret.Inst.MaintenanceTextRegex = d.Get("maintenance_text_regex").(string)
	ret.Inst.PostPath = d.Get("post_path").(string)
	ret.Inst.PostType = d.Get("post_type").(string)
	ret.Inst.ResponseCodeRegex = d.Get("response_code_regex").(string)
	ret.Inst.Sni = d.Get("sni").(int)
	ret.Inst.TextRegex = d.Get("text_regex").(string)
	ret.Inst.UrlPath = d.Get("url_path").(string)
	ret.Inst.UrlType = d.Get("url_type").(string)
	//omit uuid
	ret.Inst.WebPort = d.Get("web_port").(int)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
