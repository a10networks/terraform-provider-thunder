package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugIpv6Ospf() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ipv6_ospf`: Open Shortest Path First (OSPF) for IPv6\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugIpv6OspfCreate,
		UpdateContext: resourceDebugIpv6OspfUpdate,
		ReadContext:   resourceDebugIpv6OspfRead,
		DeleteContext: resourceDebugIpv6OspfDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dumy": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dummy",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"bfd": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dumy": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dummy",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"dumy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dummy",
			},
			"events": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"abr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPF ABR events",
						},
						"asbr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPF ASBR events",
						},
						"os": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "OS events",
						},
						"router": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Other router events",
						},
						"vlink": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Virtual-Link event",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"ifsm": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"events": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "IFSM Event Information",
						},
						"status": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "IFSM Status Information",
						},
						"timers": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "IFSM Timer Information",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"lsa": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"flooding": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "LSA Flooding",
						},
						"gererate": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "LSA Generation",
						},
						"install": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "LSA Installation",
						},
						"maxage": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "LSA MaxAge processing",
						},
						"refresh": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "LSA Refreshment",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"nfsm": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"events": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "NFSM Event Information",
						},
						"status": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "NFSM Status Information",
						},
						"timers": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "NFSM Timer Information",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"nsm": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "NSM interface",
						},
						"redistribute": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "NSM redistribute",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"packet": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dd": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPFv3 Database Description",
						},
						"detail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Detail information",
						},
						"hello": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPFv3 Hello",
						},
						"ls_ack": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPFv3 Link State Acknowledgment",
						},
						"ls_request": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPFv3 Link State Request",
						},
						"ls_update": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPFv3 Link State Update",
						},
						"recv": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet received",
						},
						"send": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet sent",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"route": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ase": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "External route calculation information",
						},
						"ia": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Inter-Area route calculation information",
						},
						"install": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Route installation information",
						},
						"spf": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "SPF calculation information",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
		},
	}
}
func resourceDebugIpv6OspfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6Ospf(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6OspfRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugIpv6OspfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6Ospf(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugIpv6OspfRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugIpv6OspfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6Ospf(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugIpv6OspfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugIpv6OspfRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugIpv6Ospf(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDebugIpv6OspfAll325(d []interface{}) edpt.DebugIpv6OspfAll325 {

	count1 := len(d)
	var ret edpt.DebugIpv6OspfAll325
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dumy = in["dumy"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDebugIpv6OspfBfd326(d []interface{}) edpt.DebugIpv6OspfBfd326 {

	count1 := len(d)
	var ret edpt.DebugIpv6OspfBfd326
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dumy = in["dumy"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDebugIpv6OspfEvents327(d []interface{}) edpt.DebugIpv6OspfEvents327 {

	count1 := len(d)
	var ret edpt.DebugIpv6OspfEvents327
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Abr = in["abr"].(int)
		ret.Asbr = in["asbr"].(int)
		ret.Os = in["os"].(int)
		ret.Router = in["router"].(int)
		ret.Vlink = in["vlink"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDebugIpv6OspfIfsm328(d []interface{}) edpt.DebugIpv6OspfIfsm328 {

	count1 := len(d)
	var ret edpt.DebugIpv6OspfIfsm328
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Events = in["events"].(int)
		ret.Status = in["status"].(int)
		ret.Timers = in["timers"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDebugIpv6OspfLsa329(d []interface{}) edpt.DebugIpv6OspfLsa329 {

	count1 := len(d)
	var ret edpt.DebugIpv6OspfLsa329
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Flooding = in["flooding"].(int)
		ret.Gererate = in["gererate"].(int)
		ret.Install = in["install"].(int)
		ret.Maxage = in["maxage"].(int)
		ret.Refresh = in["refresh"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDebugIpv6OspfNfsm330(d []interface{}) edpt.DebugIpv6OspfNfsm330 {

	count1 := len(d)
	var ret edpt.DebugIpv6OspfNfsm330
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Events = in["events"].(int)
		ret.Status = in["status"].(int)
		ret.Timers = in["timers"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDebugIpv6OspfNsm331(d []interface{}) edpt.DebugIpv6OspfNsm331 {

	count1 := len(d)
	var ret edpt.DebugIpv6OspfNsm331
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interface = in["interface"].(int)
		ret.Redistribute = in["redistribute"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDebugIpv6OspfPacket332(d []interface{}) edpt.DebugIpv6OspfPacket332 {

	count1 := len(d)
	var ret edpt.DebugIpv6OspfPacket332
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dd = in["dd"].(int)
		ret.Detail = in["detail"].(int)
		ret.Hello = in["hello"].(int)
		ret.LsAck = in["ls_ack"].(int)
		ret.LsRequest = in["ls_request"].(int)
		ret.LsUpdate = in["ls_update"].(int)
		ret.Recv = in["recv"].(int)
		ret.Send = in["send"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDebugIpv6OspfRoute333(d []interface{}) edpt.DebugIpv6OspfRoute333 {

	count1 := len(d)
	var ret edpt.DebugIpv6OspfRoute333
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ase = in["ase"].(int)
		ret.Ia = in["ia"].(int)
		ret.Install = in["install"].(int)
		ret.Spf = in["spf"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointDebugIpv6Ospf(d *schema.ResourceData) edpt.DebugIpv6Ospf {
	var ret edpt.DebugIpv6Ospf
	ret.Inst.All = getObjectDebugIpv6OspfAll325(d.Get("all").([]interface{}))
	ret.Inst.Bfd = getObjectDebugIpv6OspfBfd326(d.Get("bfd").([]interface{}))
	ret.Inst.Dumy = d.Get("dumy").(int)
	ret.Inst.Events = getObjectDebugIpv6OspfEvents327(d.Get("events").([]interface{}))
	ret.Inst.Ifsm = getObjectDebugIpv6OspfIfsm328(d.Get("ifsm").([]interface{}))
	ret.Inst.Lsa = getObjectDebugIpv6OspfLsa329(d.Get("lsa").([]interface{}))
	ret.Inst.Nfsm = getObjectDebugIpv6OspfNfsm330(d.Get("nfsm").([]interface{}))
	ret.Inst.Nsm = getObjectDebugIpv6OspfNsm331(d.Get("nsm").([]interface{}))
	ret.Inst.Packet = getObjectDebugIpv6OspfPacket332(d.Get("packet").([]interface{}))
	ret.Inst.Route = getObjectDebugIpv6OspfRoute333(d.Get("route").([]interface{}))
	return ret
}
