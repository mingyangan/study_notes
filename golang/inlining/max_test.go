package main

import (
	"testing"
	//"github.com/agiledragon/gomonkey"
	"github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_F(t *testing.T) {
	Convey("Should Print b", t, func() {
		patch := gomonkey.ApplyFunc(Max, func(a, b int) int {
			return b
		})
		defer patch.Reset()
		m := F()
		So(m, ShouldEqual, 20)
	})
}
