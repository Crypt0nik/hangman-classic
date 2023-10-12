package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

var nombrelignealire = 8
var positiondedepart = 0

func main() {
	// On choisis un mot au pif dans words.txt

	fmt.Print("\033[H\033[2J")
	data, err := ioutil.ReadFile("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(data), "\n")

	// On séléctionne un mot au hasard

	rand.Seed(time.Now().UnixNano())
	word := words[rand.Intn(len(words))]

	// On créer une variable pour le nombre de tentatives

	maxattempts := 10
	attempts := maxattempts

	print("Bienvenue dans le jeu du pendu", "\n")

	// Créez un tableau pour suivre les lettres correctement devinées

	lettresDevinees := make([]bool, len(word))

	// Créez le mot partiel initial avec des "_"

	motPartiel := make([]string, len(word))
	for i := range motPartiel {
		motPartiel[i] = "_"
	}

	lettresEssayees := make(map[string]bool)

	for attempts > 0 {
		fmt.Printf("il vous reste %d tentatives\n", attempts)

		// Demander une lettre à l'utilisateur

		var input string
		println("Entrez une lettre : ")
		if input == word {
		}
		if lettresEssayees[input] {
			fmt.Printf("Vous avez déjà essayé la lettre %s.\n", input)
			continue
		}
		_, err = fmt.Scan(&input)
		if err != nil {
			log.Fatal(err)
			return
		}
		// On met la lettre en minuscule

		input = strings.ToLower(input)

		fmt.Print("\033[H\033[2J")

		// Si la lettre est dans le mot, on affiche le mot avec les lettres trouvées

		lettreTrouvee := false

		for i, lettre := range word {
			if strings.ToLower(string(lettre)) == input {
				lettreTrouvee = true
				motPartiel[i] = string(lettre)
				lettresDevinees[i] = true
				hangman(positiondedepart)
			}
		}

		lettresEssayees[input] = true

		// Vérifiez si toutes les lettres ont été devinées

		toutesLettresTrouvees := true
		for _, trouvee := range lettresDevinees {
			if !trouvee {
				toutesLettresTrouvees = false
				break
			}
		}
		fmt.Println("Mot à deviner :", strings.Join(motPartiel, " "))

		if toutesLettresTrouvees {
			fmt.Println("Bravo, vous avez trouvé le mot :", word)
			break
		}

		// Si la lettre n'est pas dans le mot, on affiche un message d'erreur et on décrémente le nombre de tentatives

		if !lettreTrouvee {
			fmt.Println("Lettre incorrecte :", input)
			attempts--
			hangman(positiondedepart)
			positiondedepart += nombrelignealire
		}
	}
	if attempts == 0 {
		fmt.Println("Vous n'avez plus de tentatives, vous avez perdu !", "\n", "Le mot était :", word)
	}
}

func hangman(startPosition int) {
	filePath := "hangman.txt" // Remplacez par le chemin de votre fichier
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Créez un reader à partir du contenu du fichier
	reader := strings.NewReader(string(content))

	// scanner
	scanner := bufio.NewScanner(reader)

	for i := 0; i < startPosition; i++ {
		if !scanner.Scan() {
			log.Fatal(scanner.Err())
		}
	}

	// Lire les 8 lignes suivantes
	for i := 0; i < nombrelignealire; i++ {
		if scanner.Scan() {
			ligne := scanner.Text()
			fmt.Println(ligne)
		} else if scanner.Err() != nil {
			log.Fatal(scanner.Err())
		}
	}
}
