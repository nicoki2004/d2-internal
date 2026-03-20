package models

import (
	"testing"
	"time"
)

func TestIsExpiredSoon_TrueWhenWithin60s(t *testing.T) {
	tok := &Token{
		ExpiresIn:  10,
		ReceivedAt: time.Now().Add(-5 * time.Second),
	}
	if !tok.IsExpiredSoon() {
		t.Fatalf("expected token to be expiring soon")
	}
}

func TestIsExpiredSoon_FalseWhenPlentyOfTime(t *testing.T) {
	tok := &Token{
		ExpiresIn:  3600,
		ReceivedAt: time.Now(),
	}
	if tok.IsExpiredSoon() {
		t.Fatalf("expected token not to be expiring soon")
	}
}

func TestIsExpiredSoon_TrueWhenAlreadyExpired(t *testing.T) {
	tok := &Token{
		ExpiresIn:  1,
		ReceivedAt: time.Now().Add(-10 * time.Second),
	}
	if !tok.IsExpiredSoon() {
		t.Fatalf("expected expired token to be expiring soon")
	}
}
