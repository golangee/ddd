package validation

import (
	"fmt"
	"github.com/golangee/architecture/ddd/v1"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var urlRegex = regexp.MustCompile(`/((\w+)/*|:\w+)*[^/]`)

const (
	minDescriptionLength = 5
	pKeyDescription      = "description"
	pKeyName             = "name"
	pKeyComment          = "comment"
	maxValueCitateLength = 40
	returnParamPosName   = "Return"
)

type withName interface {
	Pos() ddd.Pos
	Name() string
}

type withPos interface {
	Pos() ddd.Pos
}

type withDescription interface {
	Pos() ddd.Pos
	Description() string
}

type withComment interface {
	Pos() ddd.Pos
	Comment() string
}

type withStory interface {
	Pos() ddd.Pos
	Story() string
}

// Validate inspects the AppSpec and bails out, if something does not taste.
func Validate(spec *ddd.AppSpec) error {
	return spec.Walk(func(obj interface{}) error {
		if obj, ok := obj.(withDescription); ok {
			if err := checkDescription(obj); err != nil {
				return err
			}
		}

		if obj, ok := obj.(withName); ok {
			if err := checkName(obj); err != nil {
				return err
			}
		}

		if obj, ok := obj.(withComment); ok {
			if err := checkComment(obj); err != nil {
				return err
			}
		}

		if obj, ok := obj.(withStory); ok {
			if _, err := CheckUserStory(obj.Story()); err != nil {
				return buildErr("story", obj.Story(), err.Error(), obj)
			}
		}

		if obj, ok := obj.(*ddd.RestServerSpec); ok {
			if len(strings.TrimSpace(obj.Url())) == 0 {
				return buildErr("url", obj.Url(), "may not be empty", obj)
			}

			if !strings.HasPrefix(obj.Url(), "http") {
				return buildErr("url", obj.Url(), "must start with http:// or https://", obj)
			}

			_, err := url.Parse(obj.Url())
			if err != nil {
				return buildErr("url", obj.Url(), err.Error(), obj)
			}
		}

		if obj, ok := obj.(*ddd.HttpResourceSpec); ok {
			if err := checkHttpRes(obj); err != nil {
				return err
			}
		}

		return nil
	})
}

func checkName(d withName) error {
	v := d.Name()
	switch v {
	case "Id":
		//see https://github.com/golang/lint/issues/124
		return buildErr(pKeyName, v, "should be ID", d)
	case "Ids":
		//see https://github.com/golang/lint/issues/124
		return buildErr(pKeyName, v, "should be IDs", d)
	}

	switch t := d.(type) {
	case *ddd.BoundedContextSpec:
		if !isGoPackageName(v) {
			return buildErr(pKeyName, v, "must be a nice go package name", d)
		}

	case *ddd.ParamSpec:
		// empty return variable is ok
		if t.Pos().Name == returnParamPosName && t.Name() == "" {
			return nil
		}

		switch t.NameValidationKind() {
		case ddd.NVGoPublicIdentifier:
			if !isPublicGoIdentifier(v) {
				return buildErr(pKeyName, v, "must be a public go identifier", d)
			}
		case ddd.NVGoPrivateIdentifier:
			if !isPrivateGoIdentifier(v) {
				return buildErr(pKeyName, v, "must be a private go identifier", d)
			}
		case ddd.NVHttpHeaderParameter:
			if !isHttpHeaderOk(v) {
				return buildErr(pKeyName, v, "must be a valid http header parameter", d)
			}
		case ddd.NVHttpQueryParameter:
			if !isHttpQueryOk(v) {
				return buildErr(pKeyName, v, "must be a valid http query parameter", d)
			}
		case ddd.NVHttpPathParameter:
			if !isHttpPathOk(v) {
				return buildErr(pKeyName, v, "must be a valid http path parameter", d)
			}
		default:
			panic("not yet implemented: " + strconv.Itoa(int(t.NameValidationKind())))
		}

	default:
		if !isPublicGoIdentifier(v) {
			return buildErr(pKeyName, v, "must be a public go identifier", d)
		}

	}

	return nil
}

func checkDescription(d withDescription) error {
	v := d.Description()
	if v == "" {
		return buildErr(pKeyDescription, v, "must not be empty", d)
	}

	if len(v) < minDescriptionLength {
		return buildErr(pKeyDescription, v, "is to short", d)
	}

	if strings.Count(d.Description(), " ") < 2 {
		return buildErr(pKeyComment, v, "not enough words.", d)
	}

	if !startsUppercase(v) {
		return buildErr(pKeyDescription, v, "must start with a capital letter", d)
	}

	if !strings.HasSuffix(v, ".") {
		return buildErr(pKeyDescription, v, "must end with a dot (.)", d)
	}

	return nil
}

func checkComment(d withComment) error {
	v := d.Comment()

	if v == "" {
		return buildErr(pKeyComment, v, "must not be empty", d)
	}

	if len(v) < minDescriptionLength {
		return buildErr(pKeyComment, v, "is to short", d)
	}

	if strings.Count(d.Comment(), " ") < 2 {
		return buildErr(pKeyComment, v, "not enough words.", d)
	}

	if !strings.HasPrefix(v, "...") {
		return buildErr(pKeyComment, v, "must start with ellipsis '...'", d)
	}

	if !strings.HasSuffix(v, ".") {
		return buildErr(pKeyComment, v, "must end with a dot '.'", d)
	}

	return nil
}

type userStory struct {
	Role   string
	Goal   string
	Reason string
}

// checkUserStory validates the first sentence to be in the form of Mike Cohns user story format as shown at
// https://www.mountaingoatsoftware.com/agile/user-stories
func CheckUserStory(story string) (userStory, error) {
	usrStory := userStory{}
	storyStart := []string{"As a", "As an"}
	goalStart := []string{"I want to", "I need to", "I must to", "I have to"}
	reasonStart := []string{"so that", "because"}
	storyEnd := []string{"."}

	sentenceIdx := strings.IndexByte(story, '.')
	if sentenceIdx < 0 {
		return usrStory, fmt.Errorf("story must end with a . (dot)")
	}

	firstSentence := story[:sentenceIdx+1]

	subString := func(src string, left, right []string) (string, error) {
		leftIdx := -1
		lenLeft := -1
		for _, s := range left {
			lenLeft = len(s)
			leftIdx = strings.Index(src, s)
			if leftIdx >= 0 {
				break
			}
		}
		if leftIdx == -1 {
			return "", fmt.Errorf("expected phrase like '%s' not found", left[0])
		}

		rightIdx := -1
		for _, s := range right {
			rightIdx = strings.Index(src, s)
			if rightIdx >= 0 {
				break
			}
		}
		if rightIdx == -1 {
			return "", fmt.Errorf("expected phrase like '%s' not found", right[0])
		}

		if leftIdx > rightIdx {
			return "", fmt.Errorf("phrases likes '%s' must come after phrases like '%s'", right[0], left[0])
		}

		return trimComma(src[leftIdx+lenLeft : rightIdx]), nil
	}

	var err error
	usrStory.Role, err = subString(firstSentence, storyStart, goalStart)
	if err != nil {
		return usrStory, err
	}

	if usrStory.Role == "" {
		return usrStory, fmt.Errorf("role cannot be empty")
	}

	usrStory.Goal, err = subString(firstSentence, goalStart, reasonStart)
	if err != nil {
		return usrStory, err
	}

	if usrStory.Goal == "" {
		return usrStory, fmt.Errorf("goal cannot be empty")
	}

	usrStory.Reason, err = subString(firstSentence, reasonStart, storyEnd)
	if err != nil {
		return usrStory, err
	}

	if usrStory.Reason == "" {
		return usrStory, fmt.Errorf("reason cannot be empty")
	}

	return usrStory, nil
}

func trimComma(str string) string {
	str = strings.TrimSpace(str)

	if strings.HasPrefix(str, ",") {
		return strings.TrimSpace(str[1:])
	}

	if strings.HasSuffix(str, ",") {
		return strings.TrimSpace(str[:len(str)-1])
	}

	return strings.TrimSpace(str)
}

func checkHttpRes(res *ddd.HttpResourceSpec) error {
	if urlRegex.FindString(res.Path()) != res.Path() {
		return buildErr("path", res.Path(), "must be a valid url path", res)
	}

	return nil
}
