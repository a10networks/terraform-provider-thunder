package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceTrunkIpv6RouterOspf() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_trunk_ipv6_router_ospf`: Open Shortest Path First for IPv6 (OSPFv3)\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceTrunkIpv6RouterOspfCreate,
		UpdateContext: resourceInterfaceTrunkIpv6RouterOspfUpdate,
		ReadContext:   resourceInterfaceTrunkIpv6RouterOspfRead,
		DeleteContext: resourceInterfaceTrunkIpv6RouterOspfDelete,

		Schema: map[string]*schema.Schema{
			"area_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"area_id_num": {
							Type: schema.TypeInt, Optional: true, Description: "OSPF area ID as a decimal value",
						},
						"area_id_addr": {
							Type: schema.TypeString, Optional: true, Description: "OSPF area ID in IP address format",
						},
						"tag": {
							Type: schema.TypeString, Optional: true, Description: "Set the OSPFv3 process tag",
						},
						"instance_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set the interface instance ID",
						},
					},
				},
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
func resourceInterfaceTrunkIpv6RouterOspfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6RouterOspfCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6RouterOspf(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpv6RouterOspfRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceTrunkIpv6RouterOspfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6RouterOspfUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6RouterOspf(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceTrunkIpv6RouterOspfRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceTrunkIpv6RouterOspfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6RouterOspfDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6RouterOspf(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceTrunkIpv6RouterOspfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceTrunkIpv6RouterOspfRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceTrunkIpv6RouterOspf(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceInterfaceTrunkIpv6RouterOspfAreaList(d []interface{}) []edpt.InterfaceTrunkIpv6RouterOspfAreaList {

	count1 := len(d)
	ret := make([]edpt.InterfaceTrunkIpv6RouterOspfAreaList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceTrunkIpv6RouterOspfAreaList
		oi.AreaIdNum = in["area_id_num"].(int)
		oi.AreaIdAddr = in["area_id_addr"].(string)
		oi.Tag = in["tag"].(string)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceTrunkIpv6RouterOspf(d *schema.ResourceData) edpt.InterfaceTrunkIpv6RouterOspf {
	var ret edpt.InterfaceTrunkIpv6RouterOspf
	ret.Inst.AreaList = getSliceInterfaceTrunkIpv6RouterOspfAreaList(d.Get("area_list").([]interface{}))
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
