package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLdapServerHostLdapHostname() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ldap_server_host_ldap_hostname`: Specify the hostname of LDAP server\n\n__PLACEHOLDER__",
		CreateContext: resourceLdapServerHostLdapHostnameCreate,
		UpdateContext: resourceLdapServerHostLdapHostnameUpdate,
		ReadContext:   resourceLdapServerHostLdapHostnameRead,
		DeleteContext: resourceLdapServerHostLdapHostnameDelete,

		Schema: map[string]*schema.Schema{
			"base": {
				Type: schema.TypeString, Optional: true, Description: "Configure the group DN which user belongs to",
			},
			"cn_value": {
				Type: schema.TypeString, Optional: true, Description: "LDAP common name identifier (i.e.: cn, uid)",
			},
			"dn_value": {
				Type: schema.TypeString, Optional: true, Description: "LDAP distinguished name (dn)",
			},
			"domain": {
				Type: schema.TypeString, Optional: true, Description: "Configure AD domain name",
			},
			"domain_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the LDAP server's port (default 389 without ssl or 636 with ssl)",
						},
						"ssl": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use SSL",
						},
						"timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the LDAP server's timeout (default 3)",
						},
					},
				},
			},
			"group": {
				Type: schema.TypeString, Optional: true, Description: "Configure the group DN which user belongs to",
			},
			"hostname": {
				Type: schema.TypeString, Required: true, Description: "Hostname of LDAP server",
			},
			"hostname_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the LDAP server's port (default 3268 without ssl or 3269 with ssl)",
						},
						"ssl": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use SSL",
						},
						"timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the LDAP server's timeout (default 3)",
						},
					},
				},
			},
			"port_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the LDAP server's port (default 389 without ssl or 636 with ssl)",
						},
						"ssl": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use SSL",
						},
						"timeout": {
							Type: schema.TypeInt, Optional: true, Description: "Specify the LDAP server's timeout (default 3)",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceLdapServerHostLdapHostnameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLdapServerHostLdapHostnameCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLdapServerHostLdapHostname(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLdapServerHostLdapHostnameRead(ctx, d, meta)
	}
	return diags
}

func resourceLdapServerHostLdapHostnameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLdapServerHostLdapHostnameUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLdapServerHostLdapHostname(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLdapServerHostLdapHostnameRead(ctx, d, meta)
	}
	return diags
}
func resourceLdapServerHostLdapHostnameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLdapServerHostLdapHostnameDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLdapServerHostLdapHostname(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLdapServerHostLdapHostnameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLdapServerHostLdapHostnameRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLdapServerHostLdapHostname(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectLdapServerHostLdapHostnameDomainCfg(d []interface{}) edpt.LdapServerHostLdapHostnameDomainCfg {

	count1 := len(d)
	var ret edpt.LdapServerHostLdapHostnameDomainCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Port = in["port"].(int)
		ret.Ssl = in["ssl"].(int)
		ret.Timeout = in["timeout"].(int)
	}
	return ret
}

func getObjectLdapServerHostLdapHostnameHostnameCfg(d []interface{}) edpt.LdapServerHostLdapHostnameHostnameCfg {

	count1 := len(d)
	var ret edpt.LdapServerHostLdapHostnameHostnameCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Port = in["port"].(int)
		ret.Ssl = in["ssl"].(int)
		ret.Timeout = in["timeout"].(int)
	}
	return ret
}

func getObjectLdapServerHostLdapHostnamePortCfg(d []interface{}) edpt.LdapServerHostLdapHostnamePortCfg {

	count1 := len(d)
	var ret edpt.LdapServerHostLdapHostnamePortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Port = in["port"].(int)
		ret.Ssl = in["ssl"].(int)
		ret.Timeout = in["timeout"].(int)
	}
	return ret
}

func dataToEndpointLdapServerHostLdapHostname(d *schema.ResourceData) edpt.LdapServerHostLdapHostname {
	var ret edpt.LdapServerHostLdapHostname
	ret.Inst.Base = d.Get("base").(string)
	ret.Inst.CnValue = d.Get("cn_value").(string)
	ret.Inst.DnValue = d.Get("dn_value").(string)
	ret.Inst.Domain = d.Get("domain").(string)
	ret.Inst.DomainCfg = getObjectLdapServerHostLdapHostnameDomainCfg(d.Get("domain_cfg").([]interface{}))
	ret.Inst.Group = d.Get("group").(string)
	ret.Inst.Hostname = d.Get("hostname").(string)
	ret.Inst.HostnameCfg = getObjectLdapServerHostLdapHostnameHostnameCfg(d.Get("hostname_cfg").([]interface{}))
	ret.Inst.PortCfg = getObjectLdapServerHostLdapHostnamePortCfg(d.Get("port_cfg").([]interface{}))
	//omit uuid
	return ret
}
