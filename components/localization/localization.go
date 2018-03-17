package localization

import (
	"fmt"
	"log"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/pt"
	ut "github.com/go-playground/universal-translator"

	validator "gopkg.in/go-playground/validator.v9"
)

func init() {

	log.Println("Localization Package")

	en := en.New()
	uni = ut.New(en, en, pt.New())

	err := uni.Import(ut.FormatJSON, "translations")
	if err != nil {
		log.Fatal(err)
	}

	err = uni.VerifyTranslations()
	if err != nil {
		log.Fatal(err)
	}

}

var uni *ut.UniversalTranslator
var validate *validator.Validate

func T(language string, key interface{}, params ...string) string {

	fmt.Printf("key = %s\n", key)

	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator(language)

	translation, _ := trans.T(key, params...)

	return translation

}
