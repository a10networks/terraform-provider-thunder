package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosEventFilterL4Type() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_event_filter_l4_type`: DDOS L4 type event\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosEventFilterL4TypeCreate,
		UpdateContext: resourceDdosEventFilterL4TypeUpdate,
		ReadContext:   resourceDdosEventFilterL4TypeRead,
		DeleteContext: resourceDdosEventFilterL4TypeDelete,

		Schema: map[string]*schema.Schema{
			"out_of_seq": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "TCP out-of-seq pkts",
			},
			"protocol": {
				Type: schema.TypeString, Required: true, Description: "'tcp': tcp; 'udp': udp;",
			},
			"retrans_syn_cfg": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retrans_syn": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "TCP SYN retransmission",
						},
						"retrans_syn_exceed": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "TCP SYN retransmission exceed",
						},
					},
				},
			},
			"tcp_auth": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tcp_auth_init": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet that inits syn-auth/action-on-ack",
						},
						"tcp_auth_pass": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet that passes syn-auth/action-on-ack",
						},
						"tcp_auth_fail": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet that fails syn-auth/action-on-ack",
						},
					},
				},
			},
			"udp_auth": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"udp_auth_init": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet that inits spoof-detect",
						},
						"udp_auth_pass": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Packet that passes spoof-detect",
						},
					},
				},
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zero_window": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "TCP zero window pkts",
			},
			"filter_name": {
				Type: schema.TypeString, Required: true, Description: "FilterName",
			},
		},
	}
}
func resourceDdosEventFilterL4TypeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosEventFilterL4TypeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosEventFilterL4Type(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosEventFilterL4TypeRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosEventFilterL4TypeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosEventFilterL4TypeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosEventFilterL4Type(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosEventFilterL4TypeRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosEventFilterL4TypeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosEventFilterL4TypeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosEventFilterL4Type(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosEventFilterL4TypeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosEventFilterL4TypeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosEventFilterL4Type(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosEventFilterL4TypeRetransSynCfg(d []interface{}) edpt.DdosEventFilterL4TypeRetransSynCfg {

	count1 := len(d)
	var ret edpt.DdosEventFilterL4TypeRetransSynCfg
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.RetransSyn = in["retrans_syn"].(int)
		ret.RetransSynExceed = in["retrans_syn_exceed"].(int)
	}
	return ret
}

func getObjectDdosEventFilterL4TypeTcpAuth(d []interface{}) edpt.DdosEventFilterL4TypeTcpAuth {

	count1 := len(d)
	var ret edpt.DdosEventFilterL4TypeTcpAuth
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.TcpAuthInit = in["tcp_auth_init"].(int)
		ret.TcpAuthPass = in["tcp_auth_pass"].(int)
		ret.TcpAuthFail = in["tcp_auth_fail"].(int)
	}
	return ret
}

func getObjectDdosEventFilterL4TypeUdpAuth(d []interface{}) edpt.DdosEventFilterL4TypeUdpAuth {

	count1 := len(d)
	var ret edpt.DdosEventFilterL4TypeUdpAuth
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.UdpAuthInit = in["udp_auth_init"].(int)
		ret.UdpAuthPass = in["udp_auth_pass"].(int)
	}
	return ret
}

func dataToEndpointDdosEventFilterL4Type(d *schema.ResourceData) edpt.DdosEventFilterL4Type {
	var ret edpt.DdosEventFilterL4Type
	ret.Inst.OutOfSeq = d.Get("out_of_seq").(int)
	ret.Inst.Protocol = d.Get("protocol").(string)
	ret.Inst.RetransSynCfg = getObjectDdosEventFilterL4TypeRetransSynCfg(d.Get("retrans_syn_cfg").([]interface{}))
	ret.Inst.TcpAuth = getObjectDdosEventFilterL4TypeTcpAuth(d.Get("tcp_auth").([]interface{}))
	ret.Inst.UdpAuth = getObjectDdosEventFilterL4TypeUdpAuth(d.Get("udp_auth").([]interface{}))
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZeroWindow = d.Get("zero_window").(int)
	ret.Inst.FilterName = d.Get("filter_name").(string)
	return ret
}
