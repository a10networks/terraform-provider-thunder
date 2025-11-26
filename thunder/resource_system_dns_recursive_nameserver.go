package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemDnsRecursiveNameserver() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_dns_recursive_nameserver`: Configure default DNS authoritative name servers for recursive DNS resolver\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemDnsRecursiveNameserverCreate,
		UpdateContext: resourceSystemDnsRecursiveNameserverUpdate,
		ReadContext:   resourceSystemDnsRecursiveNameserverRead,
		DeleteContext: resourceSystemDnsRecursiveNameserverDelete,

		Schema: map[string]*schema.Schema{
			"follow_shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use the configured name servers of shared partition",
			},
			"server_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_addr": {
							Type: schema.TypeString, Optional: true, Description: "Specify IPv4 server address",
						},
						"v4_desc": {
							Type: schema.TypeString, Optional: true, Description: "Description for this ipv4 address",
						},
						"ipv6_addr": {
							Type: schema.TypeString, Optional: true, Description: "Specify IPv6 server address",
						},
						"v6_desc": {
							Type: schema.TypeString, Optional: true, Description: "Description for this ipv6 address",
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
func resourceSystemDnsRecursiveNameserverCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDnsRecursiveNameserverCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDnsRecursiveNameserver(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemDnsRecursiveNameserverRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemDnsRecursiveNameserverUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDnsRecursiveNameserverUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDnsRecursiveNameserver(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemDnsRecursiveNameserverRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemDnsRecursiveNameserverDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDnsRecursiveNameserverDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDnsRecursiveNameserver(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemDnsRecursiveNameserverRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemDnsRecursiveNameserverRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemDnsRecursiveNameserver(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceSystemDnsRecursiveNameserverServerList(d []interface{}) []edpt.SystemDnsRecursiveNameserverServerList {

	count1 := len(d)
	ret := make([]edpt.SystemDnsRecursiveNameserverServerList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemDnsRecursiveNameserverServerList
		oi.Ipv4Addr = in["ipv4_addr"].(string)
		oi.V4Desc = in["v4_desc"].(string)
		oi.Ipv6Addr = in["ipv6_addr"].(string)
		oi.V6Desc = in["v6_desc"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemDnsRecursiveNameserver(d *schema.ResourceData) edpt.SystemDnsRecursiveNameserver {
	var ret edpt.SystemDnsRecursiveNameserver
	ret.Inst.FollowShared = d.Get("follow_shared").(int)
	ret.Inst.ServerList = getSliceSystemDnsRecursiveNameserverServerList(d.Get("server_list").([]interface{}))
	//omit uuid
	return ret
}
