package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_group`: GSLB Group\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbGroupCreate,
		UpdateContext: resourceGslbGroupUpdate,
		ReadContext:   resourceGslbGroupRead,
		DeleteContext: resourceGslbGroupDelete,

		Schema: map[string]*schema.Schema{
			"auto_map_learn": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "IP Address learned from other controller",
			},
			"auto_map_primary": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Primary Controller's IP address",
			},
			"auto_map_smart": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Choose Best IP address",
			},
			"config_anywhere": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Every member can do config",
			},
			"config_merge": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Merge old master's config when new one take over",
			},
			"config_save": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Accept config-save message from master",
			},
			"data_interface": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Data Interface IP Address",
			},
			"dns_discover": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Discover member via DNS Protocol",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Join GSLB Group",
			},
			"learn": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Learn neighbour information from other controllers",
			},
			"mgmt_interface": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Management Interface IP Address",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Specify Group domain name",
			},
			"primary_ipv6_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"primary_ipv6": {
							Type: schema.TypeString, Optional: true, Description: "Specify Primary controller's IP address",
						},
					},
				},
			},
			"primary_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"primary": {
							Type: schema.TypeString, Optional: true, Description: "Specify Primary controller's IP address",
						},
					},
				},
			},
			"priority": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "Specify Local Priority, default is 100",
			},
			"resolve_as": {
				Type: schema.TypeString, Optional: true, Default: "resolve-to-ipv4", Description: "'resolve-to-ipv4': Use A Query only to resolve FQDN (Default Query type); 'resolve-to-ipv6': Use AAAA Query only to resolve FQDN; 'resolve-to-ipv4-and-ipv6': Use A as well as AAAA Query to resolve FQDN;",
			},
			"standalone": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Run GSLB Group in standalone mode",
			},
			"suffix": {
				Type: schema.TypeString, Optional: true, Description: "Set DNS Suffix (Name)",
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
func resourceGslbGroupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbGroupCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbGroup(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbGroupRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbGroupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbGroupUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbGroup(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbGroupRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbGroupDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbGroup(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbGroupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbGroupRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbGroup(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceGslbGroupPrimaryIpv6List(d []interface{}) []edpt.GslbGroupPrimaryIpv6List {

	count1 := len(d)
	ret := make([]edpt.GslbGroupPrimaryIpv6List, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbGroupPrimaryIpv6List
		oi.PrimaryIpv6 = in["primary_ipv6"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceGslbGroupPrimaryList(d []interface{}) []edpt.GslbGroupPrimaryList {

	count1 := len(d)
	ret := make([]edpt.GslbGroupPrimaryList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GslbGroupPrimaryList
		oi.Primary = in["primary"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointGslbGroup(d *schema.ResourceData) edpt.GslbGroup {
	var ret edpt.GslbGroup
	ret.Inst.AutoMapLearn = d.Get("auto_map_learn").(int)
	ret.Inst.AutoMapPrimary = d.Get("auto_map_primary").(int)
	ret.Inst.AutoMapSmart = d.Get("auto_map_smart").(int)
	ret.Inst.ConfigAnywhere = d.Get("config_anywhere").(int)
	ret.Inst.ConfigMerge = d.Get("config_merge").(int)
	ret.Inst.ConfigSave = d.Get("config_save").(int)
	ret.Inst.DataInterface = d.Get("data_interface").(int)
	ret.Inst.DnsDiscover = d.Get("dns_discover").(int)
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.Learn = d.Get("learn").(int)
	ret.Inst.MgmtInterface = d.Get("mgmt_interface").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.PrimaryIpv6List = getSliceGslbGroupPrimaryIpv6List(d.Get("primary_ipv6_list").([]interface{}))
	ret.Inst.PrimaryList = getSliceGslbGroupPrimaryList(d.Get("primary_list").([]interface{}))
	ret.Inst.Priority = d.Get("priority").(int)
	ret.Inst.ResolveAs = d.Get("resolve_as").(string)
	ret.Inst.Standalone = d.Get("standalone").(int)
	ret.Inst.Suffix = d.Get("suffix").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
