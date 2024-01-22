package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfigureSync() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_configure_sync`: Sync operation\n\n__PLACEHOLDER__",
		CreateContext: resourceConfigureSyncCreate,
		UpdateContext: resourceConfigureSyncUpdate,
		ReadContext:   resourceConfigureSyncRead,
		DeleteContext: resourceConfigureSyncDelete,

		Schema: map[string]*schema.Schema{
			"address": {
				Type: schema.TypeString, Optional: true, Description: "Specify the destination ip address to sync",
			},
			"all_partitions": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "All partition configurations",
			},
			"auto_authentication": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Authenticate with local username and password",
			},
			"partition_name": {
				Type: schema.TypeString, Optional: true, Description: "Partition name",
			},
			"private_key": {
				Type: schema.TypeString, Optional: true, Description: "Use private key for authentication",
			},
			"pwd": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"shared": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Shared partition",
			},
			"type": {
				Type: schema.TypeString, Optional: true, Description: "'running': Sync local running to peer's running configuration; 'all': Sync local running to peer's running configuration, and local startup to peer's startup configuration;",
			},
			"usr": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
		},
	}
}
func resourceConfigureSyncCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConfigureSyncCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConfigureSync(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceConfigureSyncRead(ctx, d, meta)
	}
	return diags
}

func resourceConfigureSyncUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConfigureSyncUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConfigureSync(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceConfigureSyncRead(ctx, d, meta)
	}
	return diags
}
func resourceConfigureSyncDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConfigureSyncDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConfigureSync(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceConfigureSyncRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceConfigureSyncRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointConfigureSync(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointConfigureSync(d *schema.ResourceData) edpt.ConfigureSync {
	var ret edpt.ConfigureSync
	ret.Inst.Address = d.Get("address").(string)
	ret.Inst.AllPartitions = d.Get("all_partitions").(int)
	ret.Inst.AutoAuthentication = d.Get("auto_authentication").(int)
	ret.Inst.PartitionName = d.Get("partition_name").(string)
	ret.Inst.PrivateKey = d.Get("private_key").(string)
	ret.Inst.Pwd = d.Get("pwd").(string)
	//omit pwd_enc
	ret.Inst.Shared = d.Get("shared").(int)
	ret.Inst.Type = d.Get("type").(string)
	ret.Inst.Usr = d.Get("usr").(string)
	return ret
}
