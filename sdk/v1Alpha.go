package sdk

// V1Alpha defines an interface containing callback methods for all the specs
// We define the methods as individual interfaces as we want to enable the user to
// implement only the callbacks they need
type V1Alpha interface {
	v1AlphaAccess
	v1AlphaSplit
	v1AlphaSpecs
}

// v1Alpha2Impl is a concrete implementation of the V1Alpha2 interface
type v1AlphaImpl struct {
	userV1alpha interface{}
}

// RegisterV1Alpha2 registers user defined callback functions to the api
// This is a loose registration rather than a direct interface to allow optional
// implementation of the interface methods
func (a *v1AlphaImpl) RegisterV1Alpha(i interface{}) {
	a.userV1alpha = i
}
