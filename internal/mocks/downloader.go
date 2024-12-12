// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	application "github.com/ava-labs/avalanche-cli/pkg/application"
	mock "github.com/stretchr/testify/mock"
)

// Downloader is an autogenerated mock type for the Downloader type
type Downloader struct {
	mock.Mock
}

// Download provides a mock function with given fields: url
func (_m *Downloader) Download(url string) ([]byte, error) {
	ret := _m.Called(url)

	if len(ret) == 0 {
		panic("no return value specified for Download")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]byte, error)); ok {
		return rf(url)
	}
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllReleasesForRepo provides a mock function with given fields: org, repo, component, kind
func (_m *Downloader) GetAllReleasesForRepo(org string, repo string, component string, kind application.ReleaseKind) ([]string, error) {
	ret := _m.Called(org, repo, component, kind)

	if len(ret) == 0 {
		panic("no return value specified for GetAllReleasesForRepo")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, application.ReleaseKind) ([]string, error)); ok {
		return rf(org, repo, component, kind)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, application.ReleaseKind) []string); ok {
		r0 = rf(org, repo, component, kind)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string, application.ReleaseKind) error); ok {
		r1 = rf(org, repo, component, kind)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestPreReleaseVersion provides a mock function with given fields: org, repo, component
func (_m *Downloader) GetLatestPreReleaseVersion(org string, repo string, component string) (string, error) {
	ret := _m.Called(org, repo, component)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestPreReleaseVersion")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (string, error)); ok {
		return rf(org, repo, component)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) string); ok {
		r0 = rf(org, repo, component)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(org, repo, component)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestReleaseVersion provides a mock function with given fields: org, repo, component
func (_m *Downloader) GetLatestReleaseVersion(org string, repo string, component string) (string, error) {
	ret := _m.Called(org, repo, component)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestReleaseVersion")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (string, error)); ok {
		return rf(org, repo, component)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) string); ok {
		r0 = rf(org, repo, component)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(org, repo, component)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDownloader creates a new instance of Downloader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDownloader(t interface {
	mock.TestingT
	Cleanup(func())
}) *Downloader {
	mock := &Downloader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
