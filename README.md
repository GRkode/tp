## Traitement Batch avec Hadoop HDFS et Map Reduce

### Tools
- Apache Hadoop Version: 2.7.2.
- Java Version 1.8
- Docker version 24.0.7

```bash
## Télécharger l'image docker
docker pull totofunku/bigdata-cours:latest
## Création de réseau bridge pour les trois conteneurs
docker network create --driver=bridge hadoop
```

### Créer et lancer les trois contenaires (1 main et 2 suiveurs)
<img src="pictures/big_data1.png">

```bash
## Entrer dans le contenaire main
docker exec -it hadoop-master bash
## Lancer hadoop et yarn avec le fichier suivant
/start-hadoop.sh
```
<img src="pictures/big_data2.png">

```bash
## Création du répertoire dans HDFS
hdfs dfs –mkdir -p input
## Copie du fichier purchases dans le répertoire input
hdfs dfs –put purchases.txt input
```
<img src="pictures/big_data3.png">

### Interface de visualisation des Jobs sur l'adresse http://localhost:8088
<img src="pictures/big_data4.png">

### Création de projet java avec maven
<img src="pictures/big_data5.png">

<hr>

## Ajout de code Map et Reduce

### Ajout du mapper
<img src="pictures/big_data6.png">

### Ajout du reducer
<img src="pictures/big_data7.png">

### Ajout de la classe main
<img src="pictures/big_data8.png">

### Utilisation de fichier pour tester le code
<img src="pictures/big_data9.png">

### Création d'une configuration de type Application
<img src="pictures/big_data10.png">

###Exécution du code avec le cluster hadoop
<img src="pictures/big_data11.png">

<hr>

## Écriture de job map reduce pour déterminer le total des ventes par magasin

### Code du mapping
<img src="pictures/big_data12.png">

### code du reducer
<img src="pictures/big_data13.png">

### Résultat d'execution du code sur le cluster
<img src="pictures/big_data14.png">