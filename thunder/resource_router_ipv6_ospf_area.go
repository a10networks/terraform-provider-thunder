package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterIpv6OspfArea() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_router_ipv6_ospf_area`: OSPF area parameters\n\n__PLACEHOLDER__",
		CreateContext: resourceRouterIpv6OspfAreaCreate,
		UpdateContext: resourceRouterIpv6OspfAreaUpdate,
		ReadContext:   resourceRouterIpv6OspfAreaRead,
		DeleteContext: resourceRouterIpv6OspfAreaDelete,

		Schema: map[string]*schema.Schema{
			"area_ipv4": {
				Type: schema.TypeString, Required: true, Description: "OSPFv3 area ID in IP address format",
			},
			"area_num": {
				Type: schema.TypeInt, Required: true, Description: "OSPFv3 area ID as a decimal value",
			},
			"default_cost": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set the summary-default cost of a NSSA or stub area (Stub's advertised default summary cost)",
			},
			"no_summary": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Do not inject inter-area routes into area",
			},
			"range_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type: schema.TypeString, Optional: true, Description: "Area range for IPv6 prefix",
						},
						"option": {
							Type: schema.TypeString, Optional: true, Description: "'advertise': Advertise this range (default); 'not-advertise': DoNotAdvertise this range;",
						},
					},
				},
			},
			"stub": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Configure OSPFv3 area as stub",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"virtual_link_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type: schema.TypeString, Optional: true, Description: "ID (IP addr) associated with virtual link neighbor",
						},
						"dead_interval": {
							Type: schema.TypeInt, Optional: true, Description: "Dead router detection time (Seconds)",
						},
						"bfd": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
						},
						"hello_interval": {
							Type: schema.TypeInt, Optional: true, Description: "Hello packet interval (Seconds)",
						},
						"retransmit_interval": {
							Type: schema.TypeInt, Optional: true, Description: "LSA retransmit interval (Seconds)",
						},
						"transmit_delay": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "LSA transmission delay (Seconds)",
						},
						"instance_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPFv3 instance ID",
						},
					},
				},
			},
			"process_id": {
				Type: schema.TypeString, Required: true, Description: "ProcessId",
			},
		},
	}
}
func resourceRouterIpv6OspfAreaCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6OspfAreaCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6OspfArea(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6OspfAreaRead(ctx, d, meta)
	}
	return diags
}

func resourceRouterIpv6OspfAreaUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6OspfAreaUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6OspfArea(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRouterIpv6OspfAreaRead(ctx, d, meta)
	}
	return diags
}
func resourceRouterIpv6OspfAreaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6OspfAreaDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6OspfArea(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRouterIpv6OspfAreaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRouterIpv6OspfAreaRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRouterIpv6OspfArea(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceRouterIpv6OspfAreaRangeList(d []interface{}) []edpt.RouterIpv6OspfAreaRangeList {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfAreaRangeList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfAreaRangeList
		oi.Value = in["value"].(string)
		oi.Option = in["option"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceRouterIpv6OspfAreaVirtualLinkList(d []interface{}) []edpt.RouterIpv6OspfAreaVirtualLinkList {

	count1 := len(d)
	ret := make([]edpt.RouterIpv6OspfAreaVirtualLinkList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.RouterIpv6OspfAreaVirtualLinkList
		oi.Value = in["value"].(string)
		oi.DeadInterval = in["dead_interval"].(int)
		oi.Bfd = in["bfd"].(int)
		oi.HelloInterval = in["hello_interval"].(int)
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointRouterIpv6OspfArea(d *schema.ResourceData) edpt.RouterIpv6OspfArea {
	var ret edpt.RouterIpv6OspfArea
	ret.Inst.AreaIpv4 = d.Get("area_ipv4").(string)
	ret.Inst.AreaNum = d.Get("area_num").(int)
	ret.Inst.DefaultCost = d.Get("default_cost").(int)
	ret.Inst.NoSummary = d.Get("no_summary").(int)
	ret.Inst.RangeList = getSliceRouterIpv6OspfAreaRangeList(d.Get("range_list").([]interface{}))
	ret.Inst.Stub = d.Get("stub").(int)
	//omit uuid
	ret.Inst.VirtualLinkList = getSliceRouterIpv6OspfAreaVirtualLinkList(d.Get("virtual_link_list").([]interface{}))
	ret.Inst.ProcessId = d.Get("process_id").(string)
	return ret
}
