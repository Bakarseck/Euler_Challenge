#!/bin/bash

# Récupérer le message de commit en ligne de commande
commit_message="$1"

# Mettre à jour le dépôt local
git pull

# Ajouter les fichiers modifiés à la zone de staging
git add .

# Valider les changements avec le message de commit spécifié
git commit -m "$commit_message"

# Pousser les modifications vers le dépôt distant
git push origin master
