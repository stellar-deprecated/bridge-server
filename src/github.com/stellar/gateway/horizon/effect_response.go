package horizon

// EffectsPageResponse contains page of effects returned by Horizon
type EffectsPageResponse struct {
	Embedded struct {
		Records []EffectResponse
	} `json:"_embedded"`
}

// EffectResponse contains effect data returned by Horizon
type EffectResponse struct {
	Type   string `json:"type"`
	Amount string `json:"amount"`
}
