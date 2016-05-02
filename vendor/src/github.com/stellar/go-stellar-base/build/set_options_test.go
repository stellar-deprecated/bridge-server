package build

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stellar/go-stellar-base"
	"github.com/stellar/go-stellar-base/xdr"
)

var _ = Describe("SetOptionsBuilder Mutators", func() {

	var (
		subject SetOptionsBuilder
		mut     interface{}

		address = "GAXEMCEXBERNSRXOEKD4JAIKVECIXQCENHEBRVSPX2TTYZPMNEDSQCNQ"
		bad     = "foo"
	)

	JustBeforeEach(func() {
		subject = SetOptionsBuilder{}
		subject.Mutate(mut)
	})

	Describe("InflationDest", func() {
		Context("using a valid stellar address", func() {
			BeforeEach(func() { mut = InflationDest(address) })

			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})

			It("sets the destination to the correct xdr.AccountId", func() {
				aid, _ := stellarbase.AddressToAccountId(address)
				Expect(subject.SO.InflationDest.MustEd25519()).To(Equal(aid.MustEd25519()))
			})
		})

		Context("using an invalid value", func() {
			BeforeEach(func() { mut = InflationDest(bad) })
			It("failed", func() { Expect(subject.Err).To(HaveOccurred()) })
		})
	})

	Describe("Signer", func() {
		Context("using a valid stellar address", func() {
			BeforeEach(func() { mut = Signer{address, 5} })

			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})

			It("sets the values", func() {
				aid, _ := stellarbase.AddressToAccountId(address)
				Expect(subject.SO.Signer.PubKey.MustEd25519()).To(Equal(aid.MustEd25519()))
				Expect(subject.SO.Signer.Weight).To(Equal(xdr.Uint32(5)))
			})
		})

		Context("using an invalid PubKey", func() {
			BeforeEach(func() { mut = Signer{bad, 5} })
			It("failed", func() { Expect(subject.Err).To(HaveOccurred()) })
		})
	})

	Describe("HomeDomain", func() {
		Context("using a valid value", func() {
			BeforeEach(func() { mut = HomeDomain("stellar.org") })

			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})

			It("sets the HomeDomain to correct value", func() {
				Expect(*subject.SO.HomeDomain).To(Equal(xdr.String32("stellar.org")))
			})
		})

		Context("value too long", func() {
			BeforeEach(func() { mut = HomeDomain("123456789012345678901234567890123") })
			It("failed", func() { Expect(subject.Err).To(HaveOccurred()) })
		})
	})

	Describe("SetFlag", func() {
		Context("using a valid account flag", func() {
			BeforeEach(func() { mut = SetFlag(1) })

			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})

			It("sets the flag to the correct value", func() {
				Expect(*subject.SO.SetFlags).To(Equal(xdr.Uint32(1)))
			})
		})

		Context("using an invalid value", func() {
			BeforeEach(func() { mut = SetFlag(3) })
			It("failed", func() { Expect(subject.Err).To(HaveOccurred()) })
		})
	})

	Describe("ClearFlag", func() {
		Context("using a valid account flag", func() {
			BeforeEach(func() { mut = ClearFlag(1) })

			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})

			It("sets the flag to the correct value", func() {
				Expect(*subject.SO.ClearFlags).To(Equal(xdr.Uint32(1)))
			})
		})

		Context("using an invalid value", func() {
			BeforeEach(func() { mut = ClearFlag(3) })
			It("failed", func() { Expect(subject.Err).To(HaveOccurred()) })
		})
	})

	Describe("MasterWeight", func() {
		Context("using a valid weight", func() {
			BeforeEach(func() { mut = MasterWeight(1) })

			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})

			It("sets the weight to the correct value", func() {
				Expect(*subject.SO.MasterWeight).To(Equal(xdr.Uint32(1)))
			})
		})
	})

	Describe("Thresholds", func() {
		Context("using a valid weight", func() {
			BeforeEach(func() {
				low := uint32(1)
				med := uint32(2)
				high := uint32(3)
				mut = Thresholds{&low, &med, &high}
			})

			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})

			It("sets the thresholds to the correct value", func() {
				Expect(*subject.SO.LowThreshold).To(Equal(xdr.Uint32(1)))
				Expect(*subject.SO.MedThreshold).To(Equal(xdr.Uint32(2)))
				Expect(*subject.SO.HighThreshold).To(Equal(xdr.Uint32(3)))
			})
		})
	})

	Describe("SourceAccount", func() {
		Context("using a valid stellar address", func() {
			BeforeEach(func() { mut = SourceAccount{address} })

			It("succeeds", func() {
				Expect(subject.Err).NotTo(HaveOccurred())
			})

			It("sets the destination to the correct xdr.AccountId", func() {
				aid, _ := stellarbase.AddressToAccountId(address)
				Expect(subject.O.SourceAccount.MustEd25519()).To(Equal(aid.MustEd25519()))
			})
		})

		Context("using an invalid value", func() {
			BeforeEach(func() { mut = SourceAccount{bad} })
			It("failed", func() { Expect(subject.Err).To(HaveOccurred()) })
		})
	})
})
