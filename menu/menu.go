package menu

import (
	"errors"
	"fmt"
)

// MusicEntry define music item define
type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

func (m *MusicEntry) String() string {
	return fmt.Sprintf("[%s]: %s | %s | %s | %s", m.Id, m.Name, m.Artist, m.Type, m.Source)
}

// MusicManager define music player manager
type MusicManager struct {
	musics []MusicEntry
}

// NewMusicManager define manager
func NewMusicManager() *MusicManager {
	musics := make([]MusicEntry, 0)
	return &MusicManager{musics}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= m.Len() {
		return nil, errors.New("index out of range")
	}

	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
	for _, v := range m.musics {
		if v.Name == name {
			return &v
		}
	}

	return nil
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	music, err := m.Get(index)
	if err == nil {
		m.musics = append(m.musics[:index], m.musics[index+1:]...)
	}else{
		fmt.Printf("remove %v failed: %v", index, err)
	}

	return music
}

func (m *MusicManager) FindIndex(name string) int {
	for i, v := range m.musics {
		if v.Name == name {
			return i
		}
	}

	return -1
}

func (m *MusicManager) List() {
	for i := 0; i < m.Len(); i++ {
		if me, err := m.Get(i); err == nil {
			fmt.Println(me)
		}

	}
}
