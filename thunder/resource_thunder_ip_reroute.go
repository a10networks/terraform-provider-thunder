package thunder

//Thunder resource IpReroute

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpReroute() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpRerouteCreate,
		UpdateContext: resourceIpRerouteUpdate,
		ReadContext:   resourceIpRerouteRead,
		DeleteContext: resourceIpRerouteDelete,
		Schema: map[string]*schema.Schema{
			"suppress_protocols": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connected": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ibgp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"static": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"isis": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"rip": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ebgp": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"uuid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ospf": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceIpRerouteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating IpReroute (Inside resourceIpRerouteCreate) ")

		data := dataToIpReroute(d)
		logger.Println("[INFO] received V from method data to IpReroute --")
		d.SetId("1")
		err := go_thunder.PostIpReroute(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpRerouteRead(ctx, d, meta)

	}
	return diags
}

func resourceIpRerouteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading IpReroute (Inside resourceIpRerouteRead)")

	if client.Host != "" {
		data, err := go_thunder.GetIpReroute(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceIpRerouteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpRerouteRead(ctx, d, meta)
}

func resourceIpRerouteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpRerouteRead(ctx, d, meta)
}

func dataToIpReroute(d *schema.ResourceData) go_thunder.Reroute {
	var vc go_thunder.Reroute
	var c go_thunder.RerouteInstance

	var obj1 go_thunder.SuppressProtocols
	prefix := "suppress_protocols.0."
	obj1.Ibgp = d.Get(prefix + "ibgp").(int)
	obj1.Ospf = d.Get(prefix + "ospf").(int)
	obj1.Connected = d.Get(prefix + "connected").(int)
	obj1.Rip = d.Get(prefix + "rip").(int)
	obj1.Ebgp = d.Get(prefix + "ebgp").(int)
	obj1.Isis = d.Get(prefix + "isis").(int)
	obj1.Static = d.Get(prefix + "static").(int)
	c.Static = obj1

	vc.UUID = c
	return vc
}
