package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDebugSip() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_debug_sip`: Debug SIP\n\n__PLACEHOLDER__",
		CreateContext: resourceDebugSipCreate,
		UpdateContext: resourceDebugSipUpdate,
		ReadContext:   resourceDebugSipRead,
		DeleteContext: resourceDebugSipDelete,

		Schema: map[string]*schema.Schema{
			"ack": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Method ACK",
			},
			"bye": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SIP Method BYE",
			},
			"cancel": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SIP Method CANCEL",
			},
			"info": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SIP Method INFO",
			},
			"invite": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Method INVITE",
			},
			"message": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SIP Method MESSAGE",
			},
			"method": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Set filter with SIP method types",
			},
			"notify": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SIP Method NOTIFY",
			},
			"options": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SIP Method OPTIONS",
			},
			"prack": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SIP Method PRACK",
			},
			"publish": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SIP Method PUBLISH",
			},
			"refer": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SIP Method REFER",
			},
			"register": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SIP Method REGISTER",
			},
			"subscribe": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SIP Method SUBSCRIBE",
			},
			"update": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "SIP Method UPDATE",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDebugSipCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugSipCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugSip(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugSipRead(ctx, d, meta)
	}
	return diags
}

func resourceDebugSipUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugSipUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugSip(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDebugSipRead(ctx, d, meta)
	}
	return diags
}
func resourceDebugSipDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugSipDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugSip(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDebugSipRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDebugSipRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDebugSip(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDebugSip(d *schema.ResourceData) edpt.DebugSip {
	var ret edpt.DebugSip
	ret.Inst.Ack = d.Get("ack").(int)
	ret.Inst.Bye = d.Get("bye").(int)
	ret.Inst.Cancel = d.Get("cancel").(int)
	ret.Inst.Info = d.Get("info").(int)
	ret.Inst.Invite = d.Get("invite").(int)
	ret.Inst.Message = d.Get("message").(int)
	ret.Inst.Method = d.Get("method").(int)
	ret.Inst.Notify = d.Get("notify").(int)
	ret.Inst.Options = d.Get("options").(int)
	ret.Inst.Prack = d.Get("prack").(int)
	ret.Inst.Publish = d.Get("publish").(int)
	ret.Inst.Refer = d.Get("refer").(int)
	ret.Inst.Register = d.Get("register").(int)
	ret.Inst.Subscribe = d.Get("subscribe").(int)
	ret.Inst.Update = d.Get("update").(int)
	//omit uuid
	return ret
}
