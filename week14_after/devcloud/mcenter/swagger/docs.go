package swagger

import (
	"github.com/go-openapi/spec"
	"github.com/go_8_mage/week14_after/devcloud/mcenter/version"
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
