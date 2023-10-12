package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// on choisis un mot au pif dans words.txt
	data, err := ioutil.ReadFile("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(data), "\n")

	// on séléctionne un mot au hasard
	rand.Seed(time.Now().UnixNano())
	word := words[rand.Intn(len(words))]
	println("Mot aléatoire", word)

	// on créer une variable pour le nombre de lettre du mot
	lenword := len(word)/2 - 1
	println("Nombre de lettre", lenword)

	// on créer une variable pour le nombre de tentatives
	maxattempts := 10
	attempts := maxattempts

	print("Bienvenue dans le jeu du pendu", "\n")

	for attempts > 0 {
		fmt.Printf("il vous reste %d tentatives\n", attempts)
		// demander une lettre à l'utilisateur
		var input string
		println("Entrez une lettre : ")
		if input == word {
		}
		_, err = fmt.Scan(&input)
		if err != nil {
			log.Fatal(err)
			return
		}
		// on met la lettre en minuscule
		input = strings.ToLower(input)

		// si la lettre est dans le mot, on affiche le mot avec les lettres trouvées
		lettreTrouvee := false

		for i, lettre := range word {
			if strings.ToLower(string(lettre)) == input {
				fmt.Printf("Bonne lettre trouvée à la position %d : %s\n", i, input)
				lettreTrouvee = true
			}
		}
		// si la lettre n'est pas dans le mot, on affiche un message d'erreur et on décrémente le nombre de tentatives
		if !lettreTrouvee {
			fmt.Println("Lettre incorrecte :", input)
			attempts--
		} else {
			motcomplete := true
			for _, lettre := range word {
				if !strings.Contains(input, string(lettre)) {
					motcomplete = false
					break
				}
			}
			if motcomplete {
				fmt.Println("Vous avez trouvé le mot !")
				break
			}
		}
	}
	if attempts == 0 {
		fmt.Println("Vous n'avez plus de tentatives, vous avez perdu !", "\n", "Le mot était :", word)
	}
}
