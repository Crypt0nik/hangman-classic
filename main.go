package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/fatih/color"
)

var nombrelignealire = 8
var positiondedepart = 0

func main() {

	fichier, err := os.Open("GROSNOOB.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer fichier.Close()

	// Lisez le contenu du fichier
	tnul, err := ioutil.ReadAll(fichier)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	file, err := os.Open("BRAVO.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer file.Close()

	// Lisez le contenu du fichier
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	// Définit les couleurs
	red := color.New(color.FgRed)

	// On choisit un mot au hasard dans words.txt
	level := selectDifficulty()

	// Chargez les mots en fonction du niveau de difficulté choisi
	words, err := loadWordsByDifficulty(level)
	if err != nil {
		fmt.Println("Erreur lors du chargement des mots :", err)
		return
	}

	// On sélectionne un mot au hasard
	rand.Seed(time.Now().UnixNano())
	word := words[rand.Intn(len(words))]
	print(word)
	// On crée une variable pour le nombre de tentatives
	maxattempts := 10
	attempts := maxattempts

	// Créez un tableau pour suivre les lettres correctement devinées
	lettresDevinees := make([]bool, len(word))

	// Créez le mot partiel initial avec des "_"
	motPartiel := make([]string, len(word))
	for i := range motPartiel {
		motPartiel[i] = "_"
	}

	lettresEssayees := make(map[string]bool)

	for attempts > 0 {
		fmt.Print("\033[H\033[2J") // Effacer l'écran
		fmt.Print("\n")
		fmt.Print("\n")
		red.Print(" ----------------------------------\n")
		fmt.Println("  Bienvenue dans le jeu du pendu !")
		red.Print(" ----------------------------------\n")
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Printf("Il vous reste %d tentatives\n", attempts)
		fmt.Print("\n")
		fmt.Print("\n")

		// Affichez le hangman
		hangman(positiondedepart)

		// Affichez le mot à deviner
		red.Print("Mot à deviner: ")
		for i, lettre := range motPartiel {
			if lettresDevinees[i] {
				fmt.Print(lettre)
			} else {
				fmt.Print("_")
			}
		}
		fmt.Print("\n")

		// Demandez une lettre à l'utilisateur
		var input string
		fmt.Println("Entrez une lettre : ")
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
		// Normalisez la lettre entrée par l'utilisateur
		input = normalizeString(input)

		// Mettez la lettre en minuscule
		input = strings.ToLower(input)

		// Si la lettre est dans le mot, affichez le mot avec les lettres trouvées
		lettreTrouvee := false
		for i, lettre := range word {
			if strings.ToLower(string(lettre)) == input {
				lettreTrouvee = true
				motPartiel[i] = string(lettre)
				lettresDevinees[i] = true
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

		if toutesLettresTrouvees {
			fmt.Print("\033[H\033[2J") // Effacer l'écran
			fmt.Print("Mot à deviner: ")
			for i, lettre := range motPartiel {
				if lettresDevinees[i] {
					fmt.Print(lettre)
				} else {
					fmt.Print("_")
				}
			}
			fmt.Print("\033[H\033[2J")
			fmt.Print("\n")
			fmt.Print("\n")
			fmt.Print("\n")
			fmt.Print("\n")
			red.Print("                        -------------------------------------\n")
			fmt.Println("                       Bravo, vous avez trouvé le mot :", word)
			red.Print("                        -------------------------------------\n")
			fmt.Print("\n")
			fmt.Print("\n")
			fmt.Print("\n")
			fmt.Println(string(content))
			fmt.Print("\n")
			fmt.Print("\n")
			fmt.Print("\n")
			fmt.Print("\n")
			break
		}

		// Si la lettre n'est pas dans le mot, affichez un message d'erreur et décrémentez le nombre de tentatives
		if !lettreTrouvee {
			fmt.Print("\033[H\033[2J") // Effacer l'écran
			fmt.Printf("Mot à deviner: ")
			for i, lettre := range motPartiel {
				if lettresDevinees[i] {
					fmt.Print(lettre)
				} else {
					fmt.Print("_")
				}
			}
			fmt.Print("\n")
			fmt.Println("Lettre incorrecte :", input)
			attempts--
			positiondedepart += nombrelignealire
		}
	}

	if attempts == 0 {
		fmt.Print("\033[H\033[2J") // Effacer l'écran
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print("\n")
		red.Print("                     -----------------------------------------\n")
		fmt.Println("                  Vous n'avez plus de tentatives, vous avez perdu !")
		fmt.Println("                              Le mot était :", word)
		red.Print("                     -----------------------------------------\n")
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Println(string(tnul))
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print("\n")
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

	// Lisez les 8 lignes suivantes
	for i := 0; i < nombrelignealire; i++ {
		if scanner.Scan() {
			ligne := scanner.Text()
			fmt.Println(ligne)
		} else if scanner.Err() !=
			nil {
			if scanner.Err() != nil {
				log.Fatal(scanner.Err())
			}
		}
	}
}

func selectDifficulty() string {
	red := color.New(color.FgRed)
	yellow := color.New(color.FgYellow)
	green := color.New(color.FgHiGreen)
	var level string
	fmt.Print("\033[H\033[2J") // Effacer l'écran
	fmt.Print("\n")
	fmt.Print("\n")
	red.Print(" ----------------------------------\n")
	fmt.Println("  Bienvenue dans le jeu du pendu !")
	red.Print(" ----------------------------------\n")
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("Sélectionnez un ")
	red.Print("niveau ")
	fmt.Println("de difficulté :")
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("\n")
	green.Println("                    1. Facile")
	fmt.Print("\n")
	fmt.Print("\n")
	yellow.Println("                    2. Moyen")
	fmt.Print("\n")
	fmt.Print("\n")
	red.Println("                    3. Difficile")
	fmt.Print("\n")

	for {
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print("Entrez le ")
		red.Print("numéro ")
		fmt.Print("du niveau de difficulté : ")
		fmt.Print("\n")

		_, err := fmt.Scan(&level)
		if err != nil {
			log.Fatal(err)
			return ""
		}

		switch level {
		case "1":
			return "facile"
		case "2":
			return "moyen"
		case "3":
			return "difficile"
		default:
			fmt.Println("Niveau invalide. Veuillez réessayer.")
		}
	}
}
func loadWordsByDifficulty(level string) ([]string, error) {
	filePath := ""
	switch level {
	case "facile":
		filePath = "mots_faciles.txt"
	case "moyen":
		filePath = "mots_moyens.txt"
	case "difficile":
		filePath = "mots_difficiles.txt"
	default:
		return nil, fmt.Errorf("Niveau de difficulté non reconnu")
	}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	words := strings.Split(string(data), "\n")
	return words, nil
}
func normalizeString(input string) string {
	var result strings.Builder

	for _, r := range input {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		result.WriteRune(r)
	}

	return result.String()
}
