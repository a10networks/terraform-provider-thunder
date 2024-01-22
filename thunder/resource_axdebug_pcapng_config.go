package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAxdebugPcapngConfig() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_axdebug_pcapng_config`: Configure pacpng\n\n__PLACEHOLDER__",
		CreateContext: resourceAxdebugPcapngConfigCreate,
		UpdateContext: resourceAxdebugPcapngConfigUpdate,
		ReadContext:   resourceAxdebugPcapngConfigRead,
		DeleteContext: resourceAxdebugPcapngConfigDelete,

		Schema: map[string]*schema.Schema{
			"exit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Exit from axdebug pcapng mode",
			},
			"pcapng_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable pcapng",
			},
			"ssl_key_enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable ssl key tracking",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceAxdebugPcapngConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugPcapngConfigCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugPcapngConfig(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAxdebugPcapngConfigRead(ctx, d, meta)
	}
	return diags
}

func resourceAxdebugPcapngConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugPcapngConfigUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugPcapngConfig(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceAxdebugPcapngConfigRead(ctx, d, meta)
	}
	return diags
}
func resourceAxdebugPcapngConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugPcapngConfigDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugPcapngConfig(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceAxdebugPcapngConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceAxdebugPcapngConfigRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointAxdebugPcapngConfig(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointAxdebugPcapngConfig(d *schema.ResourceData) edpt.AxdebugPcapngConfig {
	var ret edpt.AxdebugPcapngConfig
	ret.Inst.Exit = d.Get("exit").(int)
	ret.Inst.PcapngEnable = d.Get("pcapng_enable").(int)
	ret.Inst.SslKeyEnable = d.Get("ssl_key_enable").(int)
	//omit uuid
	return ret
}
