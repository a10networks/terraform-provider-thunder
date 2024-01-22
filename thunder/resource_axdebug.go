package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAxdebug() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_axdebug`: Packet Trace Options\n\n__PLACEHOLDER__",
		CreateContext: resourceAxdebugCreate,
		UpdateContext: resourceAxdebugUpdate,
		ReadContext:   resourceAxdebugRead,
		DeleteContext: resourceAxdebugDelete1,

		Schema: map[string]*schema.Schema{
			"apply_config": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config_file": {
							Type: schema.TypeString, Optional: true, Description: "config file name",
						},
					},
				},
			},
			"capture": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"brief": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Print basic packet information",
						},
						"detail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Include packet payload",
						},
						"save": {
							Type: schema.TypeString, Optional: true, Description: "Save packets into file (Specify filename to save packets)",
						},
						"current_slot": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Only for current-slot of chassis",
						},
						"no_stop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Non-stop execution",
						},
					},
				},
			},
			"count1": {
				Type: schema.TypeInt, Optional: true, Default: 3000, Description: "Maximum packets to capture per cpu. Default is 3000. (Maximum packets to capture. For umlimited, specify 0)",
			},
			"delete": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"capture_file": {
							Type: schema.TypeString, Optional: true, Description: "Delete a capture file (Specify target filename to change)",
						},
						"config_file": {
							Type: schema.TypeString, Optional: true, Description: "Delete AXDebug config file (Specify target filename to change)",
						},
					},
				},
			},
			"exit": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stop_capture": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "stop capture traffic",
						},
					},
				},
			},
			"filter_config_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"number": {
							Type: schema.TypeInt, Required: true, Description: "Specify filter id",
						},
						"l3_proto": {
							Type: schema.TypeString, Optional: true, Description: "'arp': arp; 'neighbor': neighbor;",
						},
						"dst": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Destination",
						},
						"src": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Src",
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
						"mac": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "mac address",
						},
						"mac_addr": {
							Type: schema.TypeString, Optional: true, Description: "mac address",
						},
						"port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "port configurations",
						},
						"dst_ip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "dest IP",
						},
						"dst_ipv4_address": {
							Type: schema.TypeString, Optional: true, Description: "dest ip address",
						},
						"src_ip": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "src IP",
						},
						"src_ipv4_address": {
							Type: schema.TypeString, Optional: true, Description: "src ip address",
						},
						"dst_mac": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "dest mac address",
						},
						"dst_mac_addr": {
							Type: schema.TypeString, Optional: true, Description: "dest mac address",
						},
						"src_mac": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "src mac address",
						},
						"src_mac_addr": {
							Type: schema.TypeString, Optional: true, Description: "src mac address",
						},
						"dst_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "dest port number",
						},
						"dst_port_num": {
							Type: schema.TypeInt, Optional: true, Description: "dest Port number",
						},
						"src_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "src port number",
						},
						"src_port_num": {
							Type: schema.TypeInt, Optional: true, Description: "src Port number",
						},
						"port_num_min": {
							Type: schema.TypeInt, Optional: true, Description: "min port number",
						},
						"port_num_max": {
							Type: schema.TypeInt, Optional: true, Description: "max port number",
						},
						"proto": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "ip protocol number",
						},
						"proto_val": {
							Type: schema.TypeString, Optional: true, Description: "'icmp': icmp; 'icmpv6': icmpv6; 'tcp': tcp; 'udp': udp;",
						},
						"prot_num": {
							Type: schema.TypeInt, Optional: true, Description: "protocol number",
						},
						"offset": {
							Type: schema.TypeInt, Optional: true, Description: "byte offset",
						},
						"length": {
							Type: schema.TypeInt, Optional: true, Description: "byte length",
						},
						"oper_range": {
							Type: schema.TypeString, Optional: true, Description: "'gt': greater than; 'gte': greater than or equal to; 'se': smaller than or equal to; 'st': smaller than; 'eq': equal to; 'range': select a range;",
						},
						"hex": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Define hex value",
						},
						"min_hex": {
							Type: schema.TypeString, Optional: true, Description: "min value",
						},
						"max_hex": {
							Type: schema.TypeString, Optional: true, Description: "max value",
						},
						"comp_hex": {
							Type: schema.TypeString, Optional: true, Description: "value to compare",
						},
						"integer": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Define decimal value",
						},
						"integer_min": {
							Type: schema.TypeInt, Optional: true, Description: "min value",
						},
						"integer_max": {
							Type: schema.TypeInt, Optional: true, Description: "max value",
						},
						"integer_comp": {
							Type: schema.TypeInt, Optional: true, Description: "value to compare",
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
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
						"user_tag": {
							Type: schema.TypeString, Optional: true, Description: "Customized tag",
						},
					},
				},
			},
			"inc_port_num": {
				Type: schema.TypeString, Optional: true, Description: "Port Numbers separated by commas(,) and hyphens(-) without spaces. ex: 4,5,10-30, or separated by spaces and double-quoted(\")",
			},
			"incoming": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Incoming interface. (For all ports, don't specify port number.)",
			},
			"length": {
				Type: schema.TypeInt, Optional: true, Default: 1518, Description: "Packet length to capture",
			},
			"maxfile": {
				Type: schema.TypeInt, Optional: true, Default: 100, Description: "Maximum number of debug packet files. Default is 100",
			},
			"out_port_num": {
				Type: schema.TypeString, Optional: true, Description: "Port Numbers separated by commas(,) and hyphens(-) without spaces. ex: 4,5,10-30, or separated by spaces and double-quoted(\")",
			},
			"outgoing": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Outgoing interface (For all ports, don't specify port number.)",
			},
			"pcapng_config": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pcapng_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable pcapng",
						},
						"ssl_key_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ssl key tracking",
						},
						"exit": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Exit from axdebug pcapng mode",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"save_config": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config_file": {
							Type: schema.TypeString, Optional: true, Description: "config file name",
						},
						"default": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "save to default config file",
						},
					},
				},
			},
			"sess_filter_dis": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Disable session based filter",
			},
			"timeout": {
				Type: schema.TypeInt, Optional: true, Default: 5, Description: "Maximum number of minutes for a capture. Default is 5 minutes. For unlimited, specify 0",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAxdebugCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebug(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAxdebugRead(ctx, d, meta)
	}
	return diags
}

func resourceAxdebugUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebug(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAxdebugRead(ctx, d, meta)
	}
	return diags
}
func resourceAxdebugDelete1(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugDelete1()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebug(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAxdebugRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebug(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectAxdebugApplyConfig75(d []interface{}) edpt.AxdebugApplyConfig75 {

	count1 := len(d)
	var ret edpt.AxdebugApplyConfig75
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConfigFile = in["config_file"].(string)
	}
	return ret
}

func getObjectAxdebugCapture76(d []interface{}) edpt.AxdebugCapture76 {

	count1 := len(d)
	var ret edpt.AxdebugCapture76
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Brief = in["brief"].(int)
		ret.Detail = in["detail"].(int)
		ret.Save = in["save"].(string)
		ret.CurrentSlot = in["current_slot"].(int)
		ret.NoStop = in["no_stop"].(int)
	}
	return ret
}

func getObjectAxdebugDelete77(d []interface{}) edpt.AxdebugDelete77 {

	count1 := len(d)
	var ret edpt.AxdebugDelete77
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CaptureFile = in["capture_file"].(string)
		ret.ConfigFile = in["config_file"].(string)
	}
	return ret
}

func getObjectAxdebugExit78(d []interface{}) edpt.AxdebugExit78 {

	count1 := len(d)
	var ret edpt.AxdebugExit78
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.StopCapture = in["stop_capture"].(int)
	}
	return ret
}

func getSliceAxdebugFilterConfigList(d []interface{}) []edpt.AxdebugFilterConfigList {

	count1 := len(d)
	ret := make([]edpt.AxdebugFilterConfigList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.AxdebugFilterConfigList
		oi.Number = in["number"].(int)
		oi.L3Proto = in["l3_proto"].(string)
		oi.Dst = in["dst"].(int)
		oi.Src = in["src"].(int)
		oi.Ip = in["ip"].(int)
		oi.Ipv4Address = in["ipv4_address"].(string)
		oi.Ipv4Netmask = in["ipv4_netmask"].(string)
		oi.Ipv6 = in["ipv6"].(int)
		oi.Ipv6Address = in["ipv6_address"].(string)
		oi.Mac = in["mac"].(int)
		oi.MacAddr = in["mac_addr"].(string)
		oi.Port = in["port"].(int)
		oi.DstIp = in["dst_ip"].(int)
		oi.DstIpv4Address = in["dst_ipv4_address"].(string)
		oi.SrcIp = in["src_ip"].(int)
		oi.SrcIpv4Address = in["src_ipv4_address"].(string)
		oi.DstMac = in["dst_mac"].(int)
		oi.DstMacAddr = in["dst_mac_addr"].(string)
		oi.SrcMac = in["src_mac"].(int)
		oi.SrcMacAddr = in["src_mac_addr"].(string)
		oi.DstPort = in["dst_port"].(int)
		oi.DstPortNum = in["dst_port_num"].(int)
		oi.SrcPort = in["src_port"].(int)
		oi.SrcPortNum = in["src_port_num"].(int)
		oi.PortNumMin = in["port_num_min"].(int)
		oi.PortNumMax = in["port_num_max"].(int)
		oi.Proto = in["proto"].(int)
		oi.ProtoVal = in["proto_val"].(string)
		oi.ProtNum = in["prot_num"].(int)
		oi.Offset = in["offset"].(int)
		oi.Length = in["length"].(int)
		oi.OperRange = in["oper_range"].(string)
		oi.Hex = in["hex"].(int)
		oi.MinHex = in["min_hex"].(string)
		oi.MaxHex = in["max_hex"].(string)
		oi.CompHex = in["comp_hex"].(string)
		oi.Integer = in["integer"].(int)
		oi.IntegerMin = in["integer_min"].(int)
		oi.IntegerMax = in["integer_max"].(int)
		oi.IntegerComp = in["integer_comp"].(int)
		oi.Word = in["word"].(int)
		oi.Word0 = in["word0"].(string)
		oi.Word1 = in["word1"].(string)
		oi.Word2 = in["word2"].(string)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectAxdebugPcapngConfig79(d []interface{}) edpt.AxdebugPcapngConfig79 {

	count1 := len(d)
	var ret edpt.AxdebugPcapngConfig79
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PcapngEnable = in["pcapng_enable"].(int)
		ret.SslKeyEnable = in["ssl_key_enable"].(int)
		ret.Exit = in["exit"].(int)
		//omit uuid
	}
	return ret
}

func getObjectAxdebugSaveConfig80(d []interface{}) edpt.AxdebugSaveConfig80 {

	count1 := len(d)
	var ret edpt.AxdebugSaveConfig80
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ConfigFile = in["config_file"].(string)
		ret.Default = in["default"].(int)
	}
	return ret
}

func dataToEndpointAxdebug(d *schema.ResourceData) edpt.Axdebug {
	var ret edpt.Axdebug
	ret.Inst.ApplyConfig = getObjectAxdebugApplyConfig75(d.Get("apply_config").([]interface{}))
	ret.Inst.Capture = getObjectAxdebugCapture76(d.Get("capture").([]interface{}))
	ret.Inst.Count1 = d.Get("count1").(int)
	ret.Inst.Delete = getObjectAxdebugDelete77(d.Get("delete").([]interface{}))
	ret.Inst.Exit = getObjectAxdebugExit78(d.Get("exit").([]interface{}))
	ret.Inst.FilterConfigList = getSliceAxdebugFilterConfigList(d.Get("filter_config_list").([]interface{}))
	ret.Inst.IncPortNum = d.Get("inc_port_num").(string)
	ret.Inst.Incoming = d.Get("incoming").(int)
	ret.Inst.Length = d.Get("length").(int)
	ret.Inst.Maxfile = d.Get("maxfile").(int)
	ret.Inst.OutPortNum = d.Get("out_port_num").(string)
	ret.Inst.Outgoing = d.Get("outgoing").(int)
	ret.Inst.PcapngConfig = getObjectAxdebugPcapngConfig79(d.Get("pcapng_config").([]interface{}))
	ret.Inst.SaveConfig = getObjectAxdebugSaveConfig80(d.Get("save_config").([]interface{}))
	ret.Inst.SessFilterDis = d.Get("sess_filter_dis").(int)
	ret.Inst.Timeout = d.Get("timeout").(int)
	//omit uuid
	return ret
}
