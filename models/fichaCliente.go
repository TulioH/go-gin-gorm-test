package models

import "time"

type FichaCliente struct {
	Id              int `gorm:"primaryKey"`
	Nombrenegocio   string
	Estrato         string
	Nit             string
	Telefono        string
	Contacto        string
	Cargo           string
	Correo          string
	Pais            string
	Departamento    string
	Ciudadpoblacion string
	Territorio      string
	Nrocomunazona   string
	Nombrecomzon    string
	Barrio          string
	Clasificacion   string
	Tipologia       string
	Ordenruta       string
	Comentarios     string
	Direccion       string
	Creadopor       string
	Creacionfh      time.Time `gorm:"autoCreateTime"`
	Modificadopor   string
	Modificacionfh  time.Time `gorm:"autoUpdateTime"`
	Pubexte         string
	Pubinte         string
	Exhibi          string
	Lat             string
	Lng             string
	Goositioid      string
	Foto            string
	Origendatos     string
	Prodabrasivos   string
	Lineass         string
	Promcompr       string
	Uncortef        string
	Marcas          string
	Unflapd         string
	Marcasfd        string
	Promls          string
	Malijs          string
	Numverparper    string
	Canal           string
	Habeasdata      string
}

func (FichaCliente) TableName() string {
	return "fichacliente"
}
