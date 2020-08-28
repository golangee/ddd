package validation

import (
	"github.com/golangee/architecture/ddd/v1"
	"strings"
)

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

		if !isPrivateGoIdentifier(v) {
			return buildErr(pKeyName, v, "must be a private go identifier", d)
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

	if !strings.HasPrefix(v, "...") {
		return buildErr(pKeyComment, v, "must start with ellipsis '...'", d)
	}

	if !strings.HasSuffix(v, ".") {
		return buildErr(pKeyComment, v, "must end with a dot '.'", d)
	}

	return nil
}
