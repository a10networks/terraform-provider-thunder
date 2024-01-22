package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpRerouteSuppressProtocols() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ip_reroute_suppress_protocols`: suppress protocols\n\n__PLACEHOLDER__",
		CreateContext: resourceIpRerouteSuppressProtocolsCreate,
		UpdateContext: resourceIpRerouteSuppressProtocolsUpdate,
		ReadContext:   resourceIpRerouteSuppressProtocolsRead,
		DeleteContext: resourceIpRerouteSuppressProtocolsDelete,

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
func resourceIpRerouteSuppressProtocolsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRerouteSuppressProtocolsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRerouteSuppressProtocols(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpRerouteSuppressProtocolsRead(ctx, d, meta)
	}
	return diags
}

func resourceIpRerouteSuppressProtocolsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRerouteSuppressProtocolsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRerouteSuppressProtocols(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceIpRerouteSuppressProtocolsRead(ctx, d, meta)
	}
	return diags
}
func resourceIpRerouteSuppressProtocolsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRerouteSuppressProtocolsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRerouteSuppressProtocols(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceIpRerouteSuppressProtocolsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceIpRerouteSuppressProtocolsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointIpRerouteSuppressProtocols(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointIpRerouteSuppressProtocols(d *schema.ResourceData) edpt.IpRerouteSuppressProtocols {
	var ret edpt.IpRerouteSuppressProtocols
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
