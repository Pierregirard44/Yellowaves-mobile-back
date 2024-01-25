# Bienvenue sur Yellowaves !  
Application mobile qui répertorie les meilleurs spots de surf !

## Conditions  
Un mois de projet.  
Equipe de 3 ([Quentin](https://www.linkedin.com/in/quentin-kerzerho/), [Pierre](https://www.linkedin.com/in/pierre-girard-468214246/), [Marie](https://www.linkedin.com/in/marie-c-97665394/)).  
Deux semaines front, deux semaines back.  
Rétrospective/Démonstration/MVP.  

# Projet Collectif - Backend en Go

Bienvenue dans le backend de notre projet collectif ! Ce projet est développé en langage Go et nécessite quelques étapes pour être correctement configuré.

## Prérequis

- [MAMP](https://www.mamp.info/) installé pour le serveur MySQL.
- [Visual Studio Code (VSC)](https://code.visualstudio.com/) installé pour le développement.
- Accès à [PHPMyAdmin](http://localhost:8888/phpmyadmin/) via MAMP pour importer le fichier CSV.

## Techno  
- Go
- Visual Studio Code
- MySql
- MAMP

## Configuration de la base de données

1. Lancez MAMP et assurez-vous que le serveur MySQL est en cours d'exécution.
2. Accédez à PHPMyAdmin en utilisant l'URL fournie par MAMP.
3. Importez le fichier CSV fourni dans le dossier `database` à l'aide de PHPMyAdmin.

## Configuration du backend

1. Ouvrez le projet dans Visual Studio Code (VSC).
2. Se déplacer dans le dossier /go-rest-api

   ```bash
   cd /go-rest-api
   
4. Exécutez la commande suivante pour construire l'application Go :

   ```bash
   go build

5. Exécutez la commande suivante pour lancer le serveur Go :
   
   ```bash
   go run . 

## Configuration du frontend

   Passons maintenant au frontend en suivant les étapes sur le [repository](https://github.com/adatechschool/projet-collectif-mobile-front-yellowaves) du frontend.
