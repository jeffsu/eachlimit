package eachlimit_test

import (
	"math/rand"
	"time"

	"github.com/jeffsu/eachlimit"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Eachlimit", func() {
	It("should work outside of loop", func() {
		concurrent := 4
		total := 10
		el := eachlimit.New(concurrent)
		j := 0

		for i := 0; i < total; i++ {
			if waited := el.Wait(); waited {
				j++
			}
			go func() {
				time.Sleep(time.Duration(rand.Int31n(100)) * time.Millisecond)
				el.Done()
			}()
		}

		el.WaitAll()

		Expect(j).To(Equal(total - concurrent))
	})
})
