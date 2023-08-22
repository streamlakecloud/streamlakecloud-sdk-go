package encrypt

type KeyParameter struct {
	Key []byte
}

func NewKeyParameter(key []byte) *KeyParameter {
	return &KeyParameter{
		Key: key,
	}
}

func NewKeyParameterWithOffset(key []byte, keyOff, keyLen int) *KeyParameter {
	keyCopy := make([]byte, keyLen)
	copy(keyCopy, key[keyOff:keyOff+keyLen])
	return &KeyParameter{
		Key: keyCopy,
	}
}

func (kp *KeyParameter) GetKey() []byte {
	return kp.Key
}
