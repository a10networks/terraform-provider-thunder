package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodHttp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_http`: HTTP type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodHttpCreate,
		UpdateContext: resourceHealthMonitorMethodHttpUpdate,
		ReadContext:   resourceHealthMonitorMethodHttpRead,
		DeleteContext: resourceHealthMonitorMethodHttpDelete,

		Schema: map[string]*schema.Schema{
			"http": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "HTTP type",
			},
			"http_expect": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify what you expect from the response message",
			},
			"http_host": {
				Type: schema.TypeString, Optional: true, Description: "Specify \"Host:\" header used in request (enclose IPv6 address in [])",
			},
			"http_kerberos_auth": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Http Kerberos Auth",
			},
			"http_kerberos_kdc": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"http_kerberos_hostip": {
							Type: schema.TypeString, Optional: true, Description: "Kdc's hostname(length:1-31) or IP address",
						},
						"http_kerberos_hostipv6": {
							Type: schema.TypeString, Optional: true, Description: "Server's IPV6 address",
						},
						"http_kerberos_port": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the kdc port",
						},
						"http_kerberos_portv6": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the kdc port",
						},
					},
				},
			},
			"http_kerberos_realm": {
				Type: schema.TypeString, Optional: true, Description: "Specify realm of Kerberos server",
			},
			"http_maintenance_code": {
				Type: schema.TypeString, Optional: true, Description: "Specify response code for maintenance (Format is xx,xx-xx (xx between [100, 899]))",
			},
			"http_password": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the user password",
			},
			"http_password_string": {
				Type: schema.TypeString, Optional: true, Description: "Specify password, '' means empty password",
			},
			"http_port": {
				Type: schema.TypeInt, Optional: true, Default: 80, Description: "Specify HTTP Port (Specify port number (default 80))",
			},
			"http_postdata": {
				Type: schema.TypeString, Optional: true, Description: "Specify the HTTP post data (Input post data here)",
			},
			"http_postfile": {
				Type: schema.TypeString, Optional: true, Description: "Specify the HTTP post data (Input post data file name here)",
			},
			"http_response_code": {
				Type: schema.TypeString, Optional: true, Description: "Specify response code range (e.g. 200,400-430) (Format is xx,xx-xx (xx between [100, 899]))",
			},
			"http_text": {
				Type: schema.TypeString, Optional: true, Description: "Specify text expected",
			},
			"http_url": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify URL string, default is GET /",
			},
			"http_username": {
				Type: schema.TypeString, Optional: true, Description: "Specify the username",
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
			"version2": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify HTTP version2 (Specify http version 2)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceHealthMonitorMethodHttpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodHttpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodHttp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodHttpRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodHttpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodHttpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodHttp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodHttpRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodHttpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodHttpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodHttp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodHttpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodHttpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodHttp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectHealthMonitorMethodHttpHttpKerberosKdc(d []interface{}) edpt.HealthMonitorMethodHttpHttpKerberosKdc {

	count1 := len(d)
	var ret edpt.HealthMonitorMethodHttpHttpKerberosKdc
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HttpKerberosHostip = in["http_kerberos_hostip"].(string)
		ret.HttpKerberosHostipv6 = in["http_kerberos_hostipv6"].(string)
		ret.HttpKerberosPort = in["http_kerberos_port"].(int)
		ret.HttpKerberosPortv6 = in["http_kerberos_portv6"].(int)
	}
	return ret
}

func dataToEndpointHealthMonitorMethodHttp(d *schema.ResourceData) edpt.HealthMonitorMethodHttp {
	var ret edpt.HealthMonitorMethodHttp
	ret.Inst.Http = d.Get("http").(int)
	//omit http_encrypted
	ret.Inst.HttpExpect = d.Get("http_expect").(int)
	ret.Inst.HttpHost = d.Get("http_host").(string)
	ret.Inst.HttpKerberosAuth = d.Get("http_kerberos_auth").(int)
	ret.Inst.HttpKerberosKdc = getObjectHealthMonitorMethodHttpHttpKerberosKdc(d.Get("http_kerberos_kdc").([]interface{}))
	ret.Inst.HttpKerberosRealm = d.Get("http_kerberos_realm").(string)
	ret.Inst.HttpMaintenanceCode = d.Get("http_maintenance_code").(string)
	ret.Inst.HttpPassword = d.Get("http_password").(int)
	ret.Inst.HttpPasswordString = d.Get("http_password_string").(string)
	ret.Inst.HttpPort = d.Get("http_port").(int)
	ret.Inst.HttpPostdata = d.Get("http_postdata").(string)
	ret.Inst.HttpPostfile = d.Get("http_postfile").(string)
	ret.Inst.HttpResponseCode = d.Get("http_response_code").(string)
	ret.Inst.HttpText = d.Get("http_text").(string)
	ret.Inst.HttpUrl = d.Get("http_url").(int)
	ret.Inst.HttpUsername = d.Get("http_username").(string)
	ret.Inst.Maintenance = d.Get("maintenance").(int)
	ret.Inst.MaintenanceText = d.Get("maintenance_text").(string)
	ret.Inst.MaintenanceTextRegex = d.Get("maintenance_text_regex").(string)
	ret.Inst.PostPath = d.Get("post_path").(string)
	ret.Inst.PostType = d.Get("post_type").(string)
	ret.Inst.ResponseCodeRegex = d.Get("response_code_regex").(string)
	ret.Inst.TextRegex = d.Get("text_regex").(string)
	ret.Inst.UrlPath = d.Get("url_path").(string)
	ret.Inst.UrlType = d.Get("url_type").(string)
	//omit uuid
	ret.Inst.Version2 = d.Get("version2").(int)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
