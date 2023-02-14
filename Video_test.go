package Aaaa

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Name string `valid:"required"`
	Url  string  `valid:"url"`
}
func TestValid(t*testing.T){
	g := NewGomegaWithT(t)

	u := Video{
		Name: "Kn",
		Url: "http://www.youtube.com",
	}

	ok, err := govalidator.ValidateStruct(u)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())

}

func Test_Name_cannot_be_blank(t*testing.T) {
	g := NewGomegaWithT(t)

	u := Video{
		Name: "",
		Url: "https://www.youtube.com",
	}
	ok, err := govalidator.ValidateStruct(u)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Name: non zero value required"))
}

func Test_URL(t*testing.T) {
	g := NewGomegaWithT(t)

	u := Video{
		Name: "kn",
		Url: "youtube",
	}
	ok, err := govalidator.ValidateStruct(u)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Url: youtube does not validate as url"))
}
