package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtTemplateTcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_template_tcp`: TCP header\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtTemplateTcpCreate,
		UpdateContext: resourceSysUtTemplateTcpUpdate,
		ReadContext:   resourceSysUtTemplateTcpRead,
		DeleteContext: resourceSysUtTemplateTcpDelete,

		Schema: map[string]*schema.Schema{
			"ack_seq_number": {
				Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
			},
			"checksum": {
				Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
			},
			"dest_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Dest port",
			},
			"dest_port_value": {
				Type: schema.TypeInt, Optional: true, Description: "Dest port value",
			},
			"flags": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"syn": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Syn",
						},
						"ack": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ack",
						},
						"fin": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Fin",
						},
						"rst": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Rst",
						},
						"psh": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Psh",
						},
						"ece": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Ece",
						},
						"urg": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Urg",
						},
						"cwr": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Cwr",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"nat_pool": {
				Type: schema.TypeString, Optional: true, Description: "Nat pool port",
			},
			"options": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mss": {
							Type: schema.TypeInt, Optional: true, Description: "TCP MSS",
						},
						"wscale": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"sack_type": {
							Type: schema.TypeString, Optional: true, Description: "'permitted': permitted; 'block': block;",
						},
						"time_stamp_enable": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "adds Time Stamp to options",
						},
						"nop": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "No Op",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"seq_number": {
				Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
			},
			"src_port_range": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"src_port_start": {
							Type: schema.TypeInt, Optional: true, Description: "Source port value",
						},
						"src_port_end": {
							Type: schema.TypeInt, Optional: true, Description: "Src port end value",
						},
					},
				},
			},
			"urgent": {
				Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"window": {
				Type: schema.TypeString, Optional: true, Default: "valid", Description: "'valid': valid; 'invalid': invalid;",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSysUtTemplateTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateTcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateTcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateTcpRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtTemplateTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateTcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateTcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateTcpRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtTemplateTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateTcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateTcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtTemplateTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateTcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateTcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSysUtTemplateTcpFlags1553(d []interface{}) edpt.SysUtTemplateTcpFlags1553 {

	count1 := len(d)
	var ret edpt.SysUtTemplateTcpFlags1553
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Syn = in["syn"].(int)
		ret.Ack = in["ack"].(int)
		ret.Fin = in["fin"].(int)
		ret.Rst = in["rst"].(int)
		ret.Psh = in["psh"].(int)
		ret.Ece = in["ece"].(int)
		ret.Urg = in["urg"].(int)
		ret.Cwr = in["cwr"].(int)
		//omit uuid
	}
	return ret
}

func getObjectSysUtTemplateTcpOptions1554(d []interface{}) edpt.SysUtTemplateTcpOptions1554 {

	count1 := len(d)
	var ret edpt.SysUtTemplateTcpOptions1554
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Mss = in["mss"].(int)
		ret.Wscale = in["wscale"].(int)
		ret.SackType = in["sack_type"].(string)
		ret.TimeStampEnable = in["time_stamp_enable"].(int)
		ret.Nop = in["nop"].(int)
		//omit uuid
	}
	return ret
}

func getSliceSysUtTemplateTcpSrcPortRange(d []interface{}) []edpt.SysUtTemplateTcpSrcPortRange {

	count1 := len(d)
	ret := make([]edpt.SysUtTemplateTcpSrcPortRange, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SysUtTemplateTcpSrcPortRange
		oi.SrcPortStart = in["src_port_start"].(int)
		oi.SrcPortEnd = in["src_port_end"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSysUtTemplateTcp(d *schema.ResourceData) edpt.SysUtTemplateTcp {
	var ret edpt.SysUtTemplateTcp
	ret.Inst.AckSeqNumber = d.Get("ack_seq_number").(string)
	ret.Inst.Checksum = d.Get("checksum").(string)
	ret.Inst.DestPort = d.Get("dest_port").(int)
	ret.Inst.DestPortValue = d.Get("dest_port_value").(int)
	ret.Inst.Flags = getObjectSysUtTemplateTcpFlags1553(d.Get("flags").([]interface{}))
	ret.Inst.NatPool = d.Get("nat_pool").(string)
	ret.Inst.Options = getObjectSysUtTemplateTcpOptions1554(d.Get("options").([]interface{}))
	ret.Inst.SeqNumber = d.Get("seq_number").(string)
	ret.Inst.SrcPortRange = getSliceSysUtTemplateTcpSrcPortRange(d.Get("src_port_range").([]interface{}))
	ret.Inst.Urgent = d.Get("urgent").(string)
	//omit uuid
	ret.Inst.Window = d.Get("window").(string)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
