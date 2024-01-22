package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLdapServerHostIpv4() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ldap_server_host_ipv4`: Specify the hostname of ldap server\n\n__PLACEHOLDER__",
		CreateContext: resourceLdapServerHostIpv4Create,
		UpdateContext: resourceLdapServerHostIpv4Update,
		ReadContext:   resourceLdapServerHostIpv4Read,
		DeleteContext: resourceLdapServerHostIpv4Delete,

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
			"ipv4_addr": {
				Type: schema.TypeString, Required: true, Description: "IPV4 address of ldap server",
			},
			"ipv4_addr_cfg": {
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
func resourceLdapServerHostIpv4Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLdapServerHostIpv4Create()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLdapServerHostIpv4(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLdapServerHostIpv4Read(ctx, d, meta)
	}
	return diags
}

func resourceLdapServerHostIpv4Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLdapServerHostIpv4Update()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLdapServerHostIpv4(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceLdapServerHostIpv4Read(ctx, d, meta)
	}
	return diags
}
func resourceLdapServerHostIpv4Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLdapServerHostIpv4Delete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLdapServerHostIpv4(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceLdapServerHostIpv4Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceLdapServerHostIpv4Read()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointLdapServerHostIpv4(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectLdapServerHostIpv4DomainCfg(d []interface{}) edpt.LdapServerHostIpv4DomainCfg {

	count1 := len(d)
	var ret edpt.LdapServerHostIpv4DomainCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Port = in["port"].(int)
		ret.Ssl = in["ssl"].(int)
		ret.Timeout = in["timeout"].(int)
	}
	return ret
}

func getObjectLdapServerHostIpv4Ipv4AddrCfg(d []interface{}) edpt.LdapServerHostIpv4Ipv4AddrCfg {

	count1 := len(d)
	var ret edpt.LdapServerHostIpv4Ipv4AddrCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Port = in["port"].(int)
		ret.Ssl = in["ssl"].(int)
		ret.Timeout = in["timeout"].(int)
	}
	return ret
}

func getObjectLdapServerHostIpv4PortCfg(d []interface{}) edpt.LdapServerHostIpv4PortCfg {

	count1 := len(d)
	var ret edpt.LdapServerHostIpv4PortCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Port = in["port"].(int)
		ret.Ssl = in["ssl"].(int)
		ret.Timeout = in["timeout"].(int)
	}
	return ret
}

func dataToEndpointLdapServerHostIpv4(d *schema.ResourceData) edpt.LdapServerHostIpv4 {
	var ret edpt.LdapServerHostIpv4
	ret.Inst.Base = d.Get("base").(string)
	ret.Inst.CnValue = d.Get("cn_value").(string)
	ret.Inst.DnValue = d.Get("dn_value").(string)
	ret.Inst.Domain = d.Get("domain").(string)
	ret.Inst.DomainCfg = getObjectLdapServerHostIpv4DomainCfg(d.Get("domain_cfg").([]interface{}))
	ret.Inst.Group = d.Get("group").(string)
	ret.Inst.Ipv4Addr = d.Get("ipv4_addr").(string)
	ret.Inst.Ipv4AddrCfg = getObjectLdapServerHostIpv4Ipv4AddrCfg(d.Get("ipv4_addr_cfg").([]interface{}))
	ret.Inst.PortCfg = getObjectLdapServerHostIpv4PortCfg(d.Get("port_cfg").([]interface{}))
	//omit uuid
	return ret
}
