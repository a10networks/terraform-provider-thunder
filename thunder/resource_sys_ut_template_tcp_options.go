package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSysUtTemplateTcpOptions() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sys_ut_template_tcp_options`: Tcp Flags\n\n__PLACEHOLDER__",
		CreateContext: resourceSysUtTemplateTcpOptionsCreate,
		UpdateContext: resourceSysUtTemplateTcpOptionsUpdate,
		ReadContext:   resourceSysUtTemplateTcpOptionsRead,
		DeleteContext: resourceSysUtTemplateTcpOptionsDelete,

		Schema: map[string]*schema.Schema{
			"mss": {
				Type: schema.TypeInt, Optional: true, Description: "TCP MSS",
			},
			"nop": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "No Op",
			},
			"sack_type": {
				Type: schema.TypeString, Optional: true, Description: "'permitted': permitted; 'block': block;",
			},
			"time_stamp_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "adds Time Stamp to options",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"wscale": {
				Type: schema.TypeInt, Optional: true, Description: "",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "Name",
			},
		},
	}
}
func resourceSysUtTemplateTcpOptionsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateTcpOptionsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateTcpOptions(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateTcpOptionsRead(ctx, d, meta)
	}
	return diags
}

func resourceSysUtTemplateTcpOptionsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateTcpOptionsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateTcpOptions(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSysUtTemplateTcpOptionsRead(ctx, d, meta)
	}
	return diags
}
func resourceSysUtTemplateTcpOptionsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateTcpOptionsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateTcpOptions(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSysUtTemplateTcpOptionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSysUtTemplateTcpOptionsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSysUtTemplateTcpOptions(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSysUtTemplateTcpOptions(d *schema.ResourceData) edpt.SysUtTemplateTcpOptions {
	var ret edpt.SysUtTemplateTcpOptions
	ret.Inst.Mss = d.Get("mss").(int)
	ret.Inst.Nop = d.Get("nop").(int)
	ret.Inst.SackType = d.Get("sack_type").(string)
	ret.Inst.TimeStampEnable = d.Get("time_stamp_enable").(int)
	//omit uuid
	ret.Inst.Wscale = d.Get("wscale").(int)
	ret.Inst.Name = d.Get("name").(string)
	return ret
}
