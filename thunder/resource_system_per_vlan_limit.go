package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemPerVlanLimit() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_per_vlan_limit`: Per vlan flooding packet limit\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemPerVlanLimitCreate,
		UpdateContext: resourceSystemPerVlanLimitUpdate,
		ReadContext:   resourceSystemPerVlanLimitRead,
		DeleteContext: resourceSystemPerVlanLimitDelete,

		Schema: map[string]*schema.Schema{
			"bcast": {
				Type: schema.TypeInt, Optional: true, Default: 1000, Description: "broadcast packets (per second limit)",
			},
			"ipmcast": {
				Type: schema.TypeInt, Optional: true, Default: 1000, Description: "IP multicast packets (per second limit)",
			},
			"mcast": {
				Type: schema.TypeInt, Optional: true, Default: 1000, Description: "multicast packets (per second limit)",
			},
			"unknown_ucast": {
				Type: schema.TypeInt, Optional: true, Default: 1000, Description: "unknown unicast packets (per second limit)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemPerVlanLimitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPerVlanLimitCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPerVlanLimit(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemPerVlanLimitRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemPerVlanLimitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPerVlanLimitUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPerVlanLimit(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemPerVlanLimitRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemPerVlanLimitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPerVlanLimitDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPerVlanLimit(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemPerVlanLimitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPerVlanLimitRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPerVlanLimit(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemPerVlanLimit(d *schema.ResourceData) edpt.SystemPerVlanLimit {
	var ret edpt.SystemPerVlanLimit
	ret.Inst.Bcast = d.Get("bcast").(int)
	ret.Inst.Ipmcast = d.Get("ipmcast").(int)
	ret.Inst.Mcast = d.Get("mcast").(int)
	ret.Inst.UnknownUcast = d.Get("unknown_ucast").(int)
	//omit uuid
	return ret
}
