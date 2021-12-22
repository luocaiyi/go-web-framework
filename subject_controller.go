package main

import (
	"github.com/luocaiyi/go-web-framework/framework"
	"net/http"
)

func SubjectAddController(c *framework.Context) error {
	c.Json(http.StatusOK, "ok, SubjectAddController")
	return nil
}

func SubjectListController(c *framework.Context) error {
	c.Json(http.StatusOK, "ok, SubjectListController")
	return nil
}

func SubjectDelController(c *framework.Context) error {
	c.Json(http.StatusOK, "ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	c.Json(http.StatusOK, "ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	c.Json(http.StatusOK, "ok, SubjectGetController")
	return nil
}

func SubjectNameController(c *framework.Context) error {
	c.Json(http.StatusOK, "ok, SubjectNameController")
	return nil
}
