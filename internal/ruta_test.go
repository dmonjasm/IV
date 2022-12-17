package internal_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"example.com/internal"
)

var _ = Describe("Ruta", func() {
	var m_30, a_4, a_44, gr_30  internal.Tramo
	var madrid_granada internal.Ruta

	BeforeEach(func() {
		m_30, _ = 	internal.NewTramo(7, 70)
		a_4, _ = internal.NewTramo(360, 120)
		a_44, _ = internal.NewTramo(117, 100)
		gr_30, _ = internal.NewTramo(45, 90)
		madrid_granada, _ = internal.NewRuta([]internal.Tramo{m_30, a_4, a_44, gr_30}, 0.0)
	})

	Describe("Calculating total distance", func() {
		It("should be the sum of the tramo distances", func() {
			var total_distance, err = internal.EstimatedDistance(madrid_granada)

			Expect(err).NotTo(HaveOccurred())
			Expect(total_distance).To(Equal(uint(529)))
		})
	})

	Describe("Calculating total route time", func() {
		It("should be the sum of the times of the tramos", func() {
			var total_time, err = internal.EstimatedTime(madrid_granada)

			Expect(err).NotTo(HaveOccurred())
			Expect(total_time).To(Equal(float32(4.77)))
		})
	})
})
