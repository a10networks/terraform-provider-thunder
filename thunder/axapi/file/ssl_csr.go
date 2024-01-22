package file

//pseudo endpoint for create and delete SSL CSR

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi/oper"
)

type SslCsr struct {
	CertType      string
	CommonName    string
	Country       string
	Digest        string
	Division      string
	Email         string
	KeySize       string
	Locality      string
	Name          string
	Organization  string
	Secured       int
	StateProvince string
	ValidDays     int
}

func (p *SslCsr) GetId() string {
	return p.Name
}

func (p *SslCsr) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("file.SslCsr::Post")

	realInst := edpt.PkiCreateOper{}
	realInst.Inst.Bits = p.KeySize
	realInst.Inst.CertType = p.CertType
	realInst.Inst.CommonName = p.CommonName
	realInst.Inst.Country = p.Country
	realInst.Inst.Digest = p.Digest
	realInst.Inst.Division = p.Division
	realInst.Inst.Email = p.Email
	realInst.Inst.Filename = p.Name
	realInst.Inst.Locality = p.Locality
	realInst.Inst.Organization = p.Organization
	realInst.Inst.Secured = p.Secured
	realInst.Inst.StateProvince = p.StateProvince
	realInst.Inst.V3Request = 1
	realInst.Inst.ValidDays = p.ValidDays

	return realInst.Post(authToken, host, logger)
}

func (p *SslCsr) Get(authToken string, host, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.SslCsr::Get")
	err := oper.FileCsrGet(authToken, host, instId, logger)
	return err
}

func (p *SslCsr) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.SslCsr::Delet")
	obj := edpt.PkiDelete{}
	obj.Inst.Csr = instId
	obj.Inst.PrivateKey = instId //A private-key of same name was created with CSR
	return obj.Post(authToken, host, logger)
}
