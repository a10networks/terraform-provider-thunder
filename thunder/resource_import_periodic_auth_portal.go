package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportPeriodicAuthPortal() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_periodic_auth_portal`: Portal file for http authentication\n\n__PLACEHOLDER__",
		CreateContext: resourceImportPeriodicAuthPortalCreate,
		UpdateContext: resourceImportPeriodicAuthPortalUpdate,
		ReadContext:   resourceImportPeriodicAuthPortalRead,
		DeleteContext: resourceImportPeriodicAuthPortalDelete,

		Schema: map[string]*schema.Schema{
			"auth_portal": {
				Type: schema.TypeString, Required: true, Description: "Portal file for http authentication",
			},
			"period": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the period in second",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceImportPeriodicAuthPortalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicAuthPortalCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicAuthPortal(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicAuthPortalRead(ctx, d, meta)
	}
	return diags
}

func resourceImportPeriodicAuthPortalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicAuthPortalUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicAuthPortal(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicAuthPortalRead(ctx, d, meta)
	}
	return diags
}
func resourceImportPeriodicAuthPortalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicAuthPortalDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicAuthPortal(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportPeriodicAuthPortalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicAuthPortalRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicAuthPortal(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportPeriodicAuthPortal(d *schema.ResourceData) edpt.ImportPeriodicAuthPortal {
	var ret edpt.ImportPeriodicAuthPortal
	ret.Inst.AuthPortal = d.Get("auth_portal").(string)
	ret.Inst.Period = d.Get("period").(int)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	return ret
}
