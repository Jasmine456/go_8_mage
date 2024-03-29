package swagger

import (
	"github.com/Jasmine456/go_8_mage/week14_after/devcloud/mcenter/version"
	"github.com/go-openapi/spec"
)

func Docs(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "BookService",
			Description: "Resource for managing Books",
			Contact: &spec.ContactInfo{
				ContactInfoProps: spec.ContactInfoProps{
					Name:  "john",
					Email: "john@doe.rp",
					URL:   "auth://johndoe.org",
				},
			},
			License: &spec.License{
				LicenseProps: spec.LicenseProps{
					Name: "MIT",
					URL:  "auth://mit.org",
				},
			},
			Version: version.Short(),
		},
	}

}
