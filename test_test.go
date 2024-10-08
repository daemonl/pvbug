package pvbug

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	"github.com/daemonl/pvbug/gen/test/v1/test_pb"
	"google.golang.org/protobuf/proto"
)

func mustFailValidate(t *testing.T, msg proto.Message) {
	t.Helper()
	pv, err := protovalidate.New()
	if err != nil {
		t.Errorf("Failed to create validator: %s", err)
	}

	err = pv.Validate(msg)
	if err == nil {
		t.Errorf("Fresh Validation should have failed")
	}
}

func mustFailWarmedUp(t *testing.T, pv *protovalidate.Validator, msg proto.Message) {
	t.Helper()
	err := pv.Validate(msg)
	if err == nil {
		t.Errorf("Warmed Up Validation should have failed")
	}
}

func invalidObject() *test_pb.F1 {
	return &test_pb.F1{
		NeedThis: "Foo",
		Field:    &test_pb.FieldWithIssue{},
	}
}

func TestLocalJ5(t *testing.T) {

	warmedUp, err := protovalidate.New()
	if err != nil {
		t.Errorf("Failed to create validator: %s", err)
	}

	t.Run("Value by itself", func(t *testing.T) {
		obj := invalidObject()

		// Both Correctly Fail Validation
		mustFailWarmedUp(t, warmedUp, obj)
		mustFailValidate(t, obj)
	})

	t.Run("Wrapped in order 1 2", func(t *testing.T) {
		obj := invalidObject()
		oneTwo := &test_pb.OneTwo{
			Field1: obj,
		}

		// Both Correctly Fail Validation
		mustFailWarmedUp(t, warmedUp, obj)
		mustFailValidate(t, oneTwo)
	})

	t.Run("Wrapped in order 2 1", func(t *testing.T) {
		obj := invalidObject()
		twoOne := &test_pb.TwoOne{
			Field1: obj,
		}

		// Correctly fails validation
		mustFailWarmedUp(t, warmedUp, obj)

		// Does not fail validation
		mustFailValidate(t, twoOne)
	})
}
