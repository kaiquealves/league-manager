package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type competidor struct {
	nome string
}

type match struct {
	mandante        competidor
	visitante       competidor
	rodada          int
	placarMandante  int
	placarVisitante int
}

func reverseIntSlice(s []int) []int {
	a := make([]int, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func shuffleStringSlice(s []string) []string {

	a := make([]string, len(s))
	copy(a, s)

	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	return a
}

func minPlayoffsRounds(qtdeTimes int, repescagem bool) int {
	rounds := 1

	if qtdeTimes > 2 {

		if repescagem {
			rounds++
		}

		for qtdeTimes > 1 {
			rounds++
			qtdeTimes /= 2
		}
	}

	return rounds
}

func jogosPorRodada(qtdeTimes int) []int {

	rounds := minPlayoffsRounds(qtdeTimes, false)

	var jogosRodada []int

	jogosRodada = append(jogosRodada, 1)

	for i := 1; i < rounds-1; i++ {

		jogosRodada = append(jogosRodada, jogosRodada[i-1]*2)
	}

	jogosRodada = append(jogosRodada, 1)

	return reverseIntSlice(jogosRodada)
}

func criaConfrontos(times []string) []string {

	times = shuffleStringSlice(times)
	jogosRodada := jogosPorRodada(len(times))
	rodadas := 5

	fmt.Println("Sorteio Times: ", times)

	var confrontos []string

	for r := 1; r <= rodadas; r++ {

		for j := 1; j <= jogosRodada[r-1]; j++ {

			if len(times) >= 2 {
				confrontos = append(confrontos, "Rodada "+strconv.Itoa(r)+" - Jogo "+strconv.Itoa(j)+" - "+times[0]+" x "+times[1])
				a := len(times)
				times = times[2:a]
				times = append(times, "Venc. Jogo "+strconv.Itoa(j)+" - Rodada "+strconv.Itoa(r))
			}

		}
	}

	return confrontos

}

func sorteiaGrupos(times []string, qtdeGrupos int) [][]string {

	times = shuffleStringSlice(times)

	//TO-DO: Inicializar este slice 2D de maneira dinâmica
	grupos := [][]string{{"GRUPO 1"}, {"GRUPO 2"}, {"GRUPO 3"}, {"GRUPO 4"}}
	grupo := 0

	for i := 0; i < len(times); i++ {
		grupo = i % qtdeGrupos
		grupos[grupo] = append(grupos[grupo], times[i])
	}

	return grupos

}

func main() {

	times := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q"}
	rand.Seed(time.Now().UnixNano())

	competidores := len(times)
	rounds := minPlayoffsRounds(competidores, false)
	var qtdeJogosPorRodada []int
	qtdeJogosPorRodada = jogosPorRodada(competidores)

	fmt.Println("Qtde de Times: ", competidores)
	fmt.Println("Qtde Rounds: ", rounds)
	fmt.Println("Jogos Por Rodada: ", qtdeJogosPorRodada)
	fmt.Println("Confrontos: ", criaConfrontos(times))
	fmt.Println("------ Iniciando Sorteio dos Grupos - Equipes Participantes: ", competidores)
	fmt.Println("Grupos: ", sorteiaGrupos(times, 4))

}
