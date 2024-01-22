package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtStateNextStateCaseActionTcp() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_state_next_state_case_action_tcp`: TCP header\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtStateNextStateCaseActionTcpCreate,
		UpdateContext: resourceSysUtStateNextStateCaseActionTcpUpdate,
		ReadContext:   resourceSysUtStateNextStateCaseActionTcpRead,
		DeleteContext: resourceSysUtStateNextStateCaseActionTcpDelete,

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
			"src_port": {
				Type: schema.TypeInt, Optional: true, Description: "Source port value",
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
			"case_number": {
				Type: schema.TypeString, Required: true, Description: "CaseNumber",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
			"direction": {
				Type: schema.TypeString, Required: true, Description: "Direction",
			},
		},
	}
}
func resourceSysUtStateNextStateCaseActionTcpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionTcpCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionTcp(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtStateNextStateCaseActionTcpRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtStateNextStateCaseActionTcpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionTcpUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionTcp(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtStateNextStateCaseActionTcpRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtStateNextStateCaseActionTcpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionTcpDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionTcp(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtStateNextStateCaseActionTcpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtStateNextStateCaseActionTcpRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtStateNextStateCaseActionTcp(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectSysUtStateNextStateCaseActionTcpFlags1540(d []interface{}) edpt.SysUtStateNextStateCaseActionTcpFlags1540 {

	count1 := len(d)
	var ret edpt.SysUtStateNextStateCaseActionTcpFlags1540
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

func getObjectSysUtStateNextStateCaseActionTcpOptions1541(d []interface{}) edpt.SysUtStateNextStateCaseActionTcpOptions1541 {

	count1 := len(d)
	var ret edpt.SysUtStateNextStateCaseActionTcpOptions1541
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

func dataToEndpointSysUtStateNextStateCaseActionTcp(d *schema.ResourceData) edpt.SysUtStateNextStateCaseActionTcp {
	var ret edpt.SysUtStateNextStateCaseActionTcp
	ret.Inst.AckSeqNumber = d.Get("ack_seq_number").(string)
	ret.Inst.Checksum = d.Get("checksum").(string)
	ret.Inst.DestPort = d.Get("dest_port").(int)
	ret.Inst.DestPortValue = d.Get("dest_port_value").(int)
	ret.Inst.Flags = getObjectSysUtStateNextStateCaseActionTcpFlags1540(d.Get("flags").([]interface{}))
	ret.Inst.NatPool = d.Get("nat_pool").(string)
	ret.Inst.Options = getObjectSysUtStateNextStateCaseActionTcpOptions1541(d.Get("options").([]interface{}))
	ret.Inst.SeqNumber = d.Get("seq_number").(string)
	ret.Inst.SrcPort = d.Get("src_port").(int)
	ret.Inst.Urgent = d.Get("urgent").(string)
	//omit uuid
	ret.Inst.Window = d.Get("window").(string)
	ret.Inst.CaseNumber = d.Get("case_number").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.Direction = d.Get("direction").(string)
	return ret
}
