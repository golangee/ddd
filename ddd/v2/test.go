package ddd

func main() {
	// a user story can only be referenced once per layer, but may be shared vertical
	// (through an entire bounded context) and horizontal (between different bounded contexts).
	us1 := AsA("moderator").
		IWantTo("create a new game").
		By("entering a name").And("an optional description").
		SoThat("I can start inviting estimators")

	us2 := AsA("moderator").
		IWantTo("invite estimators"). // story to large: 1) create invitation URL and 2) then the invite method with other users
		By("giving them an url where they can access the game"). // wrong declaration -> passive mode turns result into parameter
		SoThat("we can start the game")

	us3 := AsAn("User").
		IWantTo("authorize").
		By("username").And("password").
		SoThat("I can prove that I'm the one I pretend to be")

	us4 := Story("As a moderator I want to create a new game by entering a name and an optional description so that I can start inviting estimators.")
	us5 := Story("As a moderator I want to invite estimators by giving them an url where they can access the game so that we can start the game.")
	epic1 := Epic("As a moderator, I want to play a game.", us4, us5)

	Application("BookLibrary",
		BoundedContext("Search",
			DomainLayer(
				Introduce(us1, "Moderator", "CreateGame"),       // + domain specific parameters
				Introduce(us2, "Moderator", "InviteEstimators"), // + domain specific parameters

				SPI("GameRepo", "FindAll"),
				SPI("ModeratorRepo", "Insert"),
			),
			AppLayer(
				Introduce(us1, "AppService", "CreateGame"),       // + only serialized primitive parameters + BC dependencies or just SPI?
				Introduce(us2, "AppService", "InviteEstimators"), // + only serialized primitive parameters + BC dependencies or just SPI?
			),

			PresentationLayer(
				Rest("v1.0.1"),
				Rest("v2.0.3"),
			),

			PersistenceLayer(
				MySQL(),
			),
		),

		BoundedContext("Loaning",
			DomainLayer(
				Introduce(us1, "User", "Rent"),
			),
			AppLayer(
				Introduce(us1, "AppService", "RentIt"),
			),
			PresentationLayer(),

			PersistenceLayer(),
		),

		BoundedContext("IdentityAndAccess",
			DomainLayer(
				Introduce(us3, "User", "authorize"),
			),
			AppLayer(
				Introduce(us3, "Service", "login"),
			),

			PresentationLayer(
				Rest("v1.0.1"),
				Rest("v2.0.3"),
			),

			PersistenceLayer(
				MySQL(),
			),
		),

	)

}

type any interface{}

type Moderator interface {
	CreateANewGame(enteringAName any, anOptionalDescription any) any // "so that" always must return something or cause a side effect
	InviteEstimators(givingThemAnUrl any)                            // <- should return URL
}
