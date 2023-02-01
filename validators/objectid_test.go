package validators_test

import (
	"github.com/go-playground/validator/v10"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/vulpes-ferrilata/custom-validator/validators"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ = Describe("ObjectID", Ordered, func() {
	var v *validator.Validate

	BeforeAll(func() {
		v = validator.New()
		err := validators.RegisterObjectIDValidator(v)
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("validate with variable", func() {
		var id string

		When("variable is valid", func() {
			BeforeEach(func() {
				id = primitive.NewObjectID().Hex()
			})

			It("must not occured any error", func(ctx SpecContext) {
				err := v.VarCtx(ctx, id, "objectid")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		When("variable is invalid", func() {
			BeforeEach(func() {
				id = "abc1234"
			})

			It("must return error", func(ctx SpecContext) {
				err := v.VarCtx(ctx, id, "objectid")
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Context("validate with struct", func() {
		type Model struct {
			ID string `validate:"objectid"`
		}

		var model *Model

		When("struct is valid", func() {
			BeforeEach(func() {
				model = &Model{
					ID: primitive.NewObjectID().Hex(),
				}
			})

			It("must not occured any error", func(ctx SpecContext) {
				err := v.StructCtx(ctx, model)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		When("struct is invalid", func() {
			BeforeEach(func() {
				model = &Model{
					ID: "abc1234",
				}
			})

			It("must return error", func(ctx SpecContext) {
				err := v.StructCtx(ctx, model)
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})
