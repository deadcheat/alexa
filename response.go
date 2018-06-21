package alexa

// ResponseEnvelope general wrapper struct
type ResponseEnvelope struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Response          `json:"response"`
}

// Response response section
type Response struct {
	OutputSpeech     *OutputSpeech `json:"outputSpeech,omitempty"`
	Card             *Card         `json:"card,omitempty"`
	Reprompt         *Reprompt     `json:"reprompt,omitempty"`
	Directives       []Directive   `json:"directives,omitempty"`
	ShouldEndSession bool          `json:"shouldEndSession"`
}

// OutputSpeech outputSpeech section
type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
	SSML string `json:"ssml,omitempty"`
}

// Card card section
type Card struct {
	Type    string `json:"type"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Text    string `json:"text,omitempty"`
	Image   *Image `json:"image,omitempty"`
}

// Image image section
type Image struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}

// Reprompt reprompt section
type Reprompt struct {
	OutputSpeech *OutputSpeech `json:"outputSpeech,omitempty"`
}

// Directive each directive element of directives section
type Directive struct {
	Type          string     `json:"type"`
	PlayBehavior  string     `json:"playBehavior,omitempty"`
	AudioItem     *AudioItem `json:"audioItem,omitempty"`
	SlotToElicit  string     `json:"slotToElicit,omitempty"`
	SlotToConfirm string     `json:"slotToConfirm,omitempty"`
	UpdatedIntent *Intent    `json:"updatedIntent,omitempty"`
}

// AudioItem audioItem section
type AudioItem struct {
	Stream `json:"stream,omitempty"`
}

// Stream stream section
type Stream struct {
	Token                string `json:"token"`
	URL                  string `json:"url"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
}
