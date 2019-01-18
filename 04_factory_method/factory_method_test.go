package factory_method

import "testing"

func TestFactoryMethod(t *testing.T) {
	bf := BulbFactory{}
	bulb := bf.Create()
	bulb.UseFor()

	tf := TubeFactory{}
	tube := tf.Create()
	tube.UseFor()
}
