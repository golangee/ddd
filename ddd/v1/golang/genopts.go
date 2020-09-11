package golang

import (
	"github.com/golangee/architecture/ddd/v1"
	"github.com/golangee/architecture/ddd/v1/internal/text"
	"github.com/golangee/src"
	"strconv"
	"strings"
)

// generateSetDefault appends a method to reset each struct value to either their natural default
// or to its configured default.
func generateSetDefault(name string, rec *src.TypeBuilder, origin *ddd.StructSpec) error {
	fun := src.NewFunc(name).SetPointerReceiver(true)
	comment := "...restores this instance to the default state.\n"
	rec.AddMethods(fun)
	body := src.NewBlock()
	for i, field := range rec.Fields() {
		oType := origin.Fields()[i]
		comment += " * The default value of " + oType.Name() + " is '"

		if oType.Default() != "" {
			comment += oType.Default()
			body.AddLine(fun.ReceiverName(), ".", field.Name(), " = ", oType.Default())
		} else {
			switch oType.TypeName() {
			case "string":
				body.AddLine(fun.ReceiverName(), ".", field.Name(), " = \"\"")
			case "byte":
				fallthrough
			case "int8":
				fallthrough
			case "int16":
				fallthrough
			case "int32":
				fallthrough
			case "int64":
				fallthrough
			case "float32":
				fallthrough
			case "float64":
				fallthrough
			case "time.Duration":
				body.AddLine(fun.ReceiverName(), ".", field.Name(), " = 0")
				comment += "0"
			case "bool":
				body.AddLine(fun.ReceiverName(), ".", field.Name(), " = false")
				comment += "false"
			default:
				return buildErr("TypeName", string(oType.TypeName()), "has no supported default value. You need to define a default literal by using SetDefault()", oType.Pos())
			}
		}
		comment += "'. \n"
	}

	fun.AddBody(body)
	fun.SetDoc(comment)
	return nil
}

// generateFlagsConfigure appends a method to setup the go flags package to be parsed by the according struct instance.
// The naming is <a>-<b>-<c> for the flags. On unix, camel case is discouraged, so we have only the alternatives
// of . _ or - and we decided for now, to use -.
func generateFlagsConfigure(envPrefix, name string, rec *src.TypeBuilder, origin *ddd.StructSpec) error {
	fun := src.NewFunc(name).SetPointerReceiver(true)
	rec.AddMethods(fun)
	body := src.NewBlock()
	comment := "... configures the flags to be ready to get evaluated. The default values are taken from the struct at calling time.\nAfter calling, use flag.Parse() to load the values. You can only use it once, otherwise the flag package will panic.\nThe following flags will be tied to this instance:\n"
	for i, field := range rec.Fields() {
		oType := origin.Fields()[i]
		envName := strings.ReplaceAll(strings.ToLower(envPrefix+field.Name()), ".", "-")
		fieldComment := field.Name() + " " + text.TrimComment(field.Doc())
		comment += " * " + field.Name() + " is parsed from flag '" + envName + "'\n"
		switch oType.TypeName() {
		case "string":
			body.AddLine(src.NewTypeDecl("flag.StringVar"), "(&", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(envName), ",", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(fieldComment), ")")
		case "bool":
			body.AddLine(src.NewTypeDecl("flag.BoolVar"), "(&", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(envName), ",", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(fieldComment), ")")
		case "int64":
			body.AddLine(src.NewTypeDecl("flag.Int64Var"), "(&", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(envName), ",", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(fieldComment), ")")
		case "float64":
			body.AddLine(src.NewTypeDecl("flag.Float64Var"), "(&", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(envName), ",", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(fieldComment), ")")
		case "time.Duration":
			body.AddLine(src.NewTypeDecl("flag.DurationVar"), "(&", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(envName), ",", fun.ReceiverName(), ".", field.Name(), ",", strconv.Quote(fieldComment), ")")
		default:
			return buildErr("TypeName", string(oType.TypeName()), "is not supported to be parsed as a flag", oType.Pos())
		}
	}
	fun.AddBody(body)
	fun.SetDoc(comment)
	return nil
}

// generateParseEnv appends a parse method for environment variables.
func generateParseEnv(envPrefix, name string, rec *src.TypeBuilder, origin *ddd.StructSpec) error {
	fun := src.NewFunc(name).SetPointerReceiver(true)
	fun.AddResults(src.NewParameter("", src.NewTypeDecl("error")))
	rec.AddMethods(fun)
	body := src.NewBlock()

	comment := "... tries to parse the environment variables into this instance. It will only set those values, which have been actually defined. If values cannot be parsed, an error is returned.\n"
	for i, field := range rec.Fields() {
		oType := origin.Fields()[i]
		envName := strings.ReplaceAll(strings.ToUpper(envPrefix+field.Name()), ".", "_")
		comment += " * " + field.Name() + " is parsed from flag '" + envName + "'\n"

		body.AddLine("if value,ok := ", src.NewTypeDecl("os.LookupEnv"), "(\"", envName, "\");ok{")
		var myTypeDecl *src.TypeDecl
		switch oType.TypeName() {
		case "string":
			body.AddLine(fun.ReceiverName(), ".", field.Name(), " = value")
		case "bool":
			myTypeDecl = src.NewTypeDecl("bool")
		case "int64":
			myTypeDecl = src.NewTypeDecl("int64")
		case "float64":
			myTypeDecl = src.NewTypeDecl("float64")
		case "time.Duration":
			myTypeDecl = src.NewTypeDecl("time.Duration")
		default:
			return buildErr("TypeName", string(oType.TypeName()), "is not supported to be parsed as a flag", oType.Pos())
		}

		// this is currently only if no parser is needed, just like plain string
		if myTypeDecl != nil {
			code, err := genParseStr("value", myTypeDecl)
			if err != nil {
				return err
			}
			body.AddLine("v,err := ", code, "")
			body.AddLine("if err != nil {")
			body.AddLine("return ", src.NewTypeDecl("fmt.Errorf"), "(\"cannot parse environment variable '", envName, "': %w\"", ",err)")
			body.AddLine("}")
			body.NewLine()
			body.AddLine(fun.ReceiverName(), ".", field.Name(), " = v")
		}

		body.AddLine("}")
		body.NewLine()
	}
	body.Add("return nil")

	fun.AddBody(body)
	fun.SetDoc(comment)
	return nil
}
