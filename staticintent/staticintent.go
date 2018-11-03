package staticintent

import (
  "github.com/arienmalec/alexa-go"
  "github.com/agilesyndrome/go-alexa-i18n/alexai18n"
)

func Simple(request alexa.Request, intent_name string) alexa.Response {
  request.Body.Intent.Name = intent_name
  return Handler(request)
}

func Handler(request alexa.Request) (alexa.Response) {

        intent_name := request.Body.Intent.Name
	if ( intent_name == "" ) {
	  intent_name = "not-specified"
	}
	title := alexai18n.WorldString(request, intent_name + ".title")
        text := alexai18n.WorldString(request, intent_name + ".text")
        return alexa.NewSimpleResponse(title, text)
}
