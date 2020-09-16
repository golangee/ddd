package validation

import (
	"fmt"
	"github.com/golangee/architecture/ddd/v1"
	"github.com/xwb1989/sqlparser"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
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

	for _, bc := range spec.BoundedContexts() {
		core := -10
		mysql := -9
		usecase := -8
		rest := -7

		for i, layer := range bc.Layers() {
			switch t := layer.(type) {
			case *ddd.CoreLayerSpec:
				if core >= 0 {
					return buildErr("BoundedContexts", "core", "multiple core definitions", t)
				}
				core = i
			case *ddd.UseCaseLayerSpec:
				if usecase >= 0 {
					return buildErr("BoundedContexts", "core", "multiple usecase definitions", t)
				}
				usecase = i
			case *ddd.RestLayerSpec:
				rest = i
			case *ddd.MySQLLayerSpec:
				if mysql >= 0 {
					return buildErr("BoundedContexts", "mysql", "multiple mysql definitions", t)
				}
				mysql = i
			default:
				panic("not yet implemented: " + reflect.TypeOf(t).String())
			}
		}

		if core > usecase {
			return buildErr("BoundedContexts", "core vs usecase", "core must be defined before the use case layer", bc.Layers()[core])
		}

		if rest > 0 && rest < usecase {
			return buildErr("BoundedContexts", "rest vs usecase", "usecase must be defined before the rest layer", bc.Layers()[rest])
		}

		if mysql >= 0 && mysql < core {
			return buildErr("BoundedContexts", "core vs mysql", "mysql must be defined after core layer", bc.Layers()[core])
		}

		for _, layer := range bc.Layers() {
			switch t := layer.(type) {
			case *ddd.MySQLLayerSpec:
				for _, repoSpec := range t.Repositories() {
					if err := validateSqlMigration(bc, repoSpec); err != nil {
						return err
					}
				}
			}
		}
	}

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

		if obj, ok := obj.(*ddd.MigrationSpec); ok {
			format := "2006-01-02T15:04:05"
			if _, err := time.Parse(format, obj.DateTime()); err != nil {
				return buildErr("dateTime", obj.DateTime(), err.Error(), obj)
			}

			for _, statement := range obj.RawStatements() {
				if err := validateMySQLSyntax(string(statement)); err != nil {
					return buildErr("statement", string(statement), err.Error(), obj)
				}
			}
		}

		if obj, ok := obj.(*ddd.GenFuncSpec); ok {
			stmt := string(obj.RawStatement())
			if err := validateMySQLSyntax(stmt); err != nil {
				return buildErr("statement", stmt, err.Error(), obj)
			}
		}

		return nil
	})
}

func validateMySQLSyntax(stmt string) error {
	_, err := sqlparser.Parse(stmt)
	if err != nil {
		return err
	}

	return nil
}

// validateSqlMigration checks if type references are correct and the funcs contains valid signature bits, like
// context and error.
func validateSqlMigration(bc *ddd.BoundedContextSpec, repo *ddd.RepoSpec) error {
	var ifaceSpec *ddd.InterfaceSpec
	for _, spi := range bc.SPIServices() {
		if spi.Name() == repo.InterfaceName() {
			ifaceSpec = spi
			break
		}
	}

	if ifaceSpec == nil {
		return buildErr("interfaceName", repo.InterfaceName(), "is not defined as an SPI interface in core", repo)
	}

	// every method from spec has a context and error return?
	for _, funcSpec := range ifaceSpec.Funcs() {
		if len(funcSpec.In()) == 0 {
			return buildErr("In", ifaceSpec.Name(), "method '"+funcSpec.Name()+"' must at least provide a 'context.Context' parameter", funcSpec)
		}

		if funcSpec.In()[0].TypeName() != ddd.Ctx {
			return buildErr("In", ifaceSpec.Name(), "the first parameter of method '"+funcSpec.Name()+"' must be of type 'context.Context'", funcSpec)
		}

		if len(funcSpec.Out()) == 0 {
			return buildErr("Out", ifaceSpec.Name(), "method '"+funcSpec.Name()+"' must at least provide an 'error' return parameter", funcSpec)
		}

		if funcSpec.Out()[len(funcSpec.Out())-1].TypeName() != ddd.Error {
			return buildErr("Out", ifaceSpec.Name(), "the last return parameter of method '"+funcSpec.Name()+"' must be of type 'error'", funcSpec)
		}

		if len(funcSpec.Out()) > 2 {
			return buildErr("Out", ifaceSpec.Name(), "the return parameters of method '"+funcSpec.Name()+"' must match (<[]T>, error) or (<T>, error) or (error)", funcSpec)
		}

	}

	// every method from spec in impl?
	for _, funcSpec := range ifaceSpec.Funcs() {
		implSpec := repo.ImplementationByName(funcSpec.Name())
		if implSpec == nil {
			return buildErr("interfaceName", ifaceSpec.Name(), "declared method '"+funcSpec.Name()+"' still needs a mapping", repo)
		}

		// check if all parameters have been used. Because the dev can use any imported struct and nest them deeply we cannot verify correct usage anyway,
		// but at least the go compiler will check and bail out later in the generated code, so that is not that bad at all.
		for _, specParam := range funcSpec.In() {
			if specParam.TypeName() == ddd.Ctx {
				continue
			}

			hasUsedSpecParam := false
			for _, inParam := range implSpec.Params() {
				path := strings.Split(string(inParam), ".")
				actualParam := path[0]
				if specParam.Name() == actualParam {
					hasUsedSpecParam = true
					break
				}
			}

			if !hasUsedSpecParam {
				return buildErr("Prepare", ifaceSpec.Name()+"."+funcSpec.Name(), "the core parameter '"+specParam.Name()+"' is unused", implSpec)
			}

		}

		for _, inParam := range implSpec.Params() {
			path := strings.Split(string(inParam), ".")
			actualParam := path[0]
			if funcSpec.InByName(actualParam) == nil {
				return buildErr("Prepare", ifaceSpec.Name()+"."+funcSpec.Name(), "the parameter '"+string(inParam)+"' is undeclared in core", implSpec)
			}
		}

		// check also return parameter definitions. funcSpec must be one of (code above guarantees that):
		// 1. (<[]T>, error)
		// 2. (<T>, error)
		// 3. (error)
		switch len(funcSpec.Out()) {
		case 1:
			if len(implSpec.Row()) != 0 {
				return buildErr("Row", ifaceSpec.Name()+"."+funcSpec.Name(), "core does not declare any output, so '"+string(implSpec.Row()[0])+"' superfluous", implSpec)
			}
		case 2:
			myType := funcSpec.Out()[0].TypeName()
			if strings.HasPrefix(string(myType), "[]") {
				myType = myType[2:]
			}

			// this is also to complicated:
			// - the dev may decide to fill only a subset
			// - we cannot validate that, because the struct may be nested deeply and from external dependency
			// - we cannot even validate the returned columns, because the query may use tables out of our scope
			// - the only thing we can say for sure is, that an empty row mapping is wrong
			if len(implSpec.Row()) == 0 {
				return buildErr("Row", ifaceSpec.Name()+"."+funcSpec.Name(), "core requires a mapping to '"+string(myType)+"'", implSpec)
			}
		default:
			panic("internal error")
		}

	}

	// no extra method in impl?
	for _, implSpec := range repo.Implementations() {
		if ifaceSpec.FuncByName(implSpec.Name()) == nil {
			return buildErr("interfaceName", ifaceSpec.Name(), "extra method '"+implSpec.Name()+"' is not declared in core", repo)
		}
	}

	return nil
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
