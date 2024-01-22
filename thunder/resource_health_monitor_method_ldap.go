package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHealthMonitorMethodLdap() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_health_monitor_method_ldap`: LDAP type\n\n__PLACEHOLDER__",
		CreateContext: resourceHealthMonitorMethodLdapCreate,
		UpdateContext: resourceHealthMonitorMethodLdapUpdate,
		ReadContext:   resourceHealthMonitorMethodLdapRead,
		DeleteContext: resourceHealthMonitorMethodLdapDelete,

		Schema: map[string]*schema.Schema{
			"acceptnotfound": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mark server up on receiving a not-found response",
			},
			"acceptresref": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mark server up on receiving a search result reference response",
			},
			"basedn": {
				Type: schema.TypeString, Optional: true, Description: "Specify LDAP DN distinguished name",
			},
			"ldap": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "LDAP type",
			},
			"ldap_binddn": {
				Type: schema.TypeString, Optional: true, Description: "Specify the distinguished name for bindRequest (LDAP DN distinguished name)",
			},
			"ldap_password": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the user password",
			},
			"ldap_password_string": {
				Type: schema.TypeString, Optional: true, Description: "Configure password, '' means empty password",
			},
			"ldap_port": {
				Type: schema.TypeInt, Optional: true, Default: 389, Description: "Specify the LDAP port (Speciry port number, default is 389, or 636 if LDAP over SSL)",
			},
			"ldap_query": {
				Type: schema.TypeString, Optional: true, Description: "LDAP query to be excuted",
			},
			"ldap_run_search": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify a query to be executed",
			},
			"ldap_security": {
				Type: schema.TypeString, Optional: true, Description: "'overssl': Set LDAP over SSL; 'StartTLS': LDAP switch to TLS;",
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
func resourceHealthMonitorMethodLdapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodLdapCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodLdap(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodLdapRead(ctx, d, meta)
	}
	return diags
}

func resourceHealthMonitorMethodLdapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodLdapUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodLdap(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHealthMonitorMethodLdapRead(ctx, d, meta)
	}
	return diags
}
func resourceHealthMonitorMethodLdapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodLdapDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodLdap(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHealthMonitorMethodLdapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHealthMonitorMethodLdapRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHealthMonitorMethodLdap(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointHealthMonitorMethodLdap(d *schema.ResourceData) edpt.HealthMonitorMethodLdap {
	var ret edpt.HealthMonitorMethodLdap
	ret.Inst.Acceptnotfound = d.Get("acceptnotfound").(int)
	ret.Inst.Acceptresref = d.Get("acceptresref").(int)
	ret.Inst.Basedn = d.Get("basedn").(string)
	ret.Inst.Ldap = d.Get("ldap").(int)
	ret.Inst.LdapBinddn = d.Get("ldap_binddn").(string)
	//omit ldap_encrypted
	ret.Inst.LdapPassword = d.Get("ldap_password").(int)
	ret.Inst.LdapPasswordString = d.Get("ldap_password_string").(string)
	ret.Inst.LdapPort = d.Get("ldap_port").(int)
	ret.Inst.LdapQuery = d.Get("ldap_query").(string)
	ret.Inst.LdapRunSearch = d.Get("ldap_run_search").(int)
	ret.Inst.LdapSecurity = d.Get("ldap_security").(string)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
