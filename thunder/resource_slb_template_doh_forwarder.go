package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSlbTemplateDohForwarder() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_slb_template_doh_forwarder`: DNS over HTTP(s) template forwarding policy\n\n__PLACEHOLDER__",
		CreateContext: resourceSlbTemplateDohForwarderCreate,
		UpdateContext: resourceSlbTemplateDohForwarderUpdate,
		ReadContext:   resourceSlbTemplateDohForwarderRead,
		DeleteContext: resourceSlbTemplateDohForwarderDelete,

		Schema: map[string]*schema.Schema{
			"bypass_doh": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Forward valid DoH HTTP request as is, no DNS packet extraction (Bypass DoH)",
			},
			"forwarding_ipv4": {
				Type: schema.TypeString, Optional: true, Description: "SLB VIP IPv4 address to forward DOH query (IP address)",
			},
			"forwarding_ipv6": {
				Type: schema.TypeString, Optional: true, Description: "SLB VIP IPv6 address to forward DOH query (IP address)",
			},
			"tcp_service_group": {
				Type: schema.TypeString, Optional: true, Description: "Bind a TCP Service Group to the template (Service Group Name)",
			},
			"udp_service_group": {
				Type: schema.TypeString, Optional: true, Description: "Bind a UDP Service Group to the template (Service Group Name)",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"v4_internal": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Try to find this IP as a VIP in this L3v Partition and forward it internally to the VIP",
			},
			"v4_l4_proto": {
				Type: schema.TypeString, Optional: true, Default: "both", Description: "'tcp': Use TCP only when forwarding DNS traffic; 'udp': Use UDP only when forwarding DNS traffic; 'both': Use UDP 1st and if unreachable, retry with TCP when forwarding DNS traffic;",
			},
			"v4_port": {
				Type: schema.TypeInt, Optional: true, Default: 53, Description: "Forwarding port number, Default is 53",
			},
			"v6_internal": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Try to find this IP as a VIP in this L3v Partition and forward it internally to the VIP",
			},
			"v6_l4_proto": {
				Type: schema.TypeString, Optional: true, Default: "both", Description: "'tcp': Use TCP only when forwarding DNS traffic; 'udp': Use UDP only when forwarding DNS traffic; 'both': Use UDP 1st and if unreachable, retry with TCP when forwarding DNS traffic;",
			},
			"v6_port": {
				Type: schema.TypeInt, Optional: true, Default: 53, Description: "Forwarding port number, Default is 53",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSlbTemplateDohForwarderCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDohForwarderCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDohForwarder(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDohForwarderRead(ctx, d, meta)
	}
	return diags
}

func resourceSlbTemplateDohForwarderUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDohForwarderUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDohForwarder(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSlbTemplateDohForwarderRead(ctx, d, meta)
	}
	return diags
}
func resourceSlbTemplateDohForwarderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDohForwarderDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDohForwarder(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSlbTemplateDohForwarderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSlbTemplateDohForwarderRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSlbTemplateDohForwarder(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSlbTemplateDohForwarder(d *schema.ResourceData) edpt.SlbTemplateDohForwarder {
	var ret edpt.SlbTemplateDohForwarder
	ret.Inst.BypassDoh = d.Get("bypass_doh").(int)
	ret.Inst.ForwardingIpv4 = d.Get("forwarding_ipv4").(string)
	ret.Inst.ForwardingIpv6 = d.Get("forwarding_ipv6").(string)
	ret.Inst.TcpServiceGroup = d.Get("tcp_service_group").(string)
	ret.Inst.UdpServiceGroup = d.Get("udp_service_group").(string)
	//omit uuid
	ret.Inst.V4Internal = d.Get("v4_internal").(int)
	ret.Inst.V4L4Proto = d.Get("v4_l4_proto").(string)
	ret.Inst.V4Port = d.Get("v4_port").(int)
	ret.Inst.V6Internal = d.Get("v6_internal").(int)
	ret.Inst.V6L4Proto = d.Get("v6_l4_proto").(string)
	ret.Inst.V6Port = d.Get("v6_port").(int)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
