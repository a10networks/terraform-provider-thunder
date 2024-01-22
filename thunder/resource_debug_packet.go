package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugPacket() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_packet`: Debug packet\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugPacketCreate,
		UpdateContext: resourceDebugPacketUpdate,
		ReadContext:   resourceDebugPacketRead,
		DeleteContext: resourceDebugPacketDelete,

		Schema: map[string]*schema.Schema{
			"all_ipv4": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "All",
			},
			"all_ipv6": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "All",
			},
			"all_sctp_ports": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "All",
			},
			"all_tcp_ports": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "All",
			},
			"all_udp_ports": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "All",
			},
			"arp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "ARP",
			},
			"count1": {
				Type: schema.TypeInt, Optional: true, Description: "Maximum packets to capture. Default is 3000 (Specify maximum packet number. For unlimited, specify 0)",
			},
			"detail": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print packet content",
			},
			"ethernet": {
				Type: schema.TypeInt, Optional: true, Description: "Ethernet interface number",
			},
			"icmp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"icmpv6": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"interface": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Interface to debug",
			},
			"ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP",
			},
			"ipv4ad": {
				Type: schema.TypeString, Optional: true, Description: "IP Address",
			},
			"ipv6": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IPV6",
			},
			"ipv6ad": {
				Type: schema.TypeString, Optional: true, Description: "IPV6 Address",
			},
			"l3_protocol": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Layer 3 protocol",
			},
			"l4_protocol": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Layer 4 protocol",
			},
			"neighbor": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IPv6 Neighbor/Router",
			},
			"port_range": {
				Type: schema.TypeInt, Optional: true, Description: "Port Number",
			},
			"sctp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"tcp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"timestamp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print timestamp instead of jiffies",
			},
			"udp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"ve": {
				Type: schema.TypeInt, Optional: true, Description: "ve number",
			},
		},
	}
}
func resourceDebugPacketCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugPacketCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugPacket(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugPacketRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugPacketUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugPacketUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugPacket(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugPacketRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugPacketDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugPacketDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugPacket(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugPacketRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugPacketRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugPacket(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugPacket(d *schema.ResourceData) edpt.DebugPacket {
	var ret edpt.DebugPacket
	ret.Inst.AllIpv4 = d.Get("all_ipv4").(int)
	ret.Inst.AllIpv6 = d.Get("all_ipv6").(int)
	ret.Inst.AllSctpPorts = d.Get("all_sctp_ports").(int)
	ret.Inst.AllTcpPorts = d.Get("all_tcp_ports").(int)
	ret.Inst.AllUdpPorts = d.Get("all_udp_ports").(int)
	ret.Inst.Arp = d.Get("arp").(int)
	ret.Inst.Count1 = d.Get("count1").(int)
	ret.Inst.Detail = d.Get("detail").(int)
	ret.Inst.Ethernet = d.Get("ethernet").(int)
	ret.Inst.Icmp = d.Get("icmp").(int)
	ret.Inst.Icmpv6 = d.Get("icmpv6").(int)
	ret.Inst.Interface = d.Get("interface").(int)
	ret.Inst.Ip = d.Get("ip").(int)
	ret.Inst.Ipv4ad = d.Get("ipv4ad").(string)
	ret.Inst.Ipv6 = d.Get("ipv6").(int)
	ret.Inst.Ipv6ad = d.Get("ipv6ad").(string)
	ret.Inst.L3Protocol = d.Get("l3_protocol").(int)
	ret.Inst.L4Protocol = d.Get("l4_protocol").(int)
	ret.Inst.Neighbor = d.Get("neighbor").(int)
	ret.Inst.PortRange = d.Get("port_range").(int)
	ret.Inst.Sctp = d.Get("sctp").(int)
	ret.Inst.Tcp = d.Get("tcp").(int)
	ret.Inst.Timestamp = d.Get("timestamp").(int)
	ret.Inst.Udp = d.Get("udp").(int)
	//omit uuid
	ret.Inst.Ve = d.Get("ve").(int)
	return ret
}
