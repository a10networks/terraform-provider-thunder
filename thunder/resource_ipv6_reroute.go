package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6Reroute() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_reroute`: reroute sessions\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6RerouteCreate,
		UpdateContext: resourceIpv6RerouteUpdate,
		ReadContext:   resourceIpv6RerouteRead,
		DeleteContext: resourceIpv6RerouteDelete,

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
func resourceIpv6RerouteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RerouteCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6Reroute(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RerouteRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6RerouteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RerouteUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6Reroute(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RerouteRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6RerouteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RerouteDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6Reroute(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6RerouteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RerouteRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6Reroute(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectIpv6RerouteSuppressProtocols1044(d []interface{}) edpt.Ipv6RerouteSuppressProtocols1044 {

	count1 := len(d)
	var ret edpt.Ipv6RerouteSuppressProtocols1044
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

func dataToEndpointIpv6Reroute(d *schema.ResourceData) edpt.Ipv6Reroute {
	var ret edpt.Ipv6Reroute
	ret.Inst.SuppressProtocols = getObjectIpv6RerouteSuppressProtocols1044(d.Get("suppress_protocols").([]interface{}))
	//omit uuid
	return ret
}
