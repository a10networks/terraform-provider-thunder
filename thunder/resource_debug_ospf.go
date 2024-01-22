package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugOspf() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_ospf`: Debug Open Shortest Path First (OSPF)\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugOspfCreate,
		UpdateContext: resourceDebugOspfUpdate,
		ReadContext:   resourceDebugOspfRead,
		DeleteContext: resourceDebugOspfDelete,

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
func resourceDebugOspfCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspf(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugOspfUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspf(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugOspfRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugOspfDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspf(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugOspfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugOspfRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugOspf(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDebugOspfAll334(d []interface{}) edpt.DebugOspfAll334 {

	count1 := len(d)
	var ret edpt.DebugOspfAll334
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dumy = in["dumy"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDebugOspfBfd335(d []interface{}) edpt.DebugOspfBfd335 {

	count1 := len(d)
	var ret edpt.DebugOspfBfd335
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Dumy = in["dumy"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDebugOspfEvents336(d []interface{}) edpt.DebugOspfEvents336 {

	count1 := len(d)
	var ret edpt.DebugOspfEvents336
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

func getObjectDebugOspfIfsm337(d []interface{}) edpt.DebugOspfIfsm337 {

	count1 := len(d)
	var ret edpt.DebugOspfIfsm337
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Events = in["events"].(int)
		ret.Status = in["status"].(int)
		ret.Timers = in["timers"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDebugOspfLsa338(d []interface{}) edpt.DebugOspfLsa338 {

	count1 := len(d)
	var ret edpt.DebugOspfLsa338
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

func getObjectDebugOspfNfsm339(d []interface{}) edpt.DebugOspfNfsm339 {

	count1 := len(d)
	var ret edpt.DebugOspfNfsm339
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Events = in["events"].(int)
		ret.Status = in["status"].(int)
		ret.Timers = in["timers"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDebugOspfNsm340(d []interface{}) edpt.DebugOspfNsm340 {

	count1 := len(d)
	var ret edpt.DebugOspfNsm340
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Interface = in["interface"].(int)
		ret.Redistribute = in["redistribute"].(int)
		//omit uuid
	}
	return ret
}

func getObjectDebugOspfPacket341(d []interface{}) edpt.DebugOspfPacket341 {

	count1 := len(d)
	var ret edpt.DebugOspfPacket341
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

func getObjectDebugOspfRoute342(d []interface{}) edpt.DebugOspfRoute342 {

	count1 := len(d)
	var ret edpt.DebugOspfRoute342
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

func dataToEndpointDebugOspf(d *schema.ResourceData) edpt.DebugOspf {
	var ret edpt.DebugOspf
	ret.Inst.All = getObjectDebugOspfAll334(d.Get("all").([]interface{}))
	ret.Inst.Bfd = getObjectDebugOspfBfd335(d.Get("bfd").([]interface{}))
	ret.Inst.Dumy = d.Get("dumy").(int)
	ret.Inst.Events = getObjectDebugOspfEvents336(d.Get("events").([]interface{}))
	ret.Inst.Ifsm = getObjectDebugOspfIfsm337(d.Get("ifsm").([]interface{}))
	ret.Inst.Lsa = getObjectDebugOspfLsa338(d.Get("lsa").([]interface{}))
	ret.Inst.Nfsm = getObjectDebugOspfNfsm339(d.Get("nfsm").([]interface{}))
	ret.Inst.Nsm = getObjectDebugOspfNsm340(d.Get("nsm").([]interface{}))
	ret.Inst.Packet = getObjectDebugOspfPacket341(d.Get("packet").([]interface{}))
	ret.Inst.Route = getObjectDebugOspfRoute342(d.Get("route").([]interface{}))
	return ret
}
