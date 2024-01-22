package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCgnv6SctpPermitPayloadProtocolProtocolId() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_cgnv6_sctp_permit_payload_protocol_protocol_id`: Configure SCTP permitted payload protocol ID\n\n__PLACEHOLDER__",
		CreateContext: resourceCgnv6SctpPermitPayloadProtocolProtocolIdCreate,
		UpdateContext: resourceCgnv6SctpPermitPayloadProtocolProtocolIdUpdate,
		ReadContext:   resourceCgnv6SctpPermitPayloadProtocolProtocolIdRead,
		DeleteContext: resourceCgnv6SctpPermitPayloadProtocolProtocolIdDelete,

		Schema: map[string]*schema.Schema{
			"id1": {
				Type: schema.TypeInt, Required: true, Description: "Specify SCTP permitted payload protocol IDs",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceCgnv6SctpPermitPayloadProtocolProtocolIdCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpPermitPayloadProtocolProtocolIdCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpPermitPayloadProtocolProtocolId(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6SctpPermitPayloadProtocolProtocolIdRead(ctx, d, meta)
	}
	return diags
}

func resourceCgnv6SctpPermitPayloadProtocolProtocolIdUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpPermitPayloadProtocolProtocolIdUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpPermitPayloadProtocolProtocolId(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceCgnv6SctpPermitPayloadProtocolProtocolIdRead(ctx, d, meta)
	}
	return diags
}
func resourceCgnv6SctpPermitPayloadProtocolProtocolIdDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpPermitPayloadProtocolProtocolIdDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpPermitPayloadProtocolProtocolId(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceCgnv6SctpPermitPayloadProtocolProtocolIdRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceCgnv6SctpPermitPayloadProtocolProtocolIdRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointCgnv6SctpPermitPayloadProtocolProtocolId(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointCgnv6SctpPermitPayloadProtocolProtocolId(d *schema.ResourceData) edpt.Cgnv6SctpPermitPayloadProtocolProtocolId {
	var ret edpt.Cgnv6SctpPermitPayloadProtocolProtocolId
	ret.Inst.Id1 = d.Get("id1").(int)
	//omit uuid
	return ret
}
