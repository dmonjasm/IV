package internal_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/dmonjasm/RouteCheck/internal"
)

var _ = Describe("Ruta", func() {
	var m_30, a_4, a_44, gr_30  internal.Tramo
	var madrid_granada_1 internal.Ruta

	var a_42, cm_42, a_4_corta internal.Tramo
	var madrid_granada_2 internal.Ruta

	var m_301, cm_4005, n_401, n_420, n_432 internal.Tramo
	var madrid_granada_3 internal.Ruta

	var different_routes []internal.Ruta


	BeforeEach(func() {
		m_30, _ = 	internal.NewTramo(7, 70)
		a_4, _ = internal.NewTramo(360, 120)
		a_44, _ = internal.NewTramo(117, 100)
		gr_30, _ = internal.NewTramo(45, 90)
		madrid_granada_1, _ = internal.NewRuta([]internal.Tramo{m_30, a_4, a_44, gr_30}, 42.83)
		
		a_42, _ = internal.NewTramo(107, 80)
		cm_42, _ = internal.NewTramo(85, 90)
		a_4_corta, _ = internal.NewTramo(165, 120)
		madrid_granada_2, _ = internal.NewRuta([]internal.Tramo{m_30, a_42, cm_42, a_44, gr_30, a_4_corta}, 5.0)

		m_301, _ = internal.NewTramo(58, 60)
		cm_4005, _ = internal.NewTramo(56, 80)
		n_401, _ = internal.NewTramo(103, 90)
		n_420, _ = internal.NewTramo(140, 90)
		n_432, _ = internal.NewTramo(130, 90)
		madrid_granada_3, _ = internal.NewRuta([]internal.Tramo{m_301, cm_4005, n_401, n_420, n_432}, 0.0)

		different_routes = []internal.Ruta{madrid_granada_1, madrid_granada_2, madrid_granada_3}
	})

	Describe("Calculating total distance", func() {
		It("should be the sum of the tramo distances", func() {
			var total_distance, err = internal.EstimatedDistance(madrid_granada_1)

			Expect(err).NotTo(HaveOccurred())
			Expect(total_distance).To(Equal(uint(529)))
		})
	})

	Describe("Calculating total route time", func() {
		It("should be the sum of the times of the tramos", func() {
			var total_time, err = internal.EstimatedTime(madrid_granada_1)

			Expect(err).NotTo(HaveOccurred())
			Expect(total_time).To(Equal(float32(4.77)))
		})
	})

	Describe("Choosing optimal route for", func() {
		Context("fastest route", func() {
			It("should be the fastest route", func(){
				var fastest_route, err = internal.FastestRoute(different_routes)

				Expect(err).NotTo(HaveOccurred())
				Expect(fastest_route).To(Equal(madrid_granada_1))
			})
		})

		Context("cheapest route", func() {
			It("should be the cheapest route", func(){
				var cheapest_route, err = internal.CheapestRoute(different_routes)

				Expect(err).NotTo(HaveOccurred())
				Expect(cheapest_route).To(Equal(madrid_granada_3))
			})
		})

		Context("best cost/time route", func() {
			It("should be the best cost/time route", func(){
				var cost_time_route, err = internal.CostTimeRoute(different_routes)

				Expect(err).NotTo(HaveOccurred())
				Expect(cost_time_route).To(Equal(madrid_granada_2))
			})
		})
	})
})
