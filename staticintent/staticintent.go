package staticintent

import (
  "github.com/arienmalec/alexa-go"
  "github.com/agilesyndrome/go-alexa-i18n/alexai18n"
  "strings"
)

func Simple(request alexa.Request, intent_name string) alexa.Response {
  request.Body.Intent.Name = intent_name
  return Handler(request)
}

func Handler(request alexa.Request) (alexa.Response) {

        intent_name := request.Body.Intent.Name
	if ( intent_name == "" ) {
	  intent_name = "skill-name"
	}


	title := alexai18n.WorldString(request, intent_name + ".title")
        text := alexai18n.WorldString(request, intent_name + ".text")

        //Should the title be missing, and we end up with an id instead of actual text
	//then fallback to the skill-name as the title
        if ( strings.HasSuffix(title, ".title") ) {
	  title = alexai18n.WorldString(request, "skill-name")
	}

        return alexa.NewSimpleResponse(title, text)
}
