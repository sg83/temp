// Code generated by mockery v2.23.2. DO NOT EDIT.

package mocks
/*
import (
	models "github.com/sg83/go-microservice/article-api/models"
	mock "github.com/stretchr/testify/mock"
)

// ArticlesData is an autogenerated mock type for the ArticlesData type
type ArticlesData struct {
	mock.Mock
}

// AddArticle provides a mock function with given fields: ar
func (_m *ArticlesData) AddArticle(ar models.Article) error {
	ret := _m.Called(ar)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Article) error); ok {
		r0 = rf(ar)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *ArticlesData) Close() {
	_m.Called()
}

// GetArticleByID provides a mock function with given fields: id
func (_m *ArticlesData) GetArticleByID(id int) (*models.Article, error) {
	ret := _m.Called(id)

	var r0 *models.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.Article, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *models.Article); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// init provides a mock function with given fields:
func (_m *ArticlesData) init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewArticlesData interface {
	mock.TestingT
	Cleanup(func())
}

// NewArticlesData creates a new instance of ArticlesData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewArticlesData(t mockConstructorTestingTNewArticlesData) *ArticlesData {
	mock := &ArticlesData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
*/