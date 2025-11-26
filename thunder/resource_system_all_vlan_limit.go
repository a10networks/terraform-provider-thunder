package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemAllVlanLimit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_all_vlan_limit`: all vlan flooding packet limit\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemAllVlanLimitCreate,
		UpdateContext: resourceSystemAllVlanLimitUpdate,
		ReadContext:   resourceSystemAllVlanLimitRead,
		DeleteContext: resourceSystemAllVlanLimitDelete,

		Schema: map[string]*schema.Schema{
			"bcast": {
				Type: schema.TypeInt, Optional: true, Default: 5000, Description: "broadcast packets (per second limit)",
			},
			"ipmcast": {
				Type: schema.TypeInt, Optional: true, Default: 5000, Description: "IP multicast packets (per second limit)",
			},
			"mcast": {
				Type: schema.TypeInt, Optional: true, Default: 5000, Description: "multicast packets (per second limit)",
			},
			"unknown_ucast": {
				Type: schema.TypeInt, Optional: true, Default: 5000, Description: "unknown unicast packets (per second limit)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemAllVlanLimitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAllVlanLimitCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAllVlanLimit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemAllVlanLimitRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemAllVlanLimitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAllVlanLimitUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAllVlanLimit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemAllVlanLimitRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemAllVlanLimitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAllVlanLimitDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAllVlanLimit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemAllVlanLimitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAllVlanLimitRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAllVlanLimit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemAllVlanLimit(d *schema.ResourceData) edpt.SystemAllVlanLimit {
	var ret edpt.SystemAllVlanLimit
	ret.Inst.Bcast = d.Get("bcast").(int)
	ret.Inst.Ipmcast = d.Get("ipmcast").(int)
	ret.Inst.Mcast = d.Get("mcast").(int)
	ret.Inst.UnknownUcast = d.Get("unknown_ucast").(int)
	//omit uuid
	return ret
}
