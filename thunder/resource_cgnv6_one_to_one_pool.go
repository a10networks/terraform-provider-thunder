package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6OneToOnePool() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_one_to_one_pool`: Configure CGNv6 one-to-one pool\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6OneToOnePoolCreate,
		UpdateContext: resourceCgnv6OneToOnePoolUpdate,
		ReadContext:   resourceCgnv6OneToOnePoolRead,
		DeleteContext: resourceCgnv6OneToOnePoolDelete,

		Schema: map[string]*schema.Schema{
			"end_address": {
				Type: schema.TypeString, Optional: true, Description: "Configure end IP address of NAT pool",
			},
			"group": {
				Type: schema.TypeString, Optional: true, Description: "Share with a partition group (Partition Group Name)",
			},
			"netmask": {
				Type: schema.TypeString, Optional: true, Description: "Configure mask for pool",
			},
			"partition": {
				Type: schema.TypeString, Optional: true, Description: "Share with a single partition (Partition Name)",
			},
			"pool_name": {
				Type: schema.TypeString, Required: true, Description: "Specify pool name or pool group",
			},
			"shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Share this pool with other partitions (default: not shared)",
			},
			"start_address": {
				Type: schema.TypeString, Optional: true, Description: "Configure start IP address of NAT pool",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vrid": {
				Type: schema.TypeInt, Optional: true, Description: "Configure VRRP-A vrid (Specify ha VRRP-A vrid)",
			},
		},
	}
}
func resourceCgnv6OneToOnePoolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOnePoolCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOnePool(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6OneToOnePoolRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6OneToOnePoolUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOnePoolUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOnePool(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6OneToOnePoolRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6OneToOnePoolDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOnePoolDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOnePool(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6OneToOnePoolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6OneToOnePoolRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6OneToOnePool(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6OneToOnePool(d *schema.ResourceData) edpt.Cgnv6OneToOnePool {
	var ret edpt.Cgnv6OneToOnePool
	ret.Inst.EndAddress = d.Get("end_address").(string)
	ret.Inst.Group = d.Get("group").(string)
	ret.Inst.Netmask = d.Get("netmask").(string)
	ret.Inst.Partition = d.Get("partition").(string)
	ret.Inst.PoolName = d.Get("pool_name").(string)
	ret.Inst.Shared = d.Get("shared").(int)
	ret.Inst.StartAddress = d.Get("start_address").(string)
	//omit uuid
	ret.Inst.Vrid = d.Get("vrid").(int)
	return ret
}
