package vi_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"

	vi_locales "github.com/go-playground/locales/vi"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/vulpes-ferrilata/custom-validator/translations/vi"
	"github.com/vulpes-ferrilata/custom-validator/validators"
)

var _ = Describe("Vi", Ordered, func() {
	var v *validator.Validate
	var translator ut.Translator

	BeforeAll(func() {
		var ok bool

		v = validator.New()

		err := validators.RegisterObjectIDValidator(v)
		Expect(err).ShouldNot(HaveOccurred())

		viLocale := vi_locales.New()
		universalTranslator := ut.New(viLocale, viLocale)
		translator, ok = universalTranslator.GetTranslator(viLocale.Locale())
		if !ok {
			Fail("unable to get translator")
		}

		err = vi.RegisterDefaultTranslations(v, translator)
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
				Expect(message).Should(BeEquivalentTo(" phải là ObjectID hợp lệ"))
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
				Expect(message).Should(BeEquivalentTo("ID phải là ObjectID hợp lệ"))
			})
		})
	})
})
