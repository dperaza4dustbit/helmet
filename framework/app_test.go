package framework

import (
	"os"
	"testing"

	"github.com/redhat-appstudio/helmet/api"
	"github.com/redhat-appstudio/helmet/internal/chartfs"
)

func testAppContext() *api.AppContext {
	return api.NewAppContext(
		"helmet-test",
		api.WithNamespace("test-ns"),
	)
}

func testChartFS(t *testing.T) *chartfs.ChartFS {
	t.Helper()
	return chartfs.New(os.DirFS("../test/charts"))
}

func TestNewApp_WithoutImage_Succeeds(t *testing.T) {
	t.Parallel()
	app, err := NewApp(testAppContext(), testChartFS(t))
	if err != nil {
		t.Fatalf("NewApp without WithImage should succeed, got: %v", err)
	}
	if app == nil {
		t.Fatal("expected non-nil App")
	}
}

func TestWithImage_SetsImage(t *testing.T) {
	t.Parallel()
	app, err := NewApp(testAppContext(), testChartFS(t),
		WithImage("test:latest"),
	)
	if err != nil {
		t.Fatalf("NewApp with WithImage should succeed, got: %v", err)
	}
	if app == nil {
		t.Fatal("expected non-nil App")
	}
	if app.image != "test:latest" {
		t.Fatalf("expected image %q, got %q", "test:latest", app.image)
	}
}

func TestWithVerifyRetries_SetsFlags(t *testing.T) {
	t.Parallel()
	app, err := NewApp(testAppContext(), testChartFS(t),
		WithVerifyRetries(1),
	)
	if err != nil {
		t.Fatalf("NewApp with WithVerifyRetries should succeed, got: %v", err)
	}
	if app.flags.VerifyRetries != 1 {
		t.Fatalf("expected VerifyRetries 1, got %d", app.flags.VerifyRetries)
	}
	// Default delay unchanged.
	if app.flags.VerifyRetryDelay <= 0 {
		t.Fatalf("expected default VerifyRetryDelay > 0, got %v", app.flags.VerifyRetryDelay)
	}
}

func TestWithVerifyRetries_ClampsBelowOne(t *testing.T) {
	t.Parallel()
	app, err := NewApp(testAppContext(), testChartFS(t),
		WithVerifyRetries(0),
	)
	if err != nil {
		t.Fatalf("NewApp should succeed, got: %v", err)
	}
	if app.flags.VerifyRetries != 1 {
		t.Fatalf("expected VerifyRetries clamped to 1, got %d", app.flags.VerifyRetries)
	}
}

func TestNewApp_DefaultVerifyRetries(t *testing.T) {
	t.Parallel()
	app, err := NewApp(testAppContext(), testChartFS(t))
	if err != nil {
		t.Fatalf("NewApp should succeed, got: %v", err)
	}
	if app.flags.VerifyRetries != 3 {
		t.Fatalf("expected default VerifyRetries 3, got %d", app.flags.VerifyRetries)
	}
}
