package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNtpNtpGlobal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ntp_ntp_global`: ntp global settings\n\n__PLACEHOLDER__",
		CreateContext: resourceNtpNtpGlobalCreate,
		UpdateContext: resourceNtpNtpGlobalUpdate,
		ReadContext:   resourceNtpNtpGlobalRead,
		DeleteContext: resourceNtpNtpGlobalDelete,

		Schema: map[string]*schema.Schema{
			"allow_data_ports": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceNtpNtpGlobalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNtpNtpGlobalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNtpNtpGlobal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNtpNtpGlobalRead(ctx, d, meta)
	}
	return diags
}

func resourceNtpNtpGlobalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNtpNtpGlobalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNtpNtpGlobal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceNtpNtpGlobalRead(ctx, d, meta)
	}
	return diags
}
func resourceNtpNtpGlobalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNtpNtpGlobalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNtpNtpGlobal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceNtpNtpGlobalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceNtpNtpGlobalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointNtpNtpGlobal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointNtpNtpGlobal(d *schema.ResourceData) edpt.NtpNtpGlobal {
	var ret edpt.NtpNtpGlobal
	ret.Inst.AllowDataPorts = d.Get("allow_data_ports").(int)
	//omit uuid
	return ret
}
