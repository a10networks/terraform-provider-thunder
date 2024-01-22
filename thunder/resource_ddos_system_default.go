package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosSystemDefault() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_system_default`: Configure system default values for DDOS\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosSystemDefaultCreate,
		UpdateContext: resourceDdosSystemDefaultUpdate,
		ReadContext:   resourceDdosSystemDefaultRead,
		DeleteContext: resourceDdosSystemDefaultDelete,

		Schema: map[string]*schema.Schema{
			"limit_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"limit_type": {
							Type: schema.TypeString, Required: true, Description: "'dst-entry': dst-entry; 'dst-icmp': dst-icmp; 'dst-other': dst-other; 'dst-tcp': dst-tcp; 'dst-udp': dst-udp; 'src-entry': src-entry; 'src-icmp': src-icmp; 'src-other': src-other; 'src-tcp': src-tcp; 'src-udp': src-udp;",
						},
						"default_over_limit_action": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"drop": {
										Type: schema.TypeInt, Optional: true, Default: 0, Description: "Silently Drop the new connection / new packet when it exceeds limit",
									},
								},
							},
						},
						"default_pkt_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Configure Default Packet rate limit",
						},
						"default_bit_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Configure Default Kibit (kibibit / 1024-bit) rate limit",
						},
						"default_frag_pkt_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Configure Default Fragmented packet rate limit",
						},
						"default_conn_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Configure Default Connection limit",
						},
						"default_conn_rate_limit": {
							Type: schema.TypeInt, Optional: true, Description: "Configure Default Connection rate limit",
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
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosSystemDefaultCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSystemDefaultCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSystemDefault(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSystemDefaultRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosSystemDefaultUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSystemDefaultUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSystemDefault(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosSystemDefaultRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosSystemDefaultDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSystemDefaultDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSystemDefault(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosSystemDefaultRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosSystemDefaultRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosSystemDefault(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceDdosSystemDefaultLimitList(d []interface{}) []edpt.DdosSystemDefaultLimitList {

	count1 := len(d)
	ret := make([]edpt.DdosSystemDefaultLimitList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.DdosSystemDefaultLimitList
		oi.LimitType = in["limit_type"].(string)
		oi.DefaultOverLimitAction = getObjectDdosSystemDefaultLimitListDefaultOverLimitAction(in["default_over_limit_action"].([]interface{}))
		oi.DefaultPktRateLimit = in["default_pkt_rate_limit"].(int)
		oi.DefaultBitRateLimit = in["default_bit_rate_limit"].(int)
		oi.DefaultFragPktRateLimit = in["default_frag_pkt_rate_limit"].(int)
		oi.DefaultConnLimit = in["default_conn_limit"].(int)
		oi.DefaultConnRateLimit = in["default_conn_rate_limit"].(int)
		//omit uuid
		oi.UserTag = in["user_tag"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getObjectDdosSystemDefaultLimitListDefaultOverLimitAction(d []interface{}) edpt.DdosSystemDefaultLimitListDefaultOverLimitAction {

	count1 := len(d)
	var ret edpt.DdosSystemDefaultLimitListDefaultOverLimitAction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Drop = in["drop"].(int)
	}
	return ret
}

func dataToEndpointDdosSystemDefault(d *schema.ResourceData) edpt.DdosSystemDefault {
	var ret edpt.DdosSystemDefault
	ret.Inst.LimitList = getSliceDdosSystemDefaultLimitList(d.Get("limit_list").([]interface{}))
	//omit uuid
	return ret
}
