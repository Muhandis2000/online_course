package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Mock controllers and middleware
type mockAuthController struct {
	registerCalled bool
	loginCalled    bool
}

func (m *mockAuthController) Register(c *gin.Context) {
	m.registerCalled = true
	c.Status(http.StatusCreated)
}
func (m *mockAuthController) Login(c *gin.Context) {
	m.loginCalled = true
	c.Status(http.StatusOK)
}

type mockCourseController struct {
	getCoursesCalled   bool
	createCourseCalled bool
}

func (m *mockCourseController) GetCourses(c *gin.Context) {
	m.getCoursesCalled = true
	c.Status(http.StatusOK)
}
func (m *mockCourseController) CreateCourse(c *gin.Context) {
	m.createCourseCalled = true
	c.Status(http.StatusCreated)
}

// Mock AuthMiddleware that always allows
func mockAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

// Patch the middleware.AuthMiddleware for testing
func init() {
	// Replace the real middleware with the mock
	middleware.AuthMiddleware = mockAuthMiddleware
}

func TestSetupRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	authCtrl := &mockAuthController{}
	courseCtrl := &mockCourseController{}

	SetupRoutes(router, authCtrl, courseCtrl)

	// Test /register
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(`{}`))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.True(t, authCtrl.registerCalled)

	// Test /login
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/login", strings.NewReader(`{}`))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, authCtrl.loginCalled)

	// Test /api/courses GET
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/courses", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, courseCtrl.getCoursesCalled)

	// Test /api/courses POST
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/courses", strings.NewReader(`{}`))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.True(t, courseCtrl.createCourseCalled)
}