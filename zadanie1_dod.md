> Zbudować obrazy kontenera z aplikacją opracowaną w punkcie nr 1, które będą pracował na architekturach: linux/arm/v7, linux/arm64/v8 oraz linux/amd64 wykorzystując sterownik docker-container.

To możemy osiągnąć przy użyciu buildx

`docker buildx create --name builder`

`docker buildx inspect --bootstrap`

A następnie przy procesie budowania wskazać na jakie platformy chcemy budować aplikację.

---

> Dockerfile powinien wykorzystywać rozszerzony frontend, zawierać deklaracje wykorzystania cache

Podczas procesu budowania użyto `--cache-to type=inline --cache-from type=inline` aby zapisać w cache 
poszczególne etapy budowania, oraz podzielone budowe obrazu na dwa stage.

---

> umożliwiać bezpośrednie wykorzystanie kodów aplikacji umieszczonych w swoim repozytorium publicznym na GitHub

Aby aplikacja właściwie przechwytywała adres ip użytkownika użyto flagi `--net=host`

Aplikacja jest dostępna pod adresem 54.38.52.140:3456