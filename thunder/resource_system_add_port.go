package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemAddPort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_add_port`: Add port to Active list\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemAddPortCreate,
		UpdateContext: resourceSystemAddPortUpdate,
		ReadContext:   resourceSystemAddPortRead,
		DeleteContext: resourceSystemAddPortDelete,

		Schema: map[string]*schema.Schema{
			"port_index": {
				Type: schema.TypeInt, Optional: true, Description: "port index to be configured (Specify port index)",
			},
		},
	}
}
func resourceSystemAddPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAddPortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAddPort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemAddPortRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemAddPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAddPortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAddPort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemAddPortRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemAddPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAddPortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAddPort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemAddPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemAddPortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemAddPort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemAddPort(d *schema.ResourceData) edpt.SystemAddPort {
	var ret edpt.SystemAddPort
	ret.Inst.PortIndex = d.Get("port_index").(int)
	return ret
}
