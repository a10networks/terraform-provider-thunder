package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemResourceUsage() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_resource_usage`: Configure System Resource Usage\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemResourceUsageCreate,
		UpdateContext: resourceSystemResourceUsageUpdate,
		ReadContext:   resourceSystemResourceUsageRead,
		DeleteContext: resourceSystemResourceUsageDelete,

		Schema: map[string]*schema.Schema{
			"aflex_table_entry_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total aFleX table entry in the system (Total aFlex entry in the system)",
			},
			"auth_portal_html_file_size": {
				Type: schema.TypeInt, Optional: true, Default: 20, Description: "Specify maximum html file size for each html page in auth portal (in KB)",
			},
			"auth_portal_image_file_size": {
				Type: schema.TypeInt, Optional: true, Default: 6, Description: "Specify maximum image file size for default portal (in KB)",
			},
			"auth_session_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total auth sessions in the system",
			},
			"authz_policy_number": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the maximum number of authorization policies",
			},
			"class_list_ac_entry_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total entries for AC class-list",
			},
			"class_list_entry_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total entries for class-list",
			},
			"class_list_ipv6_addr_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total IPv6 addresses for class-list",
			},
			"ipsec_sa_number": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the maximum number of IPsec SA",
			},
			"l4_session_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total Sessions in the System",
			},
			"max_aflex_authz_collection_number": {
				Type: schema.TypeInt, Optional: true, Default: 512, Description: "Specify the maximum number of collections supported by aFleX authorization",
			},
			"max_aflex_file_size": {
				Type: schema.TypeInt, Optional: true, Default: 32, Description: "Set maximum aFleX file size (Maximum file size in KBytes, default is 32K)",
			},
			"nat_pool_addr_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable NAT Pool addresses in the System",
			},
			"radius_table_size": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable CGNV6 RADIUS Table entries",
			},
			"ram_cache_memory_limit": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the maximum memory used by ram cache",
			},
			"ssl_context_memory": {
				Type: schema.TypeInt, Optional: true, Default: 2048, Description: "Total SSL context memory needed in units of MB. Will be rounded to closest multiple of 2MB",
			},
			"ssl_dma_memory": {
				Type: schema.TypeInt, Optional: true, Default: 256, Description: "Total SSL DMA memory needed in units of MB. Will be rounded to closest multiple of 2MB",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"visibility": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitored_entity_count": {
							Type: schema.TypeInt, Optional: true, Description: "Total number of monitored entities for visibility",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
		},
	}
}
func resourceSystemResourceUsageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceUsageCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceUsage(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemResourceUsageRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemResourceUsageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceUsageUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceUsage(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemResourceUsageRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemResourceUsageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceUsageDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceUsage(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemResourceUsageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemResourceUsageRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemResourceUsage(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSystemResourceUsageVisibility1645(d []interface{}) edpt.SystemResourceUsageVisibility1645 {

	count1 := len(d)
	var ret edpt.SystemResourceUsageVisibility1645
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.MonitoredEntityCount = in["monitored_entity_count"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointSystemResourceUsage(d *schema.ResourceData) edpt.SystemResourceUsage {
	var ret edpt.SystemResourceUsage
	ret.Inst.AflexTableEntryCount = d.Get("aflex_table_entry_count").(int)
	ret.Inst.AuthPortalHtmlFileSize = d.Get("auth_portal_html_file_size").(int)
	ret.Inst.AuthPortalImageFileSize = d.Get("auth_portal_image_file_size").(int)
	ret.Inst.AuthSessionCount = d.Get("auth_session_count").(int)
	ret.Inst.AuthzPolicyNumber = d.Get("authz_policy_number").(int)
	ret.Inst.ClassListAcEntryCount = d.Get("class_list_ac_entry_count").(int)
	ret.Inst.ClassListEntryCount = d.Get("class_list_entry_count").(int)
	ret.Inst.ClassListIpv6AddrCount = d.Get("class_list_ipv6_addr_count").(int)
	ret.Inst.IpsecSaNumber = d.Get("ipsec_sa_number").(int)
	ret.Inst.L4SessionCount = d.Get("l4_session_count").(int)
	ret.Inst.MaxAflexAuthzCollectionNumber = d.Get("max_aflex_authz_collection_number").(int)
	ret.Inst.MaxAflexFileSize = d.Get("max_aflex_file_size").(int)
	ret.Inst.NatPoolAddrCount = d.Get("nat_pool_addr_count").(int)
	ret.Inst.RadiusTableSize = d.Get("radius_table_size").(int)
	ret.Inst.RamCacheMemoryLimit = d.Get("ram_cache_memory_limit").(int)
	ret.Inst.SslContextMemory = d.Get("ssl_context_memory").(int)
	ret.Inst.SslDmaMemory = d.Get("ssl_dma_memory").(int)
	//omit uuid
	ret.Inst.Visibility = getObjectSystemResourceUsageVisibility1645(d.Get("visibility").([]interface{}))
	return ret
}
