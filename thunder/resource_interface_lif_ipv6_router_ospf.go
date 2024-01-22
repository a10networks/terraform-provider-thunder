package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLifIpv6RouterOspf() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_lif_ipv6_router_ospf`: Open Shortest Path First for IPv6 (OSPFv3)\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLifIpv6RouterOspfCreate,
		UpdateContext: resourceInterfaceLifIpv6RouterOspfUpdate,
		ReadContext:   resourceInterfaceLifIpv6RouterOspfRead,
		DeleteContext: resourceInterfaceLifIpv6RouterOspfDelete,

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
			"ifname": {
				Type: schema.TypeString, Required: true, Description: "Ifname",
			},
		},
	}
}
func resourceInterfaceLifIpv6RouterOspfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpv6RouterOspfCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpv6RouterOspf(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIpv6RouterOspfRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLifIpv6RouterOspfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpv6RouterOspfUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpv6RouterOspf(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLifIpv6RouterOspfRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLifIpv6RouterOspfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpv6RouterOspfDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpv6RouterOspf(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLifIpv6RouterOspfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLifIpv6RouterOspfRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLifIpv6RouterOspf(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceInterfaceLifIpv6RouterOspfAreaList(d []interface{}) []edpt.InterfaceLifIpv6RouterOspfAreaList {

	count1 := len(d)
	ret := make([]edpt.InterfaceLifIpv6RouterOspfAreaList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLifIpv6RouterOspfAreaList
		oi.AreaIdNum = in["area_id_num"].(int)
		oi.AreaIdAddr = in["area_id_addr"].(string)
		oi.Tag = in["tag"].(string)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceLifIpv6RouterOspf(d *schema.ResourceData) edpt.InterfaceLifIpv6RouterOspf {
	var ret edpt.InterfaceLifIpv6RouterOspf
	ret.Inst.AreaList = getSliceInterfaceLifIpv6RouterOspfAreaList(d.Get("area_list").([]interface{}))
	//omit uuid
	ret.Inst.Ifname = d.Get("ifname").(string)
	return ret
}
