package templating

import (
	"github.com/gofiber/template/django/v3"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/template/mustache/v2"
	"github.com/livghit/linkhub/pkg/utils"
)

// The second parameter of the  new function is used to tell the engine wich ending the files have

const templatesEnding = ".html"

func MustacheEngine() *mustache.Engine {
	engine := mustache.New(utils.Views(), templatesEnding)

	return engine
}

func DjangoEngine() *django.Engine {
	engine := django.New(utils.Views(), templatesEnding)

	return engine
}

func HtmlEngine() *html.Engine {
	engine := html.New(utils.Views(), templatesEnding)

	return engine
}
