package hashtable

import (
	"testing"
)

func TestHash(t *testing.T) {
	h := NewHash(10)

	h.Add("key1", "val1")
	h.Add("key2", "val2")
	h.Add("key3", "val3")
	h.Add("key4", "val4")

	if h.fillSize != 4 {
		t.Errorf("Failed to add the element to Hash. Expecting 4, instead %d", h.fillSize)
	}

	e := h.Get("key4")
	if e != "val4" {
		t.Errorf("Failed to get the element to Hash. Expecting val4, instead %d", e)
	}

	e = h.Get("key3")
	if e != "val3" {
		t.Errorf("Failed to get the element to Hash. Expecting val3, instead %d", e)
	}

	nh := NewHash(2)

	nh.Add("key1", "val1")
	nh.Add("key2", "val2")

	if nh.fillSize != 2 {
		t.Errorf("Failed to add the element to Hash. Expecting 2, instead %d", nh.fillSize)
	}

	nh.Add("key3", "val3")

	if nh.fillSize != 3 {
		t.Errorf("Failed to add the element to Hash. Expecting 3, instead %d", nh.fillSize)
	}

	e = nh.Get("key3")
	if e != "val3" {
		t.Errorf("Failed to get the element to Hash. Expecting val3, instead %s", e)
	}

	e = nh.Get("key4")
	if e != nil {
		t.Errorf("Failed to get the element to Hash. Expecting nil, instead %s", e)
	}
}
