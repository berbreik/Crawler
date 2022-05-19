package handler

import "assignment_v0.2/models"

type serviceInterfae interface {
	Crawl(urls []string) (*models.ResponseData, error)
}