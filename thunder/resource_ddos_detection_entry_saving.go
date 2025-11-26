package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDdosDetectionEntrySaving() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_ddos_detection_entry_saving`: Detection entries and indicators saving operations\n\n__PLACEHOLDER__",
		CreateContext: resourceDdosDetectionEntrySavingCreate,
		UpdateContext: resourceDdosDetectionEntrySavingUpdate,
		ReadContext:   resourceDdosDetectionEntrySavingRead,
		DeleteContext: resourceDdosDetectionEntrySavingDelete,

		Schema: map[string]*schema.Schema{
			"clear_saved_data": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Clear all saved network-object-based detection entries and learned indicators",
			},
			"manual_restore": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Manually restore network-object-based detection entries and learned indicators",
			},
			"manual_save": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Manually save network-object-based detection entries and learned indicators",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceDdosDetectionEntrySavingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionEntrySavingCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionEntrySaving(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionEntrySavingRead(ctx, d, meta)
	}
	return diags
}

func resourceDdosDetectionEntrySavingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionEntrySavingUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionEntrySaving(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceDdosDetectionEntrySavingRead(ctx, d, meta)
	}
	return diags
}
func resourceDdosDetectionEntrySavingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionEntrySavingDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionEntrySaving(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceDdosDetectionEntrySavingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceDdosDetectionEntrySavingRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointDdosDetectionEntrySaving(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointDdosDetectionEntrySaving(d *schema.ResourceData) edpt.DdosDetectionEntrySaving {
	var ret edpt.DdosDetectionEntrySaving
	ret.Inst.ClearSavedData = d.Get("clear_saved_data").(int)
	ret.Inst.ManualRestore = d.Get("manual_restore").(int)
	ret.Inst.ManualSave = d.Get("manual_save").(int)
	//omit uuid
	return ret
}
