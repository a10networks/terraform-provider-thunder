package thunder

//Thunder resource IPRouteStaticBfd

import (
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIPRouteStaticBfd() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPRouteStaticBfdCreate,
		Update: resourceIPRouteStaticBfdUpdate,
		Read:   resourceIPRouteStaticBfdRead,
		Delete: resourceIPRouteStaticBfdDelete,
		Schema: map[string]*schema.Schema{
			"local_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"nexthop_ip": {
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

func resourceIPRouteStaticBfdCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating IPRouteStaticBfd (Inside resourceIPRouteStaticBfdCreate) ")
		name := d.Get("local_ip").(string) + d.Get("nexthop_ip").(string)
		data := dataToIPRouteStaticBfd(d)
		logger.Println("[INFO] received V from method data to IPRouteStaticBfd --")
		d.SetId(name)
		go_thunder.PostIPRouteStaticBfd(client.Token, name, data, client.Host)

		return resourceIPRouteStaticBfdRead(d, meta)

	}
	return nil
}

func resourceIPRouteStaticBfdRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)
	logger.Println("[INFO] Reading IPRouteStaticBfd (Inside resourceIPRouteStaticBfdRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetIPRouteStaticBfd(client.Token, name, client.Host)
		if data == nil {
			logger.Println("[INFO] No data found " + name)
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceIPRouteStaticBfdUpdate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		logger.Println("[INFO] Modifying IPRouteStaticBfd   (Inside resourceIPRouteStaticBfdUpdate) ")
		name := d.Get("local_ip").(string) + d.Get("nexthop_ip").(string)
		data := dataToIPRouteStaticBfd(d)
		logger.Println("[INFO] received V from method data to IPRouteStaticBfd ")
		d.SetId(name)
		go_thunder.PutIPRouteStaticBfd(client.Token, name, data, client.Host)

		return resourceIPRouteStaticBfdRead(d, meta)

	}
	return nil
}

func resourceIPRouteStaticBfdDelete(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Deleting instance (Inside resourceIPRouteStaticBfdDelete) " + name)
		err := go_thunder.DeleteIPRouteStaticBfd(client.Token, name, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name, err)
			return err
		}
		d.SetId("")
		return nil
	}
	return nil
}

func dataToIPRouteStaticBfd(d *schema.ResourceData) go_thunder.RouteStaticBfd {
	var vc go_thunder.RouteStaticBfd
	var c go_thunder.IPBfd
	c.LocalIP = d.Get("local_ip").(string)
	c.NexthopIP = d.Get("nexthop_ip").(string)

	vc.LocalIP = c
	return vc
}
