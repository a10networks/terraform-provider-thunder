package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpReroute() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_reroute`: reroute sessions\n\n__PLACEHOLDER__",
		CreateContext: resourceIpRerouteCreate,
		UpdateContext: resourceIpRerouteUpdate,
		ReadContext:   resourceIpRerouteRead,
		DeleteContext: resourceIpRerouteDelete,

		Schema: map[string]*schema.Schema{
			"suppress_protocols": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ospf": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPF",
						},
						"ebgp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "EBGP",
						},
						"ibgp": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "IBGP",
						},
						"static": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"isis": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "ISIS",
						},
						"rip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "RIP",
						},
						"connected": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpRerouteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRerouteCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpReroute(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpRerouteRead(ctx, d, meta)
	}
	return diags
}

func resourceIpRerouteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRerouteUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpReroute(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpRerouteRead(ctx, d, meta)
	}
	return diags
}
func resourceIpRerouteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRerouteDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpReroute(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpRerouteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRerouteRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpReroute(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectIpRerouteSuppressProtocols1042(d []interface{}) edpt.IpRerouteSuppressProtocols1042 {

	count1 := len(d)
	var ret edpt.IpRerouteSuppressProtocols1042
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Ospf = in["ospf"].(int)
		ret.Ebgp = in["ebgp"].(int)
		ret.Ibgp = in["ibgp"].(int)
		ret.Static = in["static"].(int)
		ret.Isis = in["isis"].(int)
		ret.Rip = in["rip"].(int)
		ret.Connected = in["connected"].(int)
		//omit uuid
	}
	return ret
}

func dataToEndpointIpReroute(d *schema.ResourceData) edpt.IpReroute {
	var ret edpt.IpReroute
	ret.Inst.SuppressProtocols = getObjectIpRerouteSuppressProtocols1042(d.Get("suppress_protocols").([]interface{}))
	//omit uuid
	return ret
}
