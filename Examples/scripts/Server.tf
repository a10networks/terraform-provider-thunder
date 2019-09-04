provider "vthunder" {
  address = "18.229.28.140"
  username = "admin"
  password = "a10"
}

resource "vthunder_service_group" "sg3" {
name="SG3"
protocol="TCP"
member_list {
		name="rs1"
		port=43
	}
member_list {
		name="rs1"
		port=83		
	}
}