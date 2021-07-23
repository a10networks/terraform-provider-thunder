package thunder

//Thunder resource IpNatPool

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpNatPool() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpNatPoolCreate,
		UpdateContext: resourceIpNatPoolUpdate,
		ReadContext:   resourceIpNatPoolRead,
		DeleteContext: resourceIpNatPoolDelete,
		Schema: map[string]*schema.Schema{
			"use_if_ip": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"start_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"port_overload": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"vrid": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"netmask": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"end_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ip_rr": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"chunk_netmask": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ethernet": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"scaleout_device_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"gateway": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"pool_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceIpNatPoolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating IpNatPool (Inside resourceIpNatPoolCreate) ")
		name1 := d.Get("pool_name").(string)
		data := dataToIpNatPool(d)
		logger.Println("[INFO] received formatted data from method data to IpNatPool --")
		d.SetId(name1)
		err := go_thunder.PostIpNatPool(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpNatPoolRead(ctx, d, meta)

	}
	return diags
}

func resourceIpNatPoolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading IpNatPool (Inside resourceIpNatPoolRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetIpNatPool(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return nil
}

func resourceIpNatPoolUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Modifying IpNatPool   (Inside resourceIpNatPoolUpdate) ")
		data := dataToIpNatPool(d)
		logger.Println("[INFO] received formatted data from method data to IpNatPool ")
		err := go_thunder.PutIpNatPool(client.Token, name1, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpNatPoolRead(ctx, d, meta)

	}
	return diags
}

func resourceIpNatPoolDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceIpNatPoolDelete) " + name1)
		err := go_thunder.DeleteIpNatPool(client.Token, name1, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToIpNatPool(d *schema.ResourceData) go_thunder.IpNatPool {
	var vc go_thunder.IpNatPool
	var c go_thunder.IpNatPoolInstance
	c.UseIfIP = d.Get("use_if_ip").(int)
	c.StartAddress = d.Get("start_address").(string)
	c.PortOverload = d.Get("port_overload").(int)
	c.Vrid = d.Get("vrid").(int)
	c.Netmask = d.Get("netmask").(string)
	c.EndAddress = d.Get("end_address").(string)
	c.IPRr = d.Get("ip_rr").(int)
	c.ChunkNetmask = d.Get("chunk_netmask").(string)
	c.Ethernet = d.Get("ethernet").(int)
	c.ScaleoutDeviceID = d.Get("scaleout_device_id").(int)
	c.Gateway = d.Get("gateway").(string)
	c.PoolName = d.Get("pool_name").(string)

	vc.UseIfIP = c
	return vc
}
