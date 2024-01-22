package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWriteMemory() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_write_memory`: Write memory\n\n__PLACEHOLDER__",
		CreateContext: resourceWriteMemoryCreate,
		UpdateContext: resourceWriteMemoryUpdate,
		ReadContext:   resourceWriteMemoryRead,
		DeleteContext: resourceWriteMemoryDelete,

		Schema: map[string]*schema.Schema{
			"destination": {
				Type: schema.TypeString, Optional: true, Description: "'primary': Write to default Primary Configuration; 'secondary': Write to default Secondary Configuration; 'local': Local Configuration Profile Name;",
			},
			"partition": {
				Type: schema.TypeString, Optional: true, Description: "'all': All partition configurations; 'shared': Shared partition; 'specified': Specified partition;",
			},
			"profile": {
				Type: schema.TypeString, Optional: true, Description: "Local Configuration Profile Name",
			},
			"specified_partition": {
				Type: schema.TypeString, Optional: true, Description: "Specified partition",
			},
		},
	}
}
func resourceWriteMemoryCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWriteMemoryCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWriteMemory(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWriteMemoryRead(ctx, d, meta)
	}
	return diags
}

func resourceWriteMemoryUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWriteMemoryUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWriteMemory(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceWriteMemoryRead(ctx, d, meta)
	}
	return diags
}
func resourceWriteMemoryDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWriteMemoryDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWriteMemory(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceWriteMemoryRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceWriteMemoryRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointWriteMemory(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointWriteMemory(d *schema.ResourceData) edpt.WriteMemory {
	var ret edpt.WriteMemory
	ret.Inst.Destination = d.Get("destination").(string)
	ret.Inst.Partition = d.Get("partition").(string)
	ret.Inst.Profile = d.Get("profile").(string)
	ret.Inst.SpecifiedPartition = d.Get("specified_partition").(string)
	return ret
}
