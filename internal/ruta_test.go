package internal_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"example.com/internal"
)

var _ = Describe("Ruta", func() {
	var m_30 = NewTramo(3, 70)
	var a_4 = NewTramo(285, 120)
	var a_44 = NewTramo(117, 100)
	var gr_30 = NewTramo(13, 90)
	var madrid_granada = NewRuta([m_30,a_4,a_44,gr_30], 0.0)
})
