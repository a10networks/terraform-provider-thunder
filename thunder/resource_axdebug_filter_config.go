package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAxdebugFilterConfig() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_axdebug_filter_config`: Global debug filter\n\n__PLACEHOLDER__",
		CreateContext: resourceAxdebugFilterConfigCreate,
		UpdateContext: resourceAxdebugFilterConfigUpdate,
		ReadContext:   resourceAxdebugFilterConfigRead,
		DeleteContext: resourceAxdebugFilterConfigDelete,

		Schema: map[string]*schema.Schema{
			"comp_hex": {
				Type: schema.TypeString, Optional: true, Description: "value to compare",
			},
			"dst": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Destination",
			},
			"dst_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "dest IP",
			},
			"dst_ipv4_address": {
				Type: schema.TypeString, Optional: true, Description: "dest ip address",
			},
			"dst_mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "dest mac address",
			},
			"dst_mac_addr": {
				Type: schema.TypeString, Optional: true, Description: "dest mac address",
			},
			"dst_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "dest port number",
			},
			"dst_port_num": {
				Type: schema.TypeInt, Optional: true, Description: "dest Port number",
			},
			"hex": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Define hex value",
			},
			"integer": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Define decimal value",
			},
			"integer_comp": {
				Type: schema.TypeInt, Optional: true, Description: "value to compare",
			},
			"integer_max": {
				Type: schema.TypeInt, Optional: true, Description: "max value",
			},
			"integer_min": {
				Type: schema.TypeInt, Optional: true, Description: "min value",
			},
			"ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IP",
			},
			"ipv4_address": {
				Type: schema.TypeString, Optional: true, Description: "ip address",
			},
			"ipv4_netmask": {
				Type: schema.TypeString, Optional: true, Description: "IP subnet mask",
			},
			"ipv6": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "IPV6",
			},
			"ipv6_address": {
				Type: schema.TypeString, Optional: true, Description: "ipv6 address",
			},
			"l3_proto": {
				Type: schema.TypeString, Optional: true, Description: "'arp': arp; 'neighbor': neighbor;",
			},
			"length": {
				Type: schema.TypeInt, Optional: true, Description: "byte length",
			},
			"mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "mac address",
			},
			"mac_addr": {
				Type: schema.TypeString, Optional: true, Description: "mac address",
			},
			"max_hex": {
				Type: schema.TypeString, Optional: true, Description: "max value",
			},
			"min_hex": {
				Type: schema.TypeString, Optional: true, Description: "min value",
			},
			"number": {
				Type: schema.TypeInt, Required: true, Description: "Specify filter id",
			},
			"offset": {
				Type: schema.TypeInt, Optional: true, Description: "byte offset",
			},
			"oper_range": {
				Type: schema.TypeString, Optional: true, Description: "'gt': greater than; 'gte': greater than or equal to; 'se': smaller than or equal to; 'st': smaller than; 'eq': equal to; 'range': select a range;",
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "port configurations",
			},
			"port_num_max": {
				Type: schema.TypeInt, Optional: true, Description: "max port number",
			},
			"port_num_min": {
				Type: schema.TypeInt, Optional: true, Description: "min port number",
			},
			"prot_num": {
				Type: schema.TypeInt, Optional: true, Description: "protocol number",
			},
			"proto": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "ip protocol number",
			},
			"proto_val": {
				Type: schema.TypeString, Optional: true, Description: "'icmp': icmp; 'icmpv6': icmpv6; 'tcp': tcp; 'udp': udp;",
			},
			"src": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Src",
			},
			"src_ip": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "src IP",
			},
			"src_ipv4_address": {
				Type: schema.TypeString, Optional: true, Description: "src ip address",
			},
			"src_mac": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "src mac address",
			},
			"src_mac_addr": {
				Type: schema.TypeString, Optional: true, Description: "src mac address",
			},
			"src_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "src port number",
			},
			"src_port_num": {
				Type: schema.TypeInt, Optional: true, Description: "src Port number",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"word": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Define hex value",
			},
			"word0": {
				Type: schema.TypeString, Optional: true, Description: "WORD0 to compare",
			},
			"word1": {
				Type: schema.TypeString, Optional: true, Description: "WORD min value",
			},
			"word2": {
				Type: schema.TypeString, Optional: true, Description: "WORD max value",
			},
		},
	}
}
func resourceAxdebugFilterConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugFilterConfigCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugFilterConfig(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAxdebugFilterConfigRead(ctx, d, meta)
	}
	return diags
}

func resourceAxdebugFilterConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugFilterConfigUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugFilterConfig(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAxdebugFilterConfigRead(ctx, d, meta)
	}
	return diags
}
func resourceAxdebugFilterConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugFilterConfigDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugFilterConfig(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAxdebugFilterConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugFilterConfigRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugFilterConfig(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAxdebugFilterConfig(d *schema.ResourceData) edpt.AxdebugFilterConfig {
	var ret edpt.AxdebugFilterConfig
	ret.Inst.CompHex = d.Get("comp_hex").(string)
	ret.Inst.Dst = d.Get("dst").(int)
	ret.Inst.DstIp = d.Get("dst_ip").(int)
	ret.Inst.DstIpv4Address = d.Get("dst_ipv4_address").(string)
	ret.Inst.DstMac = d.Get("dst_mac").(int)
	ret.Inst.DstMacAddr = d.Get("dst_mac_addr").(string)
	ret.Inst.DstPort = d.Get("dst_port").(int)
	ret.Inst.DstPortNum = d.Get("dst_port_num").(int)
	ret.Inst.Hex = d.Get("hex").(int)
	ret.Inst.Integer = d.Get("integer").(int)
	ret.Inst.IntegerComp = d.Get("integer_comp").(int)
	ret.Inst.IntegerMax = d.Get("integer_max").(int)
	ret.Inst.IntegerMin = d.Get("integer_min").(int)
	ret.Inst.Ip = d.Get("ip").(int)
	ret.Inst.Ipv4Address = d.Get("ipv4_address").(string)
	ret.Inst.Ipv4Netmask = d.Get("ipv4_netmask").(string)
	ret.Inst.Ipv6 = d.Get("ipv6").(int)
	ret.Inst.Ipv6Address = d.Get("ipv6_address").(string)
	ret.Inst.L3Proto = d.Get("l3_proto").(string)
	ret.Inst.Length = d.Get("length").(int)
	ret.Inst.Mac = d.Get("mac").(int)
	ret.Inst.MacAddr = d.Get("mac_addr").(string)
	ret.Inst.MaxHex = d.Get("max_hex").(string)
	ret.Inst.MinHex = d.Get("min_hex").(string)
	ret.Inst.Number = d.Get("number").(int)
	ret.Inst.Offset = d.Get("offset").(int)
	ret.Inst.OperRange = d.Get("oper_range").(string)
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.PortNumMax = d.Get("port_num_max").(int)
	ret.Inst.PortNumMin = d.Get("port_num_min").(int)
	ret.Inst.ProtNum = d.Get("prot_num").(int)
	ret.Inst.Proto = d.Get("proto").(int)
	ret.Inst.ProtoVal = d.Get("proto_val").(string)
	ret.Inst.Src = d.Get("src").(int)
	ret.Inst.SrcIp = d.Get("src_ip").(int)
	ret.Inst.SrcIpv4Address = d.Get("src_ipv4_address").(string)
	ret.Inst.SrcMac = d.Get("src_mac").(int)
	ret.Inst.SrcMacAddr = d.Get("src_mac_addr").(string)
	ret.Inst.SrcPort = d.Get("src_port").(int)
	ret.Inst.SrcPortNum = d.Get("src_port_num").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Word = d.Get("word").(int)
	ret.Inst.Word0 = d.Get("word0").(string)
	ret.Inst.Word1 = d.Get("word1").(string)
	ret.Inst.Word2 = d.Get("word2").(string)
	return ret
}
