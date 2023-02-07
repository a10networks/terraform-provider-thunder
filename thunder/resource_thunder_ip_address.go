package thunder

//Thunder resource IPAddress

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIPAddress() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIPAddressCreate,
		UpdateContext: resourceIPAddressUpdate,
		ReadContext:   resourceIPAddressRead,
		DeleteContext: resourceIPAddressDelete,
		Schema: map[string]*schema.Schema{
			"ip_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ip_mask": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceIPAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating IPAddress (Inside resourceIPAddressCreate) ")
		data := dataToIPAddress(d)
		logger.Println("[INFO] received V from method data to IPAddress --")
		d.SetId("1")
		err := go_thunder.PostIPAddress(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIPAddressRead(ctx, d, meta)

	}
	return diags
}

func resourceIPAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading IPAddress (Inside resourceIPAddressRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetIPAddress(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceIPAddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIPAddressRead(ctx, d, meta)
}

func resourceIPAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIPAddressRead(ctx, d, meta)
}
func dataToIPAddress(d *schema.ResourceData) go_thunder.IPAddress {
	var vc go_thunder.IPAddress
	var c go_thunder.IPAddressInstance
	c.IPAddr = d.Get("ip_addr").(string)
	c.IPMask = d.Get("ip_mask").(string)

	vc.UUID = c
	return vc
}
