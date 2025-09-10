package caimport

import (
	"testing"
)

// TestImportOptions tests the ImportOptions struct
func TestImportOptions(t *testing.T) {
	options := ImportOptions{
		CertPath:   "/path/to/cert.pem",
		DockerHost: "registry.example.com",
		Force:      true,
	}

	if options.CertPath != "/path/to/cert.pem" {
		t.Errorf("Expected CertPath to be '/path/to/cert.pem', got '%s'", options.CertPath)
	}

	if options.DockerHost != "registry.example.com" {
		t.Errorf("Expected DockerHost to be 'registry.example.com', got '%s'", options.DockerHost)
	}

	if options.Force != true {
		t.Errorf("Expected Force to be true, got %v", options.Force)
	}
}

// TestVerifyCertificate tests the VerifyCertificate function
func TestVerifyCertificate(t *testing.T) {
	// This is a placeholder test since we don't have a real certificate file in the test environment
	// In a real scenario, we would test with a valid certificate file
	
	// Test with a non-existent file
	_, err := VerifyCertificate("/non/existent/file.crt")
	if err == nil {
		t.Error("Expected error when verifying non-existent certificate file, got nil")
	}
}

// TestImportCA tests the ImportCA function
func TestImportCA(t *testing.T) {
	// This is a placeholder test since we can't actually import certificates in the test environment
	// In a real scenario, we would test with a valid certificate file
	
	options := ImportOptions{
		CertPath:   "/non/existent/file.crt",
		DockerHost: "",
		Force:      false,
	}

	err := ImportCA(options)
	if err == nil {
		t.Error("Expected error when importing non-existent certificate file, got nil")
	}
}