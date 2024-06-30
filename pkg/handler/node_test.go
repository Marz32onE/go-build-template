package handler_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	handler "github.com/marz32one/go-build-template/pkg/handler"
	"github.com/stretchr/testify/assert"
	mock_echo "github.com/yourusername/yourproject/pkg/mocks/echo"
	mock_kubernetes "github.com/yourusername/yourproject/pkg/mocks/kubernetes"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetNodeResources(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockContext := mock_echo.NewMockContext(ctrl)
	mockContext.EXPECT().QueryParam("context").Return("test-context")
	mockContext.EXPECT().JSON(gomock.Any(), gomock.Any()).Return(nil)

	mockClientset := mock_kubernetes.NewMockInterface(ctrl)
	mockNodes := mock_kubernetes.NewMockNodeInterface(ctrl)
	mockNodes.EXPECT().List(context.Background(), metav1.ListOptions{}).Return(&corev1.NodeList{}, nil)
	mockClientset.EXPECT().CoreV1().Return(&mockCoreV1)
	mockCoreV1.EXPECT().Nodes().Return(mockNodes)

	err := handler.GetNodeResources(mockContext, mockClientset)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
