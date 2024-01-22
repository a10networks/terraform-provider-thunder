package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSflowPollingEthernetExt() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_sflow_polling_ethernet_ext`: Configure sFlow extended counter polling on ethernet interfaces\n\n__PLACEHOLDER__",
		CreateContext: resourceSflowPollingEthernetExtCreate,
		UpdateContext: resourceSflowPollingEthernetExtUpdate,
		ReadContext:   resourceSflowPollingEthernetExtRead,
		DeleteContext: resourceSflowPollingEthernetExtDelete,

		Schema: map[string]*schema.Schema{
			"start": {
				Type: schema.TypeInt, Required: true, Description: "Ethernet interface to poll",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSflowPollingEthernetExtCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingEthernetExtCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingEthernetExt(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowPollingEthernetExtRead(ctx, d, meta)
	}
	return diags
}

func resourceSflowPollingEthernetExtUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingEthernetExtUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingEthernetExt(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSflowPollingEthernetExtRead(ctx, d, meta)
	}
	return diags
}
func resourceSflowPollingEthernetExtDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingEthernetExtDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingEthernetExt(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSflowPollingEthernetExtRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSflowPollingEthernetExtRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSflowPollingEthernetExt(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSflowPollingEthernetExt(d *schema.ResourceData) edpt.SflowPollingEthernetExt {
	var ret edpt.SflowPollingEthernetExt
	ret.Inst.Start = d.Get("start").(int)
	//omit uuid
	return ret
}
