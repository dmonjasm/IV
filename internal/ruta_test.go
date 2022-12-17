package internal_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/dmonjasm/RouteCheck"
)

var _ = Describe("Ruta", func() {
	var m_30, a_4, a_44, gr_30  Tramo
	var madrid_granada Ruta

	BeforeEach(func() {
		m_30, _ = NewTramo(7, 70)
		a_4, _ = NewTramo(360, 120)
		a_44, _ = NewTramo(117, 100)
		gr_30, _ = NewTramo(45, 90)
		madrid_granada, _ = NewRuta([4]Tramo{m_30, a_4, a_44, gr_30}, 0.0)
	})

	Describe("Calculating total distance", func() {
		It("should be the sum of the tramo distances", func() {
			var total_distance, err = madrid_granada.EstimatedDistance()

			Expect(err).To(NotToHaveOcurred())
			Expect(total_distance).To(Equal(529))
		})
	})

	Describe("Calculating total route time", func() {
		It("should be the sum of the times of the tramos", func() {
			var total_time, err = madrid_granada.EstimatedTime()

			Expect(err).To(NotToHaveOcurred())
			Expect(total_time).To(Equal(4.77))
		})
	})
})
