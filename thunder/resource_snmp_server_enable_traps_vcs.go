package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsVcs() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_vcs`: Enable vcs traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsVcsCreate,
		UpdateContext: resourceSnmpServerEnableTrapsVcsUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsVcsRead,
		DeleteContext: resourceSnmpServerEnableTrapsVcsDelete,

		Schema: map[string]*schema.Schema{
			"state_change": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable VCS state change trap",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsVcsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsVcsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsVcs(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsVcsRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsVcsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsVcsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsVcs(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsVcsRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsVcsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsVcsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsVcs(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsVcsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsVcsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsVcs(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerEnableTrapsVcs(d *schema.ResourceData) edpt.SnmpServerEnableTrapsVcs {
	var ret edpt.SnmpServerEnableTrapsVcs
	ret.Inst.StateChange = d.Get("state_change").(int)
	//omit uuid
	return ret
}
