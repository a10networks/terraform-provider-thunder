package vthunder

//vThunder resource IpReroute

import (
	"util"

	go_vthunder "github.com/go_vthunder/vthunder"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIpReroute() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpRerouteCreate,
		Update: resourceIpRerouteUpdate,
		Read:   resourceIpRerouteRead,
		Delete: resourceIpRerouteDelete,
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

func resourceIpRerouteCreate(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)

	if client.Host != "" {
		logger.Println("[INFO] Creating IpReroute (Inside resourceIpRerouteCreate) ")

		data := dataToIpReroute(d)
		logger.Println("[INFO] received V from method data to IpReroute --")
		d.SetId("1")
		go_vthunder.PostIpReroute(client.Token, data, client.Host)

		return resourceIpRerouteRead(d, meta)

	}
	return nil
}

func resourceIpRerouteRead(d *schema.ResourceData, meta interface{}) error {
	logger := util.GetLoggerInstance()
	client := meta.(vThunder)
	logger.Println("[INFO] Reading IpReroute (Inside resourceIpRerouteRead)")

	if client.Host != "" {
		data, err := go_vthunder.GetIpReroute(client.Token, client.Host)
		if data == nil {
			d.SetId("")
			return nil
		}
		return err
	}
	return nil
}

func resourceIpRerouteUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceIpRerouteRead(d, meta)
}

func resourceIpRerouteDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceIpRerouteRead(d, meta)
}

func dataToIpReroute(d *schema.ResourceData) go_vthunder.Reroute {
	var vc go_vthunder.Reroute
	var c go_vthunder.RerouteInstance

	var obj1 go_vthunder.SuppressProtocols
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
