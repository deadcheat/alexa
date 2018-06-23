package alexa

import "errors"

const (
	Version = "1.0"
)

// const value of request.type
const (
	TypeLaunchRequest       = "LaunchRequest"
	TypeIntentRequest       = "IntentRequest"
	TypeSessionEndedRequest = "SessionEndedRequest"
)

// const value of request.intent.name
const (
	IntentAMAZONCancelIntent     = "AMAZON.CancelIntent"
	IntentAMAZONHelpIntent       = "AMAZON.HelpIntent"
	IntentAMAZONLoopOffIntent    = "AMAZON.LoopOffIntent"
	IntentAMAZONLoopOnIntent     = "AMAZON.LoopOnIntent"
	IntentAMAZONNextIntent       = "AMAZON.NextIntent"
	IntentAMAZONNoIntent         = "AMAZON.NoIntent"
	IntentAMAZONPauseIntent      = "AMAZON.PauseIntent"
	IntentAMAZONPreviousIntent   = "AMAZON.PreviousIntent"
	IntentAMAZONRepeatIntent     = "AMAZON.RepeatIntent"
	IntentAMAZONResumeIntent     = "AMAZON.ResumeIntent"
	IntentAMAZONShuffleOffIntent = "AMAZON.ShuffleOffIntent"
	IntentAMAZONShuffleOnIntent  = "AMAZON.ShuffleOnIntent"
	IntentAMAZONStartOverIntent  = "AMAZON.StartOverIntent"
	IntentAMAZONStopIntent       = "AMAZON.StopIntent"
	IntentAMAZONYesIntent        = "AMAZON.YesIntent"
)

// OutputSpeech type
const (
	TypePlainText = "PlainText"
	TypeSSML      = "SSML"
)

// Slot type
const (
	SlotAMAZONDate            = "AMAZON.DATE"
	SlotAMAZONDuration        = "AMAZON.DURATION"
	SlotAMAZONFourDigitNumber = "AMAZON.FOUR_DIGIT_NUMBER"
	SlotAMAZONNumber          = "AMAZON.NUMBER"
	SlotAMAZONPhoneNumber     = "AMAZON.PhoneNumber"
	SlotAMAZONTime            = "AMAZON.TIME"
	SlotAMAZONSearchQuery     = "AMAZON.SearchQuery"
	SlotAMAZONActor           = "AMAZON.Actor"
	SlotAMAZONAnimal          = "AMAZON.Animal"
	SlotAMAZONArtist          = "AMAZON.Artist"
	SlotAMAZONAuthor          = "AMAZON.Author"
	SlotAMAZONBook            = "AMAZON.Book"
	SlotAMAZONCity            = "AMAZON.City"
	SlotAMAZONColor           = "AMAZON.Color"
	SlotAMAZONCorporation     = "AMAZON.Corporation"
	SlotAMAZONCountry         = "AMAZON.Country"
	SlotAMAZONFirstName       = "AMAZON.FirstName"
	SlotAMAZONFood            = "AMAZON.Food"
	SlotAMAZONLanguage        = "AMAZON.Language"
	SlotAMAZONMovie           = "AMAZON.Movie"
	SlotAMAZONPerson          = "AMAZON.Person"
	SlotAMAZONRegion          = "AMAZON.Region"
	SlotAMAZONRoom            = "AMAZON.Room"
)

// variables
var (
	EmptyResponse = ResponseEnvelope{}

	ErrNoHandler                       = errors.New("theare are no handlers assigned with specified request")
	ErrNoIntentToHandle                = errors.New("specified intent was not found")
	ErrNecessaryHandlersAreNotAssigned = errors.New("handlers for launch request and session end request are both needed")
)
