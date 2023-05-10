![Logo](https://go.dev/images/go-logo-white.svg)
## ASCII-ART-WEB

* Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of the project [ascii-art](https://learn.zone01dakar.sn/git/nifaye/ascii-art)
.


## Installation

 - To use this program clone this Git repository on your local machine

```bash
  git clone https://learn.zone01dakar.sn/git/nifaye/ascii-art-web
```
-  Open Terminal and install go packages
```bash
  apt install golang
```
-  Run program using the next command
```bash
  go run .
```
- Click on this link on your favorite navigator : 
-  http://localhost:8080

- Enjoy 😇🕺🤸🏼‍♂️

## Author

- [@nifaye](https://learn.zone01dakar.sn/git/nifaye)
- [@igueye](https://learn.zone01dakar.sn/git/igueye)

## Implementation details
### English

- The program sets up a web server that serves a static file directory and handles HTTP requests using various handlers. The main function starts the server, listens on port 8080, and prints a message to the console with a link to the server.

- There are two handler functions defined: ServerHandler and HomeHandler. ServerHandler serves an HTML template for the index page, and HomeHandler serves an HTML template for the home page. When a user submits a form on the home page, the input is passed to a function called AsciiArt, which converts the text input into ASCII art using a banner selected by the user. The resulting ASCII art is then displayed on the home page.

- There are also some error handling functions that serve an error page if there is an issue with the user's input or if the server encounters an error.

- Overall, this program allows users to convert text into ASCII art using a web interface.
#
### French
- Le programme configure un serveur web qui sert un répertoire de fichiers statiques et gère les requêtes HTTP à l'aide de divers gestionnaires. La fonction principale lance le serveur, écoute sur le port 8080, et affiche un message sur la console avec un lien vers le serveur.

- Deux fonctions de gestionnaire sont définies : ServerHandler et HomeHandler. ServerHandler sert un modèle HTML pour la page d'index, et HomeHandler sert un modèle HTML pour la page d'accueil. Lorsqu'un utilisateur soumet un formulaire sur la page d'accueil, l'entrée est passée à une fonction appelée AsciiArt, qui convertit l'entrée de texte en art ASCII en utilisant une bannière sélectionnée par l'utilisateur. L'art ASCII résultant est ensuite affiché sur la page d'accueil.

- Il y a également des fonctions de gestion des erreurs qui servent une page d'erreur en cas de problème avec l'entrée de l'utilisateur ou si le serveur rencontre une erreur.

- Dans l'ensemble, ce programme permet aux utilisateurs de convertir du texte en art ASCII en utilisant une interface web.