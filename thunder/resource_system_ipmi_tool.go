package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemIpmiTool() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ipmi_tool`: Command to execute in double quotes\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemIpmiToolCreate,
		UpdateContext: resourceSystemIpmiToolUpdate,
		ReadContext:   resourceSystemIpmiToolRead,
		DeleteContext: resourceSystemIpmiToolDelete,

		Schema: map[string]*schema.Schema{
			"cmd": {
				Type: schema.TypeString, Optional: true, Description: "Command to execute in double quotes",
			},
		},
	}
}
func resourceSystemIpmiToolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiToolCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmiTool(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpmiToolRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemIpmiToolUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiToolUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmiTool(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemIpmiToolRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemIpmiToolDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiToolDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmiTool(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemIpmiToolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemIpmiToolRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemIpmiTool(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemIpmiTool(d *schema.ResourceData) edpt.SystemIpmiTool {
	var ret edpt.SystemIpmiTool
	ret.Inst.Cmd = d.Get("cmd").(string)
	return ret
}
