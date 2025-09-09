package main

import (
	"sync"
	"testing"
	"time"
)

func TestSoma(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			t.Log("rodando em paralelo...")
			time.Sleep(200 * time.Millisecond)
		}
	}()

	resultado := Soma(2, 3)
	if resultado != 5 {
		t.Errorf("esperando 5 recebeu %d", resultado)
	}

	// Espera a goroutine acabar
	wg.Wait()
}