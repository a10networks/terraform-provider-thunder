package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVcsUnicastElectionMembers() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_vcs_unicast_election_members`: members for vcs unicast election\n\n__PLACEHOLDER__",
		CreateContext: resourceVcsUnicastElectionMembersCreate,
		UpdateContext: resourceVcsUnicastElectionMembersUpdate,
		ReadContext:   resourceVcsUnicastElectionMembersRead,
		DeleteContext: resourceVcsUnicastElectionMembersDelete,

		Schema: map[string]*schema.Schema{
			"ip_address_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type: schema.TypeString, Optional: true, Description: "IP Address of the member",
						},
						"use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port for connections in different subnet",
						},
					},
				},
			},
			"ipv6_address_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_address": {
							Type: schema.TypeString, Optional: true, Description: "IPV6 address of the member",
						},
						"use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port for connections in different subnet",
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
func resourceVcsUnicastElectionMembersCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsUnicastElectionMembersCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsUnicastElectionMembers(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsUnicastElectionMembersRead(ctx, d, meta)
	}
	return diags
}

func resourceVcsUnicastElectionMembersUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsUnicastElectionMembersUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsUnicastElectionMembers(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceVcsUnicastElectionMembersRead(ctx, d, meta)
	}
	return diags
}
func resourceVcsUnicastElectionMembersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsUnicastElectionMembersDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsUnicastElectionMembers(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVcsUnicastElectionMembersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVcsUnicastElectionMembersRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVcsUnicastElectionMembers(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceVcsUnicastElectionMembersIpAddressCfg(d []interface{}) []edpt.VcsUnicastElectionMembersIpAddressCfg {

	count1 := len(d)
	ret := make([]edpt.VcsUnicastElectionMembersIpAddressCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsUnicastElectionMembersIpAddressCfg
		oi.IpAddress = in["ip_address"].(string)
		oi.UseMgmtPort = in["use_mgmt_port"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceVcsUnicastElectionMembersIpv6AddressCfg(d []interface{}) []edpt.VcsUnicastElectionMembersIpv6AddressCfg {

	count1 := len(d)
	ret := make([]edpt.VcsUnicastElectionMembersIpv6AddressCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.VcsUnicastElectionMembersIpv6AddressCfg
		oi.Ipv6Address = in["ipv6_address"].(string)
		oi.UseMgmtPort = in["use_mgmt_port"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointVcsUnicastElectionMembers(d *schema.ResourceData) edpt.VcsUnicastElectionMembers {
	var ret edpt.VcsUnicastElectionMembers
	ret.Inst.IpAddressCfg = getSliceVcsUnicastElectionMembersIpAddressCfg(d.Get("ip_address_cfg").([]interface{}))
	ret.Inst.Ipv6AddressCfg = getSliceVcsUnicastElectionMembersIpv6AddressCfg(d.Get("ipv6_address_cfg").([]interface{}))
	//omit uuid
	return ret
}
