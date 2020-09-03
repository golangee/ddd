package golang

import (
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/architecture/ddd/v1/validation"
	"github.com/golangee/plantuml"
	"sort"
)

func addUseCaseDiagram(md *Markdown, useCase *ddd.UseCaseSpec) {
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
	ucDiag := md.UML("use case " + useCase.Name())
	for _, a := range sortedActors {
		actor := plantuml.NewActor(a)
		actors[a].ucActor = actor
		ucDiag.Add(actor)
	}

	// iterate again and put stories
	ucRect := plantuml.NewRectangle(camelCaseToWords(useCase.Name()))
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
