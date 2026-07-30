Ascii-art:

Killian Hoarau #khoarau

Théo Bordaix-Barroso Ferreira #botheo

Rafaela Pessanha Souza #rpessanh

Objectif :

Créer un programme qui consiste à prendre un "string" dans un argument, le ressortir en caractères de la table ASCII de manière graphique. Il doit pouvoir gérer une entrée contenant des chiffres, des lettres, des espaces, des caractères spéciaux et \n.

##Structure du projet:
```
ascii-art/
├── main.go          # Point d'entrée du programme
├── standard.txt     # Fichier de police ASCII (requis)
└── README.md        # Ce fichier
```

##Notre projet est constitué de 3 fonctions :

func main() :

 Elle effectue les étapes suivantes :
 1. Récupère le premier argument passé en ligne de commande (le texte à afficher).
 2. Vérifie qu'exactement un argument est fourni — sinon le programme quitte sans message.
 3. Lit le fichier de police "standard.txt" et le découpe ligne par ligne.
 4. Découpe l'argument sur les séquences littérales `\n` pour gérer les sauts de ligne demandés par l'utilisateur.
 5. Appelle PrintFormated pour construire le rendu ASCII, puis l'affiche.

func OpenFile(fileName string) []byte :

 OpenFile ouvre un fichier en lecture seule et retourne son contenu brut en octets.

 Comportement :
  - Si le fichier est introuvable ou illisible, le programme s'arrête immédiatement
     via log.Fatal (affiche l'erreur et appelle os.Exit(1)).
  - Le fichier est fermé automatiquement à la fin de la fonction grâce au defer.

 Retour :
  - []byte : le contenu intégral du fichier sous forme de slice d'octets.

  func PrintFormated(words []string, lines []string) string : 
 1. Pour chaque mot dans words :
    - Si le mot est vide (correspond à un `\n` dans l'entrée), on ajoute
      un simple saut de ligne et on passe au mot suivant.
    - Sinon, on itère sur les 8 rangées visuelles du caractère (index 1 à 8).
      La rangée 0 est un séparateur vide dans le fichier, on la saute donc.
 2. Pour chaque rangée, on parcourt chaque lettre du mot :
    - On calcule la position du caractère dans la police :
      position = valeurASCII(lettre) - 32
     (soustrait 32 pour que l'espace, ASCII 32, soit à l'index 0)
    - On calcule la ligne exacte dans le fichier :
      ligneCourante = (position × 9) + index
    - On concatène cette ligne au résultat.
 3. Après chaque rangée complète (tous les caractères du mot), on ajoute "\n".

##clonage & Lancement

Clonez le projet :  https://zone01normandie.org/git/khoarau/ascii-art.git
cd ascii-art 

Lancement: 
-go run . "<texte>"

Exemple:
- go run . "hello"

résultat:
```
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
```
- go run . "hello\nthere"
- go run . "hello\n\nthere"
- go run . ".:#~"

pour afficher les caractere de fin ligne ($ dans cat -e):
-go run . "<texte>" | cat -e

Exemple:
- go run . "hello" | cat -e

résultat:
```
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
```
- go run . "hello\nhello" | cat -e
- go run . "hello\n\nhello" | cat -e
- go run . ",;:%&" | cat -e

##Cas spécifique:

La commande go run ."<texte>" ne fonctionne avec "!!" ou avec "!<texte>", Pour pouvoir les lancer il faut remplacé les guillemets (") par des apostrophe (').
