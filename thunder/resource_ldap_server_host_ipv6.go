package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLdapServerHostIpv6() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ldap_server_host_ipv6`: Specify the hostname of ldap server\n\n__PLACEHOLDER__",
		CreateContext: resourceLdapServerHostIpv6Create,
		UpdateContext: resourceLdapServerHostIpv6Update,
		ReadContext:   resourceLdapServerHostIpv6Read,
		DeleteContext: resourceLdapServerHostIpv6Delete,

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
			"ipv6_addr": {
				Type: schema.TypeString, Required: true, Description: "IPV6 address of ldap server",
			},
			"ipv6_addr_cfg": {
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
func resourceLdapServerHostIpv6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLdapServerHostIpv6Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLdapServerHostIpv6(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLdapServerHostIpv6Read(ctx, d, meta)
	}
	return diags
}

func resourceLdapServerHostIpv6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLdapServerHostIpv6Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLdapServerHostIpv6(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLdapServerHostIpv6Read(ctx, d, meta)
	}
	return diags
}
func resourceLdapServerHostIpv6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLdapServerHostIpv6Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLdapServerHostIpv6(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLdapServerHostIpv6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLdapServerHostIpv6Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLdapServerHostIpv6(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectLdapServerHostIpv6DomainCfg(d []interface{}) edpt.LdapServerHostIpv6DomainCfg {

	count1 := len(d)
	var ret edpt.LdapServerHostIpv6DomainCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Port = in["port"].(int)
		ret.Ssl = in["ssl"].(int)
		ret.Timeout = in["timeout"].(int)
	}
	return ret
}

func getObjectLdapServerHostIpv6Ipv6AddrCfg(d []interface{}) edpt.LdapServerHostIpv6Ipv6AddrCfg {

	count1 := len(d)
	var ret edpt.LdapServerHostIpv6Ipv6AddrCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Port = in["port"].(int)
		ret.Ssl = in["ssl"].(int)
		ret.Timeout = in["timeout"].(int)
	}
	return ret
}

func getObjectLdapServerHostIpv6PortCfg(d []interface{}) edpt.LdapServerHostIpv6PortCfg {

	count1 := len(d)
	var ret edpt.LdapServerHostIpv6PortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Port = in["port"].(int)
		ret.Ssl = in["ssl"].(int)
		ret.Timeout = in["timeout"].(int)
	}
	return ret
}

func dataToEndpointLdapServerHostIpv6(d *schema.ResourceData) edpt.LdapServerHostIpv6 {
	var ret edpt.LdapServerHostIpv6
	ret.Inst.Base = d.Get("base").(string)
	ret.Inst.CnValue = d.Get("cn_value").(string)
	ret.Inst.DnValue = d.Get("dn_value").(string)
	ret.Inst.Domain = d.Get("domain").(string)
	ret.Inst.DomainCfg = getObjectLdapServerHostIpv6DomainCfg(d.Get("domain_cfg").([]interface{}))
	ret.Inst.Group = d.Get("group").(string)
	ret.Inst.Ipv6Addr = d.Get("ipv6_addr").(string)
	ret.Inst.Ipv6AddrCfg = getObjectLdapServerHostIpv6Ipv6AddrCfg(d.Get("ipv6_addr_cfg").([]interface{}))
	ret.Inst.PortCfg = getObjectLdapServerHostIpv6PortCfg(d.Get("port_cfg").([]interface{}))
	//omit uuid
	return ret
}
