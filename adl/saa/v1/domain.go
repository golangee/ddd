package spec

// A Domain follows the DDD definition and declares the universe of the problem space (e.g. an entire
// organisation). In contrast to classical enterprise models, one should not try to model an entire
// enterprise into a common model. Instead it must be broke down into bounded contexts which consists of
// subdomains (or subsets of them, see Vaughn Vernon in "Implementing Domain Driven Design", page 45).
type Domain struct {
	Name            Identifier        // Name identifies whatever the domain is about. A company name may be a good choice, like "worldiety".
	Comment         String            // Comment should describe why this fragment exists.
	BoundedContexts []*BoundedContext // BoundedContext represent all isolated areas of the domain.

}
