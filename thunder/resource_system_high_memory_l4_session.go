package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemHighMemoryL4Session() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_high_memory_l4_session`: Enable/Disable high memory L4 session support\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemHighMemoryL4SessionCreate,
		UpdateContext: resourceSystemHighMemoryL4SessionUpdate,
		ReadContext:   resourceSystemHighMemoryL4SessionRead,
		DeleteContext: resourceSystemHighMemoryL4SessionDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable/Disable high memory l4 session support",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemHighMemoryL4SessionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHighMemoryL4SessionCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHighMemoryL4Session(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemHighMemoryL4SessionRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemHighMemoryL4SessionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHighMemoryL4SessionUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHighMemoryL4Session(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemHighMemoryL4SessionRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemHighMemoryL4SessionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHighMemoryL4SessionDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHighMemoryL4Session(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemHighMemoryL4SessionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHighMemoryL4SessionRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHighMemoryL4Session(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemHighMemoryL4Session(d *schema.ResourceData) edpt.SystemHighMemoryL4Session {
	var ret edpt.SystemHighMemoryL4Session
	ret.Inst.Enable = d.Get("enable").(int)
	//omit uuid
	return ret
}
