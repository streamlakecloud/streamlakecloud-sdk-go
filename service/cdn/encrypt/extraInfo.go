package encrypt

type ExtraInfo struct {
	ClientIP string
	TTL      int
}

func NewExtraInfo(builder *ExtraInfoBuilder) *ExtraInfo {
	return &ExtraInfo{
		ClientIP: builder.clientIP,
		TTL:      builder.ttl,
	}
}

func (e *ExtraInfo) GetClientIP() string {
	return e.ClientIP
}

func (e *ExtraInfo) SetClientIP(clientIP string) {
	e.ClientIP = clientIP
}

func (e *ExtraInfo) GetTTL() int {
	return e.TTL
}

func (e *ExtraInfo) SetTTL(ttl int) {
	e.TTL = ttl
}

type ExtraInfoBuilder struct {
	clientIP string
	ttl      int
}

func NewExtraInfoBuilder() *ExtraInfoBuilder {
	return &ExtraInfoBuilder{
		clientIP: "0.0.0.0",
		ttl:      24 * 3600,
	}
}

func (b *ExtraInfoBuilder) SetClientIP(clientIP string) *ExtraInfoBuilder {
	b.clientIP = clientIP
	return b
}

func (b *ExtraInfoBuilder) SetTTL(ttl int) *ExtraInfoBuilder {
	b.ttl = ttl
	return b
}

func (b *ExtraInfoBuilder) Build() *ExtraInfo {
	return NewExtraInfo(b)
}
