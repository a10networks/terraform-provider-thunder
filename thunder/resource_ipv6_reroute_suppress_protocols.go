package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpv6RerouteSuppressProtocols() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ipv6_reroute_suppress_protocols`: suppress protocols\n\n__PLACEHOLDER__",
		CreateContext: resourceIpv6RerouteSuppressProtocolsCreate,
		UpdateContext: resourceIpv6RerouteSuppressProtocolsUpdate,
		ReadContext:   resourceIpv6RerouteSuppressProtocolsRead,
		DeleteContext: resourceIpv6RerouteSuppressProtocolsDelete,

		Schema: map[string]*schema.Schema{
			"connected": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"ebgp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "EBGP",
			},
			"ibgp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IBGP",
			},
			"isis": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "ISIS",
			},
			"ospf": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "OSPF",
			},
			"rip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "RIP",
			},
			"static": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceIpv6RerouteSuppressProtocolsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RerouteSuppressProtocolsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RerouteSuppressProtocols(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RerouteSuppressProtocolsRead(ctx, d, meta)
	}
	return diags
}

func resourceIpv6RerouteSuppressProtocolsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RerouteSuppressProtocolsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RerouteSuppressProtocols(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpv6RerouteSuppressProtocolsRead(ctx, d, meta)
	}
	return diags
}
func resourceIpv6RerouteSuppressProtocolsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RerouteSuppressProtocolsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RerouteSuppressProtocols(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpv6RerouteSuppressProtocolsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpv6RerouteSuppressProtocolsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpv6RerouteSuppressProtocols(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpv6RerouteSuppressProtocols(d *schema.ResourceData) edpt.Ipv6RerouteSuppressProtocols {
	var ret edpt.Ipv6RerouteSuppressProtocols
	ret.Inst.Connected = d.Get("connected").(int)
	ret.Inst.Ebgp = d.Get("ebgp").(int)
	ret.Inst.Ibgp = d.Get("ibgp").(int)
	ret.Inst.Isis = d.Get("isis").(int)
	ret.Inst.Ospf = d.Get("ospf").(int)
	ret.Inst.Rip = d.Get("rip").(int)
	ret.Inst.Static = d.Get("static").(int)
	//omit uuid
	return ret
}
