package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRunningConfig() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_running_config`: Configure the behaviour of show running config\n\n__PLACEHOLDER__",
		CreateContext: resourceRunningConfigCreate,
		UpdateContext: resourceRunningConfigUpdate,
		ReadContext:   resourceRunningConfigRead,
		DeleteContext: resourceRunningConfigDelete,

		Schema: map[string]*schema.Schema{
			"aflex": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Show aFleX scripts",
			},
			"class_list": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Show class-list files",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceRunningConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRunningConfigCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRunningConfig(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRunningConfigRead(ctx, d, meta)
	}
	return diags
}

func resourceRunningConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRunningConfigUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRunningConfig(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceRunningConfigRead(ctx, d, meta)
	}
	return diags
}
func resourceRunningConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRunningConfigDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRunningConfig(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceRunningConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceRunningConfigRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointRunningConfig(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointRunningConfig(d *schema.ResourceData) edpt.RunningConfig {
	var ret edpt.RunningConfig
	ret.Inst.Aflex = d.Get("aflex").(int)
	ret.Inst.ClassList = d.Get("class_list").(int)
	//omit uuid
	return ret
}
