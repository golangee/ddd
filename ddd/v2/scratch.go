package ddd

type Requirement interface {
}

type iwantto interface {
	IWantTo(action string) soThatOrBy
}

type soThatOrBy interface {
	By(input string) and
	SoThat(goal string)
}

type sothat interface {
	SoThat(goal string) Requirement
}

type and interface {
	And(input string) and
	sothat
}

func AsA(role string) iwantto {
	return nil
}

func AsAn(role string) iwantto {
	return nil
}

type DeclaredFunc interface {
}

func Introduce(justifiedBy Requirement, domainType, domainMethod string) DeclaredFunc {
	return nil
}

type DomainCoreSpec interface {
}

func DomainLayer(fun DeclaredFunc, funcs ...DeclaredFunc) DomainCoreSpec {
	return nil
}

func SPI(domainType, domainMethod string) DeclaredFunc {
	return nil
}

func AppLayer(fun DeclaredFunc, funcs ...DeclaredFunc) ApplicationSpec {
	return nil
}

type PresentationSpec interface {
}

func PresentationLayer(p ...SpecificPresentation) PresentationSpec {
	return nil
}

type SpecificPresentation interface {
}

type ApplicationSpec interface {
}

func Application(name string, bc BC, bcs ...BC) *ApplicationSpec {
	return nil
}

type BC interface {
}

func BoundedContext(name string, domain DomainCoreSpec, app ApplicationSpec, presentation PresentationSpec, persistence PersistenceLayerSpec) BC {
	return nil
}

func Rest(version string) SpecificPresentation {
	return nil
}

type PersistenceLayerSpec interface {
}

func PersistenceLayer(p ...SpecificPersistence) PersistenceLayerSpec {
	return nil
}

type SpecificPersistence interface {
}

func MySQL() SpecificPersistence {
	return nil
}

type StorySpec struct {
}

func Story(story string) *StorySpec {
	return &StorySpec{}
}

type EpicSpec struct {
}

func Epic(name string, story *StorySpec, others ...*StorySpec) *EpicSpec {
	return &EpicSpec{}
}
