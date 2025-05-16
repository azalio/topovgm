package lsblk

import (
	"context"
	"testing"

	"github.com/azalio/lvm2go"
)

func TestLSBLK(t *testing.T) {

	device, err := lvm2go.NewLoopbackDevice(lvm2go.MustParseSize("10M"))
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := device.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	lsblk, err := LSBLK(context.Background(), ColumnPath)
	if err != nil {
		t.Fatal(err)
	}

	found := false
	for i := range RecursiveBlockDevices(lsblk) {
		if path, _ := lsblk[i].GetString(ColumnPath); path == device.Device() {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("device %s not found in lsblk output", device.Device())
	}
}
