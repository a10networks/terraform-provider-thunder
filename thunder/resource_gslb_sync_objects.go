package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGslbSyncObjects() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_gslb_sync_objects`: GSLB dynamic sync objects\n\n__PLACEHOLDER__",
		CreateContext: resourceGslbSyncObjectsCreate,
		UpdateContext: resourceGslbSyncObjectsUpdate,
		ReadContext:   resourceGslbSyncObjectsRead,
		DeleteContext: resourceGslbSyncObjectsDelete,

		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeString, Optional: true, Default: "disable", Description: "'disable': Disable dynamic sync (default); 'enable': Enable dynamic sync;",
			},
			"object_list": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"lineage": {
							Type: schema.TypeString, Optional: true, Description: "Lineage of object e.g.: slb, slb.service-group",
						},
					},
				},
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceGslbSyncObjectsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSyncObjectsCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSyncObjects(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSyncObjectsRead(ctx, d, meta)
	}
	return diags
}

func resourceGslbSyncObjectsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSyncObjectsUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSyncObjects(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGslbSyncObjectsRead(ctx, d, meta)
	}
	return diags
}
func resourceGslbSyncObjectsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSyncObjectsDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSyncObjects(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGslbSyncObjectsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGslbSyncObjectsRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGslbSyncObjects(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectGslbSyncObjectsObjectList(d []interface{}) edpt.GslbSyncObjectsObjectList {

	count1 := len(d)
	var ret edpt.GslbSyncObjectsObjectList
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Lineage = in["lineage"].(string)
	}
	return ret
}

func dataToEndpointGslbSyncObjects(d *schema.ResourceData) edpt.GslbSyncObjects {
	var ret edpt.GslbSyncObjects
	ret.Inst.Action = d.Get("action").(string)
	ret.Inst.ObjectList = getObjectGslbSyncObjectsObjectList(d.Get("object_list").([]interface{}))
	//omit uuid
	return ret
}
