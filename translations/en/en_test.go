package en_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/vulpes-ferrilata/custom-validator/translations/en"
	"github.com/vulpes-ferrilata/custom-validator/validators"
	"go.mongodb.org/mongo-driver/bson/primitive"

	en_locales "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var _ = Describe("En", Ordered, func() {
	var v *validator.Validate
	var translator ut.Translator

	BeforeAll(func() {
		var ok bool

		v = validator.New()

		err := validators.RegisterObjectIDValidator(v)
		Expect(err).ShouldNot(HaveOccurred())

		enLocale := en_locales.New()
		universalTranslator := ut.New(enLocale, enLocale)
		translator, ok = universalTranslator.GetTranslator(enLocale.Locale())
		if !ok {
			Fail("unable to get translator")
		}

		err = en.RegisterDefaultTranslations(v, translator)
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
				validationErrs := err.(validator.ValidationErrors)
				message := validationErrs[0].Translate(translator)
				Expect(message).Should(BeEquivalentTo(" must be a valid ObjectID"))
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
				validationErrs := err.(validator.ValidationErrors)
				message := validationErrs[0].Translate(translator)
				Expect(message).Should(BeEquivalentTo("ID must be a valid ObjectID"))
			})
		})
	})
})
