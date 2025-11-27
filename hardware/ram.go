package hardware

type Memoria struct {
	Marca      string
	Frequencia int
	Latencia   int
	Voltagem   float64
}

func (m Memoria) EhBdie() bool {
	if m.Latencia <= 14 {
		return true
	} else {
		return false
	}
}
