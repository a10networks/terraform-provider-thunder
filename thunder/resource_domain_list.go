package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDomainList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_domain_list`: Configure Domain classification list\n\n__PLACEHOLDER__",
		CreateContext: resourceDomainListCreate,
		UpdateContext: resourceDomainListUpdate,
		ReadContext:   resourceDomainListRead,
		DeleteContext: resourceDomainListDelete,

		Schema: map[string]*schema.Schema{
			"domain_name_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_name": {
							Type: schema.TypeString, Optional: true, Description: "Domain name to be added to this domain list",
						},
						"interval": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "DNS query interval (in minute, default is 10)",
						},
					},
				},
			},
			"file": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Create/Edit a domain-list stored as a file",
			},
			"match_type_axfr": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"axfr_domain": {
							Type: schema.TypeString, Optional: true, Description: "Import the list of domains via zone-transfer",
						},
						"axfr_ip_address": {
							Type: schema.TypeString, Optional: true, Description: "IP address of the listening DNS server",
						},
						"axfr_ipv6_address": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 address of the listening DNS server",
						},
						"ip_axfr_port_num": {
							Type: schema.TypeInt, Optional: true, Default: 53, Description: "Port Number",
						},
						"ipv6_axfr_port_num": {
							Type: schema.TypeInt, Optional: true, Default: 53, Description: "Port Number",
						},
						"ip_refresh_intvl": {
							Type: schema.TypeInt, Optional: true, Default: 30, Description: "Poll every x minutes to check for an Updated axfr default",
						},
						"ipv6_refresh_intvl": {
							Type: schema.TypeInt, Optional: true, Default: 30, Description: "Poll every x minutes to check for an Updated axfr default",
						},
					},
				},
			},
			"match_type_equals": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"equals": {
							Type: schema.TypeString, Optional: true, Description: "Specify exact match for the Domain Name",
						},
					},
				},
			},
			"match_type_suffix": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"suffix": {
							Type: schema.TypeString, Optional: true, Description: "Specify suffix matching the Domain Name",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the domain list",
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
func resourceDomainListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDomainListRead(ctx, d, meta)
	}
	return diags
}

func resourceDomainListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDomainListRead(ctx, d, meta)
	}
	return diags
}
func resourceDomainListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDomainListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDomainListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDomainList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDomainListDomainNameList(d []interface{}) []edpt.DomainListDomainNameList {

	count1 := len(d)
	ret := make([]edpt.DomainListDomainNameList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DomainListDomainNameList
		oi.DomainName = in["domain_name"].(string)
		oi.Interval = in["interval"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDomainListMatchTypeAxfr(d []interface{}) []edpt.DomainListMatchTypeAxfr {

	count1 := len(d)
	ret := make([]edpt.DomainListMatchTypeAxfr, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DomainListMatchTypeAxfr
		oi.AxfrDomain = in["axfr_domain"].(string)
		oi.AxfrIpAddress = in["axfr_ip_address"].(string)
		oi.AxfrIpv6Address = in["axfr_ipv6_address"].(string)
		oi.IpAxfrPortNum = in["ip_axfr_port_num"].(int)
		oi.Ipv6AxfrPortNum = in["ipv6_axfr_port_num"].(int)
		oi.IpRefreshIntvl = in["ip_refresh_intvl"].(int)
		oi.Ipv6RefreshIntvl = in["ipv6_refresh_intvl"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDomainListMatchTypeEquals(d []interface{}) []edpt.DomainListMatchTypeEquals {

	count1 := len(d)
	ret := make([]edpt.DomainListMatchTypeEquals, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DomainListMatchTypeEquals
		oi.Equals = in["equals"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceDomainListMatchTypeSuffix(d []interface{}) []edpt.DomainListMatchTypeSuffix {

	count1 := len(d)
	ret := make([]edpt.DomainListMatchTypeSuffix, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DomainListMatchTypeSuffix
		oi.Suffix = in["suffix"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointDomainList(d *schema.ResourceData) edpt.DomainList {
	var ret edpt.DomainList
	ret.Inst.DomainNameList = getSliceDomainListDomainNameList(d.Get("domain_name_list").([]interface{}))
	ret.Inst.File = d.Get("file").(int)
	ret.Inst.MatchTypeAxfr = getSliceDomainListMatchTypeAxfr(d.Get("match_type_axfr").([]interface{}))
	ret.Inst.MatchTypeEquals = getSliceDomainListMatchTypeEquals(d.Get("match_type_equals").([]interface{}))
	ret.Inst.MatchTypeSuffix = getSliceDomainListMatchTypeSuffix(d.Get("match_type_suffix").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
