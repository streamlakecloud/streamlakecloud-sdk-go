package encrypt

type ParametersWithIV struct {
	IV         []byte
	Parameters KeyParameter
}

func NewParametersWithIV(parameters KeyParameter, iv []byte) *ParametersWithIV {
	return NewParametersWithIVWithOffset(parameters, iv, 0, len(iv))
}

func NewParametersWithIVWithOffset(parameters KeyParameter, iv []byte, ivOff, ivLen int) *ParametersWithIV {
	ivCopy := make([]byte, ivLen)
	copy(ivCopy, iv[ivOff:ivOff+ivLen])
	return &ParametersWithIV{
		IV:         ivCopy,
		Parameters: parameters,
	}
}

func (pwiv *ParametersWithIV) GetIV() []byte {
	return pwiv.IV
}

func (pwiv *ParametersWithIV) GetParameters() KeyParameter {
	return pwiv.Parameters
}
