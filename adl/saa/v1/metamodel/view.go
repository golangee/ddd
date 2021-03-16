package metamodel

// MVP is the passive view pattern described by Martin Fowler. It is nice to test and a variation of the classical
// MVC approach. Core concept is that the presenter does all the heavy lifting and synchronization things internally
// and just call through to the view by using an interface. All of this can only be stubs and the actual work
// must be done by the presenter alone.
// Here, by convention, the Presenter should be run in its own looper (or thread) so that it does never block the UI.
// This can be achieved by isolating delegates which can be also used to release the reference into both direction
// so that a long-running (bad behaving) presenter cannot leak the UI and cause other side effects there. To enforce
// an entire async behavior, a Presenter method is not allowed to provide return values. Also any parameter must
// be a value type or a DTO. Some model recommendations are to use an async (domain) model, but I certainly don't want
// to burden that into the domain or use-case level.
//
// Injection logic
//
// Due to the mutual dependent instances, the logic should be as follows:
//   viewCtrl := newPostViewDelegator()      // post into UI thread, implements View.Control interface
//   presCtrl := newPostPresenterDelegator() // post into presenter thread, implements Presenter.Control interface
//
//   view := newView(presCtrl)               // just the view, implements View.Control interface
//   presenter := newPresenter(viewCtrl)     // just the presenter, implements Presenter.Control interface
//
//   viewCtrl.setDelegate(view)
//   presCtrl.setDelegate(presenter)
//
//   presenter.onAttach()
//   view.onAttach()
type MVP struct {
	// Name of the MVP group.
	Name Identifier

	// Comment in various translations for this MPV group.
	Comment Text

	// ViewCtrl is no real view but just a method contract to update call through. Later this is either a testing
	// mock or the actual view. An implicit dependency is always Example:
	//   + SetFirstName()
	//   + SetEntireViewModel()
	//   + ShowProgressBar()
	//   + HideProgressBar()
	View struct {
		// Dependencies are interfaces which are required by the factory/constructor stub creation.
		Dependencies []*Interface

		// Flags is actually also a dependency, imagine feature flags. This is not a domain model.
		Flags *Flags

		// Control defines the Interface contracts.
		Control *Interface
	}

	// PresenterCtrl isolates the actual View from the actual Presenter implementation in the same way.
	// Presenter.Control is
	// Example:
	//  + Login(user,pwd)
	//  + Logout()
	Presenter struct {
		// UseCases are dependencies of the presenter.
		// Example:
		//   + CreateSession(token, login, pwd) Session
		//   + DeleteSession(id)
		UseCases []*UseCase

		// Flags is actually also a dependency, imagine feature flags. This is not a domain model.
		Flags *Flags

		// Control defines the Interface contracts.
		Control *Interface
	}

	// Models contains custom view models, which are unique in this view context. The presenter can only use
	// use-cases and is not allowed to interact with other elements.
	Models []*DTO
}
