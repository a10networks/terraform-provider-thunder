package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodKerberosKdc() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_kerberos_kdc`: Kerberos KDC type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodKerberosKdcCreate,
		UpdateContext: resourceHealthMonitorMethodKerberosKdcUpdate,
		ReadContext:   resourceHealthMonitorMethodKerberosKdcRead,
		DeleteContext: resourceHealthMonitorMethodKerberosKdcDelete,

		Schema: map[string]*schema.Schema{
			"kerberos_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kinit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Kerberos KDC",
						},
						"kinit_pricipal_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify the principal name",
						},
						"kinit_password": {
							Type: schema.TypeString, Optional: true, Description: "Password",
						},
						"kinit_kdc": {
							Type: schema.TypeString, Optional: true, Description: "Specify the kdc server, host|ip [:port]",
						},
						"tcp_only": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the kerberos tcp only",
						},
						"kadmin": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Kerberos admin",
						},
						"kadmin_realm": {
							Type: schema.TypeString, Optional: true, Description: "Specify the realm",
						},
						"kadmin_pricipal_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify the principal name",
						},
						"kadmin_password": {
							Type: schema.TypeString, Optional: true, Description: "Password",
						},
						"kadmin_server": {
							Type: schema.TypeString, Optional: true, Description: "Specify the admin server, host|ip [:port]",
						},
						"kadmin_kdc": {
							Type: schema.TypeString, Optional: true, Description: "Specify the kdc server, host|ip [:port]",
						},
						"kpasswd": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Kerberos change passwd",
						},
						"kpasswd_pricipal_name": {
							Type: schema.TypeString, Optional: true, Description: "Specify the principal name",
						},
						"kpasswd_password": {
							Type: schema.TypeString, Optional: true, Description: "Password",
						},
						"kpasswd_server": {
							Type: schema.TypeString, Optional: true, Description: "Specify the Kerberos password server, host|ip [:port]",
						},
						"kpasswd_kdc": {
							Type: schema.TypeString, Optional: true, Description: "Specify the kdc server, host|ip [:port]",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceHealthMonitorMethodKerberosKdcCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodKerberosKdcCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodKerberosKdc(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodKerberosKdcRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodKerberosKdcUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodKerberosKdcUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodKerberosKdc(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodKerberosKdcRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodKerberosKdcDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodKerberosKdcDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodKerberosKdc(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodKerberosKdcRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodKerberosKdcRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodKerberosKdc(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectHealthMonitorMethodKerberosKdcKerberosCfg(d []interface{}) edpt.HealthMonitorMethodKerberosKdcKerberosCfg {

	count1 := len(d)
	var ret edpt.HealthMonitorMethodKerberosKdcKerberosCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Kinit = in["kinit"].(int)
		ret.KinitPricipalName = in["kinit_pricipal_name"].(string)
		ret.KinitPassword = in["kinit_password"].(string)
		//omit kinit_encrypted
		ret.KinitKdc = in["kinit_kdc"].(string)
		ret.TcpOnly = in["tcp_only"].(int)
		ret.Kadmin = in["kadmin"].(int)
		ret.KadminRealm = in["kadmin_realm"].(string)
		ret.KadminPricipalName = in["kadmin_pricipal_name"].(string)
		ret.KadminPassword = in["kadmin_password"].(string)
		//omit kadmin_encrypted
		ret.KadminServer = in["kadmin_server"].(string)
		ret.KadminKdc = in["kadmin_kdc"].(string)
		ret.Kpasswd = in["kpasswd"].(int)
		ret.KpasswdPricipalName = in["kpasswd_pricipal_name"].(string)
		ret.KpasswdPassword = in["kpasswd_password"].(string)
		//omit kpasswd_encrypted
		ret.KpasswdServer = in["kpasswd_server"].(string)
		ret.KpasswdKdc = in["kpasswd_kdc"].(string)
	}
	return ret
}

func dataToEndpointHealthMonitorMethodKerberosKdc(d *schema.ResourceData) edpt.HealthMonitorMethodKerberosKdc {
	var ret edpt.HealthMonitorMethodKerberosKdc
	ret.Inst.KerberosCfg = getObjectHealthMonitorMethodKerberosKdcKerberosCfg(d.Get("kerberos_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
