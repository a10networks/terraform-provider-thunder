package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAcosEventsLogdb() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_acos_events_logdb`: Configure global logging template\n\n__PLACEHOLDER__",
		CreateContext: resourceAcosEventsLogdbCreate,
		UpdateContext: resourceAcosEventsLogdbUpdate,
		ReadContext:   resourceAcosEventsLogdbRead,
		DeleteContext: resourceAcosEventsLogdbDelete,

		Schema: map[string]*schema.Schema{
			"enable_all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable logging for all widgets",
			},
			"enable_cgn": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable CGN logging",
			},
			"enable_fw": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Firewall logging",
			},
			"enable_http_forward_proxy": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable HTTP forward proxy logging",
			},
			"enable_link_cost": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable link-cost logging",
			},
			"enable_mqtt": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable MQTT logging",
			},
			"enable_smtp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SMTP logging",
			},
			"enable_ssli": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SSLi logging",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAcosEventsLogdbCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogdbCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogdb(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsLogdbRead(ctx, d, meta)
	}
	return diags
}

func resourceAcosEventsLogdbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogdbUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogdb(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAcosEventsLogdbRead(ctx, d, meta)
	}
	return diags
}
func resourceAcosEventsLogdbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogdbDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogdb(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAcosEventsLogdbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAcosEventsLogdbRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAcosEventsLogdb(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAcosEventsLogdb(d *schema.ResourceData) edpt.AcosEventsLogdb {
	var ret edpt.AcosEventsLogdb
	ret.Inst.EnableAll = d.Get("enable_all").(int)
	ret.Inst.EnableCgn = d.Get("enable_cgn").(int)
	ret.Inst.EnableFw = d.Get("enable_fw").(int)
	ret.Inst.EnableHttpForwardProxy = d.Get("enable_http_forward_proxy").(int)
	ret.Inst.EnableLinkCost = d.Get("enable_link_cost").(int)
	ret.Inst.EnableMqtt = d.Get("enable_mqtt").(int)
	ret.Inst.EnableSmtp = d.Get("enable_smtp").(int)
	ret.Inst.EnableSsli = d.Get("enable_ssli").(int)
	//omit uuid
	return ret
}
