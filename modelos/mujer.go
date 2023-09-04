package modelos

// La estructura de la Mujer, HEREDA los rasgos del Hombre
type Mujer struct {
	Hombre
}

// Adaptamos las funciones al puntero en memoria de la mujer
func (m *Mujer) Respirar()    { m.respirando = true }
func (m *Mujer) Comer()       { m.comiendo = true }
func (m *Mujer) Pensar()      { m.pensando = true }
func (m *Mujer) Sexo() string { return "Mujer" }
