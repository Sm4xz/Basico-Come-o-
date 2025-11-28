package hardware

type Memoria struct {
	Marca      string `json:"marca"`
	Frequencia int    `json:"freq_mhz"`
	Latencia   int    `json:"cas_latency"`
	Voltagem   float64 `json:"voltagem"`
}

func (m Memoria) EhBdie() bool {
	if m.Frequencia >= 3200 && m.Latencia <=14  {
		return true }
    if m.Frequencia >= 3600 && m.Latencia <=16 {
		return true 
	} else { 
		return false }
}
