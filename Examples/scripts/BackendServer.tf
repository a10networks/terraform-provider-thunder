provider "vthunder" {
  address = ""
  username = ""
  password = ""
}

resource "vthunder_service_group" "sg3" {
name="SG3"
protocol=""
member_list {
		name="rs1"
		port=
	}
member_list {
		name="rs1"
		port=	
	}
}