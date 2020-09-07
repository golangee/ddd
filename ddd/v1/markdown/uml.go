package markdown

import (
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/architecture/ddd/v1/internal/text"
	"github.com/golangee/architecture/ddd/v1/validation"
	"github.com/golangee/plantuml"
	"sort"
	"strings"
)

func generateUMLForInterface(t *ddd.InterfaceSpec) *plantuml.Class {
	class := plantuml.NewInterface(t.Name())

	for _, fun := range t.Funcs() {
		pTmp := ""
		for i, p := range fun.In() {
			pTmp += p.Name() + " " + string(p.TypeName())
			if i < len(fun.In())-1 {
				pTmp += ", "
			}
		}

		rMultiple := false
		rTmp := ""
		for i, p := range fun.Out() {
			rTmp += p.Name() + " " + string(p.TypeName())
			if i < len(fun.Out())-1 {
				rTmp += ", "
				rMultiple = true
			}
		}

		myResLiteral := strings.TrimSpace(rTmp)
		if rMultiple {
			myResLiteral = "(" + myResLiteral + ")"
		}

		class.AddAttrs(plantuml.Attr{
			Visibility: plantuml.Public,
			Abstract:   true,
			Static:     false,
			Name:       fun.Name() + "(" + strings.TrimSpace(pTmp) + ")",
			Type:       myResLiteral,
		})
	}

	return class
}

func generateUMLForStruct(t *ddd.StructSpec) *plantuml.Class {
	class := plantuml.NewClass(t.Name())
	for _, field := range t.Fields() {
		class.AddAttrs(plantuml.Attr{
			Visibility: plantuml.Public,
			Abstract:   false,
			Static:     false,
			Name:       field.Name(),
			Type:       string(field.TypeName()),
		})
	}

	return class
}

func addRestAPI(md *Markdown, rest *ddd.RestLayerSpec) {
	md.H4("REST API *" + rest.Version() + "*")
	md.P(rest.Description())
	for _, resource := range rest.Resources() {
		md.H5(resource.Path())
		md.P(resource.Description())

		for _, verb := range resource.Verbs() {
			md.H6("*" + verb.Method() + "* " + resource.Path())
			md.P(verb.Description())
			tmp := "curl -v -X " + verb.Method() + " "
			tmp += text.JoinSlashes(rest.PrimaryUrl(), rest.Prefix(), resource.Path())
			tmp += "\n"
			md.Code("bash", tmp)
		}

	}
}

func addUseCaseDiagram(ucDiag *plantuml.Diagram, useCase *ddd.EpicSpec) {
	type umlStory struct {
		name    string
		usStory *plantuml.UseCase
	}

	type umlActor struct {
		name    string
		ucActor *plantuml.Actor
		stories []*umlStory
	}

	// create our model
	actors := map[string]*umlActor{}
	for _, usecase := range useCase.Stories() {
		story, err := validation.CheckUserStory(usecase.Story())
		if err != nil {
			panic("illegal state")
		}

		if _, ok := actors[story.Role]; !ok {
			actors[story.Role] = &umlActor{name: story.Role}
		}

		actor := actors[story.Role]
		actor.stories = append(actor.stories, &umlStory{
			name: story.Goal,
		})
	}

	// get deterministic order
	sortedActors := make([]string, 0, len(actors))
	for key, _ := range actors {
		sortedActors = append(sortedActors, key)
	}
	sort.Strings(sortedActors)

	// iterate again and create uml actors first

	for _, a := range sortedActors {
		actor := plantuml.NewActor(a)
		actors[a].ucActor = actor
		ucDiag.Add(actor)
	}

	// iterate again and put stories
	ucRect := plantuml.NewRectangle(text.CamelCaseToWords(useCase.Name()))
	ucDiag.Add(ucRect)
	for _, a := range sortedActors {
		for _, s := range actors[a].stories {
			story := plantuml.NewUseCase(s.name)
			s.usStory = story
			ucRect.Add(story)
		}
	}

	// iterate again and connect actors with stories
	for _, a := range sortedActors {
		actor := actors[a]
		for _, s := range actor.stories {
			ucDiag.Add(plantuml.NewPointer(actor.ucActor.Id(), s.usStory.Id()))
		}
	}

}
