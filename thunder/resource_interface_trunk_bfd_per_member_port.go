package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTrunkBfdPerMemberPort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_trunk_bfd_per_member_port`: Configure a Micro BFD\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTrunkBfdPerMemberPortCreate,
		UpdateContext: resourceInterfaceTrunkBfdPerMemberPortUpdate,
		ReadContext:   resourceInterfaceTrunkBfdPerMemberPortRead,
		DeleteContext: resourceInterfaceTrunkBfdPerMemberPortDelete,

		Schema: map[string]*schema.Schema{
			"ipv6_local": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 local-address",
			},
			"ipv6_nbr": {
				Type: schema.TypeString, Optional: true, Description: "IPv6 neighbor-address",
			},
			"local_address": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 local-address",
			},
			"neighbor_address": {
				Type: schema.TypeString, Optional: true, Description: "IPv4 neighbor address",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ifnum": {
				Type: schema.TypeString, Required: true, Description: "Ifnum",
			},
		},
	}
}
func resourceInterfaceTrunkBfdPerMemberPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkBfdPerMemberPortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkBfdPerMemberPort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkBfdPerMemberPortRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTrunkBfdPerMemberPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkBfdPerMemberPortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkBfdPerMemberPort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkBfdPerMemberPortRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTrunkBfdPerMemberPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkBfdPerMemberPortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkBfdPerMemberPort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTrunkBfdPerMemberPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkBfdPerMemberPortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkBfdPerMemberPort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointInterfaceTrunkBfdPerMemberPort(d *schema.ResourceData) edpt.InterfaceTrunkBfdPerMemberPort {
	var ret edpt.InterfaceTrunkBfdPerMemberPort
	ret.Inst.Ipv6Local = d.Get("ipv6_local").(string)
	ret.Inst.Ipv6Nbr = d.Get("ipv6_nbr").(string)
	ret.Inst.LocalAddress = d.Get("local_address").(string)
	ret.Inst.NeighborAddress = d.Get("neighbor_address").(string)
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
