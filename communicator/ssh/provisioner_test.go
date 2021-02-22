package ssh

import (
	"testing"

	"github.com/zclconf/go-cty/cty"
)

func TestProvisioner_connInfo(t *testing.T) {
	v := cty.ObjectVal(map[string]cty.Value{
		"type":         cty.StringVal("ssh"),
		"user":         cty.StringVal("root"),
		"password":     cty.StringVal("supersecret"),
		"private_key":  cty.StringVal("someprivatekeycontents"),
		"certificate":  cty.StringVal("somecertificate"),
		"host":         cty.StringVal("127.0.0.1"),
		"port":         cty.StringVal("22"),
		"timeout":      cty.StringVal("30s"),
		"bastion_host": cty.StringVal("127.0.1.1"),
	})

	conf, err := parseConnectionInfo(v)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if conf.User != "root" {
		t.Fatalf("bad: %v", conf)
	}
	if conf.Password != "supersecret" {
		t.Fatalf("bad: %v", conf)
	}
	if conf.PrivateKey != "someprivatekeycontents" {
		t.Fatalf("bad: %v", conf)
	}
	if conf.Certificate != "somecertificate" {
		t.Fatalf("bad: %v", conf)
	}
	if conf.Host != "127.0.0.1" {
		t.Fatalf("bad: %v", conf)
	}
	if conf.Port != 22 {
		t.Fatalf("bad: %v", conf)
	}
	if conf.Timeout != "30s" {
		t.Fatalf("bad: %v", conf)
	}
	if conf.ScriptPath != DefaultUnixScriptPath {
		t.Fatalf("bad: %v", conf)
	}
	if conf.TargetPlatform != TargetPlatformUnix {
		t.Fatalf("bad: %v", conf)
	}
	if conf.BastionHost != "127.0.1.1" {
		t.Fatalf("bad: %v", conf)
	}
	if conf.BastionPort != 22 {
		t.Fatalf("bad: %v", conf)
	}
	if conf.BastionUser != "root" {
		t.Fatalf("bad: %v", conf)
	}
	if conf.BastionPassword != "supersecret" {
		t.Fatalf("bad: %v", conf)
	}
	if conf.BastionPrivateKey != "someprivatekeycontents" {
		t.Fatalf("bad: %v", conf)
	}
}

func TestProvisioner_connInfoIpv6(t *testing.T) {
	v := cty.ObjectVal(map[string]cty.Value{
		"type":         cty.StringVal("ssh"),
		"user":         cty.StringVal("root"),
		"password":     cty.StringVal("supersecret"),
		"private_key":  cty.StringVal("someprivatekeycontents"),
		"host":         cty.StringVal("::1"),
		"port":         cty.StringVal("22"),
		"timeout":      cty.StringVal("30s"),
		"bastion_host": cty.StringVal("::1"),
	})

	conf, err := parseConnectionInfo(v)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if conf.Host != "[::1]" {
		t.Fatalf("bad: %v", conf)
	}

	if conf.BastionHost != "[::1]" {
		t.Fatalf("bad %v", conf)
	}
}

func TestProvisioner_connInfoHostname(t *testing.T) {
	v := cty.ObjectVal(map[string]cty.Value{
		"type":         cty.StringVal("ssh"),
		"user":         cty.StringVal("root"),
		"password":     cty.StringVal("supersecret"),
		"private_key":  cty.StringVal("someprivatekeycontents"),
		"host":         cty.StringVal("example.com"),
		"port":         cty.StringVal("22"),
		"timeout":      cty.StringVal("30s"),
		"bastion_host": cty.StringVal("example.com"),
	})

	conf, err := parseConnectionInfo(v)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if conf.Host != "example.com" {
		t.Fatalf("bad: %v", conf)
	}

	if conf.BastionHost != "example.com" {
		t.Fatalf("bad %v", conf)
	}
}

func TestProvisioner_connInfoEmptyHostname(t *testing.T) {
	v := cty.ObjectVal(map[string]cty.Value{
		"type":        cty.StringVal("ssh"),
		"user":        cty.StringVal("root"),
		"password":    cty.StringVal("supersecret"),
		"private_key": cty.StringVal("someprivatekeycontents"),
		"port":        cty.StringVal("22"),
		"timeout":     cty.StringVal("30s"),
	})

	_, err := parseConnectionInfo(v)
	if err == nil {
		t.Fatalf("bad: should not allow empty host")
	}
}
