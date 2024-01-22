package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_list`: Configure ip list\n\n__PLACEHOLDER__",
		CreateContext: resourceIpListCreate,
		UpdateContext: resourceIpListUpdate,
		ReadContext:   resourceIpListRead,
		DeleteContext: resourceIpListDelete,

		Schema: map[string]*schema.Schema{
			"ipv4_config": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_start_addr": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 Range Start Address / IPv4 Address",
						},
						"ipv4_end_addr": {
							Type: schema.TypeString, Optional: true, Description: "IPv4 Range End Address",
						},
					},
				},
			},
			"ipv6_config": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_start_addr": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 Range Start Address / IPv6 Address",
						},
						"ipv6_end_addr": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 Range End Address",
						},
					},
				},
			},
			"ipv6_prefix_config": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_addr_prefix": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 Prefix Range Start / IPv6 Prefix",
						},
						"ipv6_prefix_to": {
							Type: schema.TypeString, Optional: true, Description: "IPv6 Prefix Range End",
						},
						"count1": {
							Type: schema.TypeInt, Optional: true, Description: "Number of IPv6 prefixes",
						},
					},
				},
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify name of the ip list",
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
func resourceIpListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpListRead(ctx, d, meta)
	}
	return diags
}

func resourceIpListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpListRead(ctx, d, meta)
	}
	return diags
}
func resourceIpListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceIpListIpv4Config(d []interface{}) []edpt.IpListIpv4Config {

	count1 := len(d)
	ret := make([]edpt.IpListIpv4Config, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpListIpv4Config
		oi.Ipv4StartAddr = in["ipv4_start_addr"].(string)
		oi.Ipv4EndAddr = in["ipv4_end_addr"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpListIpv6Config(d []interface{}) []edpt.IpListIpv6Config {

	count1 := len(d)
	ret := make([]edpt.IpListIpv6Config, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpListIpv6Config
		oi.Ipv6StartAddr = in["ipv6_start_addr"].(string)
		oi.Ipv6EndAddr = in["ipv6_end_addr"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceIpListIpv6PrefixConfig(d []interface{}) []edpt.IpListIpv6PrefixConfig {

	count1 := len(d)
	ret := make([]edpt.IpListIpv6PrefixConfig, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.IpListIpv6PrefixConfig
		oi.Ipv6AddrPrefix = in["ipv6_addr_prefix"].(string)
		oi.Ipv6PrefixTo = in["ipv6_prefix_to"].(string)
		oi.Count1 = in["count1"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointIpList(d *schema.ResourceData) edpt.IpList {
	var ret edpt.IpList
	ret.Inst.Ipv4Config = getSliceIpListIpv4Config(d.Get("ipv4_config").([]interface{}))
	ret.Inst.Ipv6Config = getSliceIpListIpv6Config(d.Get("ipv6_config").([]interface{}))
	ret.Inst.Ipv6PrefixConfig = getSliceIpListIpv6PrefixConfig(d.Get("ipv6_prefix_config").([]interface{}))
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
