# Projet API en Golang

Ce projet utilise Go pour interroger une API dont le port change dynamiquement. Il parcourt tous les ports possibles et effectue des requêtes GET et POST sur différents chemins d'API.  
L'objectif final est de pouvoir inscrire sur nom sur le chemin spécifié.

## Fonctionnalités

- Parcours tous les ports de `1` à `65535` pour trouver le port actif de l'API.
- Envoie une requête GET à `/ping` pour tester la connexion au serveur.
- Envoie des requêtes POST à plusieurs chemins, tels que `/signup`, `/check`, `/getUserSecret`, `/getUserLevel`, `/getUserPoints`, `/iNeedAHint`, et `/enterChallenge` avec les corps de requête appropriés.
- Affiche les réponses des requêtes pour confirmer le succès.


