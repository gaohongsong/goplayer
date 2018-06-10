package menu

import "testing"

func TestMusicManager_Opts(t *testing.T) {
	mm := NewMusicManager()

	if mm == nil {
		t.Errorf("new MenuManager failed")
	}

	if mm.Len() != 0 {
		t.Errorf("new MenuManager failed, not empty")
	}

	m0 := &MusicEntry{"1", "my heart will go on", "celli dion", "http://ssdfs/232", "mp3"}
	mm.Add(m0)
	mm.List()

	if mm.Len() != 1 {
		t.Errorf("add failed")
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Errorf("find failed")
	}

	if m.Id != m0.Id || m.Artist != m0.Artist || m.Source != m0.Source || m.Type != m0.Type {
		t.Error("find item mismatch")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("get failed", err)
	}


	m1 := mm.FindIndex(m0.Name)
	if m1 == -1 {
		t.Errorf("find index failed")
	}

	m = mm.Remove(0)

	if m == nil || mm.Len() !=0 {
		t.Error("remove failed", err)
	}

}
