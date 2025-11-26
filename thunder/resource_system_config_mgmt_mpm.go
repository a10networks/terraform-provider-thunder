package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemConfigMgmtMpm() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_config_mgmt_mpm`: mdm worker count options\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemConfigMgmtMpmCreate,
		UpdateContext: resourceSystemConfigMgmtMpmUpdate,
		ReadContext:   resourceSystemConfigMgmtMpmRead,
		DeleteContext: resourceSystemConfigMgmtMpmDelete,

		Schema: map[string]*schema.Schema{
			"max_workers": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set max workers count. Default is 1",
			},
			"min_idle_workers": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set minimum idle workers count. Default is 1",
			},
			"start_workers": {
				Type: schema.TypeInt, Optional: true, Default: 1, Description: "Set starting workers count. Default is 1",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemConfigMgmtMpmCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemConfigMgmtMpmCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemConfigMgmtMpm(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemConfigMgmtMpmRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemConfigMgmtMpmUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemConfigMgmtMpmUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemConfigMgmtMpm(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemConfigMgmtMpmRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemConfigMgmtMpmDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemConfigMgmtMpmDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemConfigMgmtMpm(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemConfigMgmtMpmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemConfigMgmtMpmRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemConfigMgmtMpm(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemConfigMgmtMpm(d *schema.ResourceData) edpt.SystemConfigMgmtMpm {
	var ret edpt.SystemConfigMgmtMpm
	ret.Inst.MaxWorkers = d.Get("max_workers").(int)
	ret.Inst.MinIdleWorkers = d.Get("min_idle_workers").(int)
	ret.Inst.StartWorkers = d.Get("start_workers").(int)
	//omit uuid
	return ret
}
