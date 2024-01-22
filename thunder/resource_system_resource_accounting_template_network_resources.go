package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemResourceAccountingTemplateNetworkResources() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_resource_accounting_template_network_resources`: Enter the network resource limits\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemResourceAccountingTemplateNetworkResourcesCreate,
		UpdateContext: resourceSystemResourceAccountingTemplateNetworkResourcesUpdate,
		ReadContext:   resourceSystemResourceAccountingTemplateNetworkResourcesRead,
		DeleteContext: resourceSystemResourceAccountingTemplateNetworkResourcesDelete,

		Schema: map[string]*schema.Schema{
			"ipv4_acl_line_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_acl_line_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of ACL lines allowed (IPV4 ACL lines (default is max-value))",
						},
						"ipv4_acl_line_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"ipv6_acl_line_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv6_acl_line_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of ACL lines allowed (IPV6 ACL lines (default is max-value))",
						},
						"ipv6_acl_line_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"object_group_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"object_group_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of object groups allowed (Object group (default is max-value))",
						},
						"object_group_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"object_group_clause_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"object_group_clause_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of object group clauses allowed (Object group clauses (default is max-value))",
						},
						"object_group_clause_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"static_arp_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static_arp_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of static arp entries allowed (Static arp (default is max-value))",
						},
						"static_arp_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"static_ipv4_route_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static_ipv4_route_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of static ipv4 routes allowed (Static ipv4 routes (default is max-value))",
						},
						"static_ipv4_route_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"static_ipv6_route_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static_ipv6_route_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of static ipv6 routes allowed (Static ipv6 routes (default is max-value))",
						},
						"static_ipv6_route_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"static_mac_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static_mac_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of static MAC entries allowed (Static MACs (default is max-value))",
						},
						"static_mac_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"static_neighbor_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"static_neighbor_max": {
							Type: schema.TypeInt, Optional: true, Description: "Enter the number of static neighbor entries allowed (Static neighbors (default is max-value))",
						},
						"static_neighbor_min_guarantee": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Minimum guaranteed value ( Minimum guaranteed value)",
						},
					},
				},
			},
			"threshold": {
				Type: schema.TypeInt, Optional: true, Description: "Enter the threshold as a percentage (Threshold in percentage(default is 100%))",
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
func resourceSystemResourceAccountingTemplateNetworkResourcesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingTemplateNetworkResourcesCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccountingTemplateNetworkResources(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemResourceAccountingTemplateNetworkResourcesRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemResourceAccountingTemplateNetworkResourcesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingTemplateNetworkResourcesUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccountingTemplateNetworkResources(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemResourceAccountingTemplateNetworkResourcesRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemResourceAccountingTemplateNetworkResourcesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingTemplateNetworkResourcesDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccountingTemplateNetworkResources(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemResourceAccountingTemplateNetworkResourcesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceAccountingTemplateNetworkResourcesRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceAccountingTemplateNetworkResources(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesIpv4AclLineCfg(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesIpv4AclLineCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesIpv4AclLineCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv4AclLineMax = in["ipv4_acl_line_max"].(int)
		ret.Ipv4AclLineMinGuarantee = in["ipv4_acl_line_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesIpv6AclLineCfg(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesIpv6AclLineCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesIpv6AclLineCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ipv6AclLineMax = in["ipv6_acl_line_max"].(int)
		ret.Ipv6AclLineMinGuarantee = in["ipv6_acl_line_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesObjectGroupCfg(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesObjectGroupCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesObjectGroupCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ObjectGroupMax = in["object_group_max"].(int)
		ret.ObjectGroupMinGuarantee = in["object_group_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesObjectGroupClauseCfg(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesObjectGroupClauseCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesObjectGroupClauseCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ObjectGroupClauseMax = in["object_group_clause_max"].(int)
		ret.ObjectGroupClauseMinGuarantee = in["object_group_clause_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesStaticArpCfg(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesStaticArpCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesStaticArpCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticArpMax = in["static_arp_max"].(int)
		ret.StaticArpMinGuarantee = in["static_arp_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesStaticIpv4RouteCfg(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesStaticIpv4RouteCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesStaticIpv4RouteCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticIpv4RouteMax = in["static_ipv4_route_max"].(int)
		ret.StaticIpv4RouteMinGuarantee = in["static_ipv4_route_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesStaticIpv6RouteCfg(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesStaticIpv6RouteCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesStaticIpv6RouteCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticIpv6RouteMax = in["static_ipv6_route_max"].(int)
		ret.StaticIpv6RouteMinGuarantee = in["static_ipv6_route_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesStaticMacCfg(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesStaticMacCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesStaticMacCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticMacMax = in["static_mac_max"].(int)
		ret.StaticMacMinGuarantee = in["static_mac_min_guarantee"].(int)
	}
	return ret
}

func getObjectSystemResourceAccountingTemplateNetworkResourcesStaticNeighborCfg(d []interface{}) edpt.SystemResourceAccountingTemplateNetworkResourcesStaticNeighborCfg {

	count1 := len(d)
	var ret edpt.SystemResourceAccountingTemplateNetworkResourcesStaticNeighborCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StaticNeighborMax = in["static_neighbor_max"].(int)
		ret.StaticNeighborMinGuarantee = in["static_neighbor_min_guarantee"].(int)
	}
	return ret
}

func dataToEndpointSystemResourceAccountingTemplateNetworkResources(d *schema.ResourceData) edpt.SystemResourceAccountingTemplateNetworkResources {
	var ret edpt.SystemResourceAccountingTemplateNetworkResources
	ret.Inst.Ipv4AclLineCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesIpv4AclLineCfg(d.Get("ipv4_acl_line_cfg").([]interface{}))
	ret.Inst.Ipv6AclLineCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesIpv6AclLineCfg(d.Get("ipv6_acl_line_cfg").([]interface{}))
	ret.Inst.ObjectGroupCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesObjectGroupCfg(d.Get("object_group_cfg").([]interface{}))
	ret.Inst.ObjectGroupClauseCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesObjectGroupClauseCfg(d.Get("object_group_clause_cfg").([]interface{}))
	ret.Inst.StaticArpCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesStaticArpCfg(d.Get("static_arp_cfg").([]interface{}))
	ret.Inst.StaticIpv4RouteCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesStaticIpv4RouteCfg(d.Get("static_ipv4_route_cfg").([]interface{}))
	ret.Inst.StaticIpv6RouteCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesStaticIpv6RouteCfg(d.Get("static_ipv6_route_cfg").([]interface{}))
	ret.Inst.StaticMacCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesStaticMacCfg(d.Get("static_mac_cfg").([]interface{}))
	ret.Inst.StaticNeighborCfg = getObjectSystemResourceAccountingTemplateNetworkResourcesStaticNeighborCfg(d.Get("static_neighbor_cfg").([]interface{}))
	ret.Inst.Threshold = d.Get("threshold").(int)
	//omit uuid
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
