# Formula-Calendar

L'objectif de cette application, est de Savoir pour chaque Grand Prix (circuit), quelles seront les gatégories qui rouleront sur le circuit en meme temps que la F1 lors de la saison 2024 de Formule 1.

Savoir pour un circuit donné, quelle catégorie de support rouleront en même temps que la F1.

Ex: Je souhaite savoir quelle catégorie sera présente lors du Grand Prix de Belgique, ainsi j'aurais en retour la date du Grand Prix ainsi que les catégories en support de la F1 (F2, F3)

Possible amélioration : Le faire pour d'autre catégorie de sport auto comme le GT, ELMS, WEC. Ainsi cela permet de savoir quelle catégorie rouleront sur les circuit ou la F1 roule mais également les dates de chacunes des catégories en fonction du circuit demandé.

## Mise en place 
Le voir sous forme de micro-service. Le faire avec une base de données, une partie back, et pourquoi pas une partie front (en plus).
Une base par catégorie (entité). 

## Entités

### Les circuits
Un circuit est utilisé lors d'un week-end de Grand Prix. Celui-ci se verra obligatoirement utilisé par les F1 lors d'un week-end, mais il sera possible que d'autres catégories l'utilisent également en support de la F1.
Un circuit, c'est un nom, une date, un lieu et une liste de catégories de monoplace qui l'utilisent durant le week-end de course.


#### Les catégories de monoplace
Il existe plusieurs catégories de monoplace en plus de la F1 qui sont considérés comme des catégories de promotions afin de se diriger vers la F1. Ici on ne se focalisera que sur les catégories de monoplace qui seront présente en les mêmes week-end de Grand Prix que la F1 pour la saison 2024.
Il y aura donc la Formule 1, la Formule 2, la Formule 3 et la Formule 1 Academy.