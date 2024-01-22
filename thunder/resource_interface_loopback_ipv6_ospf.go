package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterfaceLoopbackIpv6Ospf() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_interface_loopback_ipv6_ospf`: Open Shortest Path First for IPv6 (OSPFv3)\n\n__PLACEHOLDER__",
		CreateContext: resourceInterfaceLoopbackIpv6OspfCreate,
		UpdateContext: resourceInterfaceLoopbackIpv6OspfUpdate,
		ReadContext:   resourceInterfaceLoopbackIpv6OspfRead,
		DeleteContext: resourceInterfaceLoopbackIpv6OspfDelete,

		Schema: map[string]*schema.Schema{
			"bfd": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Bidirectional Forwarding Detection (BFD)",
			},
			"cost_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cost": {
							Type: schema.TypeInt, Optional: true, Description: "Interface cost",
						},
						"instance_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
						},
					},
				},
			},
			"dead_interval_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dead_interval": {
							Type: schema.TypeInt, Optional: true, Default: 40, Description: "Interval after which a neighbor is declared dead (Seconds)",
						},
						"instance_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
						},
					},
				},
			},
			"disable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable BFD",
			},
			"hello_interval_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hello_interval": {
							Type: schema.TypeInt, Optional: true, Default: 10, Description: "Time between HELLO packets (Seconds)",
						},
						"instance_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
						},
					},
				},
			},
			"mtu_ignore_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mtu_ignore": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ignores the MTU in DBD packets",
						},
						"instance_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
						},
					},
				},
			},
			"priority_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Router priority",
						},
						"instance_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
						},
					},
				},
			},
			"retransmit_interval_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retransmit_interval": {
							Type: schema.TypeInt, Optional: true, Default: 5, Description: "Time between retransmitting lost link state advertisements (Seconds)",
						},
						"instance_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
						},
					},
				},
			},
			"transmit_delay_cfg": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"transmit_delay": {
							Type: schema.TypeInt, Optional: true, Default: 1, Description: "Link state transmit delay (Seconds)",
						},
						"instance_id": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify the interface instance ID",
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
func resourceInterfaceLoopbackIpv6OspfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpv6OspfCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpv6Ospf(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLoopbackIpv6OspfRead(ctx, d, meta)
	}
	return diags
}

func resourceInterfaceLoopbackIpv6OspfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpv6OspfUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpv6Ospf(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceInterfaceLoopbackIpv6OspfRead(ctx, d, meta)
	}
	return diags
}
func resourceInterfaceLoopbackIpv6OspfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpv6OspfDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpv6Ospf(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceInterfaceLoopbackIpv6OspfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceInterfaceLoopbackIpv6OspfRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointInterfaceLoopbackIpv6Ospf(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceInterfaceLoopbackIpv6OspfCostCfg(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfCostCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfCostCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfCostCfg
		oi.Cost = in["cost"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfDeadIntervalCfg(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfDeadIntervalCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfDeadIntervalCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfDeadIntervalCfg
		oi.DeadInterval = in["dead_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfHelloIntervalCfg(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfHelloIntervalCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfHelloIntervalCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfHelloIntervalCfg
		oi.HelloInterval = in["hello_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfMtuIgnoreCfg(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfMtuIgnoreCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfMtuIgnoreCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfMtuIgnoreCfg
		oi.MtuIgnore = in["mtu_ignore"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfPriorityCfg(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfPriorityCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfPriorityCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfPriorityCfg
		oi.Priority = in["priority"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfRetransmitIntervalCfg(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfRetransmitIntervalCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfRetransmitIntervalCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfRetransmitIntervalCfg
		oi.RetransmitInterval = in["retransmit_interval"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceInterfaceLoopbackIpv6OspfTransmitDelayCfg(d []interface{}) []edpt.InterfaceLoopbackIpv6OspfTransmitDelayCfg {

	count1 := len(d)
	ret := make([]edpt.InterfaceLoopbackIpv6OspfTransmitDelayCfg, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.InterfaceLoopbackIpv6OspfTransmitDelayCfg
		oi.TransmitDelay = in["transmit_delay"].(int)
		oi.InstanceId = in["instance_id"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointInterfaceLoopbackIpv6Ospf(d *schema.ResourceData) edpt.InterfaceLoopbackIpv6Ospf {
	var ret edpt.InterfaceLoopbackIpv6Ospf
	ret.Inst.Bfd = d.Get("bfd").(int)
	ret.Inst.CostCfg = getSliceInterfaceLoopbackIpv6OspfCostCfg(d.Get("cost_cfg").([]interface{}))
	ret.Inst.DeadIntervalCfg = getSliceInterfaceLoopbackIpv6OspfDeadIntervalCfg(d.Get("dead_interval_cfg").([]interface{}))
	ret.Inst.Disable = d.Get("disable").(int)
	ret.Inst.HelloIntervalCfg = getSliceInterfaceLoopbackIpv6OspfHelloIntervalCfg(d.Get("hello_interval_cfg").([]interface{}))
	ret.Inst.MtuIgnoreCfg = getSliceInterfaceLoopbackIpv6OspfMtuIgnoreCfg(d.Get("mtu_ignore_cfg").([]interface{}))
	ret.Inst.PriorityCfg = getSliceInterfaceLoopbackIpv6OspfPriorityCfg(d.Get("priority_cfg").([]interface{}))
	ret.Inst.RetransmitIntervalCfg = getSliceInterfaceLoopbackIpv6OspfRetransmitIntervalCfg(d.Get("retransmit_interval_cfg").([]interface{}))
	ret.Inst.TransmitDelayCfg = getSliceInterfaceLoopbackIpv6OspfTransmitDelayCfg(d.Get("transmit_delay_cfg").([]interface{}))
	//omit uuid
	ret.Inst.Ifnum = d.Get("ifnum").(string)
	return ret
}
