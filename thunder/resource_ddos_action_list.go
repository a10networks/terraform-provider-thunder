package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosActionList() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_action_list`: Action List Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosActionListCreate,
		UpdateContext: resourceDdosActionListUpdate,
		ReadContext:   resourceDdosActionListRead,
		DeleteContext: resourceDdosActionListDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type: schema.TypeString, Optional: true, Description: "'drop': Drop Packet (Default); 'ignore': Continue processing the packet; 'reset': Reset the connection; 'authenticate-src': Authenticate the source IP; 'blacklist-src': Black-list the source IP; 'tunnel-encap-packet': Encapsulate packet for tunneling. encap template need to be bound;",
						},
						"blacklist_src_value": {
							Type: schema.TypeInt, Optional: true, Description: "blacklist duration in minutes",
						},
						"stateless": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "encapsulate all packests",
						},
						"scrub_packet": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "allow packets to go through other DDoS checks before sent out",
						},
					},
				},
			},
			"capture_config": {
				Type: schema.TypeString, Optional: true, Description: "capture-config name",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "DDOS action-list name",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone_template": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logging": {
							Type: schema.TypeString, Optional: true, Description: "DDOS logging zone-template",
						},
						"encap": {
							Type: schema.TypeString, Optional: true, Description: "DDOS encap template",
						},
					},
				},
			},
		},
	}
}
func resourceDdosActionListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosActionListCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosActionList(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosActionListRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosActionListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosActionListUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosActionList(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosActionListRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosActionListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosActionListDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosActionList(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosActionListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosActionListRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosActionList(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectDdosActionListAction(d []interface{}) edpt.DdosActionListAction {

	count1 := len(d)
	var ret edpt.DdosActionListAction
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Action = in["action"].(string)
		ret.BlacklistSrcValue = in["blacklist_src_value"].(int)
		ret.Stateless = in["stateless"].(int)
		ret.ScrubPacket = in["scrub_packet"].(int)
	}
	return ret
}

func getObjectDdosActionListZoneTemplate(d []interface{}) edpt.DdosActionListZoneTemplate {

	count1 := len(d)
	var ret edpt.DdosActionListZoneTemplate
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Logging = in["logging"].(string)
		ret.Encap = in["encap"].(string)
	}
	return ret
}

func dataToEndpointDdosActionList(d *schema.ResourceData) edpt.DdosActionList {
	var ret edpt.DdosActionList
	ret.Inst.Action = getObjectDdosActionListAction(d.Get("action").([]interface{}))
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.ZoneTemplate = getObjectDdosActionListZoneTemplate(d.Get("zone_template").([]interface{}))
	return ret
}
