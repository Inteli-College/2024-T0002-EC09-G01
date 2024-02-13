package main

import (
	"reflect"
	"testing"
)

func structFieldsEqual(a, b interface{}) bool {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	if va.Kind() != reflect.Struct || vb.Kind() != reflect.Struct {
		return false
	}

	if va.Type() != vb.Type() {
		return false
	}

	numFields := va.NumField()
	for i := 0; i < numFields; i++ {
		fieldA := va.Field(i)
		fieldB := vb.Field(i)

		if !reflect.DeepEqual(fieldA.Interface(), fieldB.Interface()) {
			return false
		}
	}

	return true
}

func TestNewSensor(t *testing.T) {
	t.Run("Create new Sensor", func(t *testing.T) {
		sensor := NewSensor("Sensor1", 51.0, 0.0, 0.0, 60)
		compare := Sensor{name: "Sensor1", latitude: 51.0, longitude: 0.0, measurement: 0.0, rate: 60}

		if structFieldsEqual(sensor, compare) {
			t.Errorf("The sensor was not created successfully...")
		}
	})
}
