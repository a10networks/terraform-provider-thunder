package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCaptureConfig() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_capture_config`: Packet Capture-Configuration\n\n__PLACEHOLDER__",
		CreateContext: resourceCaptureConfigCreate,
		UpdateContext: resourceCaptureConfigUpdate,
		ReadContext:   resourceCaptureConfigRead,
		DeleteContext: resourceCaptureConfigDelete,

		Schema: map[string]*schema.Schema{
			"count1": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify maximum packet number. (default 0 for unlimited)",
			},
			"enable": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable capture-config",
			},
			"file_history_size": {
				Type: schema.TypeInt, Optional: true, Description: "(Default) Specify pcapng file history size in MB (default 10)",
			},
			"file_size": {
				Type: schema.TypeInt, Optional: true, Description: "Specify pcapng filesize in MB (default 10)",
			},
			"filter": {
				Type: schema.TypeString, Optional: true, Description: "Filter packets to save using Berkeley Packet Filter syntax",
			},
			"length": {
				Type: schema.TypeInt, Optional: true, Description: "Packet length Bytes to capture (Default 128)",
			},
			"name": {
				Type: schema.TypeString, Required: true, Description: "capture-config name",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCaptureConfigCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCaptureConfigCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCaptureConfig(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCaptureConfigRead(ctx, d, meta)
	}
	return diags
}

func resourceCaptureConfigUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCaptureConfigUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCaptureConfig(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCaptureConfigRead(ctx, d, meta)
	}
	return diags
}
func resourceCaptureConfigDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCaptureConfigDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCaptureConfig(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCaptureConfigRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCaptureConfigRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCaptureConfig(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCaptureConfig(d *schema.ResourceData) edpt.CaptureConfig {
	var ret edpt.CaptureConfig
	ret.Inst.Count1 = d.Get("count1").(int)
	ret.Inst.Enable = d.Get("enable").(int)
	ret.Inst.FileHistorySize = d.Get("file_history_size").(int)
	ret.Inst.FileSize = d.Get("file_size").(int)
	ret.Inst.Filter = d.Get("filter").(string)
	ret.Inst.Length = d.Get("length").(int)
	ret.Inst.Name = d.Get("name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	return ret
}
