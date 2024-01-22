package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6ResourceUsage() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_resource_usage`: Configure CGNV6 Resource Usage\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6ResourceUsageCreate,
		UpdateContext: resourceCgnv6ResourceUsageUpdate,
		ReadContext:   resourceCgnv6ResourceUsageRead,
		DeleteContext: resourceCgnv6ResourceUsageDelete,

		Schema: map[string]*schema.Schema{
			"fixed_nat_inside_user_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable CGNV6 Fixed NAT inside users",
			},
			"fixed_nat_ip_addr_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable CGNV6 Fixed NAT addresses",
			},
			"lsn_nat_addr_count": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable CGNV6 NAT Pool addresses",
			},
			"radius_table_size": {
				Type: schema.TypeInt, Optional: true, Description: "Total configurable CGNV6 RADIUS Table entries",
			},
			"stateless_entries": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"l4_session_count": {
							Type: schema.TypeInt, Optional: true, Description: "Helper size for CGN Stateless Technologies",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
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
func resourceCgnv6ResourceUsageCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ResourceUsageCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ResourceUsage(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6ResourceUsageRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6ResourceUsageUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ResourceUsageUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ResourceUsage(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6ResourceUsageRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6ResourceUsageDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ResourceUsageDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ResourceUsage(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6ResourceUsageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6ResourceUsageRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6ResourceUsage(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectCgnv6ResourceUsageStatelessEntries110(d []interface{}) edpt.Cgnv6ResourceUsageStatelessEntries110 {

	count1 := len(d)
	var ret edpt.Cgnv6ResourceUsageStatelessEntries110
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.L4SessionCount = in["l4_session_count"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointCgnv6ResourceUsage(d *schema.ResourceData) edpt.Cgnv6ResourceUsage {
	var ret edpt.Cgnv6ResourceUsage
	ret.Inst.FixedNatInsideUserCount = d.Get("fixed_nat_inside_user_count").(int)
	ret.Inst.FixedNatIpAddrCount = d.Get("fixed_nat_ip_addr_count").(int)
	ret.Inst.LsnNatAddrCount = d.Get("lsn_nat_addr_count").(int)
	ret.Inst.RadiusTableSize = d.Get("radius_table_size").(int)
	ret.Inst.StatelessEntries = getObjectCgnv6ResourceUsageStatelessEntries110(d.Get("stateless_entries").([]interface{}))
	//omit uuid
	return ret
}
