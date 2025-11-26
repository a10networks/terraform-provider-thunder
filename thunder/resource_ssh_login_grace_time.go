package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSshLoginGraceTime() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ssh_login_grace_time`: The grace time during which a connection can exist without successful authentication\n\n__PLACEHOLDER__",
		CreateContext: resourceSshLoginGraceTimeCreate,
		UpdateContext: resourceSshLoginGraceTimeUpdate,
		ReadContext:   resourceSshLoginGraceTimeRead,
		DeleteContext: resourceSshLoginGraceTimeDelete,

		Schema: map[string]*schema.Schema{
			"grace_time": {
				Type: schema.TypeInt, Optional: true, Default: 120, Description: "specify number of seconds",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSshLoginGraceTimeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSshLoginGraceTimeCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSshLoginGraceTime(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSshLoginGraceTimeRead(ctx, d, meta)
	}
	return diags
}

func resourceSshLoginGraceTimeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSshLoginGraceTimeUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSshLoginGraceTime(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSshLoginGraceTimeRead(ctx, d, meta)
	}
	return diags
}
func resourceSshLoginGraceTimeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSshLoginGraceTimeDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSshLoginGraceTime(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSshLoginGraceTimeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSshLoginGraceTimeRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSshLoginGraceTime(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSshLoginGraceTime(d *schema.ResourceData) edpt.SshLoginGraceTime {
	var ret edpt.SshLoginGraceTime
	ret.Inst.GraceTime = d.Get("grace_time").(int)
	//omit uuid
	return ret
}
