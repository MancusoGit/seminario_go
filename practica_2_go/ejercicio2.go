package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type billetera struct {
	id       string
	nombre   string
	apellido string
}

type transaccion struct {
	monto      float64
	emisorID   string
	receptorID string
	timestamp  time.Time
}

type bloque struct {
	hash         string
	hashAnterior string
	timestamp    time.Time
	data         transaccion
	sig          *bloque
}

type blockchain struct {
	inicio *bloque
}

func generarHash(hashAnt string, timestamp time.Time, data transaccion) string {
	input := hashAnt + timestamp.String()
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

func (b *billetera) crearBilletera(id, nombre, apellido string) {
	b.id = id
	b.nombre = nombre
	b.apellido = apellido
}

func (bc *blockchain) insertarBloque(data transaccion) {
	var hashAnt string
	var ult *bloque

	if bc.inicio == nil {
		hashAnt = "0"
	} else {
		ult = bc.inicio
		for ult.sig != nil {
			ult = ult.sig
		}
		hashAnt = ult.hash
	}

	timestamp := time.Now()
	hash := generarHash(hashAnt, timestamp, data)

	nuevoBloque := bloque{
		hash:         hash,
		hashAnterior: hashAnt,
		timestamp:    timestamp,
		data:         data,
		sig:          nil,
	}

	if bc.inicio == nil {
		bc.inicio = &nuevoBloque
	} else {
		ult.sig = &nuevoBloque
	}
}

func (bc *blockchain) getBalance(billeteraID string) float64 {
	var balance float64
	balance = 0
	actual := bc.inicio
	for actual != nil {
		if actual.data.emisorID == billeteraID {
			balance -= actual.data.monto
		}
		if actual.data.receptorID == billeteraID {
			balance += actual.data.monto
		}

		actual = actual.sig
	}
	return balance
}

func (bc *blockchain) esValido() bool {
	act := bc.inicio
	for act != nil && act.sig != nil {
		hashEsperado := generarHash(act.hash, act.sig.timestamp, act.sig.data)
		if hashEsperado != act.sig.hash {
			return false
		}
		act = act.sig
	}
	return true
}

func (bc *blockchain) enviarTransaccion(emisorId, receptorId string, monto float64) bool {
	if bc.getBalance(emisorId) < monto {
		fmt.Println("saldo insuficiente.")
		return false
	}

	t := transaccion{
		monto:      monto,
		emisorID:   emisorId,
		receptorID: receptorId,
		timestamp:  time.Now(),
	}

	bc.insertarBloque(t)
	return true
}

func main() {

	bc := blockchain{}

	var b1, b2 billetera

	b1.crearBilletera("1", "Iara", "Cigalino")
	b2.crearBilletera("2", "Tomas", "Mancuso")

	bc.insertarBloque(transaccion{monto: 1000, emisorID: "init", receptorID: b1.id, timestamp: time.Now()})

	bc.enviarTransaccion(b1.id, b2.id, 200)

	bc.enviarTransaccion(b2.id, b2.id, 300)

	fmt.Printf("el perfil %s %s, con id %s es: %f \n", b1.apellido, b1.nombre, b1.id, bc.getBalance(b1.id))
	fmt.Printf("el perfil %s %s, con id %s es: %f \n", b2.apellido, b2.nombre, b2.id, bc.getBalance(b2.id))
	fmt.Println()

	fmt.Println("validez de la cadena: ", bc.esValido())

}
