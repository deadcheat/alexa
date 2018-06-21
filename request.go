package alexa

// RequestEnvelope general struct for alexa http request
type RequestEnvelope struct {
	Version string `json:"version"`
	Session `json:"session"`
	Context `json:"context"`
	Request `json:"request"`
}

// Session session section of request
type Session struct {
	New         bool   `json:"new"`
	SessionID   string `json:"sessionId"`
	Application `json:"application"`
	Attributes  map[string]interface{} `json:"attributes"`
	User        `json:"user"`
}

// Context context section of request
type Context struct {
	System      `json:"System"`
	AudioPlayer `json:"AudioPlayer"`
}

// Application application section of request
type Application struct {
	ApplicationID string `json:"applicationId"`
}

// User user section of request
type User struct {
	UserID      string `json:"userId"`
	Permissions `json:"permissions"`
	AccessToken string `json:"accessToken"`
}

// Request request section of request
type Request struct {
	Type        string  `json:"type"`
	RequestID   string  `json:"requestId"`
	Timestamp   string  `json:"timestamp"`
	DialogState string  `json:"dialogState"`
	Locale      string  `json:"locale"`
	Intent      *Intent `json:"intent,omitempty"`
}

// Intent intent section of request
type Intent struct {
	Name               string          `json:"name,omitempty"`
	ConfirmationStatus string          `json:"confirmationStatus,omitempty"`
	Slots              map[string]Slot `json:"slots,omitempty"`
}

// Slot slot section of request
type Slot struct {
	Name               string `json:"name,omitempty"`
	Value              string `json:"value,omitempty"`
	ConfirmationStatus string `json:"confirmationStatus,omitempty"`
	Resolutions        `json:"resolutions,omitempty"`
}

// Resolutions resolutions section of request
type Resolutions struct {
	ResolutionsPerAuthority []Resolution `json:"resolutionsPerAuthority,omitempty"`
}

// Resolution each element of resolutions
type Resolution struct {
	Authority string `json:"authority,omitempty"`
	Status    `json:"status,omitempty"`
	Values    []ValueWrapper `json:"values,omitempty"`
}

// Status status section of resolution
type Status struct {
	Code string `json:"code,omitempty"`
}

// ValueWrapper wrapped value section of resolution
type ValueWrapper struct {
	Value `json:"value,omitempty"`
}

// Value each inner element of values
type Value struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}

// Permissions permissions section for request
type Permissions struct {
	ConsentToken string `json:"consentToken,omitempty"`
}

// System system section for request
type System struct {
	Application `json:"application,omitempty"`
	User        `json:"user,omitempty"`
	Device      `json:"device,omitempty"`
	APIEndpoint string `json:"apiEndpoint,omitempty"`
}

// Device device section for system
type Device struct {
	DeviceID            string `json:"deviceId,omitempty"`
	SupportedInterfaces `json:"supportedInterfaces,omitempty"`
}

// SupportedInterfaces supportedInterfaces section of device
type SupportedInterfaces struct {
	AudioPlayer `json:"AudioPlayer,omitempty"`
}

// AudioPlayer AudioPlayer section of supportedIntefaces
type AudioPlayer struct {
	Token                string `json:"token,omitempty"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds,omitempty"`
	PlayerActivity       string `json:"playerActivity,omitempty"`
}
