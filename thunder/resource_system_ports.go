package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemPorts() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_ports`: System level port configurations\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemPortsCreate,
		UpdateContext: resourceSystemPortsUpdate,
		ReadContext:   resourceSystemPortsRead,
		DeleteContext: resourceSystemPortsDelete,

		Schema: map[string]*schema.Schema{
			"link_detection_interval": {
				Type: schema.TypeInt, Optional: true, Default: 1000, Description: "Link detection interval in msecs",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemPortsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPortsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPorts(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemPortsRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemPortsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPortsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPorts(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemPortsRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemPortsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPortsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPorts(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemPortsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemPortsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemPorts(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemPorts(d *schema.ResourceData) edpt.SystemPorts {
	var ret edpt.SystemPorts
	ret.Inst.LinkDetectionInterval = d.Get("link_detection_interval").(int)
	//omit uuid
	return ret
}
