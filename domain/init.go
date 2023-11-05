package domain

import "github.com/doniiawan/tokopedia-scrapper/resource"

type Domain struct {
	resource *resource.Resource
}

func New(resource *resource.Resource) *Domain {
	return &Domain{
		resource: resource,
	}
}
