package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemModifyPort() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_modify_port`: Change port number\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemModifyPortCreate,
		UpdateContext: resourceSystemModifyPortUpdate,
		ReadContext:   resourceSystemModifyPortRead,
		DeleteContext: resourceSystemModifyPortDelete,

		Schema: map[string]*schema.Schema{
			"port_index": {
				Type: schema.TypeInt, Optional: true, Description: "port index to be configured (Specify port index)",
			},
			"port_number": {
				Type: schema.TypeInt, Optional: true, Description: "port number to be configured (Specify port number)",
			},
		},
	}
}
func resourceSystemModifyPortCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemModifyPortCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemModifyPort(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemModifyPortRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemModifyPortUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemModifyPortUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemModifyPort(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemModifyPortRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemModifyPortDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemModifyPortDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemModifyPort(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemModifyPortRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemModifyPortRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemModifyPort(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemModifyPort(d *schema.ResourceData) edpt.SystemModifyPort {
	var ret edpt.SystemModifyPort
	ret.Inst.PortIndex = d.Get("port_index").(int)
	ret.Inst.PortNumber = d.Get("port_number").(int)
	return ret
}
