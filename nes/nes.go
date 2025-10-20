package nes

import (
	"os"
	"strings"
)

type NESKeys struct {
	Up     string
	Down   string
	Left   string
	Right  string
	A      string
	B      string
	Start  string
	Select string
}

func GetNESKeys() *NESKeys {
	return &NESKeys{
		Up:     getEnvKey("NES_UP", "ArrowUp"),
		Down:   getEnvKey("NES_DOWN", "ArrowDown"),
		Left:   getEnvKey("NES_LEFT", "ArrowLeft"),
		Right:  getEnvKey("NES_RIGHT", "ArrowRight"),
		A:      getEnvKey("NES_A", "KeyZ"),
		B:      getEnvKey("NES_B", "KeyX"),
		Start:  getEnvKey("NES_START", "Enter"),
		Select: getEnvKey("NES_SELECT", "Space"),
	}
}

func getEnvKey(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return strings.TrimSpace(value)
	}
	return defaultValue
}

func (n *NESKeys) GetAllKeys() []string {
	return []string{
		n.Up,
		n.Down,
		n.Left,
		n.Right,
		n.A,
		n.B,
		n.Start,
		n.Select,
	}
}

func (n *NESKeys) GetMovementKeys() []string {
	return []string{
		n.Up,
		n.Down,
		n.Left,
		n.Right,
	}
}

func (n *NESKeys) GetActionKeys() []string {
	return []string{
		n.A,
		n.B,
	}
}

func (n *NESKeys) GetSystemKeys() []string {
	return []string{
		n.Start,
		n.Select,
	}
}
