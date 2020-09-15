// Code generated by golangee/architecture. DO NOT EDIT.

package core

// LoanServiceMock is a mock implementation of LoanService.
// ...provides stuff to loan all the things.
type LoanServiceMock struct {
	// LoanItFunc mocks the LoanIt function.
	LoanItFunc func()
}

// LoanIt loans a book.
func (m LoanServiceMock) LoanIt() {
	if m.LoanItFunc != nil {
		m.LoanItFunc()
		return
	}

	panic("mock not available: LoanIt")
}