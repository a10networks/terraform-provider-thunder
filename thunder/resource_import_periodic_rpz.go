package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImportPeriodicRpz() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import_periodic_rpz`: Response Policy Zone File\n\n__PLACEHOLDER__",
		CreateContext: resourceImportPeriodicRpzCreate,
		UpdateContext: resourceImportPeriodicRpzUpdate,
		ReadContext:   resourceImportPeriodicRpzRead,
		DeleteContext: resourceImportPeriodicRpzDelete,

		Schema: map[string]*schema.Schema{
			"period": {
				Type: schema.TypeInt, Optional: true, Description: "Specify the period in second",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
			},
			"remote_file_zone_transfer": {
				Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
			},
			"rpz": {
				Type: schema.TypeString, Required: true, Description: "Response Policy Zone File",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"zone_transfer": {
				Type: schema.TypeString, Optional: true, Description: "'zone-transfer': zone-transfer;",
			},
		},
	}
}
func resourceImportPeriodicRpzCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicRpzCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicRpz(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicRpzRead(ctx, d, meta)
	}
	return diags
}

func resourceImportPeriodicRpzUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicRpzUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicRpz(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportPeriodicRpzRead(ctx, d, meta)
	}
	return diags
}
func resourceImportPeriodicRpzDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicRpzDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicRpz(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportPeriodicRpzRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportPeriodicRpzRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImportPeriodicRpz(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointImportPeriodicRpz(d *schema.ResourceData) edpt.ImportPeriodicRpz {
	var ret edpt.ImportPeriodicRpz
	ret.Inst.Period = d.Get("period").(int)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.RemoteFileZoneTransfer = d.Get("remote_file_zone_transfer").(string)
	ret.Inst.Rpz = d.Get("rpz").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	ret.Inst.ZoneTransfer = d.Get("zone_transfer").(string)
	return ret
}
