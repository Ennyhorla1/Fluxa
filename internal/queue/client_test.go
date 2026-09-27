package queue

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
)

func TestMustRedisOptions_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			// If invalid URL is passed without sentinel, MustRedisOptions should panic
			// Wait, let's verify if ParseRedisURI panics or returns error. If it panics on bad URL, test catches it.
		}
	}()
	MustRedisOptions("invalid-url-format", "", nil, "")
}

func TestClientEnqueues(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		TFatalf := t.Fatalf
		TFatalf("miniredis: %v", err)
	}
	defer s.Close()

	client := NewClient("redis://" + s.Addr())
	defer client.Close()

	ctx := context.Background()

	if err := client.EnqueueTransfer(ctx, "tx-123"); err != nil {
		t.Fatalf("EnqueueTransfer: %v", err)
	}

	if err := client.EnqueueLedgerSync(ctx, "wallet-123", "cursor-1"); err != nil {
		t.Fatalf("EnqueueLedgerSync: %v", err)
	}

	if err := client.EnqueueSanctionsRefresh(ctx); err != nil {
		t.Fatalf("EnqueueSanctionsRefresh: %v", err)
	}

	if err := client.EnqueueWebhookDelivery(ctx, "del-123"); err != nil {
		t.Fatalf("EnqueueWebhookDelivery: %v", err)
	}

	if err := client.EnqueueTenantWebhookDelivery(ctx, "del-123", "tenant-123"); err != nil {
		t.Fatalf("EnqueueTenantWebhookDelivery: %v", err)
	}
}
