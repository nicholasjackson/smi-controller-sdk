package v1alpha1

import (
	"github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha4"
	"k8s.io/apimachinery/pkg/api/resource"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// log is for logging in this package.
var traffictargetlog = logf.Log.WithName("traffictarget-resource")

/*
Our "spoke" versions need to implement the
[`Convertible`](https://pkg.go.dev/sigs.k8s.io/controller-runtime/pkg/conversion?tab=doc#Convertible)
interface.  Namely, they'll need `ConvertTo` and `ConvertFrom` methods to convert to/from
the hub version.
*/

/*
ConvertTo is expected to modify its argument to contain the converted object.
Most of the conversion is straightforward copying, except for converting our changed field.
*/
// ConvertTo converts this TrafficSplit to the Hub version (v1alpha4).
func (src *TrafficSplit) ConvertTo(dstRaw conversion.Hub) error {
	traffictargetlog.Info("ConvertTo v1alpha4 from v1alpha1")

	dst := dstRaw.(*v1alpha4.TrafficSplit)
	dst.ObjectMeta = src.ObjectMeta

	dst.TypeMeta = src.TypeMeta
	dst.APIVersion = v1alpha4.GroupVersion.Identifier()

	dst.Spec = v1alpha4.TrafficSplitSpec{
		Service: src.Spec.Service,
	}

	dst.Spec.Backends = []v1alpha4.TrafficSplitBackend{}
	for _, b := range src.Spec.Backends {
		weight := b.Weight.AsDec().UnscaledBig().Int64()

		dst.Spec.Backends = append(
			dst.Spec.Backends,
			v1alpha4.TrafficSplitBackend{
				Service: b.Service,
				Weight:  int(weight),
			},
		)

	}

	return nil
}

/*
ConvertFrom is expected to modify its receiver to contain the converted object.
Most of the conversion is straightforward copying, except for converting our changed field.
*/

// ConvertFrom converts from the Hub version (v1alpha4) to this version.
func (dst *TrafficSplit) ConvertFrom(srcRaw conversion.Hub) error {
	traffictargetlog.Info("ConvertFrom v1alpha1 to v1alpha4")

	src := srcRaw.(*v1alpha4.TrafficSplit)
	dst.ObjectMeta = src.ObjectMeta

	dst.TypeMeta = src.TypeMeta
	dst.APIVersion = GroupVersion.Identifier()

	dst.Spec = TrafficSplitSpec{
		Service: src.Spec.Service,
	}

	dst.Spec.Backends = []TrafficSplitBackend{}
	for _, b := range src.Spec.Backends {
		weight := resource.NewQuantity(int64(b.Weight), resource.DecimalSI)

		dst.Spec.Backends = append(
			dst.Spec.Backends,
			TrafficSplitBackend{
				Service: b.Service,
				Weight:  weight,
			},
		)
	}

	return nil
}
