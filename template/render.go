package template

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ihorbryk/kalyna/request"
	"github.com/ihorbryk/kalyna/response"
)

func Render(request *request.Request, templateName string, context Context) (resp response.Response) {
	defer func() {
		if r := recover(); r != nil {
			errString := fmt.Sprintf("template error: %v", r)

			log.Println(errString)

			resp = response.Html(
				errString,
				http.StatusInternalServerError,
			)
		}
	}()

	templateDirs := request.Settings.TemplatesDirs

	templatesPaths := CollectTemplates(templateDirs)
	filePath := templatesPaths[TemplateName(templateName)]
	tr := request.Settings.TemplateRenderer

	responseBody, err := tr.Render(string(filePath), context)

	if err != nil {
		log.Printf("template error: %v", err)
		return response.Html(
			fmt.Sprintf("template error: %v", err),
			http.StatusInternalServerError,
		)
	}

	return response.Html(
		responseBody,
		http.StatusOK,
	)
}
