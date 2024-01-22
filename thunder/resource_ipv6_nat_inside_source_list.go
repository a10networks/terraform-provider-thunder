package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6NatInsideSourceList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_nat_inside_source_list`: IPv6 Access-List\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6NatInsideSourceListCreate,
		UpdateContext: resourceIpv6NatInsideSourceListUpdate,
		ReadContext:   resourceIpv6NatInsideSourceListRead,
		DeleteContext: resourceIpv6NatInsideSourceListDelete,

		Schema: map[string]*schema.Schema{
			"list_name": {
				Type: schema.TypeString, Required: true, Description: "IPv6 access-list name",
			},
			"pool": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 NAT Pool (Pool Name)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpv6NatInsideSourceListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatInsideSourceListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatInsideSourceList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6NatInsideSourceListRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6NatInsideSourceListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatInsideSourceListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatInsideSourceList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6NatInsideSourceListRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6NatInsideSourceListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatInsideSourceListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatInsideSourceList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6NatInsideSourceListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6NatInsideSourceListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6NatInsideSourceList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpv6NatInsideSourceList(d *schema.ResourceData) edpt.Ipv6NatInsideSourceList {
	var ret edpt.Ipv6NatInsideSourceList
	ret.Inst.ListName = d.Get("list_name").(string)
	ret.Inst.Pool = d.Get("pool").(string)
	//omit uuid
	return ret
}
