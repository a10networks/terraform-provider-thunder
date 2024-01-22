package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemTimeoutValue() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_system_timeout_value`: set the timeout to stop transferring a file\n\n__PLACEHOLDER__",
		CreateContext: resourceSystemTimeoutValueCreate,
		UpdateContext: resourceSystemTimeoutValueUpdate,
		ReadContext:   resourceSystemTimeoutValueRead,
		DeleteContext: resourceSystemTimeoutValueDelete,

		Schema: map[string]*schema.Schema{
			"ftp": {
				Type: schema.TypeInt, Optional: true, Default: 120, Description: "set timeout to stop ftp transfer in seconds, 0 is no limit",
			},
			"http": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "set timeout to stop http transfer in seconds, 0 is no limit",
			},
			"https": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "set timeout to stop https transfer in seconds, 0 is no limit",
			},
			"scp": {
				Type: schema.TypeInt, Optional: true, Default: 300, Description: "set timeout to stop scp transfer in seconds, 0 is no limit",
			},
			"sftp": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "set timeout to stop sftp transfer in seconds, 0 is no limit",
			},
			"tftp": {
				Type: schema.TypeInt, Optional: true, Default: 300, Description: "set timeout to stop tftp transfer in seconds, 0 is no limit",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSystemTimeoutValueCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTimeoutValueCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTimeoutValue(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTimeoutValueRead(ctx, d, meta)
	}
	return diags
}

func resourceSystemTimeoutValueUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTimeoutValueUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTimeoutValue(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSystemTimeoutValueRead(ctx, d, meta)
	}
	return diags
}
func resourceSystemTimeoutValueDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTimeoutValueDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTimeoutValue(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSystemTimeoutValueRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemTimeoutValueRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemTimeoutValue(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSystemTimeoutValue(d *schema.ResourceData) edpt.SystemTimeoutValue {
	var ret edpt.SystemTimeoutValue
	ret.Inst.Ftp = d.Get("ftp").(int)
	ret.Inst.Http = d.Get("http").(int)
	ret.Inst.Https = d.Get("https").(int)
	ret.Inst.Scp = d.Get("scp").(int)
	ret.Inst.Sftp = d.Get("sftp").(int)
	ret.Inst.Tftp = d.Get("tftp").(int)
	//omit uuid
	return ret
}
